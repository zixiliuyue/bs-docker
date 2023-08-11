/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	clbv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/generated/clientset/versioned/typed/clb/v1"
	meshv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/generated/clientset/versioned/typed/mesh/v1"
	networkv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/generated/clientset/versioned/typed/network/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ClbV1() clbv1.ClbV1Interface
	MeshV1() meshv1.MeshV1Interface
	NetworkV1() networkv1.NetworkV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	clbV1     *clbv1.ClbV1Client
	meshV1    *meshv1.MeshV1Client
	networkV1 *networkv1.NetworkV1Client
}

// ClbV1 retrieves the ClbV1Client
func (c *Clientset) ClbV1() clbv1.ClbV1Interface {
	return c.clbV1
}

// MeshV1 retrieves the MeshV1Client
func (c *Clientset) MeshV1() meshv1.MeshV1Interface {
	return c.meshV1
}

// NetworkV1 retrieves the NetworkV1Client
func (c *Clientset) NetworkV1() networkv1.NetworkV1Interface {
	return c.networkV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.clbV1, err = clbv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.meshV1, err = meshv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkV1, err = networkv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.clbV1 = clbv1.NewForConfigOrDie(c)
	cs.meshV1 = meshv1.NewForConfigOrDie(c)
	cs.networkV1 = networkv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.clbV1 = clbv1.New(c)
	cs.meshV1 = meshv1.New(c)
	cs.networkV1 = networkv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
