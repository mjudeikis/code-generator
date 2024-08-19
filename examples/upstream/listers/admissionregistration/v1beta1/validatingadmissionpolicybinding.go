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
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ValidatingAdmissionPolicyBindingLister helps list ValidatingAdmissionPolicyBindings.
// All objects returned here must be treated as read-only.
type ValidatingAdmissionPolicyBindingLister interface {
	// List lists all ValidatingAdmissionPolicyBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ValidatingAdmissionPolicyBinding, err error)
	// Get retrieves the ValidatingAdmissionPolicyBinding from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ValidatingAdmissionPolicyBinding, error)
	ValidatingAdmissionPolicyBindingListerExpansion
}

// ValidatingAdmissionPolicyBindingClusterLister helps list ValidatingAdmissionPolicyBindings.
// All objects returned here must be treated as read-only.
type ValidatingAdmissionPolicyBindingClusterLister interface {
	// List lists all ValidatingAdmissionPolicyBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ValidatingAdmissionPolicyBinding, err error)
	ValidatingAdmissionPolicyBindingClusterListerExpansion
}

// validatingAdmissionPolicyBindingLister implements the ValidatingAdmissionPolicyBindingLister interface.
type validatingAdmissionPolicyBindingLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// validatingAdmissionPolicyBindingLister implements the ValidatingAdmissionPolicyBindingClusterLister interface.
type validatingAdmissionPolicyBindingClusterLister struct {
	indexer cache.Indexer
}

// Cluster scopes the lister to one workspace, allowing users to list and get ValidatingAdmissionPolicyBinding.
func (s *validatingAdmissionPolicyBindingClusterLister) Cluster(clusterName logicalcluster.Name) ValidatingAdmissionPolicyBindingLister {
	return &validatingAdmissionPolicyBindingLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all ValidatingAdmissionPolicyBindings in the indexer.
func (s *validatingAdmissionPolicyBindingLister) List(selector labels.Selector) (ret []*v1beta1.ValidatingAdmissionPolicyBinding, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.ValidatingAdmissionPolicyBinding))
	})
	return ret, err
}

// List lists all ValidatingAdmissionPolicyBindings in the indexer.
func (s *validatingAdmissionPolicyBindingClusterLister) List(selector labels.Selector) (ret []*v1beta1.ValidatingAdmissionPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ValidatingAdmissionPolicyBinding))
	})
	return ret, err
}

// NewValidatingAdmissionPolicyBindingClusterLister returns a new ValidatingAdmissionPolicyBindingClusterLister.
func NewValidatingAdmissionPolicyBindingClusterLister(indexer cache.Indexer) ValidatingAdmissionPolicyBindingClusterLister {
	return &validatingAdmissionPolicyBindingClusterLister{
		indexer: indexer,
	}
}

// Get retrieves the ValidatingAdmissionPolicyBinding from the index for a given name.
func (s validatingAdmissionPolicyBindingLister) Get(name string) (*v1beta1.ValidatingAdmissionPolicyBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("validatingadmissionpolicybinding"), name)
	}
	return obj.(*v1beta1.ValidatingAdmissionPolicyBinding), nil
}
