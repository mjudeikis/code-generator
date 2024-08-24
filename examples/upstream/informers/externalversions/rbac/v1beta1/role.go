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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	v1beta1 "k8s.io/code-generator/examples/upstream/listers/rbac/v1beta1"
	time "time"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	rbacv1beta1 "k8s.io/api/rbac/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreamrbac.authorization.k8s.iov1beta1informers "k8s.io/client-go/informers/v1beta1/rbac.authorization.k8s.io"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
)


// RoleClusterInformer provides access to a shared informer and lister for
// Roles.
type RoleClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta1.RoleLister
	Cluster(logicalcluster.Name) upstreamrbac.authorization.k8s.iov1beta1informers.RoleInformer
}

type roleClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace string
}

// NewRoleClusterInformer constructs a new informer for Role type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRoleClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRoleClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRoleClusterInformer constructs a new informer for Role type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRoleClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1beta1().Roles(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1beta1().Roles(namespace).Watch(context.TODO(), options)
			},
		},
		&rbacv1beta1.Role{},
		resyncPeriod,
		indexers,
	)
}

func (f *roleClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRoleClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *roleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacv1beta1.Role{}, f.defaultInformer)
}

func (f *roleInformer) Lister() v1beta1.RoleLister {
	return v1beta1.NewRoleLister(f.Informer().GetIndexer())
}
