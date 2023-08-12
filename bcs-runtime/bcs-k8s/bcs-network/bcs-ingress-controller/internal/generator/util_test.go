

package generator

import (
	"testing"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/constant"
)

func TestMatchLbStrWithID(t *testing.T) {
	testCases := []struct {
		cloud   string
		lbID    string
		isValid bool
	}{
		{
			cloud:   constant.CloudTencent,
			lbID:    "lb-123",
			isValid: true,
		},
		{
			cloud:   constant.CloudTencent,
			lbID:    "123",
			isValid: false,
		},
		{
			cloud:   constant.CloudTencent,
			lbID:    "ap-guangzhou:lb-213",
			isValid: true,
		},
		{
			cloud:   constant.CloudTencent,
			lbID:    "ap-guangzhou:123",
			isValid: false,
		},
		{
			cloud:   constant.CloudTencent,
			lbID:    "1:1:lb-123",
			isValid: false,
		},
		{
			cloud:   constant.CloudTencent,
			lbID:    "ap-shenzhen:ap-shenzhen:lb-123",
			isValid: false,
		},
	}
	for i, c := range testCases {
		if MatchLbStrWithID(c.cloud, c.lbID) != c.isValid {
			t.Errorf("idx: %d result is error", i)
		}
	}

}
