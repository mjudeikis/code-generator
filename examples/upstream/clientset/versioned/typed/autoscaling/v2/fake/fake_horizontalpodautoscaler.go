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
	v2 "k8s.io/api/autoscaling/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	autoscalingv2 "k8s.io/client-go/applyconfigurations/autoscaling/v2"
	upstreamautoscalingv2client "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/autoscaling/v2"
)

var horizontalpodautoscalersResource = v2.SchemeGroupVersion.WithResource("horizontalpodautoscalers")

var horizontalpodautoscalersKind = v2.SchemeGroupVersion.WithKind("HorizontalPodAutoscaler")

// horizontalPodAutoscalersClusterClient implements horizontalPodAutoscalerInterface
type horizontalPodAutoscalersClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *horizontalPodAutoscalersClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.HorizontalPodAutoscalerNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &horizontalPodAutoscalersNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of HorizontalPodAutoscalers that match those selectors.
func (c *horizontalPodAutoscalersClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v2.HorizontalPodAutoscalerList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(horizontalpodautoscalersResource, horizontalpodautoscalersKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v2.HorizontalPodAutoscalerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.HorizontalPodAutoscalerList{ListMeta: obj.(*v2.HorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v2.HorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested horizontalPodAutoscalers across all clusters.
func (c *horizontalPodAutoscalersClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(horizontalpodautoscalersResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type horizontalPodAutoscalersNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *horizontalPodAutoscalersNamespacer) Namespace(namespace string) upstreamautoscalingv2client.HorizontalPodAutoscalerInterface {
	return &horizontalPodAutoscalersClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type horizontalPodAutoscalersClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *horizontalPodAutoscalersClient) Create(ctx context.Context, horizontalPodAutoscaler *v2.HorizontalPodAutoscaler, opts metav1.CreateOptions) (*v2.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, horizontalPodAutoscaler), &v2.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) Update(ctx context.Context, horizontalPodAutoscaler *v2.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*v2.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, horizontalPodAutoscaler), &v2.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) UpdateStatus(ctx context.Context, horizontalPodAutoscaler *v2.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*v2.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(horizontalpodautoscalersResource, c.ClusterPath, "status", c.Namespace, horizontalPodAutoscaler), &v2.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, name, opts), &v2.HorizontalPodAutoscaler{})

	return err
}

func (c *horizontalPodAutoscalersClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v2.HorizontalPodAutoscalerList{})
	return err
}

func (c *horizontalPodAutoscalersClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v2.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, name), &v2.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) List(ctx context.Context, opts metav1.ListOptions) (*v2.HorizontalPodAutoscalerList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(horizontalpodautoscalersResource, horizontalpodautoscalersKind, c.ClusterPath, c.Namespace, opts), &v2.HorizontalPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.HorizontalPodAutoscalerList{ListMeta: obj.(*v2.HorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v2.HorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *horizontalPodAutoscalersClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, opts))

}

func (c *horizontalPodAutoscalersClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v2.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v2.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) Apply(ctx context.Context, applyConfiguration *autoscalingv2.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (*v2.HorizontalPodAutoscaler, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v2.HorizontalPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) ApplyStatus(ctx context.Context, applyConfiguration *autoscalingv2.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (*v2.HorizontalPodAutoscaler, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(horizontalpodautoscalersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &v2.HorizontalPodAutoscaler{})

	return obj.(*v2.HorizontalPodAutoscaler), err
}
