/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha3 "k8s.io/api/resource/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceClaimLister helps list ResourceClaims.
// All objects returned here must be treated as read-only.
type ResourceClaimLister interface {
	// List lists all ResourceClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ResourceClaim, err error)
	// ResourceClaims returns an object that can list and get ResourceClaims.
	ResourceClaims(namespace string) ResourceClaimNamespaceLister
	ResourceClaimListerExpansion
}

// ResourceClaimClusterLister helps list ResourceClaims.
// All objects returned here must be treated as read-only.
type ResourceClaimClusterLister interface {
	// List lists all ResourceClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ResourceClaim, err error)
	ResourceClaimClusterListerExpansion
}

// resourceClaimLister implements the ResourceClaimLister interface.
type resourceClaimLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// resourceClaimLister implements the ResourceClaimClusterLister interface.
type resourceClaimClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get ResourceClaim.
func (s *resourceClaimClusterLister) Cluster(clusterName logicalcluster.Name) ResourceClaimLister {
	return &resourceClaimLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all ResourceClaims in the indexer.
func (s *resourceClaimLister) List(selector labels.Selector) (ret []*v1alpha3.ResourceClaim, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1alpha3.ResourceClaim))
	})
	return ret, err
}

// List lists all ResourceClaims in the indexer.
func (s *resourceClaimClusterLister) List(selector labels.Selector) (ret []*v1alpha3.ResourceClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ResourceClaim))
	})
	return ret, err
}

// NewResourceClaimClusterLister returns a new ResourceClaimClusterLister.
func NewResourceClaimClusterLister(indexer cache.Indexer) ResourceClaimClusterLister {
	return &resourceClaimClusterLister{
		indexer: indexer,
	}
}

// ResourceClaims returns an object that can list and get ResourceClaims.
func (s *resourceClaimLister) ResourceClaims(namespace string) ResourceClaimNamespaceLister {
	return resourceClaimNamespaceLister{
		indexer:     s.indexer,
		clusterName: s.clusterName,
		namespace:   namespace,
	}
}

// ResourceClaimNamespaceLister helps list and get ResourceClaims.
// All objects returned here must be treated as read-only.
type ResourceClaimNamespaceLister interface {
	// List lists all ResourceClaims in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ResourceClaim, err error)
	// Get retrieves the ResourceClaim from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha3.ResourceClaim, error)
	ResourceClaimNamespaceListerExpansion
}

// resourceClaimNamespaceLister implements the ResourceClaimNamespaceLister
// interface.
type resourceClaimNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the ResourceClaim from the index for a given name.
func (s resourceClaimNamespaceLister) Get(name string) (*v1alpha3.ResourceClaim, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("resourceclaim"), name)
	}
	return obj.(*v1alpha3.ResourceClaim), nil
}

// Get retrieves the ResourceClaim from the index for a given name.
func (s resourceClaimNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.ResourceClaim, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1alpha3.ResourceClaim))
	})
	return ret, err
}
