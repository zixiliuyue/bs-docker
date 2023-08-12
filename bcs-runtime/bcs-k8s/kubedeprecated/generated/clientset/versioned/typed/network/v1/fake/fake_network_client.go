
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/generated/clientset/versioned/typed/network/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkV1 struct {
	*testing.Fake
}

func (c *FakeNetworkV1) CloudListeners(namespace string) v1.CloudListenerInterface {
	return &FakeCloudListeners{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
