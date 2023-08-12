

package types

import (
	"strconv"
	"strings"
)

// Backend hold Backend info for load balance
type Backend struct {
	Host   string // host name for endpoint
	IP     string // overlap ip info
	Port   int    // listen port
	Weight int    // backend weight
}

// String 用于打印
func (b *Backend) String() string {
	return strings.Replace(b.IP, ".", "_", -1) + "_" + strconv.Itoa(b.Port)
}

// BackendList define Backend list implementing sorter interface
type BackendList []Backend

// Len is the number of elements in the collection.
func (bl BackendList) Len() int {
	return len(bl)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (bl BackendList) Less(i, j int) bool {
	return bl[i].Host < bl[j].Host
}

// Swap swaps the elements with indexes i and j.
func (bl BackendList) Swap(i, j int) {
	bl[i], bl[j] = bl[j], bl[i]
}

// ServiceInfo hold http/https info
type ServiceInfo struct {
	Name            string // serviceName
	ServicePort     int    // export port
	Balance         string // loadbalance algorithm, default source
	MaxConn         int    // max connection for service
	SessionAffinity bool
}
