

package pkgs

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/pkg/msgqueue"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/cmd/config"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-alert-manager/pkg/consumer"
)

// ResourceSwitch xxx
type ResourceSwitch string

const (
	// ResourceSubOn on sub handler
	ResourceSubOn ResourceSwitch = "on"
	// ResourceSubOff off sub handler
	ResourceSubOff ResourceSwitch = "off"
)

// GetFactoryConsumers init consumer object according to resourceSubInfo
func GetFactoryConsumers(options *config.AlertManagerOptions) []consumer.Consumer {
	resourceSubInfo := parseResourceSubs(options)
	var consumers []consumer.Consumer

	for resource, switchKey := range resourceSubInfo {
		if strings.EqualFold(switchKey, string(ResourceSubOn)) {
			con := handlerFactory(resource, options)
			if con != nil {
				consumers = append(consumers, con)
			}
		}
	}

	return consumers
}

func handlerFactory(resourceKind string, options *config.AlertManagerOptions) consumer.Consumer {
	switch resourceKind {
	case msgqueue.EventSubscribeType:
		return GetEventSyncHandler(options)
	}
	return nil
}

func parseResourceSubs(options *config.AlertManagerOptions) map[string]string {
	resourceSubs := make(map[string]string)

	for _, resource := range options.ResourceSubs {
		resourceKind := resource.Category
		switchKey := resource.Switch

		if _, ok := resourceSubs[resourceKind]; !ok {
			resourceSubs[resourceKind] = switchKey
		}
	}

	blog.Infof("parseResourceSubs %v", resourceSubs)
	return resourceSubs
}
