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
	v1 "k8s.io/api/policy/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodDisruptionBudgetLister helps list PodDisruptionBudgets.
// All objects returned here must be treated as read-only.
type PodDisruptionBudgetLister interface {
	// List lists all PodDisruptionBudgets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodDisruptionBudget, err error)
	// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
	PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister
	PodDisruptionBudgetListerExpansion
}

// PodDisruptionBudgetClusterLister helps list PodDisruptionBudgets.
// All objects returned here must be treated as read-only.
type PodDisruptionBudgetClusterLister interface {
	// List lists all PodDisruptionBudgets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodDisruptionBudget, err error)
	PodDisruptionBudgetClusterListerExpansion
}

// podDisruptionBudgetLister implements the PodDisruptionBudgetLister interface.
type podDisruptionBudgetLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// podDisruptionBudgetLister implements the PodDisruptionBudgetClusterLister interface.
type podDisruptionBudgetClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get PodDisruptionBudget.
func (s *podDisruptionBudgetClusterLister) Cluster(clusterName logicalcluster.Name) PodDisruptionBudgetLister {
	return &podDisruptionBudgetLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all PodDisruptionBudgets in the indexer.
func (s *podDisruptionBudgetLister) List(selector labels.Selector) (ret []*v1.PodDisruptionBudget, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.PodDisruptionBudget))
	})
	return ret, err
}

// List lists all PodDisruptionBudgets in the indexer.
func (s *podDisruptionBudgetClusterLister) List(selector labels.Selector) (ret []*v1.PodDisruptionBudget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodDisruptionBudget))
	})
	return ret, err
}

// NewPodDisruptionBudgetClusterLister returns a new PodDisruptionBudgetClusterLister.
func NewPodDisruptionBudgetClusterLister(indexer cache.Indexer) PodDisruptionBudgetClusterLister {
	return &podDisruptionBudgetClusterLister{
		indexer: indexer,
	}
}

// PodDisruptionBudgets returns an object that can list and get PodDisruptionBudgets.
func (s *podDisruptionBudgetLister) PodDisruptionBudgets(namespace string) PodDisruptionBudgetNamespaceLister {
	return podDisruptionBudgetNamespaceLister{
		indexer:     s.indexer,
		clusterName: s.clusterName,
		namespace:   namespace,
	}
}

// PodDisruptionBudgetNamespaceLister helps list and get PodDisruptionBudgets.
// All objects returned here must be treated as read-only.
type PodDisruptionBudgetNamespaceLister interface {
	// List lists all PodDisruptionBudgets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodDisruptionBudget, err error)
	// Get retrieves the PodDisruptionBudget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PodDisruptionBudget, error)
	PodDisruptionBudgetNamespaceListerExpansion
}

// podDisruptionBudgetNamespaceLister implements the PodDisruptionBudgetNamespaceLister
// interface.
type podDisruptionBudgetNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the PodDisruptionBudget from the index for a given name.
func (s podDisruptionBudgetNamespaceLister) Get(name string) (*v1.PodDisruptionBudget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("poddisruptionbudget"), name)
	}
	return obj.(*v1.PodDisruptionBudget), nil
}

// Get retrieves the PodDisruptionBudget from the index for a given name.
func (s podDisruptionBudgetNamespaceLister) List(selector labels.Selector) (ret []*v1.PodDisruptionBudget, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.PodDisruptionBudget))
	})
	return ret, err
}
