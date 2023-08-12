

package endpoint

import (
	"errors"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	// ErrK8sConfigNotInited show K8sConfig not inited
	ErrK8sConfigNotInited = errors.New("k8sConfig not inited")
)

// K8sConfig xxx
type K8sConfig struct {
	Mater      string `json:"master"`
	KubeConfig string `json:"kubeConfig"`
}

func (c *K8sConfig) getRestConfig() (*rest.Config, error) {
	if c == nil {
		return nil, ErrK8sConfigNotInited
	}

	config, err := clientcmd.BuildConfigFromFlags(c.Mater, c.KubeConfig)
	if err != nil {
		blog.Errorf("BuildConfigFromFlags failed: %v", err)
		return nil, err
	}

	// config client qps/burst
	config.QPS = 1e6
	config.Burst = 1e6

	return config, nil
}

// GetKubernetesClient init kubernetes clientSet by k8sConfig
func (c *K8sConfig) GetKubernetesClient() (kubernetes.Interface, error) {
	if c == nil {
		return nil, ErrK8sConfigNotInited
	}

	config, err := c.getRestConfig()
	if err != nil {
		blog.Errorf("GetKubernetesClient call getRestConfig return err: %v", err)
		return nil, err
	}
	blog.Infof("GetKubernetesClient call getRestConfig successful")

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		blog.Errorf("GetKubernetesClient call NewForConfig failed: %v", err)
		return nil, err
	}

	return clientSet, nil
}
