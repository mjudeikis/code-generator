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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoleLister helps list Roles.
// All objects returned here must be treated as read-only.
type RoleLister interface {
	// List lists all Roles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Role, err error)
	// Roles returns an object that can list and get Roles.
	Roles(namespace string) RoleNamespaceLister
	RoleListerExpansion
}

// RoleClusterLister helps list Roles.
// All objects returned here must be treated as read-only.
type RoleClusterLister interface {
	// List lists all Roles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Role, err error)
	RoleClusterListerExpansion
}

// roleLister implements the RoleLister interface.
type roleLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// roleLister implements the RoleClusterLister interface.
type roleClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get Role.
func (s *roleClusterLister) Cluster(clusterName logicalcluster.Name) RoleLister {
	return &roleLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all Roles in the indexer.
func (s *roleLister) List(selector labels.Selector) (ret []*v1beta1.Role, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.Role))
	})
	return ret, err
}

// List lists all Roles in the indexer.
func (s *roleClusterLister) List(selector labels.Selector) (ret []*v1beta1.Role, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Role))
	})
	return ret, err
}

// NewRoleClusterLister returns a new RoleClusterLister.
func NewRoleClusterLister(indexer cache.Indexer) RoleClusterLister {
	return &roleClusterLister{
		indexer: indexer,
	}
}

// Roles returns an object that can list and get Roles.
func (s *roleLister) Roles(namespace string) RoleNamespaceLister {
	return roleNamespaceLister{
		indexer:     s.indexer,
		clusterName: s.clusterName,
		namespace:   namespace,
	}
}

// RoleNamespaceLister helps list and get Roles.
// All objects returned here must be treated as read-only.
type RoleNamespaceLister interface {
	// List lists all Roles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Role, err error)
	// Get retrieves the Role from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Role, error)
	RoleNamespaceListerExpansion
}

// roleNamespaceLister implements the RoleNamespaceLister
// interface.
type roleNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the Role from the index for a given name.
func (s roleNamespaceLister) Get(name string) (*v1beta1.Role, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("role"), name)
	}
	return obj.(*v1beta1.Role), nil
}

// Get retrieves the Role from the index for a given name.
func (s roleNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Role, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.Role))
	})
	return ret, err
}
