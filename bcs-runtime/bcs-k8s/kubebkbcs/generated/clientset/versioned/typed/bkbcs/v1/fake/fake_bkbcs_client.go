
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/generated/clientset/versioned/typed/bkbcs/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBkbcsV1 struct {
	*testing.Fake
}

func (c *FakeBkbcsV1) BcsDbPrivConfigs(namespace string) v1.BcsDbPrivConfigInterface {
	return &FakeBcsDbPrivConfigs{c, namespace}
}

func (c *FakeBkbcsV1) BcsLogConfigs(namespace string) v1.BcsLogConfigInterface {
	return &FakeBcsLogConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBkbcsV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
