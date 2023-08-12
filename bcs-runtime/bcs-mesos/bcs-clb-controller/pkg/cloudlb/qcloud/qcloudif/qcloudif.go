

// Package qcloudif xxx
package qcloudif

import (
	cloudListenerType "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/network/v1"
)

// ClbAdapter interface to operate clb
type ClbAdapter interface {
	// CreateLoadBalance xxx
	// loadbalance
	CreateLoadBalance(lb *cloudListenerType.CloudLoadBalancer) (lbID string, vips []string, err error)
	DescribeLoadBalance(name string) (*cloudListenerType.CloudLoadBalancer, bool, error)

	// CreateListener xxx
	// listener
	CreateListener(listener *cloudListenerType.CloudListener) (listenerID string, err error)
	DeleteListener(lbID, listenerID string) error
	DescribeListener(lbID, listenerID string, port int) (listener *cloudListenerType.CloudListener, isExisted bool,
		err error)
	ModifyListenerAttribute(listener *cloudListenerType.CloudListener) (err error)

	// CreateRules xxx
	// rule
	CreateRules(lbID, listenerID string, rules cloudListenerType.RuleList) error
	DeleteRule(lbID, listenerID, domain, url string) error
	DescribeRuleByDomainAndURL(loadBalanceID, listenerID, Domain, URL string) (rule *cloudListenerType.Rule,
		isExisted bool, err error)
	ModifyRuleAttribute(loadBalanceID, listenerID string, rule *cloudListenerType.Rule) error

	// Register7LayerBackends xxx
	// 7 layer backend
	Register7LayerBackends(lbID, listenerID, ruleID string, backendsRegister cloudListenerType.BackendList) error
	DeRegister7LayerBackends(lbID, listenerID, ruleID string, backendsDeRegister cloudListenerType.BackendList) error

	// Register4LayerBackends xxx
	// 4 layer backend
	Register4LayerBackends(lbID, listenerID string, backendsRegister cloudListenerType.BackendList) error
	DeRegister4LayerBackends(lbID, listenerID string, backendsDeRegister cloudListenerType.BackendList) error

	// ListListener xxx
	// list all listener
	ListListener(lbID string) ([]*cloudListenerType.CloudListener, error)
}
