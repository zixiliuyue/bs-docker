

package add

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-client/cmd/utils"
	v1 "github.com/Tencent/bk-bcs/bcs-services/bcs-client/pkg/usermanager/v1"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/user-manager/v1http/tke"
)

func initVpcCidr(c *utils.ClientContext) error {
	if err := c.MustSpecified(utils.OptionVpc); err != nil {
		return err
	}

	var data []byte
	var err error
	if !c.IsSet(utils.OptionFile) {
		// reading all data from stdin
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = c.FileData()
	}
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return fmt.Errorf("failed to grant: no available resource datas")
	}

	var cidrs []tke.TkeCidr
	err = json.Unmarshal(data, &cidrs)
	if err != nil {
		return err
	}
	form := tke.AddTkeCidrForm{
		Vpc:      c.String(utils.OptionVpc),
		TkeCidrs: cidrs,
	}
	data, err = json.Marshal(form)
	if err != nil {
		return err
	}
	userManager := v1.NewBcsUserManager(utils.GetClientOption())
	err = userManager.AddVpcCidrs(data)
	if err != nil {
		return fmt.Errorf("failed to init cidr to vpc %s: %v", c.String(utils.OptionVpc), err)
	}

	fmt.Printf("success to init cidr\n")
	return nil
}
