

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/Tencent/bk-bcs/bcs-scenarios/kourse/pkg/apis/tkex/v1alpha1"
	scheme "github.com/Tencent/bk-bcs/bcs-scenarios/kourse/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GameDeploymentsGetter has a method to return a GameDeploymentInterface.
// A group's client should implement this interface.
type GameDeploymentsGetter interface {
	GameDeployments(namespace string) GameDeploymentInterface
}

// GameDeploymentInterface has methods to work with GameDeployment resources.
type GameDeploymentInterface interface {
	Create(ctx context.Context, gameDeployment *v1alpha1.GameDeployment, opts v1.CreateOptions) (*v1alpha1.GameDeployment, error)
	Update(ctx context.Context, gameDeployment *v1alpha1.GameDeployment, opts v1.UpdateOptions) (*v1alpha1.GameDeployment, error)
	UpdateStatus(ctx context.Context, gameDeployment *v1alpha1.GameDeployment, opts v1.UpdateOptions) (*v1alpha1.GameDeployment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GameDeployment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GameDeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GameDeployment, err error)
	GameDeploymentExpansion
}

// gameDeployments implements GameDeploymentInterface
type gameDeployments struct {
	client rest.Interface
	ns     string
}

// newGameDeployments returns a GameDeployments
func newGameDeployments(c *TkexV1alpha1Client, namespace string) *gameDeployments {
	return &gameDeployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gameDeployment, and returns the corresponding gameDeployment object, and an error if there is any.
func (c *gameDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GameDeployment, err error) {
	result = &v1alpha1.GameDeployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gamedeployments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GameDeployments that match those selectors.
func (c *gameDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GameDeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GameDeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gamedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gameDeployments.
func (c *gameDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gamedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gameDeployment and creates it.  Returns the server's representation of the gameDeployment, and an error, if there is any.
func (c *gameDeployments) Create(ctx context.Context, gameDeployment *v1alpha1.GameDeployment, opts v1.CreateOptions) (result *v1alpha1.GameDeployment, err error) {
	result = &v1alpha1.GameDeployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gamedeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameDeployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gameDeployment and updates it. Returns the server's representation of the gameDeployment, and an error, if there is any.
func (c *gameDeployments) Update(ctx context.Context, gameDeployment *v1alpha1.GameDeployment, opts v1.UpdateOptions) (result *v1alpha1.GameDeployment, err error) {
	result = &v1alpha1.GameDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gamedeployments").
		Name(gameDeployment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameDeployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gameDeployments) UpdateStatus(ctx context.Context, gameDeployment *v1alpha1.GameDeployment, opts v1.UpdateOptions) (result *v1alpha1.GameDeployment, err error) {
	result = &v1alpha1.GameDeployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gamedeployments").
		Name(gameDeployment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gameDeployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gameDeployment and deletes it. Returns an error if one occurs.
func (c *gameDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gamedeployments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gameDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gamedeployments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gameDeployment.
func (c *gameDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GameDeployment, err error) {
	result = &v1alpha1.GameDeployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gamedeployments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
