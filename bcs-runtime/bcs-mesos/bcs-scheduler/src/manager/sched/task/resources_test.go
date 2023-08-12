

package task

import (
	"testing"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"

	"github.com/stretchr/testify/assert"
)

func TestCreateScaleResource(t *testing.T) {
	resource := createScalarResource("cpu", 0.1)
	assert.Equal(t, *resource.Name, "cpu")
	assert.Equal(t, *resource.Scalar.Value, 0.1)
}

func TestBuildResource(t *testing.T) {
	resources := BuildResources(&types.Resource{Cpus: 0.1, Mem: 16, Disk: 10})
	assert.Equal(t, *resources[0].Name, "cpus")
	assert.Equal(t, *resources[1].Name, "mem")
	assert.Equal(t, *resources[2].Name, "disk")
}
