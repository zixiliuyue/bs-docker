

package imageacceleration

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func initK8sClient() (*kubernetes.Clientset, error) {
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "get incluster rest config failed")
	}
	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, errors.Wrapf(err, "init incluster k8s client failed")
	}
	return client, nil
}
