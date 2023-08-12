

package task

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/mesosproto/mesos"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
	// "github.com/Tencent/bk-bcs/bcs-common/common/blog"
	// "github.com/golang/protobuf/proto"
)

func createScalarResource(name string, value float64) *mesos.Resource {
	return &mesos.Resource{
		Name:   &name,
		Type:   mesos.Value_SCALAR.Enum(),
		Scalar: &mesos.Value_Scalar{Value: &value},
	}
}

// BuildResources Build mesos resource format
func BuildResources(r *types.Resource) []*mesos.Resource {
	var resources = []*mesos.Resource{}

	if r.Cpus > 0 {
		resources = append(resources, createScalarResource("cpus", r.Cpus))
	}

	if r.Mem > 0 {
		resources = append(resources, createScalarResource("mem", r.Mem))
	}

	if r.Disk > 0 {
		resources = append(resources, createScalarResource("disk", r.Disk))
	}

	return resources
}
