

package add

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	"github.com/urfave/cli"
)

// NewAddCommand sub command add registration
func NewAddCommand() cli.Command {
	return cli.Command{
		Name:  "add",
		Usage: "add cidr",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Usage: "add type, value can be cidr",
			},
			cli.StringFlag{
				Name:  "from-file, f",
				Usage: "reading with configuration `FILE`",
			},
			cli.StringFlag{
				Name:  "vpcid",
				Usage: "vpc id",
			},
		},
		Action: func(c *cli.Context) error {
			return add(utils.NewClientContext(c))
		},
	}
}

func add(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionType); err != nil {
		return err
	}

	resourceType := c.String(utils.OptionType)

	switch resourceType {
	case "cidr":
		return initVpcCidr(c)
	default:
		return fmt.Errorf("invalid type: %s", resourceType)
	}
}
