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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	networkingv1beta1 "k8s.io/client-go/applyconfigurations/networking/v1beta1"
	upstreamnetworkingv1beta1client "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/networking/v1beta1"
)

var ingressesResource = v1beta1.SchemeGroupVersion.WithResource("ingresses")

var ingressesKind = v1beta1.SchemeGroupVersion.WithKind("Ingress")

// ingressesClusterClient implements ingressInterface
type ingressesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *ingressesClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.IngressNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &ingressesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Ingresses that match those selectors.
func (c *ingressesClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IngressList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(ingressesResource, ingressesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.IngressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IngressList{ListMeta: obj.(*v1beta1.IngressList).ListMeta}
	for _, item := range obj.(*v1beta1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingresss across all clusters.
func (c *ingressesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(ingressesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type ingressesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *ingressesNamespacer) Namespace(namespace string) upstreamnetworkingv1beta1client.IngressInterface {
	return &ingressesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type ingressesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *ingressesClient) Create(ctx context.Context, ingress *v1beta1.Ingress, opts metav1.CreateOptions) (*v1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(ingressesResource, c.ClusterPath, c.Namespace, ingress), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

func (c *ingressesClient) Update(ctx context.Context, ingress *v1beta1.Ingress, opts metav1.UpdateOptions) (*v1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(ingressesResource, c.ClusterPath, c.Namespace, ingress), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

func (c *ingressesClient) UpdateStatus(ctx context.Context, ingress *v1beta1.Ingress, opts metav1.UpdateOptions) (*v1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(ingressesResource, c.ClusterPath, "status", c.Namespace, ingress), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

func (c *ingressesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(ingressesResource, c.ClusterPath, c.Namespace, name, opts), &v1beta1.Ingress{})

	return err
}

func (c *ingressesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(ingressesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IngressList{})
	return err
}

func (c *ingressesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(ingressesResource, c.ClusterPath, c.Namespace, name), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

func (c *ingressesClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.IngressList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(ingressesResource, ingressesKind, c.ClusterPath, c.Namespace, opts), &v1beta1.IngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IngressList{ListMeta: obj.(*v1beta1.IngressList).ListMeta}
	for _, item := range obj.(*v1beta1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *ingressesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(ingressesResource, c.ClusterPath, c.Namespace, opts))

}

func (c *ingressesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.Ingress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(ingressesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

func (c *ingressesClient) Apply(ctx context.Context, applyConfiguration *networkingv1beta1.IngressApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.Ingress, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(ingressesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Ingress), err
}

func (c *ingressesClient) ApplyStatus(ctx context.Context, applyConfiguration *networkingv1beta1.IngressApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.Ingress, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(ingressesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &v1beta1.Ingress{})

	return obj.(*v1beta1.Ingress), err
}
