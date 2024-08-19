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

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1 "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LeaseLister helps list Leases.
// All objects returned here must be treated as read-only.
type LeaseLister interface {
	// List lists all Leases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	// Leases returns an object that can list and get Leases.
	Leases(namespace string) LeaseNamespaceLister
	LeaseListerExpansion
}

// LeaseClusterLister helps list Leases.
// All objects returned here must be treated as read-only.
type LeaseClusterLister interface {
	// List lists all Leases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	LeaseClusterListerExpansion
}

// leaseLister implements the LeaseLister interface.
type leaseLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// leaseLister implements the LeaseClusterLister interface.
type leaseClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get Lease.
func (s *leaseClusterLister) Cluster(clusterName logicalcluster.Name) LeaseLister {
	return &leaseLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all Leases in the indexer.
func (s *leaseLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.Lease))
	})
	return ret, err
}

// List lists all Leases in the indexer.
func (s *leaseClusterLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Lease))
	})
	return ret, err
}

// NewLeaseClusterLister returns a new LeaseClusterLister.
func NewLeaseClusterLister(indexer cache.Indexer) LeaseClusterLister {
	return &leaseClusterLister{
		indexer: indexer,
	}
}

// Leases returns an object that can list and get Leases.
func (s *leaseLister) Leases(namespace string) LeaseNamespaceLister {
	return leaseNamespaceLister{
		indexer:     s.indexer,
		clusterName: s.clusterName,
		namespace:   namespace,
	}
}

// LeaseNamespaceLister helps list and get Leases.
// All objects returned here must be treated as read-only.
type LeaseNamespaceLister interface {
	// List lists all Leases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	// Get retrieves the Lease from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Lease, error)
	LeaseNamespaceListerExpansion
}

// leaseNamespaceLister implements the LeaseNamespaceLister
// interface.
type leaseNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the Lease from the index for a given name.
func (s leaseNamespaceLister) Get(name string) (*v1.Lease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("lease"), name)
	}
	return obj.(*v1.Lease), nil
}

// Get retrieves the Lease from the index for a given name.
func (s leaseNamespaceLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.Lease))
	})
	return ret, err
}
