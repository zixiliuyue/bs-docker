
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	clbv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/apis/clb/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClbIngresses implements ClbIngressInterface
type FakeClbIngresses struct {
	Fake *FakeClbV1
	ns   string
}

var clbingressesResource = schema.GroupVersionResource{Group: "clb", Version: "v1", Resource: "clbingresses"}

var clbingressesKind = schema.GroupVersionKind{Group: "clb", Version: "v1", Kind: "ClbIngress"}

// Get takes name of the clbIngress, and returns the corresponding clbIngress object, and an error if there is any.
func (c *FakeClbIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *clbv1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clbingressesResource, c.ns, name), &clbv1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clbv1.ClbIngress), err
}

// List takes label and field selectors, and returns the list of ClbIngresses that match those selectors.
func (c *FakeClbIngresses) List(ctx context.Context, opts v1.ListOptions) (result *clbv1.ClbIngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clbingressesResource, clbingressesKind, c.ns, opts), &clbv1.ClbIngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clbv1.ClbIngressList{ListMeta: obj.(*clbv1.ClbIngressList).ListMeta}
	for _, item := range obj.(*clbv1.ClbIngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clbIngresses.
func (c *FakeClbIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clbingressesResource, c.ns, opts))

}

// Create takes the representation of a clbIngress and creates it.  Returns the server's representation of the clbIngress, and an error, if there is any.
func (c *FakeClbIngresses) Create(ctx context.Context, clbIngress *clbv1.ClbIngress, opts v1.CreateOptions) (result *clbv1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clbingressesResource, c.ns, clbIngress), &clbv1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clbv1.ClbIngress), err
}

// Update takes the representation of a clbIngress and updates it. Returns the server's representation of the clbIngress, and an error, if there is any.
func (c *FakeClbIngresses) Update(ctx context.Context, clbIngress *clbv1.ClbIngress, opts v1.UpdateOptions) (result *clbv1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clbingressesResource, c.ns, clbIngress), &clbv1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clbv1.ClbIngress), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClbIngresses) UpdateStatus(ctx context.Context, clbIngress *clbv1.ClbIngress, opts v1.UpdateOptions) (*clbv1.ClbIngress, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clbingressesResource, "status", c.ns, clbIngress), &clbv1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clbv1.ClbIngress), err
}

// Delete takes name of the clbIngress and deletes it. Returns an error if one occurs.
func (c *FakeClbIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clbingressesResource, c.ns, name), &clbv1.ClbIngress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClbIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clbingressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &clbv1.ClbIngressList{})
	return err
}

// Patch applies the patch and returns the patched clbIngress.
func (c *FakeClbIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clbv1.ClbIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clbingressesResource, c.ns, name, pt, data, subresources...), &clbv1.ClbIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clbv1.ClbIngress), err
}
