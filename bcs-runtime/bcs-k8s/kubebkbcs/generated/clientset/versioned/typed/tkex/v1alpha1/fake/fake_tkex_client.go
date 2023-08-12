
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/generated/clientset/versioned/typed/tkex/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTkexV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTkexV1alpha1) GameDeployments(namespace string) v1alpha1.GameDeploymentInterface {
	return &FakeGameDeployments{c, namespace}
}

func (c *FakeTkexV1alpha1) GameStatefulSets(namespace string) v1alpha1.GameStatefulSetInterface {
	return &FakeGameStatefulSets{c, namespace}
}

func (c *FakeTkexV1alpha1) HookRuns(namespace string) v1alpha1.HookRunInterface {
	return &FakeHookRuns{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTkexV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
