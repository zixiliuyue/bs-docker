
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/generated/clientset/versioned/typed/cloud/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCloudV1 struct {
	*testing.Fake
}

func (c *FakeCloudV1) CloudIPs(namespace string) v1.CloudIPInterface {
	return &FakeCloudIPs{c, namespace}
}

func (c *FakeCloudV1) CloudIPQuotas(namespace string) v1.CloudIPQuotaInterface {
	return &FakeCloudIPQuotas{c, namespace}
}

func (c *FakeCloudV1) CloudSubnets(namespace string) v1.CloudSubnetInterface {
	return &FakeCloudSubnets{c, namespace}
}

func (c *FakeCloudV1) NodeNetworks(namespace string) v1.NodeNetworkInterface {
	return &FakeNodeNetworks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCloudV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
