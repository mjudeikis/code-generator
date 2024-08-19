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
	v1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JobLister helps list Jobs.
// All objects returned here must be treated as read-only.
type JobLister interface {
	// List lists all Jobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Job, err error)
	// Jobs returns an object that can list and get Jobs.
	Jobs(namespace string) JobNamespaceLister
	JobListerExpansion
}

// JobClusterLister helps list Jobs.
// All objects returned here must be treated as read-only.
type JobClusterLister interface {
	// List lists all Jobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Job, err error)
	JobClusterListerExpansion
}

// jobLister implements the JobLister interface.
type jobLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// jobLister implements the JobClusterLister interface.
type jobClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get Job.
func (s *jobClusterLister) Cluster(clusterName logicalcluster.Name) JobLister {
	return &jobLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all Jobs in the indexer.
func (s *jobLister) List(selector labels.Selector) (ret []*v1.Job, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.Job))
	})
	return ret, err
}

// List lists all Jobs in the indexer.
func (s *jobClusterLister) List(selector labels.Selector) (ret []*v1.Job, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Job))
	})
	return ret, err
}

// NewJobClusterLister returns a new JobClusterLister.
func NewJobClusterLister(indexer cache.Indexer) JobClusterLister {
	return &jobClusterLister{
		indexer: indexer,
	}
}

// Jobs returns an object that can list and get Jobs.
func (s *jobLister) Jobs(namespace string) JobNamespaceLister {
	return jobNamespaceLister{
		indexer:     s.indexer,
		clusterName: s.clusterName,
		namespace:   namespace,
	}
}

// JobNamespaceLister helps list and get Jobs.
// All objects returned here must be treated as read-only.
type JobNamespaceLister interface {
	// List lists all Jobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Job, err error)
	// Get retrieves the Job from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Job, error)
	JobNamespaceListerExpansion
}

// jobNamespaceLister implements the JobNamespaceLister
// interface.
type jobNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the Job from the index for a given name.
func (s jobNamespaceLister) Get(name string) (*v1.Job, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("job"), name)
	}
	return obj.(*v1.Job), nil
}

// Get retrieves the Job from the index for a given name.
func (s jobNamespaceLister) List(selector labels.Selector) (ret []*v1.Job, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.Job))
	})
	return ret, err
}
