

package azure

const (
	// SystemNameInMetricAzure system name in metric for azure
	SystemNameInMetricAzure = "azure"
	// HandlerNameInMetricAzureSDK handler name in metric for azure sdk
	HandlerNameInMetricAzureSDK = "sdk"

	// AzureProtocolHTTP elb http protocol
	AzureProtocolHTTP = "HTTP"
	// AzureProtocolHTTPS elb https protocol
	AzureProtocolHTTPS = "HTTPS"
	// AzureProtocolTCP elb tcp protocol
	AzureProtocolTCP = "TCP"
	// AzureProtocolUDP elb udp protocol
	AzureProtocolUDP = "UDP"
	// AzureProtocolTLS elb tls protocol
	AzureProtocolTLS = "TLS"

	// DefaultRequestTimeout seconds that application wait for backend's response
	DefaultRequestTimeout = 20

	// DefaultLoadBalancerProbeInterval seconds that do probe
	DefaultLoadBalancerProbeInterval = 5
	// DefaultLoadBalancerProbeNumberOfProbes  The number of probes where if no response,
	//  will result in stopping further traffic from being delivered to the endpoint.
	//	This values allows endpoints to be taken out of rotation faster or slower than
	//	the typical times used in Azure.
	DefaultLoadBalancerProbeNumberOfProbes = 1

	// DefaultBackendPoolName name for default backend address pool
	DefaultBackendPoolName = "bkbcs-default-backendaddresspool"
	// DefaultBackendSettingName name for default backend setting
	DefaultBackendSettingName = "bkbcs-default-backendsetting"

	// CreateGoroutineLimit define how much goroutines can be used to create resource each time
	CreateGoroutineLimit = 10
	// DeleteGoroutineLimit  define how much goroutines can be used to delete resource each time
	DeleteGoroutineLimit = 10

	// osEnvSep sep of os env
	osEnvSep = ","
)
