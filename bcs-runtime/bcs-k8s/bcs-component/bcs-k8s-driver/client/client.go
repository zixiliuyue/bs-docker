

// Package client xxx
package client

import (
	urllib "net/url"

	glog "github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-driver/kubedriver/options"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// NewClientSet create k8s clientset
func NewClientSet(k8sMasterUrl string, tlsCfg options.TLSConfig) *kubernetes.Clientset {

	glog.V(3).Infof("k8sConfig.Master is set: %s", k8sMasterUrl)

	config := &rest.Config{
		Host:  k8sMasterUrl,
		QPS:   1e6,
		Burst: 1e6,
	}
	kubeURL, _ := urllib.Parse(k8sMasterUrl)
	if kubeURL.Scheme == options.HTTPS {
		if tlsCfg.CAFile == "" || tlsCfg.CertFile == "" || tlsCfg.KeyFile == "" {
			return nil
		}
		config.TLSClientConfig = rest.TLSClientConfig{
			CAFile:   tlsCfg.CAFile,
			CertFile: tlsCfg.CertFile,
			KeyFile:  tlsCfg.KeyFile,
		}
	}

	// 2.2 creates the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil
	}
	return clientSet
}
