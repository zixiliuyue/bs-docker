

package cpusetdevice

// #cgo LDFLAGS: -lnuma
// #include <numa.h>
import "C"
import (
	"fmt"
	"strconv"
)

// ErrNumaNotAvailable error type for NUMA
var ErrNumaNotAvailable = fmt.Errorf("No NUMA support available on this system")

// IsNumaAvailable Is Numa available
func IsNumaAvailable() bool {
	if int(C.numa_available()) >= 0 {
		return true
	}
	return false
}

// NUMANodes Get Numa Nodes
func NUMANodes() ([]string, error) {
	if !IsNumaAvailable() {
		return nil, ErrNumaNotAvailable
	}
	maxnode := int(C.numa_max_node())
	nodes := make([]string, 0, maxnode)
	for i := 0; i <= maxnode; i++ {
		if C.numa_bitmask_isbitset(C.numa_nodes_ptr, C.uint(i)) > 0 {
			nodes = append(nodes, strconv.Itoa(i))
		}
	}
	return nodes, nil
}

// NUMACPUsOfNode Get CPU slice from the specified Node
func NUMACPUsOfNode(node string) ([]string, error) {
	if !IsNumaAvailable() {
		return nil, ErrNumaNotAvailable
	}
	nodei, err := strconv.Atoi(node)
	if err != nil {
		return nil, fmt.Errorf("node %s is invalid", node)
	}

	mask := C.numa_allocate_cpumask()
	defer C.numa_free_cpumask(mask)
	rc := C.numa_node_to_cpus(C.int(nodei), mask)
	maxCpus := NUMAConfiguredCPUs()
	cpus := make([]string, 0, maxCpus)
	if rc >= 0 {
		for i := 0; i < maxCpus; i++ {
			if C.numa_bitmask_isbitset(mask, C.uint(i)) > 0 {
				cpus = append(cpus, strconv.Itoa(i))
			}
		}
	} else {
		return cpus, ErrNumaNotAvailable
	}
	return cpus, nil
}

// NUMAConfiguredCPUs get configure cpu number
func NUMAConfiguredCPUs() int {
	return int(C.numa_num_configured_cpus())
}
