

package discovery

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/queue"
	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/mesh/v1"

	"k8s.io/apimachinery/pkg/labels"
)

// Cluster maintenance resource for specified cluster, offer cluster api for
// getting & listing all discovery datas
type Cluster interface {
	// GetName get cluster name
	GetName() string
	// Run starting cluster all discovery goroutine for caching data
	Run()
	// Stop stop cluster
	Stop()
	// AppSvcs get controller of AppSvc
	AppSvcs() AppSvcController
	// AppNodes get controller of AppNode
	AppNodes() AppNodeController
}

// AppSvcController controller for AppSvc encapsulation
type AppSvcController interface {
	// GetAppSvc get specified AppSvc by namespace, name
	GetAppSvc(ns, name string) (*v1.AppSvc, error)
	// ListAppSvcs List all AppSvc datas
	ListAppSvcs(selector labels.Selector) ([]*v1.AppSvc, error)
	// RegisterAppSvcQueue register event callback for AppSvc
	RegisterAppSvcQueue(handler queue.Queue)
}

// AppNodeController controller for AppNode encapsulation
type AppNodeController interface {
	// ListAppNodes list all appNode datas
	ListAppNodes(selector labels.Selector) ([]*v1.AppNode, error)
	// GetAppNode get specified AppNode by namespace, name
	GetAppNode(ns, name string) (*v1.AppNode, error)
	// RegisterAppNodeQueue register event callback for AppNode
	RegisterAppNodeQueue(handler queue.Queue)
}
