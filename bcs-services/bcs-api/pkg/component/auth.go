

package component

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/config"
)

const (
	appCodeKey   = "X-BK-APP-CODE"
	appSecretKey = "X-BK-APP-SECRET" // nolint
)

// PaaSAuth auth configuration for bk-app
type PaaSAuth struct {
	host      string
	appCode   string
	appSecret string
	version   string
}

// NewPaaSAuth construct Auth according configuration file/command line option
func NewPaaSAuth() *PaaSAuth {
	bKIamAuth := config.BKIamAuth
	return &PaaSAuth{
		host:      bKIamAuth.BKIamAuthHost,
		appCode:   bKIamAuth.BKIamAuthAppCode,
		appSecret: bKIamAuth.BKIamAuthAppSecret,
		version:   bKIamAuth.Version,
	}
}

// VerifyAccessTokenForIeod refresh access token
func (a *PaaSAuth) VerifyAccessTokenForIeod(accessToken string) (bool, map[string]interface{}, error) {

	url := fmt.Sprintf("%s/oauth/token", a.host)

	params := map[string]string{
		"access_token": accessToken,
	}

	result, err := HTTPGet(url, params)
	if err != nil {
		return false, nil, err
	}

	return true, result.Data, nil
}

// VerifyAccessTokenForEe verify access token according bk app info
func (a *PaaSAuth) VerifyAccessTokenForEe(accessToken string) (bool, map[string]interface{}, error) {
	var url string
	if a.version == "3" {
		url = fmt.Sprintf("%s/api/v1/auth/access-tokens/verify", a.host)
	} else {
		url = fmt.Sprintf("%s/bkiam/api/v1/auth/access-tokens/verify", a.host)
	}

	data := map[string]interface{}{
		"access_token": accessToken,
	}

	header := map[string]string{
		appCodeKey:   a.appCode,
		appSecretKey: a.appSecret,
	}

	result, err := HTTPPostToBkIamAuth(url, data, header)
	if err != nil {
		return false, nil, err
	}

	return true, result.Data.Identity, nil
}
