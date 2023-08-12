

package endpoint

import (
	"context"
	"errors"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"os/user"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getHomePath() string {
	u, err := user.Current()
	if err == nil {
		return u.HomeDir
	}
	return ""
}

func getK8sConfig() *K8sConfig {
	homeDir := getHomePath()
	if homeDir == "" {
		return nil
	}

	return &K8sConfig{
		Mater:      "",
		KubeConfig: fmt.Sprintf("%s/.kube/config", homeDir),
	}
}

func getClientSet() (kubernetes.Interface, error) {
	k8sConfig := getK8sConfig()
	if k8sConfig == nil {
		return nil, errors.New("getK8sConfig failed")
	}

	clientSet, err := k8sConfig.GetKubernetesClient()
	if err != nil {
		errMsg := fmt.Errorf("GetKubernetesClient failed: %v", err)
		return nil, errMsg
	}

	return clientSet, nil
}

func TestK8sConfig_GetKubernetesClient(t *testing.T) {

	clientSet, err := getClientSet()
	if err != nil {
		t.Logf("getClientSet failed: %v", err)
		return
	}
	t.Logf("GetKubernetesClient successful")

	ep, err := clientSet.CoreV1().Endpoints("default").Get(context.Background(), "kubernetes", metav1.GetOptions{})
	if err != nil {
		t.Logf("get namespace default Endpoints kubernetes failed: %v", err)
		return
	}

	t.Logf("%+v", ep)
}
