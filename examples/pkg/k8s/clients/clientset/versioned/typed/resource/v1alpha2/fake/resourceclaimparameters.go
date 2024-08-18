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

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	resourcev1alpha2 "acme.corp/pkg/apis/resource/v1alpha2"
	applyconfigurationsresourcev1alpha2 "acme.corp/pkg/generated/applyconfigurations/resource/v1alpha2"
	resourcev1alpha2client "acme.corp/pkg/generated/clientset/versioned/typed/resource/v1alpha2"
	kcpresourcev1alpha2 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/resource/v1alpha2"
)

var resourceClaimParametersResource = schema.GroupVersionResource{Group: "resource.k8s.io", Version: "v1alpha2", Resource: "resourceclaimparameters"}
var resourceClaimParametersKind = schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1alpha2", Kind: "ResourceClaimParameters"}

type resourceClaimParametersClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *resourceClaimParametersClusterClient) Cluster(clusterPath logicalcluster.Path) kcpresourcev1alpha2.ResourceClaimParametersNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &resourceClaimParametersNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ResourceClaimParameters that match those selectors across all clusters.
func (c *resourceClaimParametersClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha2.ResourceClaimParametersList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(resourceClaimParametersResource, resourceClaimParametersKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &resourcev1alpha2.ResourceClaimParametersList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha2.ResourceClaimParametersList{ListMeta: obj.(*resourcev1alpha2.ResourceClaimParametersList).ListMeta}
	for _, item := range obj.(*resourcev1alpha2.ResourceClaimParametersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ResourceClaimParameters across all clusters.
func (c *resourceClaimParametersClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(resourceClaimParametersResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type resourceClaimParametersNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *resourceClaimParametersNamespacer) Namespace(namespace string) resourcev1alpha2client.ResourceClaimParametersInterface {
	return &resourceClaimParametersClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type resourceClaimParametersClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *resourceClaimParametersClient) Create(ctx context.Context, resourceClaimParameters *resourcev1alpha2.ResourceClaimParameters, opts metav1.CreateOptions) (*resourcev1alpha2.ResourceClaimParameters, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, resourceClaimParameters), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}

func (c *resourceClaimParametersClient) Update(ctx context.Context, resourceClaimParameters *resourcev1alpha2.ResourceClaimParameters, opts metav1.UpdateOptions) (*resourcev1alpha2.ResourceClaimParameters, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, resourceClaimParameters), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}

func (c *resourceClaimParametersClient) UpdateStatus(ctx context.Context, resourceClaimParameters *resourcev1alpha2.ResourceClaimParameters, opts metav1.UpdateOptions) (*resourcev1alpha2.ResourceClaimParameters, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(resourceClaimParametersResource, c.ClusterPath, "status", c.Namespace, resourceClaimParameters), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}

func (c *resourceClaimParametersClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(resourceClaimParametersResource, c.ClusterPath, c.Namespace, name, opts), &resourcev1alpha2.ResourceClaimParameters{})
	return err
}

func (c *resourceClaimParametersClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &resourcev1alpha2.ResourceClaimParametersList{})
	return err
}

func (c *resourceClaimParametersClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*resourcev1alpha2.ResourceClaimParameters, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, name), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}

// List takes label and field selectors, and returns the list of ResourceClaimParameters that match those selectors.
func (c *resourceClaimParametersClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha2.ResourceClaimParametersList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(resourceClaimParametersResource, resourceClaimParametersKind, c.ClusterPath, c.Namespace, opts), &resourcev1alpha2.ResourceClaimParametersList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha2.ResourceClaimParametersList{ListMeta: obj.(*resourcev1alpha2.ResourceClaimParametersList).ListMeta}
	for _, item := range obj.(*resourcev1alpha2.ResourceClaimParametersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *resourceClaimParametersClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, opts))
}

func (c *resourceClaimParametersClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*resourcev1alpha2.ResourceClaimParameters, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}

func (c *resourceClaimParametersClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha2.ResourceClaimParametersApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha2.ResourceClaimParameters, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}

func (c *resourceClaimParametersClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha2.ResourceClaimParametersApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha2.ResourceClaimParameters, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimParametersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &resourcev1alpha2.ResourceClaimParameters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaimParameters), err
}
