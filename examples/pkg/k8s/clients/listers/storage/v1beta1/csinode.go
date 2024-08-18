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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	storagev1beta1 "acme.corp/pkg/apis/storage/v1beta1"
)

// CSINodeClusterLister can list CSINodes across all workspaces, or scope down to a CSINodeLister for one workspace.
// All objects returned here must be treated as read-only.
type CSINodeClusterLister interface {
	// List lists all CSINodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1beta1.CSINode, err error)
	// Cluster returns a lister that can list and get CSINodes in one workspace.
	Cluster(clusterName logicalcluster.Name) CSINodeLister
	CSINodeClusterListerExpansion
}

type cSINodeClusterLister struct {
	indexer cache.Indexer
}

// NewCSINodeClusterLister returns a new CSINodeClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewCSINodeClusterLister(indexer cache.Indexer) *cSINodeClusterLister {
	return &cSINodeClusterLister{indexer: indexer}
}

// List lists all CSINodes in the indexer across all workspaces.
func (s *cSINodeClusterLister) List(selector labels.Selector) (ret []*storagev1beta1.CSINode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storagev1beta1.CSINode))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get CSINodes.
func (s *cSINodeClusterLister) Cluster(clusterName logicalcluster.Name) CSINodeLister {
	return &cSINodeLister{indexer: s.indexer, clusterName: clusterName}
}

// CSINodeLister can list all CSINodes, or get one in particular.
// All objects returned here must be treated as read-only.
type CSINodeLister interface {
	// List lists all CSINodes in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*storagev1beta1.CSINode, err error)
	// Get retrieves the CSINode from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*storagev1beta1.CSINode, error)
	CSINodeListerExpansion
}

// cSINodeLister can list all CSINodes inside a workspace.
type cSINodeLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all CSINodes in the indexer for a workspace.
func (s *cSINodeLister) List(selector labels.Selector) (ret []*storagev1beta1.CSINode, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*storagev1beta1.CSINode))
	})
	return ret, err
}

// Get retrieves the CSINode from the indexer for a given workspace and name.
func (s *cSINodeLister) Get(name string) (*storagev1beta1.CSINode, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storagev1beta1.Resource("csinodes"), name)
	}
	return obj.(*storagev1beta1.CSINode), nil
}

// NewCSINodeLister returns a new CSINodeLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewCSINodeLister(indexer cache.Indexer) *cSINodeScopedLister {
	return &cSINodeScopedLister{indexer: indexer}
}

// cSINodeScopedLister can list all CSINodes inside a workspace.
type cSINodeScopedLister struct {
	indexer cache.Indexer
}

// List lists all CSINodes in the indexer for a workspace.
func (s *cSINodeScopedLister) List(selector labels.Selector) (ret []*storagev1beta1.CSINode, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*storagev1beta1.CSINode))
	})
	return ret, err
}

// Get retrieves the CSINode from the indexer for a given workspace and name.
func (s *cSINodeScopedLister) Get(name string) (*storagev1beta1.CSINode, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storagev1beta1.Resource("csinodes"), name)
	}
	return obj.(*storagev1beta1.CSINode), nil
}
