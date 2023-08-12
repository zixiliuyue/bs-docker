

package create

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"
)

func createApplication(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID); err != nil {
		return err
	}

	data, err := c.FileData()
	if err != nil {
		return err
	}

	namespace, err := utils.ParseNamespaceFromJSON(data)
	if err != nil {
		return err
	}

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())
	err = scheduler.CreateApplication(c.ClusterID(), namespace, data)
	if err != nil {
		return fmt.Errorf("failed to create application: %v", err)
	}

	fmt.Printf("success to create application\n")
	return nil
}
