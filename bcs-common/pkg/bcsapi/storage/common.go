

package storage

import "time"

// CommonResponseHeader is common response for storage resource api
type CommonResponseHeader struct {
	Result   bool   `json:"result"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	PageSize int64  `json:"pageSize"`
	Offset   int64  `json:"offset"`
	Total    int64  `json:"total"`
}

// CommonDataHeader is common header for storage dynamic data api
type CommonDataHeader struct {
	Namespace          string    `json:"namespace"`
	ResourceName       string    `json:"resourceName"`
	UpdateTime         time.Time `json:"updateTime"`
	ID                 string    `json:"_id"`
	ResourceType       string    `json:"resourceType"`
	ClusterID          string    `json:"clusterId"`
	IsBcsObjectDeleted bool      `json:"_isBcsObjectDeleted"`
	CreateTime         time.Time `json:"createTime"`
}
