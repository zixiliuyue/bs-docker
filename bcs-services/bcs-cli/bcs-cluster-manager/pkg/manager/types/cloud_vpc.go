

package types

// CreateCloudVPCReq 创建云私有网络request
type CreateCloudVPCReq struct {
	CloudID     string `json:"cloudID"`
	NetworkType string `json:"networkType"`
	Region      string `json:"region"`
	VPCName     string `json:"vpcName"`
	VPCID       string `json:"vpcID"`
	Available   string `json:"available"`
}

// UpdateCloudVPCReq 更新云私有网络request
type UpdateCloudVPCReq struct {
	CloudID     string `json:"cloudID"`
	NetworkType string `json:"networkType"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	VPCName     string `json:"vpcName"`
	VPCID       string `json:"vpcID"`
	Available   string `json:"available"`
	Updater     string `json:"updater"`
}

// DeleteCloudVPCReq 删除云私有网络request
type DeleteCloudVPCReq struct {
	CloudID string `json:"cloudID"`
	VPCID   string `json:"vpcID"`
}

// ListCloudVPCReq 查询云私有网络request
type ListCloudVPCReq struct {
	NetworkType string `json:"networkType"`
}

// ListCloudRegionsReq 查询云区域列表request
type ListCloudRegionsReq struct {
	CloudID string `json:"cloudID"`
}

// GetVPCCidrReq 查询私有网络cidr request
type GetVPCCidrReq struct {
	VPCID string `json:"vpcID"`
}

// ListCloudVPCResp 查询云私有网络列表response
type ListCloudVPCResp struct {
	Data []*CloudVPC `json:"data"`
}

// ListCloudRegionsResp 查询云区域列表response
type ListCloudRegionsResp struct {
	Data []*CloudRegion `json:"data"`
}

// GetVPCCidrResp 查询私有网络cidr response
type GetVPCCidrResp struct {
	Data []*VPCCidr `json:"data"`
}

// CloudVPC 云私有网络信息
type CloudVPC struct {
	CloudID     string `json:"cloudID"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	NetworkType string `json:"networkType"`
	VPCID       string `json:"vpcID"`
	VPCName     string `json:"vpcName"`
	Available   string `json:"available"`
	Extra       string `json:"extra"`
	Creator     string `json:"creator"`
	Updater     string `json:"updater"`
	CreatTime   string `json:"creatTime"`
	UpdateTime  string `json:"updateTime"`
}

// CloudRegion 云区域信息
type CloudRegion struct {
	CloudID    string `json:"cloudID"`
	Region     string `json:"region"`
	RegionName string `json:"regionName"`
}

// VPCCidr 私有网络cidr
type VPCCidr struct {
	VPC      string `json:"vpc"`
	Cidr     string `json:"cidr"`
	IPNumber uint32 `json:"ipNumber"`
	Status   string `json:"status"`
}

// CloudVPCMgr 云私有网络管理接口
type CloudVPCMgr interface {
	// Create 创建云VPC管理信息
	Create(CreateCloudVPCReq) error
	// Update 更新云vpc信息
	Update(UpdateCloudVPCReq) error
	// Delete 删除特定cloud vpc信息
	Delete(DeleteCloudVPCReq) error
	// List 查询Cloud VPC列表
	List(ListCloudVPCReq) (ListCloudVPCResp, error)
	// ListCloudRegions 根据cloudID获取所属cloud的地域信息
	ListCloudRegions(ListCloudRegionsReq) (ListCloudRegionsResp, error)
	// GetVPCCidr 根据vpcID获取所属vpc的cidr信息
	GetVPCCidr(GetVPCCidrReq) (GetVPCCidrResp, error)
}
