

package haproxy

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"

	"github.com/haproxytech/client-native/runtime"
	"github.com/haproxytech/models"
)

// RuntimeClient haproxy runtime client
type RuntimeClient struct {
	client   *runtime.SingleRuntime
	SockPath string
}

// NewRuntimeClient create runtime client
func NewRuntimeClient(sockPath string) (*RuntimeClient, error) {
	client := &runtime.SingleRuntime{}
	err := client.Init(sockPath, 0)
	if err != nil {
		blog.Errorf("init haproxy runtime failed with sock path %s, err %s", sockPath, err.Error())
		return nil, err
	}
	return &RuntimeClient{
		client:   client,
		SockPath: sockPath,
	}, nil
}

// NewSetServerWeightCommand create new command that set server weight
func (rc *RuntimeClient) NewSetServerWeightCommand(backend, server string, weight int) string {
	return fmt.Sprintf("set server %s/%s weight %d", backend, server, weight)
}

// NewDisableServerCommand create new command that disable unused server
func (rc *RuntimeClient) NewDisableServerCommand(backend, server string) string {
	return fmt.Sprintf("disable server %s/%s", backend, server)
}

// NewEnableServerCommand create new command that enable unused server
func (rc *RuntimeClient) NewEnableServerCommand(backend, server string) string {
	return fmt.Sprintf("enable server %s/%s", backend, server)
}

// NewSetServerAddrCommand create new command that set server addr
func (rc *RuntimeClient) NewSetServerAddrCommand(backend, server, addr string, port int) string {
	return fmt.Sprintf("set server %s/%s addr %s port %d", backend, server, addr, port)
}

// GetStats get statistic of haproxy
func (rc *RuntimeClient) GetStats() *models.NativeStatsCollection {
	return rc.client.GetStats()
}

// GetInfo get info of haproxy
func (rc *RuntimeClient) GetInfo() (models.ProcessInfoHaproxy, error) {
	return rc.client.GetInfo()
}

// ExecuteCommand send command to haproxy
func (rc *RuntimeClient) ExecuteCommand(cmd string) error {
	return rc.client.Execute(cmd)
}

// ExecuteRaw execute raw command
func (rc *RuntimeClient) ExecuteRaw(cmd string) (string, error) {
	return rc.client.ExecuteRaw(cmd)
}
