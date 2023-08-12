

// Package modules xxx
package modules

const (
	// bcs-services module list

	// BCSModuleUserManager for bcs-user-manager
	BCSModuleUserManager = "usermanager"
	// BCSModuleClusterManager for bcs-cluster-manager
	BCSModuleClusterManager = "clustermanager"
	// BCSModuleGatewayDiscovery for bcs-gateway-discovery
	BCSModuleGatewayDiscovery = "gatewaydiscovery"
	// BCSModuleNetworkdetection for bcs-network-detection
	BCSModuleNetworkdetection = "networkdetection"
	// BCSModuleBKCMDBSynchronizer for bcs-bkcmdb-synchronizer
	BCSModuleBKCMDBSynchronizer = "bkcmdb-synchronizer"
	// BCSModuleStorage for bcs-storage
	BCSModuleStorage = "storage"
	// BCSModuleIPService for bcs-ipservice
	BCSModuleIPService = "ipservice"
	// BCSModuleNetService for bcs-netservice
	BCSModuleNetService = "netservice"
	// BCSModuleDNS for bcs-dns
	BCSModuleDNS = "dns"
	// BCSModuleMetricService for bcs-metricservice
	BCSModuleMetricService = "metricservice"
	// BCSModuleMetricCollector for bcs-metriccollector
	BCSModuleMetricCollector = "metriccollector"
	// end of bcs-services module list

	// bcs mesos module list

	// BCSModuleMesosdriver for bcs-mesos-driver
	BCSModuleMesosdriver = "mesosdriver"
	// BCSModuleMesoswatch for bcs-mesos-watch
	BCSModuleMesoswatch = "mesosdatawatch"
	// BCSModuleMesosWebconsole for bcs-mesoswebconcole
	BCSModuleMesosWebconsole = "mesoswebconsole"
	// BCSModuleScheduler for bcs-scheduler
	BCSModuleScheduler = "scheduler"
	// end of bcs mesos module list

	// BCSModuleKubeagent for bcs-kube-agent
	BCSModuleKubeagent = "kubeagent"
	// BCSModuleKubewatch for bcs-k8s-watch
	BCSModuleKubewatch = "kubedatawatch"

	// mode for mesosdriver & kubeagent

	// BCSConnectModeTunnel mode for tunnel
	BCSConnectModeTunnel = "websockettunnel"
	// BCSConnectModeDirect mode for direct connection
	BCSConnectModeDirect = "direct"
)
