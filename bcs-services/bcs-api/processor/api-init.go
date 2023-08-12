

package processor

import (
	// import v4 http clusterkeeper actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/clusterkeeper"
	// import v4 http k8s actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/k8s"
	// import v4 http mesos actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/mesos"
	// import v4 http metrics actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/metrics"
	// import v4 http netservice actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/netservice"
	// import v4 http storage actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/storage"
	// import v4 http detection actions
	_ "github.com/Tencent/bk-bcs/bcs-services/bcs-api/processor/http/actions/v4http/detection"
)
