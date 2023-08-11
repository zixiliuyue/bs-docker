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

	v2 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/mesosv2/apis/bkbcs/v2"
	scheme "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/mesosv2/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TaskGroupsGetter has a method to return a TaskGroupInterface.
// A group's client should implement this interface.
type TaskGroupsGetter interface {
	TaskGroups(namespace string) TaskGroupInterface
}

// TaskGroupInterface has methods to work with TaskGroup resources.
type TaskGroupInterface interface {
	Create(ctx context.Context, taskGroup *v2.TaskGroup, opts v1.CreateOptions) (*v2.TaskGroup, error)
	Update(ctx context.Context, taskGroup *v2.TaskGroup, opts v1.UpdateOptions) (*v2.TaskGroup, error)
	UpdateStatus(ctx context.Context, taskGroup *v2.TaskGroup, opts v1.UpdateOptions) (*v2.TaskGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.TaskGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.TaskGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.TaskGroup, err error)
	TaskGroupExpansion
}

// taskGroups implements TaskGroupInterface
type taskGroups struct {
	client rest.Interface
	ns     string
}

// newTaskGroups returns a TaskGroups
func newTaskGroups(c *BkbcsV2Client, namespace string) *taskGroups {
	return &taskGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the taskGroup, and returns the corresponding taskGroup object, and an error if there is any.
func (c *taskGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.TaskGroup, err error) {
	result = &v2.TaskGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("taskgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TaskGroups that match those selectors.
func (c *taskGroups) List(ctx context.Context, opts v1.ListOptions) (result *v2.TaskGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.TaskGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("taskgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested taskGroups.
func (c *taskGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("taskgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a taskGroup and creates it.  Returns the server's representation of the taskGroup, and an error, if there is any.
func (c *taskGroups) Create(ctx context.Context, taskGroup *v2.TaskGroup, opts v1.CreateOptions) (result *v2.TaskGroup, err error) {
	result = &v2.TaskGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("taskgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(taskGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a taskGroup and updates it. Returns the server's representation of the taskGroup, and an error, if there is any.
func (c *taskGroups) Update(ctx context.Context, taskGroup *v2.TaskGroup, opts v1.UpdateOptions) (result *v2.TaskGroup, err error) {
	result = &v2.TaskGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("taskgroups").
		Name(taskGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(taskGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *taskGroups) UpdateStatus(ctx context.Context, taskGroup *v2.TaskGroup, opts v1.UpdateOptions) (result *v2.TaskGroup, err error) {
	result = &v2.TaskGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("taskgroups").
		Name(taskGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(taskGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the taskGroup and deletes it. Returns an error if one occurs.
func (c *taskGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("taskgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *taskGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("taskgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched taskGroup.
func (c *taskGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.TaskGroup, err error) {
	result = &v2.TaskGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("taskgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
