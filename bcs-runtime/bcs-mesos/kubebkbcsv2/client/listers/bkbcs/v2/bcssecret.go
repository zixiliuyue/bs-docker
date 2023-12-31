

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BcsSecretLister helps list BcsSecrets.
type BcsSecretLister interface {
	// List lists all BcsSecrets in the indexer.
	List(selector labels.Selector) (ret []*v2.BcsSecret, err error)
	// BcsSecrets returns an object that can list and get BcsSecrets.
	BcsSecrets(namespace string) BcsSecretNamespaceLister
	BcsSecretListerExpansion
}

// bcsSecretLister implements the BcsSecretLister interface.
type bcsSecretLister struct {
	indexer cache.Indexer
}

// NewBcsSecretLister returns a new BcsSecretLister.
func NewBcsSecretLister(indexer cache.Indexer) BcsSecretLister {
	return &bcsSecretLister{indexer: indexer}
}

// List lists all BcsSecrets in the indexer.
func (s *bcsSecretLister) List(selector labels.Selector) (ret []*v2.BcsSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsSecret))
	})
	return ret, err
}

// BcsSecrets returns an object that can list and get BcsSecrets.
func (s *bcsSecretLister) BcsSecrets(namespace string) BcsSecretNamespaceLister {
	return bcsSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BcsSecretNamespaceLister helps list and get BcsSecrets.
type BcsSecretNamespaceLister interface {
	// List lists all BcsSecrets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.BcsSecret, err error)
	// Get retrieves the BcsSecret from the indexer for a given namespace and name.
	Get(name string) (*v2.BcsSecret, error)
	BcsSecretNamespaceListerExpansion
}

// bcsSecretNamespaceLister implements the BcsSecretNamespaceLister
// interface.
type bcsSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BcsSecrets in the indexer for a given namespace.
func (s bcsSecretNamespaceLister) List(selector labels.Selector) (ret []*v2.BcsSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.BcsSecret))
	})
	return ret, err
}

// Get retrieves the BcsSecret from the indexer for a given namespace and name.
func (s bcsSecretNamespaceLister) Get(name string) (*v2.BcsSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("bcssecret"), name)
	}
	return obj.(*v2.BcsSecret), nil
}
