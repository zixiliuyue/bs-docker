

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BcsEndpointLister helps list BcsEndpoints.
type BcsEndpointLister interface {
	// List lists all BcsEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v2.BcsEndpoint, err error)
	// BcsEndpoints returns an object that can list and get BcsEndpoints.
	BcsEndpoints(namespace string) BcsEndpointNamespaceLister
	BcsEndpointListerExpansion
}

// bcsEndpointLister implements the BcsEndpointLister interface.
type bcsEndpointLister struct {
	indexer cache.Indexer
}

// NewBcsEndpointLister returns a new BcsEndpointLister.
func NewBcsEndpointLister(indexer cache.Indexer) BcsEndpointLister {
	return &bcsEndpointLister{indexer: indexer}
}

// List lists all BcsEndpoints in the indexer.
func (s *bcsEndpointLister) List(selector labels.Selector) (ret []*v2.BcsEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsEndpoint))
	})
	return ret, err
}

// BcsEndpoints returns an object that can list and get BcsEndpoints.
func (s *bcsEndpointLister) BcsEndpoints(namespace string) BcsEndpointNamespaceLister {
	return bcsEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BcsEndpointNamespaceLister helps list and get BcsEndpoints.
type BcsEndpointNamespaceLister interface {
	// List lists all BcsEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.BcsEndpoint, err error)
	// Get retrieves the BcsEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v2.BcsEndpoint, error)
	BcsEndpointNamespaceListerExpansion
}

// bcsEndpointNamespaceLister implements the BcsEndpointNamespaceLister
// interface.
type bcsEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BcsEndpoints in the indexer for a given namespace.
func (s bcsEndpointNamespaceLister) List(selector labels.Selector) (ret []*v2.BcsEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsEndpoint))
	})
	return ret, err
}

// Get retrieves the BcsEndpoint from the indexer for a given namespace and name.
func (s bcsEndpointNamespaceLister) Get(name string) (*v2.BcsEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("bcsendpoint"), name)
	}
	return obj.(*v2.BcsEndpoint), nil
}
