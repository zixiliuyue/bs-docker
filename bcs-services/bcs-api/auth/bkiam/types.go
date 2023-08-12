

package bkiam

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/auth"
)

// QueryParam xxx
type QueryParam struct {
	PrincipalType string      `json:"principal_type"`
	PrincipalID   string      `json:"principal_id"`
	ScopeType     string      `json:"scope_type"`
	ScopeID       string      `json:"scope_id"`
	ActionID      auth.Action `json:"action_id"`
	ResourceType  string      `json:"resource_type"`
	ResourceID    string      `json:"resource_id"`
}

// ParseResource xxx
func (qp *QueryParam) ParseResource(resource auth.Resource) {
	if resource.Namespace == "" {
		qp.ResourceType = "cluster"
		qp.ResourceID = fmt.Sprintf("cluster:%s", resource.ClusterID)
		return
	}

	qp.ResourceType = "namespace"
	qp.ResourceID = fmt.Sprintf("cluster:%s/namespace:%s", resource.ClusterID, resource.Namespace)
}

// QueryResp xxx
type QueryResp struct {
	RequestID  string    `json:"request_id"`
	Result     bool      `json:"result"`
	ErrCode    int       `json:"bk_error_code"`
	ErrMessage string    `json:"bk_error_msg"`
	Data       QueryData `json:"data"`
}

// QueryData xxx
type QueryData struct {
	IsPass bool `json:"is_pass"`
}

// ApiGwData xxx
type ApiGwData struct {
	ISS     string           `json:"iss"`
	App     ApiGwDataApp     `json:"app"`
	Project ApiGwDataProject `json:"project"`
	User    ApiGwDataUser    `json:"user"`
	Exp     float64          `json:"exp"`
	NBF     float64          `json:"nbf"`
}

// ApiGwDataApp xxx
type ApiGwDataApp struct {
	Version  float64 `json:"version"`
	Verified bool    `json:"verified"`
	AppCode  string  `json:"app_code"`
}

// ApiGwDataProject xxx
type ApiGwDataProject struct {
	ProjectID   string `json:"project_id"`
	ProjectCode string `json:"project_code"`
	Verified    bool   `json:"verified"`
}

// ApiGwDataUser xxx
type ApiGwDataUser struct {
	Username string  `json:"username"`
	Version  float64 `json:"version"`
	Verified bool    `json:"verified"`
}
