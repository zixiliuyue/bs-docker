
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/apis/tkex/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HookRunLister helps list HookRuns.
type HookRunLister interface {
	// List lists all HookRuns in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HookRun, err error)
	// HookRuns returns an object that can list and get HookRuns.
	HookRuns(namespace string) HookRunNamespaceLister
	HookRunListerExpansion
}

// hookRunLister implements the HookRunLister interface.
type hookRunLister struct {
	indexer cache.Indexer
}

// NewHookRunLister returns a new HookRunLister.
func NewHookRunLister(indexer cache.Indexer) HookRunLister {
	return &hookRunLister{indexer: indexer}
}

// List lists all HookRuns in the indexer.
func (s *hookRunLister) List(selector labels.Selector) (ret []*v1alpha1.HookRun, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HookRun))
	})
	return ret, err
}

// HookRuns returns an object that can list and get HookRuns.
func (s *hookRunLister) HookRuns(namespace string) HookRunNamespaceLister {
	return hookRunNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HookRunNamespaceLister helps list and get HookRuns.
type HookRunNamespaceLister interface {
	// List lists all HookRuns in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.HookRun, err error)
	// Get retrieves the HookRun from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.HookRun, error)
	HookRunNamespaceListerExpansion
}

// hookRunNamespaceLister implements the HookRunNamespaceLister
// interface.
type hookRunNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HookRuns in the indexer for a given namespace.
func (s hookRunNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HookRun, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HookRun))
	})
	return ret, err
}

// Get retrieves the HookRun from the indexer for a given namespace and name.
func (s hookRunNamespaceLister) Get(name string) (*v1alpha1.HookRun, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hookrun"), name)
	}
	return obj.(*v1alpha1.HookRun), nil
}
