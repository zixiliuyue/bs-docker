

package scalingconfig

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/config"
)

// Options are the option of autoscaler
type Options struct {
	config.AutoscalingOptions
	BufferedCPURatio         float64
	BufferedMemRatio         float64
	BufferedResourceRatio    float64
	WebhookMode              string
	WebhookModeConfig        string
	WebhookModeToken         string
	MaxBulkScaleUpCount      int
	MaxNodeStartupTime       time.Duration
	MaxNodeStartScheduleTime time.Duration
	ScanInterval             time.Duration
	EvictLatest              bool
}
