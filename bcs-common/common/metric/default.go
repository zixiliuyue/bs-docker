

package metric

import (
	"os"
	"runtime"
	"strconv"

	"github.com/Tencent/bk-bcs/bcs-common/common/version"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	default_namespace       = "bcs"
	module_cluster_id_label = "module_cluster_id"
	module_name_label       = "module_name"
	module_ip_label         = "module_ip"
	module_pid_label        = "module_pid"
	module_os_label         = "module_os"
)

func newVersionMetric(m MetaData) *prometheus.GaugeVec {
	verMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: default_namespace,
		Name:      "version_infos",
		Help:      "version infos about this module",
	}, []string{module_cluster_id_label, module_name_label, module_ip_label, "version", "tag", "build_time", "git_hash"})

	verMetric.WithLabelValues(m.ClusterID, m.Module, m.IP, version.BcsVersion, version.BcsTag, version.BcsBuildTime,
		version.BcsGitHash).Set(1)

	return verMetric
}

func newModuleMetric(c MetaData) *prometheus.GaugeVec {
	verMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: default_namespace,
		Name:      "module_infos",
		Help:      "module infos about this module",
	}, []string{module_name_label, module_cluster_id_label, module_ip_label})

	verMetric.WithLabelValues(c.Module, c.ClusterID, c.IP).Set(1)

	return verMetric
}

func newRuntimeMetric(m MetaData) *prometheus.GaugeVec {
	verMetric := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: default_namespace,
		Name:      "runtime_infos",
		Help:      "module infos about this module",
	}, []string{module_cluster_id_label, module_name_label, module_ip_label, module_pid_label, module_os_label})

	verMetric.WithLabelValues(m.ClusterID, m.Module, m.IP, strconv.Itoa(os.Getpid()), runtime.GOOS).Set(1)

	return verMetric
}
