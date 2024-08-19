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
	"context"
	time "time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamthirdexamplev1informers "k8s.io/client-go/informers/thirdexample/v1"
	cache "k8s.io/client-go/tools/cache"
	example3iov1 "k8s.io/code-generator/examples/apiserver/apis/example3.io/v1"
	versioned "k8s.io/code-generator/examples/apiserver/clientset/versioned"
	internalinterfaces "k8s.io/code-generator/examples/apiserver/informers/externalversions/internalinterfaces"
	v1 "k8s.io/code-generator/examples/apiserver/listers/example3.io/v1"
)

// TestTypeInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1.TestTypeClusterLister
	Cluster(logicalcluster.Name) upstreamthirdexamplev1informers.TestTypeInformer
}

type testTypeClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ThirdExampleV1().TestTypes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ThirdExampleV1().TestTypes(namespace).Watch(context.TODO(), options)
			},
		},
		&example3iov1.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testTypeClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&example3iov1.TestType{}, f.defaultInformer)
}

func (f *testTypeClusterInformer) Lister() v1.TestTypeClusterLister {
	return v1.NewTestTypeClusterLister(f.Informer().GetIndexer())
}
