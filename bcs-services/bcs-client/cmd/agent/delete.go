

package agent

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"
)

func deleteAgentSetting(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID, utils.OptionIP); err != nil {
		return err
	}

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())

	ipList := utils.GetIPList(c.String(utils.OptionIP))

	err := scheduler.DeleteAgentSetting(c.ClusterID(), ipList)
	if err != nil {
		return fmt.Errorf("failed to delete agent setting: %v", err)
	}

	fmt.Printf("success to delete agent setting\n")
	return nil
}
