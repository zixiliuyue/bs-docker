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

package v2

import (
	"context"
	"time"

	v2 "github.com/Tencent/bk-bcs/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	scheme "github.com/Tencent/bk-bcs/bcs-mesos/kubebkbcsv2/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BcsConfigMapsGetter has a method to return a BcsConfigMapInterface.
// A group's client should implement this interface.
type BcsConfigMapsGetter interface {
	BcsConfigMaps(namespace string) BcsConfigMapInterface
}

// BcsConfigMapInterface has methods to work with BcsConfigMap resources.
type BcsConfigMapInterface interface {
	Create(ctx context.Context, bcsConfigMap *v2.BcsConfigMap, opts v1.CreateOptions) (*v2.BcsConfigMap, error)
	Update(ctx context.Context, bcsConfigMap *v2.BcsConfigMap, opts v1.UpdateOptions) (*v2.BcsConfigMap, error)
	UpdateStatus(ctx context.Context, bcsConfigMap *v2.BcsConfigMap, opts v1.UpdateOptions) (*v2.BcsConfigMap, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.BcsConfigMap, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.BcsConfigMapList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsConfigMap, err error)
	BcsConfigMapExpansion
}

// bcsConfigMaps implements BcsConfigMapInterface
type bcsConfigMaps struct {
	client rest.Interface
	ns     string
}

// newBcsConfigMaps returns a BcsConfigMaps
func newBcsConfigMaps(c *BkbcsV2Client, namespace string) *bcsConfigMaps {
	return &bcsConfigMaps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bcsConfigMap, and returns the corresponding bcsConfigMap object, and an error if there is any.
func (c *bcsConfigMaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.BcsConfigMap, err error) {
	result = &v2.BcsConfigMap{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BcsConfigMaps that match those selectors.
func (c *bcsConfigMaps) List(ctx context.Context, opts v1.ListOptions) (result *v2.BcsConfigMapList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.BcsConfigMapList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bcsConfigMaps.
func (c *bcsConfigMaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bcsConfigMap and creates it.  Returns the server's representation of the bcsConfigMap, and an error, if there is any.
func (c *bcsConfigMaps) Create(ctx context.Context, bcsConfigMap *v2.BcsConfigMap, opts v1.CreateOptions) (result *v2.BcsConfigMap, err error) {
	result = &v2.BcsConfigMap{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsConfigMap).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bcsConfigMap and updates it. Returns the server's representation of the bcsConfigMap, and an error, if there is any.
func (c *bcsConfigMaps) Update(ctx context.Context, bcsConfigMap *v2.BcsConfigMap, opts v1.UpdateOptions) (result *v2.BcsConfigMap, err error) {
	result = &v2.BcsConfigMap{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		Name(bcsConfigMap.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsConfigMap).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bcsConfigMaps) UpdateStatus(ctx context.Context, bcsConfigMap *v2.BcsConfigMap, opts v1.UpdateOptions) (result *v2.BcsConfigMap, err error) {
	result = &v2.BcsConfigMap{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		Name(bcsConfigMap.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsConfigMap).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bcsConfigMap and deletes it. Returns an error if one occurs.
func (c *bcsConfigMaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bcsConfigMaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bcsConfigMap.
func (c *bcsConfigMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsConfigMap, err error) {
	result = &v2.BcsConfigMap{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bcsconfigmaps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
