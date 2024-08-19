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

package v1beta3

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta3 "k8s.io/api/flowcontrol/v1beta3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PriorityLevelConfigurationLister helps list PriorityLevelConfigurations.
// All objects returned here must be treated as read-only.
type PriorityLevelConfigurationLister interface {
	// List lists all PriorityLevelConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta3.PriorityLevelConfiguration, err error)
	// Get retrieves the PriorityLevelConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta3.PriorityLevelConfiguration, error)
	PriorityLevelConfigurationListerExpansion
}

// PriorityLevelConfigurationClusterLister helps list PriorityLevelConfigurations.
// All objects returned here must be treated as read-only.
type PriorityLevelConfigurationClusterLister interface {
	// List lists all PriorityLevelConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta3.PriorityLevelConfiguration, err error)
	PriorityLevelConfigurationClusterListerExpansion
}

// priorityLevelConfigurationLister implements the PriorityLevelConfigurationLister interface.
type priorityLevelConfigurationLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// priorityLevelConfigurationLister implements the PriorityLevelConfigurationClusterLister interface.
type priorityLevelConfigurationClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get PriorityLevelConfiguration.
func (s *priorityLevelConfigurationClusterLister) Cluster(clusterName logicalcluster.Name) PriorityLevelConfigurationLister {
	return &priorityLevelConfigurationLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all PriorityLevelConfigurations in the indexer.
func (s *priorityLevelConfigurationLister) List(selector labels.Selector) (ret []*v1beta3.PriorityLevelConfiguration, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta3.PriorityLevelConfiguration))
	})
	return ret, err
}

// List lists all PriorityLevelConfigurations in the indexer.
func (s *priorityLevelConfigurationClusterLister) List(selector labels.Selector) (ret []*v1beta3.PriorityLevelConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta3.PriorityLevelConfiguration))
	})
	return ret, err
}

// NewPriorityLevelConfigurationClusterLister returns a new PriorityLevelConfigurationClusterLister.
func NewPriorityLevelConfigurationClusterLister(indexer cache.Indexer) PriorityLevelConfigurationClusterLister {
	return &priorityLevelConfigurationClusterLister{
		indexer: indexer,
	}
}

// Get retrieves the PriorityLevelConfiguration from the index for a given name.
func (s priorityLevelConfigurationLister) Get(name string) (*v1beta3.PriorityLevelConfiguration, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta3.Resource("prioritylevelconfiguration"), name)
	}
	return obj.(*v1beta3.PriorityLevelConfiguration), nil
}
