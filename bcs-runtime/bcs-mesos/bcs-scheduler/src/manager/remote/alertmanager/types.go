

package alertmanager

// CreateBusinessAlertInfoReq for alertInfo req
type CreateBusinessAlertInfoReq struct {
	Starttime          int64               `json:"starttime"`
	Endtime            int64               `json:"endtime"`
	Generatorurl       string              `json:"generatorurl"`
	AlarmType          string              `json:"alarmType"`
	ClusterID          string              `json:"clusterID"`
	AlertAnnotation    *AlertAnnotation    `json:"alertAnnotation"`
	ModuleAlertLabel   *ModuleAlertLabel   `json:"moduleAlertLabel,omitempty"`
	ResourceAlertLabel *ResourceAlertLabel `json:"resourceAlertLabel,omitempty"`
}

// CommonAlertLabel for common label
type CommonAlertLabel struct {
	AlarmType string `json:"alarmType"` // resource/module
	ClusterID string `json:"clusterID"`
}

// AlertAnnotation annotations
type AlertAnnotation struct {
	Message string `json:"message"`
	Comment string `json:"comment,omitempty"`
}

// ModuleAlertLabel module labels
type ModuleAlertLabel struct {
	ModuleName string `json:"moduleName"`
	ModuleIP   string `json:"moduleIP"`
	AlarmName  string `json:"alarmName"`
	AlarmLevel string `json:"alarmLevel"`
}

// ResourceAlertLabel resource labels
type ResourceAlertLabel struct {
	AlarmName         string `json:"alarmName"`
	NameSpace         string `json:"nameSpace"`
	AlarmResourceType string `json:"alarmResourceType"`
	AlarmResourceName string `json:"alarmResourceName"`
	AlarmID           string `json:"alarmID"`
	AlarmLevel        string `json:"alarmLevel"`
}

// CreateBusinessAlertInfoResp http response
type CreateBusinessAlertInfoResp struct {
	ErrCode uint64 `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
