

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VersionLister helps list Versions.
type VersionLister interface {
	// List lists all Versions in the indexer.
	List(selector labels.Selector) (ret []*v2.Version, err error)
	// Versions returns an object that can list and get Versions.
	Versions(namespace string) VersionNamespaceLister
	VersionListerExpansion
}

// versionLister implements the VersionLister interface.
type versionLister struct {
	indexer cache.Indexer
}

// NewVersionLister returns a new VersionLister.
func NewVersionLister(indexer cache.Indexer) VersionLister {
	return &versionLister{indexer: indexer}
}

// List lists all Versions in the indexer.
func (s *versionLister) List(selector labels.Selector) (ret []*v2.Version, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.Version))
	})
	return ret, err
}

// Versions returns an object that can list and get Versions.
func (s *versionLister) Versions(namespace string) VersionNamespaceLister {
	return versionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VersionNamespaceLister helps list and get Versions.
type VersionNamespaceLister interface {
	// List lists all Versions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.Version, err error)
	// Get retrieves the Version from the indexer for a given namespace and name.
	Get(name string) (*v2.Version, error)
	VersionNamespaceListerExpansion
}

// versionNamespaceLister implements the VersionNamespaceLister
// interface.
type versionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Versions in the indexer for a given namespace.
func (s versionNamespaceLister) List(selector labels.Selector) (ret []*v2.Version, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.Version))
	})
	return ret, err
}

// Get retrieves the Version from the indexer for a given namespace and name.
func (s versionNamespaceLister) Get(name string) (*v2.Version, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("version"), name)
	}
	return obj.(*v2.Version), nil
}
