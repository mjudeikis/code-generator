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


// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreamadmissionregistration.k8s.iov1alpha1informers "k8s.io/client-go/informers/v1alpha1/admissionregistration.k8s.io"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	admissionregistrationv1alpha1 "k8s.io/api/admissionregistration/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/code-generator/examples/upstream/listers/admissionregistration/v1alpha1"
	time "time"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
)


// ValidatingAdmissionPolicyBindingClusterInformer provides access to a shared informer and lister for
// ValidatingAdmissionPolicyBindings.
type ValidatingAdmissionPolicyBindingClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1alpha1.ValidatingAdmissionPolicyBindingLister
	Cluster(logicalcluster.Name) upstreamadmissionregistration.k8s.iov1alpha1informers.ValidatingAdmissionPolicyBindingInformer
}

type validatingAdmissionPolicyBindingClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	
}

// NewValidatingAdmissionPolicyBindingClusterInformer constructs a new informer for ValidatingAdmissionPolicyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewValidatingAdmissionPolicyBindingClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyBindingClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredValidatingAdmissionPolicyBindingClusterInformer constructs a new informer for ValidatingAdmissionPolicyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredValidatingAdmissionPolicyBindingClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1alpha1().ValidatingAdmissionPolicyBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1alpha1().ValidatingAdmissionPolicyBindings().Watch(context.TODO(), options)
			},
		},
		&admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *validatingAdmissionPolicyBindingClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyBindingClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *validatingAdmissionPolicyBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admissionregistrationv1alpha1.ValidatingAdmissionPolicyBinding{}, f.defaultInformer)
}

func (f *validatingAdmissionPolicyBindingInformer) Lister() v1alpha1.ValidatingAdmissionPolicyBindingLister {
	return v1alpha1.NewValidatingAdmissionPolicyBindingLister(f.Informer().GetIndexer())
}
