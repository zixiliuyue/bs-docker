

package available

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"

	"github.com/urfave/cli"
)

// NewDisableCommand xxx
func NewDisableCommand() cli.Command {
	return cli.Command{
		Name:  "disable",
		Usage: "disable agent by ip",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Usage: "Disable type, agent",
			},
			cli.StringFlag{
				Name:  "clusterid",
				Usage: "Cluster ID",
			},
			cli.StringFlag{
				Name:  "ip",
				Usage: "The ip of agent to disabled. Split by ,",
			},
		},
		Action: func(c *cli.Context) error {
			if err := disable(utils.NewClientContext(c)); err != nil {
				return err
			}
			return nil
		},
	}
}

func disable(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionType); err != nil {
		return err
	}

	resourceType := c.String(utils.OptionType)

	switch resourceType {
	case "agent":
		return disableAgent(c)
	default:
		return fmt.Errorf("invalid type: %s", resourceType)
	}
}

func disableAgent(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID, utils.OptionIP); err != nil {
		return err
	}

	ipList := utils.GetIPList(c.String(utils.OptionIP))

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())
	err := scheduler.DisableAgent(c.ClusterID(), ipList)
	if err != nil {
		return fmt.Errorf("failed to disable agent: %v", err)
	}

	fmt.Printf("success to disable agent\n")
	return nil
}
