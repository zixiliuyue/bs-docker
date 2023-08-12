

// Package utils xxx
package utils

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/storages/sqlstore"
)

// TurnCredentialsIntoConfig xxx
func TurnCredentialsIntoConfig(clusterCredentials *m.ClusterCredentials) *restclient.Config {
	tlsClientConfig := restclient.TLSClientConfig{
		CAData: []byte(clusterCredentials.CaCertData),
	}

	return &restclient.Config{
		Host:            clusterCredentials.ServerAddresses,
		BearerToken:     clusterCredentials.UserToken,
		TLSClientConfig: tlsClientConfig,
	}
}

// GetKubeClient xxx
func GetKubeClient(clusterId string) (*kubernetes.Clientset, error) {
	clusterCredentials := sqlstore.GetCredentials(clusterId)
	if clusterCredentials == nil {
		return nil, fmt.Errorf("cluster credentias not found for cluster: %s", clusterId)
	}
	apiServerList := clusterCredentials.GetServerAddressesList()
	var kubeClient *kubernetes.Clientset
	for _, apiServer := range apiServerList {
		hostPort := strings.TrimPrefix(apiServer, "https://")
		if err := pingEndpoint(hostPort); err == nil {
			clusterCredentials.ServerAddresses = apiServer
			kubeClient, err = makeKubeclient(clusterCredentials)
			return kubeClient, err
		}
	}
	return nil, fmt.Errorf("couldn't find an available apiserver for cluster: %s", clusterId)
}

// pingEndpoint xxx
// probe the health of the apiserver address for 3 times
func pingEndpoint(host string) error {
	var err error
	for i := 0; i < 3; i++ {
		err = dialTls(host)
		if err != nil && strings.Contains(err.Error(), "connection refused") {
			blog.Errorf("Error connecting the apiserver %s. Retrying...: %s", host, err.Error())
			time.Sleep(time.Second * 3)
			continue
		} else if err != nil {
			blog.Errorf("Error connecting the apiserver %s: %s", host, err.Error())
			return err
		}
		return nil
	}
	return err
}

func makeKubeclient(clusterCredentials *m.ClusterCredentials) (*kubernetes.Clientset, error) {
	restConfig := TurnCredentialsIntoConfig(clusterCredentials)
	kubeClient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("error when building kubeclient from restconfig: %s", err.Error())
	}
	return kubeClient, nil
}

func dialTls(host string) error {
	conf := &tls.Config{
		InsecureSkipVerify: true, // nolint
	}
	conn, err := tls.Dial("tcp", host, conf)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
