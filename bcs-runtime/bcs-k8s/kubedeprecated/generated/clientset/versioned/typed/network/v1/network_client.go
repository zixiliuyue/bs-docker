
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/apis/network/v1"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type NetworkV1Interface interface {
	RESTClient() rest.Interface
	CloudListenersGetter
}

// NetworkV1Client is used to interact with features provided by the network group.
type NetworkV1Client struct {
	restClient rest.Interface
}

func (c *NetworkV1Client) CloudListeners(namespace string) CloudListenerInterface {
	return newCloudListeners(c, namespace)
}

// NewForConfig creates a new NetworkV1Client for the given config.
func NewForConfig(c *rest.Config) (*NetworkV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkV1Client{client}, nil
}

// NewForConfigOrDie creates a new NetworkV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NetworkV1Client for the given RESTClient.
func New(c rest.Interface) *NetworkV1Client {
	return &NetworkV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NetworkV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
