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

var resourceClaimsResource = schema.GroupVersionResource{Group: "resource.k8s.io", Version: "v1alpha2", Resource: "resourceclaims"}
var resourceClaimsKind = schema.GroupVersionKind{Group: "resource.k8s.io", Version: "v1alpha2", Kind: "ResourceClaim"}

type resourceClaimsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *resourceClaimsClusterClient) Cluster(clusterPath logicalcluster.Path) kcpresourcev1alpha2.ResourceClaimsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &resourceClaimsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ResourceClaims that match those selectors across all clusters.
func (c *resourceClaimsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha2.ResourceClaimList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(resourceClaimsResource, resourceClaimsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &resourcev1alpha2.ResourceClaimList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha2.ResourceClaimList{ListMeta: obj.(*resourcev1alpha2.ResourceClaimList).ListMeta}
	for _, item := range obj.(*resourcev1alpha2.ResourceClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ResourceClaims across all clusters.
func (c *resourceClaimsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(resourceClaimsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type resourceClaimsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *resourceClaimsNamespacer) Namespace(namespace string) resourcev1alpha2client.ResourceClaimInterface {
	return &resourceClaimsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type resourceClaimsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *resourceClaimsClient) Create(ctx context.Context, resourceClaim *resourcev1alpha2.ResourceClaim, opts metav1.CreateOptions) (*resourcev1alpha2.ResourceClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(resourceClaimsResource, c.ClusterPath, c.Namespace, resourceClaim), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}

func (c *resourceClaimsClient) Update(ctx context.Context, resourceClaim *resourcev1alpha2.ResourceClaim, opts metav1.UpdateOptions) (*resourcev1alpha2.ResourceClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(resourceClaimsResource, c.ClusterPath, c.Namespace, resourceClaim), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}

func (c *resourceClaimsClient) UpdateStatus(ctx context.Context, resourceClaim *resourcev1alpha2.ResourceClaim, opts metav1.UpdateOptions) (*resourcev1alpha2.ResourceClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(resourceClaimsResource, c.ClusterPath, "status", c.Namespace, resourceClaim), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}

func (c *resourceClaimsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(resourceClaimsResource, c.ClusterPath, c.Namespace, name, opts), &resourcev1alpha2.ResourceClaim{})
	return err
}

func (c *resourceClaimsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(resourceClaimsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &resourcev1alpha2.ResourceClaimList{})
	return err
}

func (c *resourceClaimsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*resourcev1alpha2.ResourceClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(resourceClaimsResource, c.ClusterPath, c.Namespace, name), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}

// List takes label and field selectors, and returns the list of ResourceClaims that match those selectors.
func (c *resourceClaimsClient) List(ctx context.Context, opts metav1.ListOptions) (*resourcev1alpha2.ResourceClaimList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(resourceClaimsResource, resourceClaimsKind, c.ClusterPath, c.Namespace, opts), &resourcev1alpha2.ResourceClaimList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &resourcev1alpha2.ResourceClaimList{ListMeta: obj.(*resourcev1alpha2.ResourceClaimList).ListMeta}
	for _, item := range obj.(*resourcev1alpha2.ResourceClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *resourceClaimsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(resourceClaimsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *resourceClaimsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*resourcev1alpha2.ResourceClaim, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}

func (c *resourceClaimsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha2.ResourceClaimApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha2.ResourceClaim, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}

func (c *resourceClaimsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsresourcev1alpha2.ResourceClaimApplyConfiguration, opts metav1.ApplyOptions) (*resourcev1alpha2.ResourceClaim, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(resourceClaimsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &resourcev1alpha2.ResourceClaim{})
	if obj == nil {
		return nil, err
	}
	return obj.(*resourcev1alpha2.ResourceClaim), err
}
