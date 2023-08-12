
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/common/bcs-hook/apis/tkex/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHookTemplates implements HookTemplateInterface
type FakeHookTemplates struct {
	Fake *FakeTkexV1alpha1
	ns   string
}

var hooktemplatesResource = schema.GroupVersionResource{Group: "tkex.tencent.com", Version: "v1alpha1", Resource: "hooktemplates"}

var hooktemplatesKind = schema.GroupVersionKind{Group: "tkex.tencent.com", Version: "v1alpha1", Kind: "HookTemplate"}

// Get takes name of the hookTemplate, and returns the corresponding hookTemplate object, and an error if there is any.
func (c *FakeHookTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HookTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hooktemplatesResource, c.ns, name), &v1alpha1.HookTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookTemplate), err
}

// List takes label and field selectors, and returns the list of HookTemplates that match those selectors.
func (c *FakeHookTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HookTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hooktemplatesResource, hooktemplatesKind, c.ns, opts), &v1alpha1.HookTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HookTemplateList{ListMeta: obj.(*v1alpha1.HookTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.HookTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hookTemplates.
func (c *FakeHookTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hooktemplatesResource, c.ns, opts))

}

// Create takes the representation of a hookTemplate and creates it.  Returns the server's representation of the hookTemplate, and an error, if there is any.
func (c *FakeHookTemplates) Create(ctx context.Context, hookTemplate *v1alpha1.HookTemplate, opts v1.CreateOptions) (result *v1alpha1.HookTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hooktemplatesResource, c.ns, hookTemplate), &v1alpha1.HookTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookTemplate), err
}

// Update takes the representation of a hookTemplate and updates it. Returns the server's representation of the hookTemplate, and an error, if there is any.
func (c *FakeHookTemplates) Update(ctx context.Context, hookTemplate *v1alpha1.HookTemplate, opts v1.UpdateOptions) (result *v1alpha1.HookTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hooktemplatesResource, c.ns, hookTemplate), &v1alpha1.HookTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookTemplate), err
}

// Delete takes name of the hookTemplate and deletes it. Returns an error if one occurs.
func (c *FakeHookTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hooktemplatesResource, c.ns, name), &v1alpha1.HookTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHookTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hooktemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HookTemplateList{})
	return err
}

// Patch applies the patch and returns the patched hookTemplate.
func (c *FakeHookTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HookTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hooktemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HookTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HookTemplate), err
}
