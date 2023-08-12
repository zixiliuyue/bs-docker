

package storage

import (
	mesostype "github.com/Tencent/bk-bcs/bcs-common/common/types"
	core "k8s.io/api/core/v1"
)

// Pod information from
// https://bcs-api-gateway.bk.tencent.com:8443/bcsapi/v4/storage/query/k8s/dynamic/clusters/BCS-K8S-40026/pod
// data structure difinitions are referred to store structures in mongodb. they are not the the same with
// original kubernetes data structure. Module bcs-storage add some additional control information for
// multiple cluster management.

// Pod definition in mongodb
type Pod struct {
	CommonDataHeader
	Data *core.Pod `json:"data"`
}

// Taskgroup bcs-storage taskgroup data of mesos
type Taskgroup struct {
	CommonDataHeader
	Data *mesostype.BcsPodStatus `json:"data"`
}

// PodList is response for storage pod list operation
type PodList struct {
	CommonResponseHeader
	Data []Pod `json:"data"`
}

// TaskgroupList is response for storage taskgroup list operation
type TaskgroupList struct {
	CommonResponseHeader
	Data []Taskgroup `json:"data"`
}
