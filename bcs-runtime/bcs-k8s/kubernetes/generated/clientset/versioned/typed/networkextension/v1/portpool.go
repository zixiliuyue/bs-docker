
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
	scheme "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PortPoolsGetter has a method to return a PortPoolInterface.
// A group's client should implement this interface.
type PortPoolsGetter interface {
	PortPools(namespace string) PortPoolInterface
}

// PortPoolInterface has methods to work with PortPool resources.
type PortPoolInterface interface {
	Create(ctx context.Context, portPool *v1.PortPool, opts metav1.CreateOptions) (*v1.PortPool, error)
	Update(ctx context.Context, portPool *v1.PortPool, opts metav1.UpdateOptions) (*v1.PortPool, error)
	UpdateStatus(ctx context.Context, portPool *v1.PortPool, opts metav1.UpdateOptions) (*v1.PortPool, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PortPool, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PortPoolList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PortPool, err error)
	PortPoolExpansion
}

// portPools implements PortPoolInterface
type portPools struct {
	client rest.Interface
	ns     string
}

// newPortPools returns a PortPools
func newPortPools(c *NetworkextensionV1Client, namespace string) *portPools {
	return &portPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the portPool, and returns the corresponding portPool object, and an error if there is any.
func (c *portPools) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PortPool, err error) {
	result = &v1.PortPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("portpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PortPools that match those selectors.
func (c *portPools) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PortPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PortPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("portpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested portPools.
func (c *portPools) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("portpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a portPool and creates it.  Returns the server's representation of the portPool, and an error, if there is any.
func (c *portPools) Create(ctx context.Context, portPool *v1.PortPool, opts metav1.CreateOptions) (result *v1.PortPool, err error) {
	result = &v1.PortPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("portpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(portPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a portPool and updates it. Returns the server's representation of the portPool, and an error, if there is any.
func (c *portPools) Update(ctx context.Context, portPool *v1.PortPool, opts metav1.UpdateOptions) (result *v1.PortPool, err error) {
	result = &v1.PortPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("portpools").
		Name(portPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(portPool).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *portPools) UpdateStatus(ctx context.Context, portPool *v1.PortPool, opts metav1.UpdateOptions) (result *v1.PortPool, err error) {
	result = &v1.PortPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("portpools").
		Name(portPool.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(portPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the portPool and deletes it. Returns an error if one occurs.
func (c *portPools) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("portpools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *portPools) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("portpools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched portPool.
func (c *portPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PortPool, err error) {
	result = &v1.PortPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("portpools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
