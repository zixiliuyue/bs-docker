

package available

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/scheduler/v4"

	"github.com/urfave/cli"
)

// NewEnableCommand xxx
func NewEnableCommand() cli.Command {
	return cli.Command{
		Name:  "enable",
		Usage: "enable agent by ip",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Usage: "Enable type, agent",
			},
			cli.StringFlag{
				Name:  "clusterid",
				Usage: "Cluster ID",
			},
			cli.StringFlag{
				Name:  "ip",
				Usage: "The ip of agent to enabled. Split by ,",
			},
		},
		Action: func(c *cli.Context) error {
			if err := enable(utils.NewClientContext(c)); err != nil {
				return err
			}
			return nil
		},
	}
}

func enable(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionType); err != nil {
		return err
	}

	resourceType := c.String(utils.OptionType)

	switch resourceType {
	case "agent":
		return enableAgent(c)
	default:
		return fmt.Errorf("invalid type: %s", resourceType)
	}
}

func enableAgent(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionClusterID, utils.OptionIP); err != nil {
		return err
	}

	ipList := utils.GetIPList(c.String(utils.OptionIP))

	scheduler := v4.NewBcsScheduler(utils.GetClientOption())
	err := scheduler.EnableAgent(c.ClusterID(), ipList)
	if err != nil {
		return fmt.Errorf("failed to enable agent: %v", err)
	}

	fmt.Printf("success to enable agent\n")
	return nil
}
