
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePortBindings implements PortBindingInterface
type FakePortBindings struct {
	Fake *FakeNetworkextensionV1
	ns   string
}

var portbindingsResource = schema.GroupVersionResource{Group: "networkextension", Version: "v1", Resource: "portbindings"}

var portbindingsKind = schema.GroupVersionKind{Group: "networkextension", Version: "v1", Kind: "PortBinding"}

// Get takes name of the portBinding, and returns the corresponding portBinding object, and an error if there is any.
func (c *FakePortBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkextensionv1.PortBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(portbindingsResource, c.ns, name), &networkextensionv1.PortBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkextensionv1.PortBinding), err
}

// List takes label and field selectors, and returns the list of PortBindings that match those selectors.
func (c *FakePortBindings) List(ctx context.Context, opts v1.ListOptions) (result *networkextensionv1.PortBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(portbindingsResource, portbindingsKind, c.ns, opts), &networkextensionv1.PortBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkextensionv1.PortBindingList{ListMeta: obj.(*networkextensionv1.PortBindingList).ListMeta}
	for _, item := range obj.(*networkextensionv1.PortBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested portBindings.
func (c *FakePortBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(portbindingsResource, c.ns, opts))

}

// Create takes the representation of a portBinding and creates it.  Returns the server's representation of the portBinding, and an error, if there is any.
func (c *FakePortBindings) Create(ctx context.Context, portBinding *networkextensionv1.PortBinding, opts v1.CreateOptions) (result *networkextensionv1.PortBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(portbindingsResource, c.ns, portBinding), &networkextensionv1.PortBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkextensionv1.PortBinding), err
}

// Update takes the representation of a portBinding and updates it. Returns the server's representation of the portBinding, and an error, if there is any.
func (c *FakePortBindings) Update(ctx context.Context, portBinding *networkextensionv1.PortBinding, opts v1.UpdateOptions) (result *networkextensionv1.PortBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(portbindingsResource, c.ns, portBinding), &networkextensionv1.PortBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkextensionv1.PortBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePortBindings) UpdateStatus(ctx context.Context, portBinding *networkextensionv1.PortBinding, opts v1.UpdateOptions) (*networkextensionv1.PortBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(portbindingsResource, "status", c.ns, portBinding), &networkextensionv1.PortBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkextensionv1.PortBinding), err
}

// Delete takes name of the portBinding and deletes it. Returns an error if one occurs.
func (c *FakePortBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(portbindingsResource, c.ns, name), &networkextensionv1.PortBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePortBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(portbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networkextensionv1.PortBindingList{})
	return err
}

// Patch applies the patch and returns the patched portBinding.
func (c *FakePortBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkextensionv1.PortBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(portbindingsResource, c.ns, name, pt, data, subresources...), &networkextensionv1.PortBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkextensionv1.PortBinding), err
}
