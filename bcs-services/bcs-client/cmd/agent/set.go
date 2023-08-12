

package agent

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"
)

func setAgentSetting(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID); err != nil {
		return err
	}

	data, err := c.FileData()
	if err != nil {
		return fmt.Errorf("failed to parse file: %v", err)
	}

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())

	err = scheduler.SetAgentSetting(c.ClusterID(), data)
	if err != nil {
		return fmt.Errorf("failed to set agent setting: %v", err)
	}

	fmt.Printf("success to set agent setting\n")
	return nil
}
