

package types

// CreateCloudAccountReq 创建云凭证request
type CreateCloudAccountReq struct {
	CloudID     string  `json:"cloudID"`
	ProjectID   string  `json:"projectID"`
	AccountName string  `json:"accountName"`
	Desc        string  `json:"desc"`
	Account     Account `json:"account"`
}

// UpdateCloudAccountReq 更新云凭证request
type UpdateCloudAccountReq struct {
	CloudID     string  `json:"cloudID"`
	AccountID   string  `json:"accountID"`
	ProjectID   string  `json:"projectID"`
	AccountName string  `json:"accountName"`
	Desc        string  `json:"desc"`
	Account     Account `json:"account"`
}

// DeleteCloudAccountReq 删除云凭证request
type DeleteCloudAccountReq struct {
	CloudID   string `json:"cloudID"`
	AccountID string `json:"accountID"`
}

// ListCloudAccountReq 查询云凭证列表request
type ListCloudAccountReq struct {
}

// ListCloudAccountToPermReq 查询云凭证列表,主要用于权限资源查询request
type ListCloudAccountToPermReq struct {
}

// ListCloudAccountResp 查询云凭证列表response
type ListCloudAccountResp struct {
	Data []*CloudAccountInfo `json:"data"`
}

// ListCloudAccountToPermResp 查询云凭证列表,主要用于权限资源查询response
type ListCloudAccountToPermResp struct {
	Data []*CloudAccount `json:"data"`
}

// CloudAccountInfo 云凭证信息
type CloudAccountInfo struct {
	AccountID   string   `json:"accountID"`
	AccountName string   `json:"accountName"`
	ProjectID   string   `json:"projectID"`
	Desc        string   `json:"desc"`
	Account     Account  `json:"account"`
	Clusters    []string `json:"clusters"`
}

// CloudAccount 云凭证
type CloudAccount struct {
	CloudID     string  `json:"cloudID"`
	ProjectID   string  `json:"projectID"`
	AccountID   string  `json:"accountID"`
	AccountName string  `json:"accountName"`
	Account     Account `json:"account"`
	Enable      bool    `json:"enable"`
	Creator     string  `json:"creator"`
	CreatTime   string  `json:"creatTime"`
}

// Account 账号信息
type Account struct {
	SecretID          string `json:"secretID"`
	SecretKey         string `json:"secretKey"`
	SubscriptionID    string `json:"subscriptionID"`
	TenantID          string `json:"tenantID"`
	ResourceGroupName string `json:"resourceGroupName"`
	ClientID          string `json:"clientID"`
	ClientSecret      string `json:"clientSecret"`
}

// CloudAccountMgr 云凭证管理接口
type CloudAccountMgr interface {
	// Create 创建云凭证
	Create(CreateCloudAccountReq) error
	// Update 更新云凭证
	Update(UpdateCloudAccountReq) error
	// Delete 删除云凭证
	Delete(DeleteCloudAccountReq) error
	// List 查询云凭证列表
	List(ListCloudAccountReq) (ListCloudAccountResp, error)
	// ListToPerm 查询云凭证列表,主要用于权限资源查询
	ListToPerm(ListCloudAccountToPermReq) (ListCloudAccountToPermResp, error)
}
