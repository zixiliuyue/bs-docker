

package storage

var (
	dataTypeApp                = "Application"
	dataTypeTaskGroup          = "TaskGroup"
	dataTypeCfg                = "Configmap"
	dataTypeSecret             = "Secret"
	dataTypeDeploy             = "Deployment"
	dataTypeSvr                = "Service"
	dataTypeExpSVR             = "ExportService"
	dataTypeEp                 = "Endpoint"
	dataTypeIPPoolStatic       = "IPPoolStatic"
	dataTypeIPPoolStaticDetail = "IPPoolStaticDetail"

	actionDelete = "DELETE"
	actionPut    = "PUT"
	// actionPost   = "POST"

	handlerClusterNamespaceTypeName = "mesos_cluster_namespace_type_name"
	handlerClusterNamespaceType     = "mesos_cluster_namespace_type"
	handlerClusterTypeName          = "mesos_cluster_type_name"
	handlerClusterClusterType       = "mesos_cluster_type"

	handlerAllClusterType = "mesos_all_cluster_type"

	handlerWatchClusterNamespaceTypeName = "mesos_watch_cluster_namespace_type_name"
	handlerEvent                         = "events"
)
