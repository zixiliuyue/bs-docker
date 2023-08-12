

package bkcmdb

import (
	"context"
	"encoding/json"
	"fmt"

	"bscp.io/pkg/components"
	"bscp.io/pkg/config"
	"bscp.io/pkg/thirdparty/esb/cmdb"
	"bscp.io/pkg/thirdparty/esb/types"
)

// SearchBusiness 组件化的函数
func SearchBusiness(ctx context.Context, params *cmdb.SearchBizParams) (*cmdb.SearchBizResp, error) {
	url := fmt.Sprintf("%s/api/c/compapi/v2/cc/search_business/", config.G.Base.BKPaaSHost)

	// SearchBizParams is esb search cmdb business parameter.
	type esbSearchBizParams struct {
		*types.CommParams
		*cmdb.SearchBizParams
	}

	req := &esbSearchBizParams{
		CommParams: &types.CommParams{
			AppCode:   config.G.Base.AppCode,
			AppSecret: config.G.Base.AppSecret,
			UserName:  "admin",
		},
		SearchBizParams: params,
	}

	resp, err := components.GetClient().R().
		SetContext(ctx).
		SetBody(req).
		Post(url)

	if err != nil {
		return nil, err
	}

	bizList := &cmdb.SearchBizResp{}
	if err := json.Unmarshal(resp.Body(), bizList); err != nil {
		return nil, err
	}
	return bizList, nil

}
