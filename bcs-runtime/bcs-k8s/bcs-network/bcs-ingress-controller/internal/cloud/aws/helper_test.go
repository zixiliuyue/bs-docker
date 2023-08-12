

package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

func TestIsSameRuleCondition(t *testing.T) {
	conditionA := []types.RuleCondition{
		{
			Field:            toStrPtr("host-header"),
			HostHeaderConfig: &types.HostHeaderConditionConfig{Values: []string{"www.qq.com"}},
		},
		{
			Field:             toStrPtr("path-pattern"),
			PathPatternConfig: &types.PathPatternConditionConfig{Values: []string{"/"}},
		},
	}

	conditionB := []types.RuleCondition{
		{
			Field:            toStrPtr("host-header"),
			HostHeaderConfig: &types.HostHeaderConditionConfig{Values: []string{"www.qq.com"}},
		},
		{
			Field:             toStrPtr("path-pattern"),
			PathPatternConfig: &types.PathPatternConditionConfig{Values: []string{"/grafana/pracing"}},
		},
	}

	conditionC := []types.RuleCondition{
		{
			Field:             toStrPtr("path-pattern"),
			PathPatternConfig: &types.PathPatternConditionConfig{Values: []string{"/"}},
		},
		{
			Field:            toStrPtr("host-header"),
			HostHeaderConfig: &types.HostHeaderConditionConfig{Values: []string{"www.qq.com"}},
		},
	}

	if isSameRuleCondition(conditionA, conditionB) {
		t.Error("different condition")
	}
	if !isSameRuleCondition(conditionA, conditionC) {
		t.Error("same condition")
	}
}

func toStrPtr(str string) *string {
	return &str
}
