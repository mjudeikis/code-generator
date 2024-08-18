//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	networkingv1alpha1 "acme.corp/pkg/apis/networking/v1alpha1"
)

// ServiceCIDRClusterLister can list ServiceCIDRs across all workspaces, or scope down to a ServiceCIDRLister for one workspace.
// All objects returned here must be treated as read-only.
type ServiceCIDRClusterLister interface {
	// List lists all ServiceCIDRs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.ServiceCIDR, err error)
	// Cluster returns a lister that can list and get ServiceCIDRs in one workspace.
	Cluster(clusterName logicalcluster.Name) ServiceCIDRLister
	ServiceCIDRClusterListerExpansion
}

type serviceCIDRClusterLister struct {
	indexer cache.Indexer
}

// NewServiceCIDRClusterLister returns a new ServiceCIDRClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewServiceCIDRClusterLister(indexer cache.Indexer) *serviceCIDRClusterLister {
	return &serviceCIDRClusterLister{indexer: indexer}
}

// List lists all ServiceCIDRs in the indexer across all workspaces.
func (s *serviceCIDRClusterLister) List(selector labels.Selector) (ret []*networkingv1alpha1.ServiceCIDR, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*networkingv1alpha1.ServiceCIDR))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ServiceCIDRs.
func (s *serviceCIDRClusterLister) Cluster(clusterName logicalcluster.Name) ServiceCIDRLister {
	return &serviceCIDRLister{indexer: s.indexer, clusterName: clusterName}
}

// ServiceCIDRLister can list all ServiceCIDRs, or get one in particular.
// All objects returned here must be treated as read-only.
type ServiceCIDRLister interface {
	// List lists all ServiceCIDRs in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*networkingv1alpha1.ServiceCIDR, err error)
	// Get retrieves the ServiceCIDR from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*networkingv1alpha1.ServiceCIDR, error)
	ServiceCIDRListerExpansion
}

// serviceCIDRLister can list all ServiceCIDRs inside a workspace.
type serviceCIDRLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ServiceCIDRs in the indexer for a workspace.
func (s *serviceCIDRLister) List(selector labels.Selector) (ret []*networkingv1alpha1.ServiceCIDR, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*networkingv1alpha1.ServiceCIDR))
	})
	return ret, err
}

// Get retrieves the ServiceCIDR from the indexer for a given workspace and name.
func (s *serviceCIDRLister) Get(name string) (*networkingv1alpha1.ServiceCIDR, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networkingv1alpha1.Resource("servicecidrs"), name)
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), nil
}

// NewServiceCIDRLister returns a new ServiceCIDRLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewServiceCIDRLister(indexer cache.Indexer) *serviceCIDRScopedLister {
	return &serviceCIDRScopedLister{indexer: indexer}
}

// serviceCIDRScopedLister can list all ServiceCIDRs inside a workspace.
type serviceCIDRScopedLister struct {
	indexer cache.Indexer
}

// List lists all ServiceCIDRs in the indexer for a workspace.
func (s *serviceCIDRScopedLister) List(selector labels.Selector) (ret []*networkingv1alpha1.ServiceCIDR, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*networkingv1alpha1.ServiceCIDR))
	})
	return ret, err
}

// Get retrieves the ServiceCIDR from the indexer for a given workspace and name.
func (s *serviceCIDRScopedLister) Get(name string) (*networkingv1alpha1.ServiceCIDR, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(networkingv1alpha1.Resource("servicecidrs"), name)
	}
	return obj.(*networkingv1alpha1.ServiceCIDR), nil
}
