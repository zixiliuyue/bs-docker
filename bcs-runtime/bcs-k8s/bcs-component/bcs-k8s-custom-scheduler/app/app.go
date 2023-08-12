

// Package app xxx
package app

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/app/customscheduler"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/options"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run the customScheduler
func Run(conf *config.CustomSchedulerConfig) {
	customSched := customscheduler.NewCustomScheduler(conf)
	// start customSched, and http service
	err := customSched.Start()
	if err != nil {
		blog.Errorf("start processor error %s, and exit", err.Error())
		os.Exit(1)
	}

	return
}

// RunPrometheusMetricsServer starting prometheus metrics handler
func RunPrometheusMetricsServer(conf *config.CustomSchedulerConfig) {
	http.Handle("/metrics", promhttp.Handler())
	addr := conf.Address + ":" + strconv.Itoa(int(conf.MetricPort))
	go http.ListenAndServe(addr, nil)
}

// ParseConfig xxx
func ParseConfig(op *options.ServerOption) *config.CustomSchedulerConfig {
	customSchedulerConfig := config.NewCustomSchedulerConfig()

	customSchedulerConfig.Address = op.Address
	customSchedulerConfig.Port = op.Port
	customSchedulerConfig.MetricPort = op.MetricPort
	customSchedulerConfig.InsecureAddress = op.InsecureAddress
	customSchedulerConfig.InsecurePort = op.InsecurePort
	customSchedulerConfig.ZkHosts = op.BCSZk
	customSchedulerConfig.VerifyClientTLS = op.VerifyClientTLS
	customSchedulerConfig.CustomSchedulerType = op.CustomSchedulerType
	customSchedulerConfig.KubeConfig = op.Kubeconfig
	customSchedulerConfig.KubeMaster = op.KubeMaster
	customSchedulerConfig.UpdatePeriod = op.UpdatePeriod
	customSchedulerConfig.Cluster = op.Cluster
	customSchedulerConfig.CniAnnotationKey = op.CniAnnotationKey
	customSchedulerConfig.CniAnnotationValue = op.CniAnnotationValue
	customSchedulerConfig.FixedIpAnnotationKey = op.FixedIpAnnotationKey
	customSchedulerConfig.FixedIpAnnotationValue = op.FixedIpAnnotationValue
	customSchedulerConfig.CloudNetserviceEndpoints = strings.Split(op.CloudNetserviceEndpoints, ",")

	// server cert directory
	if op.CertConfig.ServerCertFile != "" && op.CertConfig.ServerKeyFile != "" {
		customSchedulerConfig.ServCert.CertFile = op.CertConfig.ServerCertFile
		customSchedulerConfig.ServCert.KeyFile = op.CertConfig.ServerKeyFile
		customSchedulerConfig.ServCert.CAFile = op.CertConfig.CAFile
		customSchedulerConfig.ServCert.IsSSL = true
	}

	// client cert directory
	if op.CertConfig.ClientCertFile != "" && op.CertConfig.ClientKeyFile != "" {
		customSchedulerConfig.ClientCert.CertFile = op.CertConfig.ClientCertFile
		customSchedulerConfig.ClientCert.KeyFile = op.CertConfig.ClientKeyFile
		customSchedulerConfig.ClientCert.CAFile = op.CertConfig.CAFile
		customSchedulerConfig.ClientCert.IsSSL = true
	}

	// cloud netservice cert
	if op.CloudNetserviceClientCertFile != "" && op.CloudNetserviceClientKeyFile != "" {
		customSchedulerConfig.CloudNetserviceCert.CertFile = op.CloudNetserviceClientCertFile
		customSchedulerConfig.CloudNetserviceCert.KeyFile = op.CloudNetserviceClientKeyFile
		customSchedulerConfig.CloudNetserviceCert.CAFile = op.CloudNetserviceClientCaFile
		customSchedulerConfig.CloudNetserviceCert.IsSSL = true
	}

	return customSchedulerConfig
}
