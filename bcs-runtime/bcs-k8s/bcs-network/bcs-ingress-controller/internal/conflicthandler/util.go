

package conflicthandler

import (
	"strings"

	mapset "github.com/deckarep/golang-set"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/common"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-network/bcs-ingress-controller/internal/constant"
)

func getRegionLBID(lbID, defaultRegion string) (string, error) {
	region, ID, err := common.GetLbRegionAndName(lbID)
	if err != nil {
		blog.Error(err.Error())
		return "", err
	}
	if region == "" {
		region = defaultRegion
	}
	return common.BuildRegionName(region, ID), nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isProtocolConflict return true if conflict
func isProtocolConflict(isTCPUDPReuse bool, protocols, otherProtocols []string) bool {
	if isTCPUDPReuse == false {
		return true
	}

	// only TCP & UDP reuse is valid
	availableProtocolSet := mapset.NewThreadUnsafeSet()
	availableProtocolSet.Add(constant.ProtocolUDP)
	availableProtocolSet.Add(constant.ProtocolTCP)

	for _, protocol := range protocols {
		protocolUpper := strings.ToUpper(protocol)
		if !availableProtocolSet.Contains(protocolUpper) {
			return true
		}
		availableProtocolSet.Remove(protocolUpper)
	}

	for _, protocol := range otherProtocols {
		protocolUpper := strings.ToUpper(protocol)
		if !availableProtocolSet.Contains(protocolUpper) {
			return true
		}
		availableProtocolSet.Remove(protocolUpper)
	}

	return false
}
