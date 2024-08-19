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

// DeviceClassLister helps list DeviceClasses.
// All objects returned here must be treated as read-only.
type DeviceClassLister interface {
	// List lists all DeviceClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.DeviceClass, err error)
	// Get retrieves the DeviceClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha3.DeviceClass, error)
	DeviceClassListerExpansion
}

// DeviceClassClusterLister helps list DeviceClasses.
// All objects returned here must be treated as read-only.
type DeviceClassClusterLister interface {
	// List lists all DeviceClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.DeviceClass, err error)
	DeviceClassClusterListerExpansion
}

// deviceClassLister implements the DeviceClassLister interface.
type deviceClassLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// deviceClassLister implements the DeviceClassClusterLister interface.
type deviceClassClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get DeviceClass.
func (s *deviceClassClusterLister) Cluster(clusterName logicalcluster.Name) DeviceClassLister {
	return &deviceClassLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all DeviceClasses in the indexer.
func (s *deviceClassLister) List(selector labels.Selector) (ret []*v1alpha3.DeviceClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1alpha3.DeviceClass))
	})
	return ret, err
}

// List lists all DeviceClasses in the indexer.
func (s *deviceClassClusterLister) List(selector labels.Selector) (ret []*v1alpha3.DeviceClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.DeviceClass))
	})
	return ret, err
}

// NewDeviceClassClusterLister returns a new DeviceClassClusterLister.
func NewDeviceClassClusterLister(indexer cache.Indexer) DeviceClassClusterLister {
	return &deviceClassClusterLister{
		indexer: indexer,
	}
}

// Get retrieves the DeviceClass from the index for a given name.
func (s deviceClassLister) Get(name string) (*v1alpha3.DeviceClass, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("deviceclass"), name)
	}
	return obj.(*v1alpha3.DeviceClass), nil
}
