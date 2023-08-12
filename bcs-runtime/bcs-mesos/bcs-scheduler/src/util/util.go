

// Package util xxx
package util

import "github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/mesosproto/mesos"

// GetMesosAgentInnerIP xxx
func GetMesosAgentInnerIP(attributes []*mesos.Attribute) (string, bool) {
	ip := ""
	ok := false

	for _, attribute := range attributes {
		if attribute.GetName() == "InnerIP" {
			ip = attribute.Text.GetValue()
			ok = true
			break
		}
	}

	return ip, ok
}

// ParseMesosResources xxx
func ParseMesosResources(resources []*mesos.Resource) (float64, float64, float64, string) {
	var cpus, mem, disk float64
	var port string
	for _, res := range resources {
		if res.GetName() == "cpus" {
			cpus += *res.GetScalar().Value
		}
		if res.GetName() == "mem" {
			mem += *res.GetScalar().Value
		}
		if res.GetName() == "disk" {
			disk += *res.GetScalar().Value
		}
		if res.GetName() == "ports" {
			port = res.GetRanges().String()
		}
	}

	return cpus, mem, disk, port
}
