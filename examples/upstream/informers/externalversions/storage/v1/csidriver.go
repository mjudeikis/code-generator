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

package v1

import (
	storagev1 "k8s.io/api/storage/v1"
	v1 "k8s.io/code-generator/examples/upstream/listers/storage/v1"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	time "time"
	"github.com/kcp-dev/logicalcluster/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	upstreamstorage.k8s.iov1informers "k8s.io/client-go/informers/v1/storage.k8s.io"
)


// CSIDriverClusterInformer provides access to a shared informer and lister for
// CSIDrivers.
type CSIDriverClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.CSIDriverLister
	Cluster(logicalcluster.Name) upstreamstorage.k8s.iov1informers.CSIDriverInformer
}

type cSIDriverClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	
}

// NewCSIDriverClusterInformer constructs a new informer for CSIDriver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCSIDriverClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCSIDriverClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCSIDriverClusterInformer constructs a new informer for CSIDriver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCSIDriverClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().CSIDrivers().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().CSIDrivers().Watch(context.TODO(), options)
			},
		},
		&storagev1.CSIDriver{},
		resyncPeriod,
		indexers,
	)
}

func (f *cSIDriverClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCSIDriverClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *cSIDriverInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1.CSIDriver{}, f.defaultInformer)
}

func (f *cSIDriverInformer) Lister() v1.CSIDriverLister {
	return v1.NewCSIDriverLister(f.Informer().GetIndexer())
}
