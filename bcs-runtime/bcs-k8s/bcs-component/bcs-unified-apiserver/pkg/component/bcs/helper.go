

package bcs

import (
	"encoding/base64"
	"encoding/json"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-unified-apiserver/pkg/component"
)

// PaginationToContinue BCS 分页转换为 K8S Continue 参数
func PaginationToContinue(pag *component.Pagination, resCount int64) (string, error) {
	// limit 已经大于 total, 没有更多数据
	if pag.PageSize >= pag.Total {
		return "", nil
	}

	// 已经返回全部数据
	if pag.Offset+resCount >= pag.Total {
		return "", nil
	}

	// 当前offset计算规则: 上一个合法的offset + 当前返回值
	pag.Offset = resCount + pag.Offset

	body, err := json.Marshal(pag)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(body), nil
}

// ContinueToOffset K8S Continue 转换为分页参数
func ContinueToOffset(continueStr string, limit int64) int64 {
	if continueStr == "" {
		return 0
	}

	body, err := base64.StdEncoding.DecodeString(continueStr)
	if err != nil {
		return 0
	}

	pag := &component.Pagination{}
	if err := json.Unmarshal(body, pag); err != nil {
		return 0
	}

	return pag.Offset
}
