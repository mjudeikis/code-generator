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

	nodev1alpha1 "acme.corp/pkg/apis/node/v1alpha1"
	applyconfigurationsnodev1alpha1 "acme.corp/pkg/generated/applyconfigurations/node/v1alpha1"
	nodev1alpha1client "acme.corp/pkg/generated/clientset/versioned/typed/node/v1alpha1"
)

var runtimeClassesResource = schema.GroupVersionResource{Group: "node.k8s.io", Version: "v1alpha1", Resource: "runtimeclasses"}
var runtimeClassesKind = schema.GroupVersionKind{Group: "node.k8s.io", Version: "v1alpha1", Kind: "RuntimeClass"}

type runtimeClassesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *runtimeClassesClusterClient) Cluster(clusterPath logicalcluster.Path) nodev1alpha1client.RuntimeClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &runtimeClassesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of RuntimeClasses that match those selectors across all clusters.
func (c *runtimeClassesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*nodev1alpha1.RuntimeClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(runtimeClassesResource, runtimeClassesKind, logicalcluster.Wildcard, opts), &nodev1alpha1.RuntimeClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nodev1alpha1.RuntimeClassList{ListMeta: obj.(*nodev1alpha1.RuntimeClassList).ListMeta}
	for _, item := range obj.(*nodev1alpha1.RuntimeClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested RuntimeClasses across all clusters.
func (c *runtimeClassesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(runtimeClassesResource, logicalcluster.Wildcard, opts))
}

type runtimeClassesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *runtimeClassesClient) Create(ctx context.Context, runtimeClass *nodev1alpha1.RuntimeClass, opts metav1.CreateOptions) (*nodev1alpha1.RuntimeClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(runtimeClassesResource, c.ClusterPath, runtimeClass), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}

func (c *runtimeClassesClient) Update(ctx context.Context, runtimeClass *nodev1alpha1.RuntimeClass, opts metav1.UpdateOptions) (*nodev1alpha1.RuntimeClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(runtimeClassesResource, c.ClusterPath, runtimeClass), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}

func (c *runtimeClassesClient) UpdateStatus(ctx context.Context, runtimeClass *nodev1alpha1.RuntimeClass, opts metav1.UpdateOptions) (*nodev1alpha1.RuntimeClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(runtimeClassesResource, c.ClusterPath, "status", runtimeClass), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}

func (c *runtimeClassesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(runtimeClassesResource, c.ClusterPath, name, opts), &nodev1alpha1.RuntimeClass{})
	return err
}

func (c *runtimeClassesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(runtimeClassesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &nodev1alpha1.RuntimeClassList{})
	return err
}

func (c *runtimeClassesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*nodev1alpha1.RuntimeClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(runtimeClassesResource, c.ClusterPath, name), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}

// List takes label and field selectors, and returns the list of RuntimeClasses that match those selectors.
func (c *runtimeClassesClient) List(ctx context.Context, opts metav1.ListOptions) (*nodev1alpha1.RuntimeClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(runtimeClassesResource, runtimeClassesKind, c.ClusterPath, opts), &nodev1alpha1.RuntimeClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nodev1alpha1.RuntimeClassList{ListMeta: obj.(*nodev1alpha1.RuntimeClassList).ListMeta}
	for _, item := range obj.(*nodev1alpha1.RuntimeClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *runtimeClassesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(runtimeClassesResource, c.ClusterPath, opts))
}

func (c *runtimeClassesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*nodev1alpha1.RuntimeClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(runtimeClassesResource, c.ClusterPath, name, pt, data, subresources...), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}

func (c *runtimeClassesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsnodev1alpha1.RuntimeClassApplyConfiguration, opts metav1.ApplyOptions) (*nodev1alpha1.RuntimeClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(runtimeClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}

func (c *runtimeClassesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsnodev1alpha1.RuntimeClassApplyConfiguration, opts metav1.ApplyOptions) (*nodev1alpha1.RuntimeClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(runtimeClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &nodev1alpha1.RuntimeClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*nodev1alpha1.RuntimeClass), err
}
