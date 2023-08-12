

// Package bcs xxx
package bcs

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	v1 "k8s.io/api/core/v1"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-unified-apiserver/pkg/component"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-unified-apiserver/pkg/config"
)

// ResourceData struct store the bcs-storage's resource message.
type ResourceData struct {
	Data         *v1.Pod `json:"data,omitempty"`
	UpdateTime   string  `json:"updateTime"`
	Id           string  `json:"_id"`
	ClusterId    string  `json:"clusterId"`
	Namespace    string  `json:"namespace"`
	ResourceName string  `json:"resourceName"`
	ResourceType string  `json:"resourceType"`
	CreateTime   string  `json:"createTime"`
}

// ListPodResources list Pod resources
func ListPodResources(ctx context.Context, clusterIdList []string, namespace string, limit,
	offset int64) ([]*ResourceData, *component.Pagination, error) {
	url := fmt.Sprintf("%s/bcsapi/v4/storage/dynamic/customresources/Pod", config.G.BCS.Host)

	clusterId := strings.Join(clusterIdList, ",")

	resp, err := component.GetClient().R().
		SetContext(ctx).
		SetAuthToken(config.G.BCS.Token).
		SetQueryParam("clusterId", clusterId).
		SetQueryParam("namespace", namespace).
		SetQueryParam("limit", strconv.FormatInt(limit, 10)).
		SetQueryParam("offset", strconv.FormatInt(offset, 10)).
		Get(url)

	if err != nil {
		return nil, nil, err
	}

	var result []*ResourceData
	pag, err := component.UnmarshalBCSStorResult(resp, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, pag, nil

}
