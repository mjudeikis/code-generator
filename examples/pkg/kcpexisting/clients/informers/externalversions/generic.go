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

package informers

import (
	"fmt"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"

	examplev1 "acme.corp/pkg/apis/example/v1"
	examplev1alpha1 "acme.corp/pkg/apis/example/v1alpha1"
	examplev1beta1 "acme.corp/pkg/apis/example/v1beta1"
	examplev2 "acme.corp/pkg/apis/example/v2"
	example3v1 "acme.corp/pkg/apis/example3/v1"
	exampledashedv1 "acme.corp/pkg/apis/exampledashed/v1"
	existinginterfacesv1 "acme.corp/pkg/apis/existinginterfaces/v1"
	secondexamplev1 "acme.corp/pkg/apis/secondexample/v1"
	upstreaminformers "acme.corp/pkg/generated/informers/externalversions"
)

type GenericClusterInformer interface {
	Cluster(logicalcluster.Name) upstreaminformers.GenericInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() kcpcache.GenericClusterLister
}

type genericClusterInformer struct {
	informer kcpcache.ScopeableSharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.informer
}

// Lister returns the GenericClusterLister.
func (f *genericClusterInformer) Lister() kcpcache.GenericClusterLister {
	return kcpcache.NewGenericClusterLister(f.Informer().GetIndexer(), f.resource)
}

// Cluster scopes to a GenericInformer.
func (f *genericClusterInformer) Cluster(clusterName logicalcluster.Name) upstreaminformers.GenericInformer {
	return &genericInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().ByCluster(clusterName),
	}
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	lister   cache.GenericLister
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return f.lister
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error) {
	switch resource {
	// Group=example3.some.corp, Version=V1
	case example3v1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example3().V1().TestTypes().Informer()}, nil
	case example3v1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example3().V1().ClusterTestTypes().Informer()}, nil
	// Group=exampledashed.some.corp, Version=V1
	case exampledashedv1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Exampledashed().V1().TestTypes().Informer()}, nil
	case exampledashedv1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Exampledashed().V1().ClusterTestTypes().Informer()}, nil
	// Group=example, Version=V1
	case examplev1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V1().TestTypes().Informer()}, nil
	case examplev1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V1().ClusterTestTypes().Informer()}, nil
	// Group=example, Version=V1alpha1
	case examplev1alpha1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().TestTypes().Informer()}, nil
	case examplev1alpha1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V1alpha1().ClusterTestTypes().Informer()}, nil
	// Group=example, Version=V1beta1
	case examplev1beta1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V1beta1().TestTypes().Informer()}, nil
	case examplev1beta1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V1beta1().ClusterTestTypes().Informer()}, nil
	// Group=example, Version=V2
	case examplev2.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V2().TestTypes().Informer()}, nil
	case examplev2.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Example().V2().ClusterTestTypes().Informer()}, nil
	// Group=existinginterfaces.acme.corp, Version=V1
	case existinginterfacesv1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Existinginterfaces().V1().TestTypes().Informer()}, nil
	case existinginterfacesv1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Existinginterfaces().V1().ClusterTestTypes().Informer()}, nil
	// Group=secondexample, Version=V1
	case secondexamplev1.SchemeGroupVersion.WithResource("testtypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Secondexample().V1().TestTypes().Informer()}, nil
	case secondexamplev1.SchemeGroupVersion.WithResource("clustertesttypes"):
		return &genericClusterInformer{resource: resource.GroupResource(), informer: f.Secondexample().V1().ClusterTestTypes().Informer()}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
