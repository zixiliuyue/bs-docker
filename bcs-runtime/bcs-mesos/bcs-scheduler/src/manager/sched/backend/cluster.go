

package backend

import (
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
	types "github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// GetClusterResources get cluster resources
func (b *backend) GetClusterResources() (*commtypes.BcsClusterResource, error) {
	return b.sched.GetClusterResource()
}

// GetClusterEndpoints get cluster endpoints
func (b *backend) GetClusterEndpoints() *commtypes.ClusterEndpoints {
	endpoints := new(commtypes.ClusterEndpoints)

	for _, sched := range b.sched.Schedulers {
		endpoints.MesosSchedulers = append(endpoints.MesosSchedulers, *sched)
	}
	for _, master := range b.sched.Memsoses {
		endpoints.MesosMasters = append(endpoints.MesosMasters, *master)
	}

	return endpoints
}

// GetCurrentOffers get current offers of cluster
func (b *backend) GetCurrentOffers() []*types.OfferWithDelta {
	return b.sched.GetCurrentOffers()
}
