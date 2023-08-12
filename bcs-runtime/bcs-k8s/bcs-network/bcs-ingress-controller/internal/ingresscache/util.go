

package ingresscache

import (
	"fmt"
)

const (
	// ingressKeyFmt namespace/name
	ingressKeyFmt = "%s/%s"
	// serviceKeyFmt namespace/name
	serviceKeyFmt = "%s/%s"
	// workloadKeyFmt kind/namespace/name
	// kind is StatefulSet or GameStatefulSet
	workloadKeyFmt = "%s/%s/%s"
)

// buildIngressKey return cache key of ingress
func buildIngressKey(namespace, name string) string {
	return fmt.Sprintf(ingressKeyFmt, namespace, name)
}

// buildServiceKey return cache key of service
func buildServiceKey(namespace, name string) string {
	return fmt.Sprintf(serviceKeyFmt, namespace, name)
}

// buildWorkloadKey return cache key of workload
func buildWorkloadKey(kind, namespace, name string) string {
	return fmt.Sprintf(workloadKeyFmt, kind, namespace, name)
}
