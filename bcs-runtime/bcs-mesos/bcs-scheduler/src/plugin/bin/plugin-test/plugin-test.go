

package plugin_test

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/plugin"
)

// plugin must implement
// func GetHostAttributes([]string) (map[string]*types.HostAttributes,error)
// func input: ip list, example: []string{"127.0.0.10","127.0.0.11","127.0.0.12"}
// func ouput: map key = ip, example: map["127.0.0.10"] = &types.HostAttributes{}
// implement func Init(para *types.InitPluginParameter) error
// func input: *types.InitPluginParameter
// func output: error

// for example

var initPara *plugin.InitPluginParameter

// Init xxx
func Init(para *plugin.InitPluginParameter) error {
	initPara = para
	return nil
}

// Uninit xxx
func Uninit() {
	// xxx
}

// GetHostAttributes xxx
func GetHostAttributes(para *plugin.HostPluginParameter) (map[string]*plugin.HostAttributes, error) {
	atrrs := make(map[string]*plugin.HostAttributes)

	for _, ip := range para.Ips {
		hostAttr := &plugin.HostAttributes{
			Ip:         ip,
			Attributes: make([]*plugin.Attribute, 0),
		}

		atrri := &plugin.Attribute{
			Name:   "ip-resources",
			Type:   plugin.ValueScalar,
			Scalar: plugin.Value_Scalar{Value: 10},
		}
		hostAttr.Attributes = append(hostAttr.Attributes, atrri)

		atrrs[ip] = hostAttr
	}

	return atrrs, nil
}
