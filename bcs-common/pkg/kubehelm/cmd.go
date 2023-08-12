

package kubehelm

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"k8s.io/klog"
)

// command to use helm struct
type cmdHelm struct{}

// NewCmdHelm new cmd helm struct, the object requires helm command-line tool
func NewCmdHelm() KubeHelm {
	return &cmdHelm{}
}

// InstallChart xxx
// helm install --name xxxx chart-dir --set k1=v1 --set k2=v2 --kube-apiserver=xxxx --kube-token=xxxxx --kubeconfig kubeconfig
func (h *cmdHelm) InstallChart(inf InstallFlags, glf GlobalFlags) error {
	gPara, err := glf.ParseParameters()
	if err != nil {
		return err
	}
	defer os.Remove(glf.Kubeconfig)

	parameters := inf.ParseParameters() + gPara
	klog.Infof("helm install%s", parameters)
	os.Remove("install.sh")
	file, err := os.OpenFile("install.sh", os.O_CREATE|os.O_RDWR, 0755) // NOCC:gas/permission(设计如此)
	if err != nil {
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		file.Close()
		return err
	}
	_, err = file.Write([]byte(fmt.Sprintf("helm install%s", parameters)))
	file.Close()
	if err != nil {
		return err
	}
	cmd := exec.Command("/bin/bash", "-c", "./install.sh") // NOCC:gas/subprocess(设计如此)
	buf := bytes.NewBuffer(make([]byte, 1024))
	cmd.Stderr = buf
	err = cmd.Run()
	if err != nil {
		klog.Errorf("helm install failed, stderr %s error %s", buf.String(), err.Error())
		return err
	}

	klog.Infof("helm install %s success", parameters)
	return nil
}
