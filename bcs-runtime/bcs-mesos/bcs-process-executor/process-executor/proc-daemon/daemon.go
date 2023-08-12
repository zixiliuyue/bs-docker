

package procdaemon

import (
	"fmt"

	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/types"
)

const (
	// ProcessDaemonEndpoint xxx
	ProcessDaemonEndpoint = "/var/run/process.sock"
)

type daemon struct {
	cli *HttpConnection
}

// NewDaemon xxx
func NewDaemon() ProcDaemon {
	return &daemon{
		cli: NewHttpConnection(ProcessDaemonEndpoint),
	}
}

// CreateProcess xxx
func (d *daemon) CreateProcess(processInfo *types.ProcessInfo) error {
	by, _ := json.Marshal(processInfo)
	_, err := d.cli.requestProcessDaemon("POST", "/process", by)
	if err != nil {
		blog.Errorf("daemon create process %s error %s", processInfo.Id, err.Error())
		return err
	}

	return nil
}

// InspectProcessStatus xxx
func (d *daemon) InspectProcessStatus(procId string) (*types.ProcessStatusInfo, error) {
	by, err := d.cli.requestProcessDaemon("GET", fmt.Sprintf("/process/%s/status", procId), nil)
	if err != nil {
		blog.Errorf("daemon inspect process %s error %s", procId, err.Error())
		return nil, err
	}

	var status *types.ProcessStatusInfo
	err = json.Unmarshal(by, &status)
	if err != nil {
		blog.Errorf("Unmarshal data %s to types.ProcessStatusInfo error %s", string(by), err.Error())
		return nil, err
	}

	return status, nil
}

// StopProcess xxx
func (d *daemon) StopProcess(procId string, timeout int) error {
	_, err := d.cli.requestProcessDaemon("PUT", fmt.Sprintf("/process/%s/stop/%d", procId, timeout), nil)
	if err != nil {
		blog.Errorf("daemon stop process %s error %s", procId, err.Error())
		return err
	}

	return nil
}

// DeleteProcess xxx
func (d *daemon) DeleteProcess(procId string) error {
	_, err := d.cli.requestProcessDaemon("DELETE", fmt.Sprintf("/process/%s", procId), nil)
	if err != nil {
		blog.Errorf("daemon delete process %s error %s", procId, err.Error())
		return err
	}

	return nil
}

// ReloadProcess xxx
func (d *daemon) ReloadProcess(procId string) error {
	return nil
}

// RestartProcess xxx
func (d *daemon) RestartProcess(procId string) error {
	return nil
}
