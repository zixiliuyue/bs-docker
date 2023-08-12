

// Package metrics xxx
package metrics

import (
	"github.com/coredns/coredns/plugin"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Success xxx
	Success = "success"
	// Failure xxx
	Failure = "failure"
	// AddOperation xxx
	AddOperation = "add"
	// UpdateOperation xxx
	UpdateOperation = "update"
	// DeleteOperation xxx
	DeleteOperation = "delete"
)

// Metrics the bcsscheduler plugin exports.
var (
	RequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "request_count_total",
		Help:      "Counter of requests to plugin bcsscheduler.",
	}, []string{"status"})

	RequestOutProxyCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "request_out_proxy_count_total",
		Help:      "Counter of requests to external dns.",
	}, []string{"status"})

	RequestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "request_latency_seconds",
		Buckets:   plugin.TimeBuckets,
		Help:      "Histogram of the time (in seconds) each request took.",
	}, []string{"status"})

	DnsTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "dns_total",
		Help:      "Total counter of dns in plugin bcsscheduler.",
	})

	ZkNotifyTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "zookeeper_notify_total",
		Help:      "counter of zookeeper notify.",
	}, []string{"operator"})

	StorageOperatorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "storage_operator_total",
		Help:      "counter of storage operator.",
	}, []string{"operator"})

	StorageOperatorLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: plugin.Namespace,
		Subsystem: "bcsscheduler",
		Name:      "storage_operator_latency_seconds",
		Buckets:   plugin.TimeBuckets,
		Help:      "Histogram of the time (in seconds) each storage operator.",
	}, []string{"status"})
)
