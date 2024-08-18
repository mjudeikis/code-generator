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

	rbacv1alpha1 "acme.corp/pkg/apis/rbac/v1alpha1"
)

// RoleClusterLister can list Roles across all workspaces, or scope down to a RoleLister for one workspace.
// All objects returned here must be treated as read-only.
type RoleClusterLister interface {
	// List lists all Roles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error)
	// Cluster returns a lister that can list and get Roles in one workspace.
	Cluster(clusterName logicalcluster.Name) RoleLister
	RoleClusterListerExpansion
}

type roleClusterLister struct {
	indexer cache.Indexer
}

// NewRoleClusterLister returns a new RoleClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewRoleClusterLister(indexer cache.Indexer) *roleClusterLister {
	return &roleClusterLister{indexer: indexer}
}

// List lists all Roles in the indexer across all workspaces.
func (s *roleClusterLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*rbacv1alpha1.Role))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Roles.
func (s *roleClusterLister) Cluster(clusterName logicalcluster.Name) RoleLister {
	return &roleLister{indexer: s.indexer, clusterName: clusterName}
}

// RoleLister can list Roles across all namespaces, or scope down to a RoleNamespaceLister for one namespace.
// All objects returned here must be treated as read-only.
type RoleLister interface {
	// List lists all Roles in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error)
	// Roles returns a lister that can list and get Roles in one workspace and namespace.
	Roles(namespace string) RoleNamespaceLister
	RoleListerExpansion
}

// roleLister can list all Roles inside a workspace or scope down to a RoleLister for one namespace.
type roleLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Roles in the indexer for a workspace.
func (s *roleLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1alpha1.Role))
	})
	return ret, err
}

// Roles returns an object that can list and get Roles in one namespace.
func (s *roleLister) Roles(namespace string) RoleNamespaceLister {
	return &roleNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// roleNamespaceLister helps list and get Roles.
// All objects returned here must be treated as read-only.
type RoleNamespaceLister interface {
	// List lists all Roles in the workspace and namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error)
	// Get retrieves the Role from the indexer for a given workspace, namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*rbacv1alpha1.Role, error)
	RoleNamespaceListerExpansion
}

// roleNamespaceLister helps list and get Roles.
// All objects returned here must be treated as read-only.
type roleNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all Roles in the indexer for a given workspace and namespace.
func (s *roleNamespaceLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1alpha1.Role))
	})
	return ret, err
}

// Get retrieves the Role from the indexer for a given workspace, namespace and name.
func (s *roleNamespaceLister) Get(name string) (*rbacv1alpha1.Role, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbacv1alpha1.Resource("roles"), name)
	}
	return obj.(*rbacv1alpha1.Role), nil
}

// NewRoleLister returns a new RoleLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
// - has the cache.NamespaceIndex as an index
func NewRoleLister(indexer cache.Indexer) *roleScopedLister {
	return &roleScopedLister{indexer: indexer}
}

// roleScopedLister can list all Roles inside a workspace or scope down to a RoleLister for one namespace.
type roleScopedLister struct {
	indexer cache.Indexer
}

// List lists all Roles in the indexer for a workspace.
func (s *roleScopedLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1alpha1.Role))
	})
	return ret, err
}

// Roles returns an object that can list and get Roles in one namespace.
func (s *roleScopedLister) Roles(namespace string) RoleNamespaceLister {
	return &roleScopedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// roleScopedNamespaceLister helps list and get Roles.
type roleScopedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Roles in the indexer for a given workspace and namespace.
func (s *roleScopedNamespaceLister) List(selector labels.Selector) (ret []*rbacv1alpha1.Role, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*rbacv1alpha1.Role))
	})
	return ret, err
}

// Get retrieves the Role from the indexer for a given workspace, namespace and name.
func (s *roleScopedNamespaceLister) Get(name string) (*rbacv1alpha1.Role, error) {
	key := s.namespace + "/" + name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(rbacv1alpha1.Resource("roles"), name)
	}
	return obj.(*rbacv1alpha1.Role), nil
}
