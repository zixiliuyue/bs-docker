
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PortBindingLister helps list PortBindings.
type PortBindingLister interface {
	// List lists all PortBindings in the indexer.
	List(selector labels.Selector) (ret []*v1.PortBinding, err error)
	// PortBindings returns an object that can list and get PortBindings.
	PortBindings(namespace string) PortBindingNamespaceLister
	PortBindingListerExpansion
}

// portBindingLister implements the PortBindingLister interface.
type portBindingLister struct {
	indexer cache.Indexer
}

// NewPortBindingLister returns a new PortBindingLister.
func NewPortBindingLister(indexer cache.Indexer) PortBindingLister {
	return &portBindingLister{indexer: indexer}
}

// List lists all PortBindings in the indexer.
func (s *portBindingLister) List(selector labels.Selector) (ret []*v1.PortBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PortBinding))
	})
	return ret, err
}

// PortBindings returns an object that can list and get PortBindings.
func (s *portBindingLister) PortBindings(namespace string) PortBindingNamespaceLister {
	return portBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PortBindingNamespaceLister helps list and get PortBindings.
type PortBindingNamespaceLister interface {
	// List lists all PortBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.PortBinding, err error)
	// Get retrieves the PortBinding from the indexer for a given namespace and name.
	Get(name string) (*v1.PortBinding, error)
	PortBindingNamespaceListerExpansion
}

// portBindingNamespaceLister implements the PortBindingNamespaceLister
// interface.
type portBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PortBindings in the indexer for a given namespace.
func (s portBindingNamespaceLister) List(selector labels.Selector) (ret []*v1.PortBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PortBinding))
	})
	return ret, err
}

// Get retrieves the PortBinding from the indexer for a given namespace and name.
func (s portBindingNamespaceLister) Get(name string) (*v1.PortBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("portbinding"), name)
	}
	return obj.(*v1.PortBinding), nil
}
