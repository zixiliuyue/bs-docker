

package reflector

import (
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
	schedtypes "github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// Reflector watches a specified resource and causes all changes to be reflected in the given store.
type Reflector interface {
	// ListAutoscalers TODO
	// list all namespace autoscaler
	ListAutoscalers() ([]*commtypes.BcsAutoscaler, error)

	// StoreAutoscaler TODO
	// store autoscaler in zk
	StoreAutoscaler(autoscaler *commtypes.BcsAutoscaler) error

	// UpdateAutoscaler TODO
	// update autoscaler in zk
	UpdateAutoscaler(autoscaler *commtypes.BcsAutoscaler) error

	// FetchDeploymentInfo TODO
	// fetch deployment info, if deployment status is not Running, then can't autoscale this deployment
	FetchDeploymentInfo(namespace, name string) (*schedtypes.Deployment, error)

	// FetchApplicationInfo TODO
	// fetch application info, if application status is not Running or Abnormal, then can't autoscale this application
	FetchApplicationInfo(namespace, name string) (*schedtypes.Application, error)

	// ListTaskgroupRefDeployment TODO
	// list selectorRef deployment taskgroup
	ListTaskgroupRefDeployment(namespace, name string) ([]*schedtypes.TaskGroup, error)

	// ListTaskgroupRefApplication TODO
	// list selectorRef application taskgroup
	ListTaskgroupRefApplication(namespace, name string) ([]*schedtypes.TaskGroup, error)
}
