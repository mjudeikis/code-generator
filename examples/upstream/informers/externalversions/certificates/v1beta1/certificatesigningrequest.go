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
	upstreamcertificates.k8s.iov1beta1informers "k8s.io/client-go/informers/v1beta1/certificates.k8s.io"
	informers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	versioned "k8s.io/code-generator/examples/upstream/clientset/versioned"
	v1beta1 "k8s.io/code-generator/examples/upstream/listers/certificates/v1beta1"
	time "time"
	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
	internalinterfaces "k8s.io/code-generator/examples/upstream/informers/externalversions/internalinterfaces"
	"github.com/kcp-dev/logicalcluster/v3"
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
)


// CertificateSigningRequestClusterInformer provides access to a shared informer and lister for
// CertificateSigningRequests.
type CertificateSigningRequestClusterInformer interface {
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() v1beta1.CertificateSigningRequestLister
	Cluster(logicalcluster.Name) upstreamcertificates.k8s.iov1beta1informers.CertificateSigningRequestInformer
}

type certificateSigningRequestClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	
}

// NewCertificateSigningRequestClusterInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateSigningRequestClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateSigningRequestClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateSigningRequestClusterInformer constructs a new informer for CertificateSigningRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateSigningRequestClusterInformer(client versioned.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return informers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1beta1().CertificateSigningRequests().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1beta1().CertificateSigningRequests().Watch(context.TODO(), options)
			},
		},
		&certificatesv1beta1.CertificateSigningRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateSigningRequestClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateSigningRequestClusterInformer(client, resyncPeriod, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
		f.tweakListOptions)
}

func (f *certificateSigningRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certificatesv1beta1.CertificateSigningRequest{}, f.defaultInformer)
}

func (f *certificateSigningRequestInformer) Lister() v1beta1.CertificateSigningRequestLister {
	return v1beta1.NewCertificateSigningRequestLister(f.Informer().GetIndexer())
}
