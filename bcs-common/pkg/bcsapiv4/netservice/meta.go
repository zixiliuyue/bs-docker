

package netservice

import "fmt"

// SvcRequest basic Request for netservice
type SvcRequest struct {
	Type    int    `json:"type"`              // type for reqeust data
	Scheme  string `json:"scheme"`            // request scheme
	Cluster string `json:"cluster,omitempty"` // cluster name
}

// Result basic definition for netservice
type Result struct {
	Code    int    `json:"code"`              // 00 success, 0x partial success, failed otherwise
	Message string `json:"message,omitempty"` // error message when failed or partial success
}

// IsSucc check Response is success
func (r *Result) IsSucc() bool {
	return r.Code == 0
}

// SvcResponse basic response for netservice
type SvcResponse struct {
	Result `json:",inline"`
	Data   interface{} `json:"data"`
}

// StorageKey get storage key for data object
type StorageKey interface {
	GetKey() string
}

// NetServiceDataKey key function format data uniq key
func NetServiceDataKey(obj interface{}) (string, error) {
	data, ok := obj.(StorageKey)
	if !ok {
		return "", fmt.Errorf("ExportService type Assert failed")
	}
	return data.GetKey(), nil
}
