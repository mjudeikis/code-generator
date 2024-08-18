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
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	extensionsv1beta1 "acme.corp/pkg/apis/extensions/v1beta1"
	scopedclientset "acme.corp/pkg/generated/clientset/versioned"
	clientset "acme.corp/pkg/kcp/clients/clientset/versioned"
	"acme.corp/pkg/kcp/clients/informers/externalversions/internalinterfaces"
	extensionsv1beta1listers "acme.corp/pkg/kcp/clients/listers/extensions/v1beta1"
)

// NetworkPolicyClusterInformer provides access to a shared informer and lister for
// NetworkPolicies.
type NetworkPolicyClusterInformer interface {
	Cluster(logicalcluster.Name) NetworkPolicyInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() extensionsv1beta1listers.NetworkPolicyClusterLister
}

type networkPolicyClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNetworkPolicyClusterInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredNetworkPolicyClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkPolicyClusterInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkPolicyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().NetworkPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().NetworkPolicies().Watch(context.TODO(), options)
			},
		},
		&extensionsv1beta1.NetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkPolicyClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredNetworkPolicyClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *networkPolicyClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&extensionsv1beta1.NetworkPolicy{}, f.defaultInformer)
}

func (f *networkPolicyClusterInformer) Lister() extensionsv1beta1listers.NetworkPolicyClusterLister {
	return extensionsv1beta1listers.NewNetworkPolicyClusterLister(f.Informer().GetIndexer())
}

// NetworkPolicyInformer provides access to a shared informer and lister for
// NetworkPolicies.
type NetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() extensionsv1beta1listers.NetworkPolicyLister
}

func (f *networkPolicyClusterInformer) Cluster(clusterName logicalcluster.Name) NetworkPolicyInformer {
	return &networkPolicyInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type networkPolicyInformer struct {
	informer cache.SharedIndexInformer
	lister   extensionsv1beta1listers.NetworkPolicyLister
}

func (f *networkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *networkPolicyInformer) Lister() extensionsv1beta1listers.NetworkPolicyLister {
	return f.lister
}

type networkPolicyScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func (f *networkPolicyScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensionsv1beta1.NetworkPolicy{}, f.defaultInformer)
}

func (f *networkPolicyScopedInformer) Lister() extensionsv1beta1listers.NetworkPolicyLister {
	return extensionsv1beta1listers.NewNetworkPolicyLister(f.Informer().GetIndexer())
}

// NewNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkPolicyInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkPolicyInformer(client, resyncPeriod, namespace, indexers, nil)
}

// NewFilteredNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkPolicyInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().NetworkPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExtensionsV1beta1().NetworkPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&extensionsv1beta1.NetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkPolicyScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkPolicyInformer(client, resyncPeriod, f.namespace, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	}, f.tweakListOptions)
}
