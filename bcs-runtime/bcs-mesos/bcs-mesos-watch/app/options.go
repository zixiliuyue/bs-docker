

package app

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/static"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-mesos-watch/types"
)

// DefaultConfig return default command line config
func DefaultConfig() *types.CmdConfig {
	return &types.CmdConfig{
		ClusterID:              "",
		ClusterInfo:            "127.0.0.1:2181/blueking",
		CAFile:                 "",
		CertFile:               "",
		KeyFile:                "",
		PassWord:               static.ClientCertPwd,
		RegDiscvSvr:            "",
		Address:                "127.0.0.1",
		ApplicationThreadNum:   100,
		TaskgroupThreadNum:     100,
		ExportserviceThreadNum: 100,
		DeploymentThreadNum:    100,
		ServerPassWord:         static.ServerCertPwd,
		IsExternal:             false,
	}
}
