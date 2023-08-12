

package app

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-cc-agent/config"
)

var mesosCacheInfo *NodeInfo

func synchronizeMesosNodeInfo(config *config.BcsCcAgentConfig) error {
	hostIp := os.Getenv("BCS_NODE_IP")
	if hostIp == "" {
		return fmt.Errorf("env [BCS_NODE_IP] is empty")
	}

	// init info cache
	mesosCacheInfo = &NodeInfo{}

	// sync info from bk-cmdb periodically
	go func() {
		ticker := time.NewTicker(time.Duration(1) * time.Minute)
		defer ticker.Stop()
		for {
			blog.Info("starting to synchronize node info...")

			nodeProperties, err := getInfoFromBkCmdb(config, hostIp)
			if err != nil {
				blog.Errorf("error synchronizing node info: %s", err.Error())
				continue
			}

			// currentNodeInfo represents the current node info
			currentNodeInfo := &NodeInfo{
				Properties: nodeProperties,
			}

			// if nodeInfo updated, then update to file and node label
			if !reflect.DeepEqual(*mesosCacheInfo, *currentNodeInfo) {
				mesosCacheInfo = currentNodeInfo
				err := updateMesosNodeInfo(k8sCacheInfo)
				if err != nil {
					blog.Errorf("error updating node info to file and node label: %s", err.Error())
					continue
				}
			}

			select {
			case <-ticker.C:
			}
		}
	}()

	return nil
}
