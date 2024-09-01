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
	v1beta2 "k8s.io/api/flowcontrol/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	flowcontrolv1beta2 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta2"
	upstreamflowcontrolv1beta2client "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	"k8s.io/client-go/testing"
)

var flowschemasResource = v1beta2.SchemeGroupVersion.WithResource("flowschemas")

var flowschemasKind = v1beta2.SchemeGroupVersion.WithKind("FlowSchema")

// flowSchemasClusterClient implements flowSchemaInterface
type flowSchemasClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *flowSchemasClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamflowcontrolv1beta2client.FlowSchemaInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &flowSchemasClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors.
func (c *flowSchemasClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.FlowSchemaList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(flowschemasResource, flowschemasKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta2.FlowSchemaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.FlowSchemaList{ListMeta: obj.(*v1beta2.FlowSchemaList).ListMeta}
	for _, item := range obj.(*v1beta2.FlowSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested flowSchemas across all clusters.
func (c *flowSchemasClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(flowschemasResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type flowSchemasClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *flowSchemasClient) Create(ctx context.Context, flowSchema *v1beta2.FlowSchema, opts metav1.CreateOptions) (*v1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(flowschemasResource, c.ClusterPath, flowSchema), &v1beta2.FlowSchema{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.FlowSchema), err
}

func (c *flowSchemasClient) Update(ctx context.Context, flowSchema *v1beta2.FlowSchema, opts metav1.UpdateOptions) (*v1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(flowschemasResource, c.ClusterPath, flowSchema), &v1beta2.FlowSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.FlowSchema), err
}

func (c *flowSchemasClient) UpdateStatus(ctx context.Context, flowSchema *v1beta2.FlowSchema, opts metav1.UpdateOptions) (*v1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(flowschemasResource, c.ClusterPath, "status", flowSchema), &v1beta2.FlowSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.FlowSchema), err
}

func (c *flowSchemasClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(flowschemasResource, c.ClusterPath, name, opts), &v1beta2.FlowSchema{})

	return err
}

func (c *flowSchemasClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(flowschemasResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.FlowSchemaList{})
	return err
}

func (c *flowSchemasClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(flowschemasResource, c.ClusterPath, name), &v1beta2.FlowSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.FlowSchema), err
}

func (c *flowSchemasClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta2.FlowSchemaList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(flowschemasResource, flowschemasKind, c.ClusterPath, opts), &v1beta2.FlowSchemaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.FlowSchemaList{ListMeta: obj.(*v1beta2.FlowSchemaList).ListMeta}
	for _, item := range obj.(*v1beta2.FlowSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *flowSchemasClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(flowschemasResource, c.ClusterPath, opts))

}

func (c *flowSchemasClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta2.FlowSchema, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(flowschemasResource, c.ClusterPath, name, pt, data, subresources...), &v1beta2.FlowSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.FlowSchema), err
}

func (c *flowSchemasClient) Apply(ctx context.Context, applyConfiguration *flowcontrolv1beta2.FlowSchemaApplyConfiguration, opts metav1.ApplyOptions) (*v1beta2.FlowSchema, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(flowschemasResource, c.ClusterPath, *name, types.ApplyPatchType, data), &v1beta2.FlowSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.FlowSchema), err
}

func (c *flowSchemasClient) ApplyStatus(ctx context.Context, applyConfiguration *flowcontrolv1beta2.FlowSchemaApplyConfiguration, opts metav1.ApplyOptions) (*v1beta2.FlowSchema, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(flowschemasResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &v1beta2.FlowSchema{})

	return obj.(*v1beta2.FlowSchema), err
}
