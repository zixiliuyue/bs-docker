

package utils

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/bcs-health/types"
)

// AlarmFactory xxx
type AlarmFactory interface {
	SendAlarm(op *AlarmOptions, source string) error
}

const (
	// VoiceMsgLabelKey xxx
	VoiceMsgLabelKey = "bcs-health-voice-msg"
	// VoiceAlarmLabelKey xxx
	VoiceAlarmLabelKey = "bcs-health-voice-alarm"
	// EndpointsNameLabelKey xxx
	EndpointsNameLabelKey = "bcs-endpoints-name"
	// LBEndpoints xxx
	LBEndpoints = "loadbalance"
	// KubeEndpoints xxx
	KubeEndpoints = "kubernetes"
	// DefaultBcsEndpoints xxx
	DefaultBcsEndpoints = "bcsdefault"
	// EndpointsEventKindLabelKey xxx
	// values should be AlarmType's values
	EndpointsEventKindLabelKey = "bcs-endpoints-event-kind"
	// DataIDLabelKey xxx
	DataIDLabelKey = "bcs-health-dataid"
)

// AlarmOptions xxx
type AlarmOptions struct {
	// which is used to do convergence operation
	AlarmID string `json:"alarmID"`
	// used to block the alarm message
	AlarmName          string    `json:"alarmName"`
	ClusterID          string    `json:"clusterID"`
	ConvergenceSeconds *uint16   `json:"convergenceSeconds"`
	AlarmKind          AlarmType `json:"alarmKind"`
	Receivers          string    `json:"-"`
	Module             string    `json:"module"`
	AlarmMsg           string    `json:"alarmMsg"`
	Namespace          string    `json:"namespace"`
	VoiceReadMsg       string    `json:"voiceReadMsg"`
	SmsMsg             string    `json:"smsMsg"`
	ResourceType       string    `json:"resource_type"`
	ResourceName       string    `json:"resource_name"`

	// new added field.
	// the user defined detailed info about this event alarm
	EventMessage  string `json:"event_message"`
	ModuleVersion string `json:"module_version"`
	ModuleIP      string `json:"module_ip"`
	// when did this event occurred.
	AtTime int64 `json:"at_time"`
	// user defined application alarm level, which is value of
	// label with the key "io.tencent.bcs.monitor.level"
	AppAlarmLevel string `json:"alarm_level"`
	// which this event affiliation is, should be one of user, platform and both.
	// both means that this event shoud be cared by both user and platform.
	Affiliation types.AffiliationType `json:"affiliation"`

	// attached labels with this event, which can be used to set special information
	// with this event.
	Labels map[string]string

	// generated when this event is received.
	UUID string `json:"-"`
}
