

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFrameworks implements FrameworkInterface
type FakeFrameworks struct {
	Fake *FakeBkbcsV2
	ns   string
}

var frameworksResource = schema.GroupVersionResource{Group: "bkbcs", Version: "v2", Resource: "frameworks"}

var frameworksKind = schema.GroupVersionKind{Group: "bkbcs", Version: "v2", Kind: "Framework"}

// Get takes name of the framework, and returns the corresponding framework object, and an error if there is any.
func (c *FakeFrameworks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.Framework, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(frameworksResource, c.ns, name), &v2.Framework{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Framework), err
}

// List takes label and field selectors, and returns the list of Frameworks that match those selectors.
func (c *FakeFrameworks) List(ctx context.Context, opts v1.ListOptions) (result *v2.FrameworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(frameworksResource, frameworksKind, c.ns, opts), &v2.FrameworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.FrameworkList{ListMeta: obj.(*v2.FrameworkList).ListMeta}
	for _, item := range obj.(*v2.FrameworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested frameworks.
func (c *FakeFrameworks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(frameworksResource, c.ns, opts))

}

// Create takes the representation of a framework and creates it.  Returns the server's representation of the framework, and an error, if there is any.
func (c *FakeFrameworks) Create(ctx context.Context, framework *v2.Framework, opts v1.CreateOptions) (result *v2.Framework, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(frameworksResource, c.ns, framework), &v2.Framework{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Framework), err
}

// Update takes the representation of a framework and updates it. Returns the server's representation of the framework, and an error, if there is any.
func (c *FakeFrameworks) Update(ctx context.Context, framework *v2.Framework, opts v1.UpdateOptions) (result *v2.Framework, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(frameworksResource, c.ns, framework), &v2.Framework{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Framework), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFrameworks) UpdateStatus(ctx context.Context, framework *v2.Framework, opts v1.UpdateOptions) (*v2.Framework, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(frameworksResource, "status", c.ns, framework), &v2.Framework{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Framework), err
}

// Delete takes name of the framework and deletes it. Returns an error if one occurs.
func (c *FakeFrameworks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(frameworksResource, c.ns, name), &v2.Framework{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFrameworks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(frameworksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.FrameworkList{})
	return err
}

// Patch applies the patch and returns the patched framework.
func (c *FakeFrameworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.Framework, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(frameworksResource, c.ns, name, pt, data, subresources...), &v2.Framework{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Framework), err
}
