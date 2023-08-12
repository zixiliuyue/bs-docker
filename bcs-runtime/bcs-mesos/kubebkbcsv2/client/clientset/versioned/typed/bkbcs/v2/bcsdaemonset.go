

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	"context"
	"time"

	v2 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	scheme "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BcsDaemonsetsGetter has a method to return a BcsDaemonsetInterface.
// A group's client should implement this interface.
type BcsDaemonsetsGetter interface {
	BcsDaemonsets(namespace string) BcsDaemonsetInterface
}

// BcsDaemonsetInterface has methods to work with BcsDaemonset resources.
type BcsDaemonsetInterface interface {
	Create(ctx context.Context, bcsDaemonset *v2.BcsDaemonset, opts v1.CreateOptions) (*v2.BcsDaemonset, error)
	Update(ctx context.Context, bcsDaemonset *v2.BcsDaemonset, opts v1.UpdateOptions) (*v2.BcsDaemonset, error)
	UpdateStatus(ctx context.Context, bcsDaemonset *v2.BcsDaemonset, opts v1.UpdateOptions) (*v2.BcsDaemonset, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.BcsDaemonset, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.BcsDaemonsetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsDaemonset, err error)
	BcsDaemonsetExpansion
}

// bcsDaemonsets implements BcsDaemonsetInterface
type bcsDaemonsets struct {
	client rest.Interface
	ns     string
}

// newBcsDaemonsets returns a BcsDaemonsets
func newBcsDaemonsets(c *BkbcsV2Client, namespace string) *bcsDaemonsets {
	return &bcsDaemonsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bcsDaemonset, and returns the corresponding bcsDaemonset object, and an error if there is any.
func (c *bcsDaemonsets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.BcsDaemonset, err error) {
	result = &v2.BcsDaemonset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BcsDaemonsets that match those selectors.
func (c *bcsDaemonsets) List(ctx context.Context, opts v1.ListOptions) (result *v2.BcsDaemonsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.BcsDaemonsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bcsDaemonsets.
func (c *bcsDaemonsets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bcsDaemonset and creates it.  Returns the server's representation of the bcsDaemonset, and an error, if there is any.
func (c *bcsDaemonsets) Create(ctx context.Context, bcsDaemonset *v2.BcsDaemonset, opts v1.CreateOptions) (result *v2.BcsDaemonset, err error) {
	result = &v2.BcsDaemonset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsDaemonset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bcsDaemonset and updates it. Returns the server's representation of the bcsDaemonset, and an error, if there is any.
func (c *bcsDaemonsets) Update(ctx context.Context, bcsDaemonset *v2.BcsDaemonset, opts v1.UpdateOptions) (result *v2.BcsDaemonset, err error) {
	result = &v2.BcsDaemonset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		Name(bcsDaemonset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsDaemonset).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bcsDaemonsets) UpdateStatus(ctx context.Context, bcsDaemonset *v2.BcsDaemonset, opts v1.UpdateOptions) (result *v2.BcsDaemonset, err error) {
	result = &v2.BcsDaemonset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		Name(bcsDaemonset.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsDaemonset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bcsDaemonset and deletes it. Returns an error if one occurs.
func (c *bcsDaemonsets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bcsDaemonsets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bcsDaemonset.
func (c *bcsDaemonsets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsDaemonset, err error) {
	result = &v2.BcsDaemonset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bcsdaemonsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
