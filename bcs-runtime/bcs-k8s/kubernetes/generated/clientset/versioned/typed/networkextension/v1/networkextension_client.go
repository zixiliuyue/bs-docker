
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type NetworkextensionV1Interface interface {
	RESTClient() rest.Interface
	IngressesGetter
	ListenersGetter
	PortBindingsGetter
	PortPoolsGetter
}

// NetworkextensionV1Client is used to interact with features provided by the networkextension group.
type NetworkextensionV1Client struct {
	restClient rest.Interface
}

func (c *NetworkextensionV1Client) Ingresses(namespace string) IngressInterface {
	return newIngresses(c, namespace)
}

func (c *NetworkextensionV1Client) Listeners(namespace string) ListenerInterface {
	return newListeners(c, namespace)
}

func (c *NetworkextensionV1Client) PortBindings(namespace string) PortBindingInterface {
	return newPortBindings(c, namespace)
}

func (c *NetworkextensionV1Client) PortPools(namespace string) PortPoolInterface {
	return newPortPools(c, namespace)
}

// NewForConfig creates a new NetworkextensionV1Client for the given config.
func NewForConfig(c *rest.Config) (*NetworkextensionV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkextensionV1Client{client}, nil
}

// NewForConfigOrDie creates a new NetworkextensionV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkextensionV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NetworkextensionV1Client for the given RESTClient.
func New(c rest.Interface) *NetworkextensionV1Client {
	return &NetworkextensionV1Client{c}
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
func (c *NetworkextensionV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
