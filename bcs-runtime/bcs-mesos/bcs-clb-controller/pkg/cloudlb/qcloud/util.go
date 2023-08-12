

package qcloud

import (
	loadbalance "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/network/v1"
	"strings"
)

// GetClusterIDPostfix get postfix from cluster id
func GetClusterIDPostfix(clusterid string) string {
	if len(clusterid) == 0 {
		return ""
	}
	strs := strings.Split(clusterid, "-")
	if len(strs) == 1 {
		return clusterid
	}
	return strs[len(strs)-1]
}

// GetBackendsSegment get segment from a big string slice
func GetBackendsSegment(strs []*loadbalance.Backend, cur, segmentLen int) []*loadbalance.Backend {
	if len(strs) == 0 || cur < 0 || segmentLen < 0 || cur > len(strs) {
		return nil
	}
	ret := make([]*loadbalance.Backend, 0)
	realLen := segmentLen
	if cur+realLen > len(strs) {
		realLen = len(strs) - cur
	}
	ret = append(ret, strs[cur:cur+realLen]...)
	return ret
}
