

package storage

import (
	"fmt"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-watch/util"
)

// NetServiceHandler handle netservice resources for storage.
type NetServiceHandler struct {
	oper      DataOperator
	dataType  string
	ClusterID string
}

// GetType returns data type.
func (h *NetServiceHandler) GetType() string {
	return h.dataType
}

// dataNode := fmt.Sprintf("/bcsstorage/v1/mesos/dynamic/cluster_resources/clusters/%s/%s", h.ClusterID, h.dataType)

// CheckDirty cleans dirty data in remote bcs-storage.
func (h *NetServiceHandler) CheckDirty() error {
	// do nothing.
	return nil
}

// Add handle add event for netservice resources.
func (h *NetServiceHandler) Add(data interface{}) error {
	// do nothing.
	return nil
}

// Delete handle delete event for netservice resources.
func (h *NetServiceHandler) Delete(data interface{}) error {
	// do nothing.
	return nil
}

// Update xxx
// Delete handle delete event for netservice resources.
func (h *NetServiceHandler) Update(data interface{}) error {
	started := time.Now()
	dataNode := fmt.Sprintf("/bcsstorage/v1/mesos/dynamic/cluster_resources/clusters/%s/%s/%s-%s",
		h.ClusterID, h.dataType, h.dataType, h.ClusterID)

	if err := h.oper.CreateDCNode(dataNode, data, "PUT"); err != nil {
		blog.V(3).Infof("%s update node %s, err %+v", h.dataType, dataNode, err)
		util.ReportStorageMetrics(h.ClusterID, h.dataType, actionPut, handlerClusterTypeName, util.StatusFailure, started)
		return err
	}

	util.ReportStorageMetrics(h.ClusterID, h.dataType, actionPut, handlerClusterTypeName, util.StatusSuccess, started)
	return nil
}
