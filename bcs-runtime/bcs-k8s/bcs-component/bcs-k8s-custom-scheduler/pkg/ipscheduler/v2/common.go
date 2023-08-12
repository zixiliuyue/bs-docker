

package v2

const (
	// BcsSystem system namespace for scheduler
	BcsSystem = "bcs-system"
	// IPLabelKeyForHost key in ip annotations for host, used to search cloud ip
	IPLabelKeyForHost = "host.cloud.bkbcs.tencent.com"
	// IPLabelKeyForEni key in ip annotations for eni, used to search cloud ip
	IPLabelKeyForEni = "eni.cloud.bkbcs.tencent.com"
	// IPLabelKeyForIsFixed key in ip annotations for if it is fixed
	IPLabelKeyForIsFixed = "fixed.cloud.bkbcs.tencent.com"
)
