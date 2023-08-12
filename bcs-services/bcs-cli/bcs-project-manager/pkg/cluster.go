

package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cluster-manager/api/clustermanager"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

const (
	clusterUrl = "/clustermanager/v1/cluster/%s"
)

type (
	// GetClusterRequest 获取集群请求
	GetClusterRequest struct {
		ClusterID string `json:"clusterID,omitempty"`
	}
)

// GetCluster Query the cluster information of the specified Cluster ID
func (p *ProjectManagerClient) GetCluster(in *GetClusterRequest) (*clustermanager.GetClusterResp, error) {
	v, err := query.Values(in)
	if err != nil {
		return nil, fmt.Errorf("slice and Array values default to encoding as multiple URL values failed: %v", err)
	}
	bs, err := p.do(fmt.Sprintf(clusterUrl, in.ClusterID), http.MethodGet, v, nil)
	if err != nil {
		return nil, fmt.Errorf("get cluster failed: %v", err)
	}
	resp := new(clustermanager.GetClusterResp)
	if err := json.Unmarshal(bs, resp); err != nil {
		return nil, errors.Wrapf(err, "get cluster unmarshal failed with response '%s'", string(bs))
	}
	if resp != nil && resp.Code != 0 {
		return nil, fmt.Errorf("get cluster response code not 0 but %d: %s", resp.Code, resp.Message)
	}
	return resp, nil
}
