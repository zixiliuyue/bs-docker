

package storage

// IPPoolDetailResponse response from storage
type IPPoolDetailResponse struct {
	ID         string    `json:"_id"`
	ClusterID  string    `json:"clusterId"`
	CreateTime string    `json:"createTime"`
	Datas      []*IPPool `json:"data"`
}

// IPPool information for cluster underlay ip resource
type IPPool struct {
	ClusterID string   `json:"cluster"`
	Net       string   `json:"net"`
	Mask      int      `json:"mask"`
	Gateway   string   `json:"gateway"`
	Created   string   `json:"created"`
	Hosts     []string `json:"hosts"`
	Reserved  []string `json:"reserved"`
	Available []string `json:"available"`
	Active    []string `json:"active"`
}
