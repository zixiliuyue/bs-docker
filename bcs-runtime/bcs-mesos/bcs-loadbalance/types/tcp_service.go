

package types

import "sort"

// NewFourLayerServiceInfo to new a FourLayerServiceInfo
func NewFourLayerServiceInfo(s ServiceInfo, bl BackendList) FourLayerServiceInfo {
	return FourLayerServiceInfo{
		ServiceInfo: s,
		Backends:    bl,
	}
}

// FourLayerServiceInfo to hold tcp and udp service info
type FourLayerServiceInfo struct {
	ServiceInfo
	Backends BackendList // tcp Backend
}

// AddBackend add backend to list
func (tsi *FourLayerServiceInfo) AddBackend(b Backend) {
	tsi.Backends = append(tsi.Backends, b)
}

// SortBackends sort backend list
func (tsi *FourLayerServiceInfo) SortBackends() {
	sort.Sort(tsi.Backends)
}

// FourLayerServiceInfoList define serviceInfo list implementing sorter interface
type FourLayerServiceInfoList []FourLayerServiceInfo

// Len is the number of elements in the collection.
func (til FourLayerServiceInfoList) Len() int {
	return len(til)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (til FourLayerServiceInfoList) Less(i, j int) bool {
	return til[i].Name < til[j].Name
}

// Swap swaps the elements with indexes i and j.
func (til FourLayerServiceInfoList) Swap(i, j int) {
	til[i], til[j] = til[j], til[i]
}
