// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apiserver/apis/datapackaging/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// packageLister helps list Packages.
// All objects returned here must be treated as read-only.
type PackageLister interface {
	// List lists all Packages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Package, err error)
	// Packages returns an object that can list and get Packages.
	Packages(namespace string) PackageNamespaceLister
	PackageListerExpansion
}

// packageLister implements the packageLister interface.
type packageLister struct {
	indexer cache.Indexer
}

// NewPackageLister returns a new packageLister.
func NewPackageLister(indexer cache.Indexer) PackageLister {
	return &packageLister{indexer: indexer}
}

// List lists all Packages in the indexer.
func (s *packageLister) List(selector labels.Selector) (ret []*v1alpha1.Package, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Package))
	})
	return ret, err
}

// Packages returns an object that can list and get Packages.
func (s *packageLister) Packages(namespace string) PackageNamespaceLister {
	return packageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// packageNamespaceLister helps list and get Packages.
// All objects returned here must be treated as read-only.
type PackageNamespaceLister interface {
	// List lists all Packages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Package, err error)
	// Get retrieves the Package from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Package, error)
	PackageNamespaceListerExpansion
}

// packageNamespaceLister implements the packageNamespaceLister
// interface.
type packageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Packages in the indexer for a given namespace.
func (s packageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Package, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Package))
	})
	return ret, err
}

// Get retrieves the Package from the indexer for a given namespace and name.
func (s packageNamespaceLister) Get(name string) (*v1alpha1.Package, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("Package"), name)
	}
	return obj.(*v1alpha1.Package), nil
}
