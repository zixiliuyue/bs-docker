

// Package ipt used to sync network policy for every container
package ipt

// RuleType used as ipSet suffix
type RuleType string

// BehaviorType used to create iptables rules
type BehaviorType string

const (
	ruleIngress RuleType = "RuleIngress"
	ruleEgress  RuleType = "RuleEgress"

	behaviorAccept BehaviorType = "ACCEPT"
	behaviorReject BehaviorType = "REJECT"
)
