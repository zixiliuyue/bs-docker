

package alert

import "time"

const (
	startKey = "startsAt"
	endKey   = "endsAt"

	annotationKey        = "annotations"
	annotationMessageKey = "message"

	labelKey        = "labels"
	labelAlertKey   = "alert_type"
	labelClusterKey = "cluster_id"
	labelNSKey      = "namespace"
	labelIPKey      = "ip"
	labelModuleKey  = "module_name"
)

func newServiceAlert(module, message, ip string) map[string]interface{} {
	return map[string]interface{}{
		startKey: time.Now(),
		endKey:   time.Now(),
		annotationKey: map[string]string{
			annotationMessageKey: message,
		},
		labelKey: map[string]string{
			labelAlertKey:   "Error",
			labelClusterKey: "bcs-service",
			labelIPKey:      ip,
			labelNSKey:      "bcs-service",
			labelModuleKey:  module,
		},
	}
}

func newClusterAlert(cluster, module, message, ip string) map[string]interface{} {
	return map[string]interface{}{
		startKey: time.Now(),
		endKey:   time.Now(),
		annotationKey: map[string]string{
			annotationMessageKey: message,
		},
		labelKey: map[string]string{
			labelAlertKey:   "Error",
			labelClusterKey: cluster,
			labelIPKey:      ip,
			labelNSKey:      "bcs-cluster",
			labelModuleKey:  module,
		},
	}
}
