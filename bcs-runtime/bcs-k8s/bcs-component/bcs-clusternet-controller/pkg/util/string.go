

package util

import (
	"strings"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-clusternet-controller/pkg/constant"
)

// MatchAnnotationsKeyPrefix 匹配annotations中带上了 AnnotationSubscriptionKeyPrefix 的key
func MatchAnnotationsKeyPrefix(annotations map[string]string) bool {
	for key := range annotations {
		if strings.HasPrefix(key, constant.AnnotationSubscriptionKeyPrefix) {
			return true
		}
	}
	return false
}

// FindAnnotationsMathKeyPrefix 返回匹配的annotations
func FindAnnotationsMathKeyPrefix(annotations map[string]string) map[string]string {
	ret := make(map[string]string)
	for key, val := range annotations {
		if strings.HasPrefix(key, constant.AnnotationSubscriptionKeyPrefix) {
			ret[key] = val
		}
	}
	return ret
}
