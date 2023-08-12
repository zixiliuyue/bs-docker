

// Package plugin xxx
package plugin

import (
	typesplugin "github.com/Tencent/bk-bcs/bcs-common/common/plugin"
)

// Plugin interface definition for bcs-scheduler
type Plugin interface {

	// GetHostAttributes xxx
	// outer scheduler function
	// input: ip list,example: []string{"127.0.0.10","127.0.0.11","127.0.0.12"}
	// ouput: map key = ip,example: map["127.0.0.10"] = &types.HostAttributes{}
	GetHostAttributes(*typesplugin.HostPluginParameter) (map[string]*typesplugin.HostAttributes, error)
}
