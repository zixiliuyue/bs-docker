

package sqlstore

import (
	"strings"

	"fmt"
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"github.com/dchest/uniuri"
)

// RegisterTokenLen xxx
const RegisterTokenLen = 128

// GetRegisterToken return the registerToken by clusterId
func GetRegisterToken(clusterId string) *m.RegisterToken {
	token := m.RegisterToken{}
	GCoreDB.Where(&m.RegisterToken{ClusterId: clusterId}).First(&token)
	if token.ID != 0 {
		return &token
	}
	return nil
}

// CreateRegisterToken creates a new registerToken for given clusterId
func CreateRegisterToken(clusterId string) error {
	token := m.RegisterToken{
		ClusterId: clusterId,
		Token:     uniuri.NewLen(RegisterTokenLen),
	}
	err := GCoreDB.Create(&token).Error
	if err == nil {
		return err
	}

	// Transform raw db error messsage
	if strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
		return fmt.Errorf("token already exists")
	}
	return err
}
