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
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ControllerRevisionLister helps list ControllerRevisions.
// All objects returned here must be treated as read-only.
type ControllerRevisionLister interface {
	// List lists all ControllerRevisions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ControllerRevision, err error)
	// ControllerRevisions returns an object that can list and get ControllerRevisions.
	ControllerRevisions(namespace string) ControllerRevisionNamespaceLister
	ControllerRevisionListerExpansion
}

// ControllerRevisionClusterLister helps list ControllerRevisions.
// All objects returned here must be treated as read-only.
type ControllerRevisionClusterLister interface {
	// List lists all ControllerRevisions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ControllerRevision, err error)
	ControllerRevisionClusterListerExpansion
}

// controllerRevisionLister implements the ControllerRevisionLister interface.
type controllerRevisionLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// controllerRevisionLister implements the ControllerRevisionClusterLister interface.
type controllerRevisionClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get ControllerRevision.
func (s *controllerRevisionClusterLister) Cluster(clusterName logicalcluster.Name) ControllerRevisionLister {
	return &controllerRevisionLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all ControllerRevisions in the indexer.
func (s *controllerRevisionLister) List(selector labels.Selector) (ret []*v1.ControllerRevision, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ControllerRevision))
	})
	return ret, err
}

// List lists all ControllerRevisions in the indexer.
func (s *controllerRevisionClusterLister) List(selector labels.Selector) (ret []*v1.ControllerRevision, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ControllerRevision))
	})
	return ret, err
}

// NewControllerRevisionClusterLister returns a new ControllerRevisionClusterLister.
func NewControllerRevisionClusterLister(indexer cache.Indexer) ControllerRevisionClusterLister {
	return &controllerRevisionClusterLister{
		indexer: indexer,
	}
}

// ControllerRevisions returns an object that can list and get ControllerRevisions.
func (s *controllerRevisionLister) ControllerRevisions(namespace string) ControllerRevisionNamespaceLister {
	return controllerRevisionNamespaceLister{
		indexer:     s.indexer,
		clusterName: s.clusterName,
		namespace:   namespace,
	}
}

// ControllerRevisionNamespaceLister helps list and get ControllerRevisions.
// All objects returned here must be treated as read-only.
type ControllerRevisionNamespaceLister interface {
	// List lists all ControllerRevisions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ControllerRevision, err error)
	// Get retrieves the ControllerRevision from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ControllerRevision, error)
	ControllerRevisionNamespaceListerExpansion
}

// controllerRevisionNamespaceLister implements the ControllerRevisionNamespaceLister
// interface.
type controllerRevisionNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the ControllerRevision from the index for a given name.
func (s controllerRevisionNamespaceLister) Get(name string) (*v1.ControllerRevision, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("controllerrevision"), name)
	}
	return obj.(*v1.ControllerRevision), nil
}

// Get retrieves the ControllerRevision from the index for a given name.
func (s controllerRevisionNamespaceLister) List(selector labels.Selector) (ret []*v1.ControllerRevision, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ControllerRevision))
	})
	return ret, err
}
