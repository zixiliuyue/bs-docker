

package resthdrs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/metric"
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/resthdrs/filters"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/server/resthdrs/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/storages/sqlstore"
	"github.com/dchest/uniuri"
	"github.com/emicklei/go-restful"
	"github.com/iancoleman/strcase"
	"time"
)

func initCluster(cluster *m.Cluster) error {
	return nil
}

func createCluster(cluster *m.Cluster) (*m.Cluster, error) {
	clusterId := cluster.ID
	if sqlstore.GetCluster(clusterId) != nil {
		return nil, utils.NewClusterAreadyExistError("create failed, cluster with this id already exists")
	}

	err := sqlstore.CreateCluster(cluster)
	if err != nil {
		return nil, utils.NewCannotCreateClusterError(fmt.Sprintf("can not create cluster, error: %s", err))
	}

	err = initCluster(cluster)
	if err != nil {
		return nil, utils.NewClusterInitFailedError(fmt.Sprintf("cluster init failed, error: %s", err))
	}
	return sqlstore.GetCluster(clusterId), nil
}

func createClusterWithExternalInfo(cluster *m.Cluster, externalClusterInfo *m.BCSClusterInfo,
	response *restful.Response) {
	cluster, err := createCluster(cluster)
	// convert type name to screaming snake
	errorCode := strcase.ToScreamingSnake(fmt.Sprint(reflect.TypeOf(cluster)))
	if err != nil {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", "POST").Inc()
		message := fmt.Sprintf("errcode: %d, can not create cluster, error: %s", common.BcsErrApiInternalDbError, err)
		WriteClientError(response, errorCode, message)
		return
	}

	err = sqlstore.CreateBCSClusterInfo(externalClusterInfo)
	if err != nil {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", "POST").Inc()
		message := fmt.Sprintf("errcode: %d, can not create external cluster info, error: %s",
			common.BcsErrApiInternalDbError, err)
		WriteServerError(response, "CANNOT_CREATE_EXTERNAL_CLUSTER_INFO", message)
		return
	}
	response.WriteEntity(*sqlstore.GetCluster(cluster.ID))

}

// PlainCluster

// CreatePlainClusterForm xxx
type CreatePlainClusterForm struct {
	ID string `json:"id" validate:"required"`
}

// CreatePlainCluster creates a "plain" cluster for current user
func CreatePlainCluster(request *restful.Request, response *restful.Response) {
	form := CreatePlainClusterForm{}
	request.ReadEntity(&form)

	err := validate.Struct(&form)
	if err != nil {
		response.WriteEntity(FormatValidationError(err))
		return
	}

	// Prepend a fixed prefix to avoid id conflict across providers
	clusterId := m.ClusterIdPrefixPlain + form.ID
	user := filters.GetUser(request)
	cluster := &m.Cluster{
		ID:        clusterId,
		Provider:  m.ClusterProviderPlain,
		CreatorId: user.ID,
	}
	cluster, err = createCluster(cluster)
	// convert type name to screaming snake
	errorCode := strcase.ToScreamingSnake(fmt.Sprint(reflect.TypeOf(cluster)))
	if err != nil {
		WriteClientError(response, errorCode, fmt.Sprintf("can not create cluster, error: %s", err))
		return
	}
	// init plain cluster permissions
	for _, name := range []string{m.ClusterPermNameView, m.ClusterPermNameEdit} {
		err := sqlstore.SaveUserClusterPerm(m.PermBackendTypeSyncOnce, user, cluster, name, true)
		if err != nil {
			blog.Errorf("error save userCluster permission: %s", err.Error())
		}
	}
	response.WriteEntity(*cluster)
}

// BCSCluster

// CreateBCSClusterForm xxx
type CreateBCSClusterForm struct {
	ID               string `json:"id" validate:"required"`
	ProjectID        string `json:"project_id" validate:"required"`
	ClusterType      string `json:"cluster_type"`
	TkeClusterID     string `json:"tke_cluster_id"`
	TkeClusterRegion string `json:"tke_cluster_region"`
}

// CreateBCSCluster creates a "BCS" cluster for current user
func CreateBCSCluster(request *restful.Request, response *restful.Response) {

	start := time.Now()

	blog.Debug(fmt.Sprintf("CreateBCSCluster begin"))
	form := CreateBCSClusterForm{}
	request.ReadEntity(&form)

	err := validate.Struct(&form)
	if err != nil {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
		blog.Debug(fmt.Sprintf("CreateBCSCluster form validate failed, %s", err))
		response.WriteEntity(FormatValidationError(err))
		return
	}

	// check the permission
	user := filters.GetUser(request)

	// check if cluster exists already
	externalClusterInfo := sqlstore.QueryBCSClusterInfo(&m.BCSClusterInfo{
		SourceProjectId: form.ProjectID,
		SourceClusterId: form.ID,
	})
	if externalClusterInfo != nil {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
		metric.RequestErrorLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
		message := fmt.Sprintf("errcode: %d, create failed, cluster with this id already exists", common.BcsErrApiBadRequest)
		WriteClientError(response, "CLUSTER_ALREADY_EXISTS", message)
		return
	}

	// Use the "{external_id}-{random-identifier}" as the real cluster id to ensure both uniqueness and readability
	// "BCS-K8S-15007" -> "bcs-bcs-k8s-15007-FvBewMk3"
	clusterId := fmt.Sprintf("%s%s-%s", m.ClusterIdPrefixBCS, strings.ToLower(form.ID), uniuri.NewLen(8))
	cluster := &m.Cluster{
		ID:        clusterId,
		Provider:  m.ClusterProviderBCS,
		CreatorId: user.ID,
	}

	var clusterType uint
	if form.ClusterType == "" || form.ClusterType == "k8s" {
		clusterType = utils.BcsK8sCluster
	} else if form.ClusterType == "mesos" {
		clusterType = utils.BcsMesosCluster
	} else if form.ClusterType == "tke" {
		clusterType = utils.BcsTkeCluster
	} else {
		metric.RequestErrorCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
		metric.RequestErrorLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
		message := fmt.Sprintf("errcode: %d, create failed, cluster type invalid", common.BcsErrApiBadRequest)
		WriteClientError(response, "CLUSTER_TYPE_INVALID", message)
		return
	}
	externalClusterInfo = &m.BCSClusterInfo{
		ClusterId:        clusterId,
		SourceProjectId:  form.ProjectID,
		SourceClusterId:  form.ID,
		ClusterType:      clusterType,
		TkeClusterId:     form.TkeClusterID,
		TkeClusterRegion: form.TkeClusterRegion,
	}
	createClusterWithExternalInfo(cluster, externalClusterInfo, response)
	metric.RequestCount.WithLabelValues("k8s_rest", request.Request.Method).Inc()
	metric.RequestLatency.WithLabelValues("k8s_rest", request.Request.Method).Observe(time.Since(start).Seconds())
}
