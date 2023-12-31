

package bcs

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bcstypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/watch"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bmsf-mesh/bmsf-mesos-adapter/discovery"
)

// NewCluster create cluster for bk-bcs scheduler
func NewCluster(clusterID string, hosts []string) (discovery.Cluster, error) {
	m := &bkbcsCluster{
		clusterName: clusterID,
	}
	// create service store
	svcCtl, err := newServiceCache(hosts)
	if err != nil {
		return nil, err
	}
	m.svcCtl = svcCtl
	// create taskgroup
	taskCtl, err := newTaskGroupCache(hosts)
	if err != nil {
		return nil, err
	}
	m.taskgroupsCtl = taskCtl
	return m, nil
}

// containerInfo hold info from BcsContainer
type containerInfo struct {
	IPAddress   string `json:"IPAddress"`
	NodeAddress string `json:"NodeAddress"`
}

// event inner event object
type svcEvent struct {
	EventType watch.EventType
	Old       *bcstypes.BcsService
	Cur       *bcstypes.BcsService
}

// type appEvent struct {
// 	EventType watch.EventType
// 	Old       *DiscoveryApp
// 	Cur       *DiscoveryApp
// }

type taskGroupEvent struct {
	EventType watch.EventType
	Old       *TaskGroup
	Cur       *TaskGroup
}

// bkbcsCluster bcs-scheduler cluster management
// discovery informations are based on BcsService.
type bkbcsCluster struct {
	clusterName   string               // cluster name
	svcCtl        *svcController       // service controller
	taskgroupsCtl *taskGroupController // taskgroup info controller
}

// GetName implementation for cluster
func (bcs *bkbcsCluster) GetName() string {
	return "bk-bcs"
}

// Run cluster event loop
func (bcs *bkbcsCluster) Run() {
	blog.Infof("bcs-scheduler cluster data plugin is ready to run...")
	// running backgroup recvLoop
	bcs.taskgroupsCtl.run()
	bcs.svcCtl.run()
}

// Stop close cluster event loop
func (bcs *bkbcsCluster) Stop() {
	// close all
	blog.Infof("bk-bcs cluster data plugin is ready to stop...")
	bcs.svcCtl.stop()
	bcs.taskgroupsCtl.stop()
}

// AppSvcs get controller of AppSvc
func (bcs *bkbcsCluster) AppSvcs() discovery.AppSvcController {
	return bcs.svcCtl
}

// AppNodes get controller of AppNode
func (bcs *bkbcsCluster) AppNodes() discovery.AppNodeController {
	return bcs.taskgroupsCtl
}
