

package kubedriver

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/types"
	bcsVersion "github.com/Tencent/bk-bcs/bcs-common/common/version"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"
	"os"
)

// GetServerInfo returns server info of current driver node
func GetServerInfo(opts *options.KubeDriverServerOptions, clusterID string) types.ServerInfo {
	hostname, _ := os.Hostname()

	var port uint
	var scheme string
	if opts.SecureServerConfigured() {
		scheme = options.HTTPS
		port = opts.SecurePort
	} else {
		scheme = options.HTTP
		port = opts.InsecurePort
	}
	if opts.CustomReportPort != 0 {
		port = opts.CustomReportPort
	}

	hostIp := opts.HostIP
	if opts.CustomReportAddress != "" {
		hostIp = opts.CustomReportAddress
	}

	return types.ServerInfo{
		IP:           hostIp,
		Port:         port,
		ExternalIp:   opts.ExternalIp,
		ExternalPort: opts.ExternalPort,
		HostName:     hostname,
		Scheme:       scheme,
		Cluster:      clusterID,
		Version:      bcsVersion.GetVersion(),
		Pid:          os.Getpid(),
	}
}
