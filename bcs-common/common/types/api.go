

package types

// RequestApi old version api request for bcs-client & bcs-api
type RequestApi struct {
	AppId    string      `json:"appid,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Request  interface{} `json:"request,omitempty"`
}

// BcsRequest request for bcs-api & bcs-client
type BcsRequest struct {
	AppId             string            `json:"appid,omitempty"`
	Operator          string            `json:"operator,omitempty"`
	DataType          BcsDataType       `json:"dataType,omitempty"`
	Pod               BcsTaskgroup      `json:"pod,omitempty"`
	ReplicaController ReplicaController `json:"replicaController,omitempty"`
	ConfigMap         BcsConfigMap      `json:"configMap,omitempty"`
	Service           BcsService        `json:"service,omitempty"`
	Secret            BcsSecret         `json:"secret,omitempty"`
	LoadBalance       BcsLoadBalance    `json:"loadBalance,omitempty"`
}

// bcs standard header keys for http api
const (
	BcsApiHeader_ClusterID = "BCS-ClusterID"
	BcsApiHeader_Operator  = "BCS-Operator"
	BcsApiHeader_UUID      = "BCS-UUID"
)

// bcs http method
const (
	HttpMethod_POST   = "POST"
	HttpMethod_PUT    = "PUT"
	HttpMethod_GET    = "GET"
	HttpMethod_DELETE = "DELETE"
	HttpMethod_PATCH  = "PATCH"
)

// APIResponse xxx
type APIResponse struct {
	Result  bool        `json:"result"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
