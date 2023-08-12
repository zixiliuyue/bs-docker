

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

// BcsClusterAgentSettingsGetter has a method to return a BcsClusterAgentSettingInterface.
// A group's client should implement this interface.
type BcsClusterAgentSettingsGetter interface {
	BcsClusterAgentSettings(namespace string) BcsClusterAgentSettingInterface
}

// BcsClusterAgentSettingInterface has methods to work with BcsClusterAgentSetting resources.
type BcsClusterAgentSettingInterface interface {
	Create(ctx context.Context, bcsClusterAgentSetting *v2.BcsClusterAgentSetting, opts v1.CreateOptions) (*v2.BcsClusterAgentSetting, error)
	Update(ctx context.Context, bcsClusterAgentSetting *v2.BcsClusterAgentSetting, opts v1.UpdateOptions) (*v2.BcsClusterAgentSetting, error)
	UpdateStatus(ctx context.Context, bcsClusterAgentSetting *v2.BcsClusterAgentSetting, opts v1.UpdateOptions) (*v2.BcsClusterAgentSetting, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.BcsClusterAgentSetting, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.BcsClusterAgentSettingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsClusterAgentSetting, err error)
	BcsClusterAgentSettingExpansion
}

// bcsClusterAgentSettings implements BcsClusterAgentSettingInterface
type bcsClusterAgentSettings struct {
	client rest.Interface
	ns     string
}

// newBcsClusterAgentSettings returns a BcsClusterAgentSettings
func newBcsClusterAgentSettings(c *BkbcsV2Client, namespace string) *bcsClusterAgentSettings {
	return &bcsClusterAgentSettings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bcsClusterAgentSetting, and returns the corresponding bcsClusterAgentSetting object, and an error if there is any.
func (c *bcsClusterAgentSettings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.BcsClusterAgentSetting, err error) {
	result = &v2.BcsClusterAgentSetting{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BcsClusterAgentSettings that match those selectors.
func (c *bcsClusterAgentSettings) List(ctx context.Context, opts v1.ListOptions) (result *v2.BcsClusterAgentSettingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.BcsClusterAgentSettingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bcsClusterAgentSettings.
func (c *bcsClusterAgentSettings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bcsClusterAgentSetting and creates it.  Returns the server's representation of the bcsClusterAgentSetting, and an error, if there is any.
func (c *bcsClusterAgentSettings) Create(ctx context.Context, bcsClusterAgentSetting *v2.BcsClusterAgentSetting, opts v1.CreateOptions) (result *v2.BcsClusterAgentSetting, err error) {
	result = &v2.BcsClusterAgentSetting{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsClusterAgentSetting).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bcsClusterAgentSetting and updates it. Returns the server's representation of the bcsClusterAgentSetting, and an error, if there is any.
func (c *bcsClusterAgentSettings) Update(ctx context.Context, bcsClusterAgentSetting *v2.BcsClusterAgentSetting, opts v1.UpdateOptions) (result *v2.BcsClusterAgentSetting, err error) {
	result = &v2.BcsClusterAgentSetting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		Name(bcsClusterAgentSetting.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsClusterAgentSetting).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bcsClusterAgentSettings) UpdateStatus(ctx context.Context, bcsClusterAgentSetting *v2.BcsClusterAgentSetting, opts v1.UpdateOptions) (result *v2.BcsClusterAgentSetting, err error) {
	result = &v2.BcsClusterAgentSetting{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		Name(bcsClusterAgentSetting.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bcsClusterAgentSetting).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bcsClusterAgentSetting and deletes it. Returns an error if one occurs.
func (c *bcsClusterAgentSettings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bcsClusterAgentSettings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bcsClusterAgentSetting.
func (c *bcsClusterAgentSettings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsClusterAgentSetting, err error) {
	result = &v2.BcsClusterAgentSetting{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bcsclusteragentsettings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
