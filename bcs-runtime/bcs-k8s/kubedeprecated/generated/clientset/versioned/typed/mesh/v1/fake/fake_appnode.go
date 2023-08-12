
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	meshv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubedeprecated/apis/mesh/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppNodes implements AppNodeInterface
type FakeAppNodes struct {
	Fake *FakeMeshV1
	ns   string
}

var appnodesResource = schema.GroupVersionResource{Group: "mesh", Version: "v1", Resource: "appnodes"}

var appnodesKind = schema.GroupVersionKind{Group: "mesh", Version: "v1", Kind: "AppNode"}

// Get takes name of the appNode, and returns the corresponding appNode object, and an error if there is any.
func (c *FakeAppNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *meshv1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appnodesResource, c.ns, name), &meshv1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppNode), err
}

// List takes label and field selectors, and returns the list of AppNodes that match those selectors.
func (c *FakeAppNodes) List(ctx context.Context, opts v1.ListOptions) (result *meshv1.AppNodeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appnodesResource, appnodesKind, c.ns, opts), &meshv1.AppNodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &meshv1.AppNodeList{ListMeta: obj.(*meshv1.AppNodeList).ListMeta}
	for _, item := range obj.(*meshv1.AppNodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appNodes.
func (c *FakeAppNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appnodesResource, c.ns, opts))

}

// Create takes the representation of a appNode and creates it.  Returns the server's representation of the appNode, and an error, if there is any.
func (c *FakeAppNodes) Create(ctx context.Context, appNode *meshv1.AppNode, opts v1.CreateOptions) (result *meshv1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appnodesResource, c.ns, appNode), &meshv1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppNode), err
}

// Update takes the representation of a appNode and updates it. Returns the server's representation of the appNode, and an error, if there is any.
func (c *FakeAppNodes) Update(ctx context.Context, appNode *meshv1.AppNode, opts v1.UpdateOptions) (result *meshv1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appnodesResource, c.ns, appNode), &meshv1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppNode), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppNodes) UpdateStatus(ctx context.Context, appNode *meshv1.AppNode, opts v1.UpdateOptions) (*meshv1.AppNode, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appnodesResource, "status", c.ns, appNode), &meshv1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppNode), err
}

// Delete takes name of the appNode and deletes it. Returns an error if one occurs.
func (c *FakeAppNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appnodesResource, c.ns, name), &meshv1.AppNode{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appnodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &meshv1.AppNodeList{})
	return err
}

// Patch applies the patch and returns the patched appNode.
func (c *FakeAppNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *meshv1.AppNode, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appnodesResource, c.ns, name, pt, data, subresources...), &meshv1.AppNode{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppNode), err
}
