

package agent

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"
)

func updateAgentSetting(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID, utils.OptionKey, utils.OptionIP); err != nil {
		return err
	}

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())

	ipList := utils.GetIPList(c.String(utils.OptionIP))
	key := c.String(utils.OptionKey)

	var err error
	if c.IsSet(utils.OptionString) {
		err = scheduler.UpdateStringAgentSetting(c.ClusterID(), ipList, key, c.String(utils.OptionString))
	} else if c.IsSet(utils.OptionScalar) {
		err = scheduler.UpdateScalarAgentSetting(c.ClusterID(), ipList, key, c.Float64(utils.OptionScalar))
	} else {
		return fmt.Errorf("string or scalar must be specified while updating")
	}

	if err != nil {
		return fmt.Errorf("failed to update agent setting: %v", err)
	}

	fmt.Printf("success to update agent setting\n")
	return nil
}
