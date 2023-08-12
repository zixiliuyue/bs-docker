

package scheduler

import (
	"time"

	types "github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// LaunchTaskgroupType TODO
	/*schedule taskgroup type*/
	LaunchTaskgroupType = "launch"
	// RescheduleTaskgroupType TODO
	RescheduleTaskgroupType = "reschedule"
	// ScaleTaskgroupType TODO
	ScaleTaskgroupType = "scale"
	// UpdateTaskgroupType TODO
	UpdateTaskgroupType = "update"

	// LaunchApplicationType TODO
	/*operate application type*/
	LaunchApplicationType = "launch"
	// DeleteApplicationType TODO
	DeleteApplicationType = "delete"
	// ScaleApplicationType TODO
	ScaleApplicationType = "scale"
	// UpdateApplicationType TODO
	UpdateApplicationType = "update"
	// RollingupdateApplicationType TODO
	RollingupdateApplicationType = "rollingupdate"
)

// Metrics the scheduler info
var (
	ScheduleTaskgroupTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: types.MetricsNamespaceScheduler,
		Subsystem: types.MetricsSubsystemScheduler,
		Name:      "schedule_taskgroup_total",
		Help:      "Total counter of schedule taskgroup",
	}, []string{"namespace", "application", "taskgroup", "type"})

	ScheduleTaskgroupLatencyMs = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: types.MetricsNamespaceScheduler,
		Subsystem: types.MetricsSubsystemScheduler,
		Name:      "schedule_taskgroup_latency_ms",
		Help:      "Schedule taskgroup latency milliseconds",
	}, []string{"namespace", "application", "taskgroup", "type"})

	OperateAppTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: types.MetricsNamespaceScheduler,
		Subsystem: types.MetricsSubsystemScheduler,
		Name:      "operate_application_total",
		Help:      "Total counter of operate application",
	}, []string{"namespace", "application", "type"})

	OperateAppLatencySecond = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: types.MetricsNamespaceScheduler,
		Subsystem: types.MetricsSubsystemScheduler,
		Name:      "operate_application_latency_second",
		Help:      "Operate application latency seconds",
	}, []string{"namespace", "application", "type"})

	TaskgroupReportTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: types.MetricsNamespaceScheduler,
		Subsystem: types.MetricsSubsystemScheduler,
		Name:      "taskgroup_report_total",
		Help:      "Total counter of report taskgroup status",
	}, []string{"namespace", "application", "taskgroup", "status"})

	TaskgroupOomTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: types.MetricsNamespaceScheduler,
		Subsystem: types.MetricsSubsystemScheduler,
		Name:      "application_oom_total",
		Help:      "Total counter of application oom killed",
	}, []string{"namespace", "application"})
)

func init() {
	prometheus.MustRegister(ScheduleTaskgroupTotal, ScheduleTaskgroupLatencyMs, OperateAppTotal, OperateAppLatencySecond,
		TaskgroupReportTotal, TaskgroupOomTotal)
}

func reportScheduleTaskgroupMetrics(ns, name, taskgroup, scheduleType string, started time.Time) {
	ScheduleTaskgroupTotal.WithLabelValues(ns, name, taskgroup, scheduleType).Inc()
	d := time.Duration(time.Since(started).Nanoseconds())
	sec := d / time.Millisecond
	nsec := d % time.Millisecond
	ms := float64(sec) + float64(nsec)/1e6
	ScheduleTaskgroupLatencyMs.WithLabelValues(ns, name, taskgroup, scheduleType).Observe(ms)
}

func reportOperateAppMetrics(ns, name, operateType string, started time.Time) {
	OperateAppTotal.WithLabelValues(ns, name, operateType).Inc()
	OperateAppLatencySecond.WithLabelValues(ns, name, operateType).Observe(time.Since(started).Seconds())
}

func reportTaskgroupReportMetrics(ns, name, taskgroup, status string) {
	TaskgroupReportTotal.WithLabelValues(ns, name, taskgroup, status).Inc()
}

func reportTaskgroupOomMetrics(ns, name, taskgroupId string) {
	TaskgroupOomTotal.WithLabelValues(ns, name, taskgroupId).Inc()
}
