
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/cloud/v1"
	scheme "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudIPsGetter has a method to return a CloudIPInterface.
// A group's client should implement this interface.
type CloudIPsGetter interface {
	CloudIPs(namespace string) CloudIPInterface
}

// CloudIPInterface has methods to work with CloudIP resources.
type CloudIPInterface interface {
	Create(ctx context.Context, cloudIP *v1.CloudIP, opts metav1.CreateOptions) (*v1.CloudIP, error)
	Update(ctx context.Context, cloudIP *v1.CloudIP, opts metav1.UpdateOptions) (*v1.CloudIP, error)
	UpdateStatus(ctx context.Context, cloudIP *v1.CloudIP, opts metav1.UpdateOptions) (*v1.CloudIP, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CloudIP, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CloudIPList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudIP, err error)
	CloudIPExpansion
}

// cloudIPs implements CloudIPInterface
type cloudIPs struct {
	client rest.Interface
	ns     string
}

// newCloudIPs returns a CloudIPs
func newCloudIPs(c *CloudV1Client, namespace string) *cloudIPs {
	return &cloudIPs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudIP, and returns the corresponding cloudIP object, and an error if there is any.
func (c *cloudIPs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CloudIP, err error) {
	result = &v1.CloudIP{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudips").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudIPs that match those selectors.
func (c *cloudIPs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CloudIPList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CloudIPList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudIPs.
func (c *cloudIPs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloudIP and creates it.  Returns the server's representation of the cloudIP, and an error, if there is any.
func (c *cloudIPs) Create(ctx context.Context, cloudIP *v1.CloudIP, opts metav1.CreateOptions) (result *v1.CloudIP, err error) {
	result = &v1.CloudIP{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudips").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudIP).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloudIP and updates it. Returns the server's representation of the cloudIP, and an error, if there is any.
func (c *cloudIPs) Update(ctx context.Context, cloudIP *v1.CloudIP, opts metav1.UpdateOptions) (result *v1.CloudIP, err error) {
	result = &v1.CloudIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudips").
		Name(cloudIP.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudIP).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloudIPs) UpdateStatus(ctx context.Context, cloudIP *v1.CloudIP, opts metav1.UpdateOptions) (result *v1.CloudIP, err error) {
	result = &v1.CloudIP{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudips").
		Name(cloudIP.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloudIP).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloudIP and deletes it. Returns an error if one occurs.
func (c *cloudIPs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudips").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudIPs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudips").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloudIP.
func (c *cloudIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CloudIP, err error) {
	result = &v1.CloudIP{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudips").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
