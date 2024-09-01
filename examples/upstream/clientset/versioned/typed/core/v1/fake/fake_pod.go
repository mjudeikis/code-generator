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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	upstreamcorev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/core/v1"
)

var podsResource = v1.SchemeGroupVersion.WithResource("pods")

var podsKind = v1.SchemeGroupVersion.WithKind("Pod")

// podsClusterClient implements podInterface
type podsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *podsClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.PodNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Pods that match those selectors.
func (c *podsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PodList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podsResource, podsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.PodList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PodList{ListMeta: obj.(*v1.PodList).ListMeta}
	for _, item := range obj.(*v1.PodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pods across all clusters.
func (c *podsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type podsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *podsNamespacer) Namespace(namespace string) upstreamcorev1client.PodInterface {
	return &podsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type podsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *podsClient) Create(ctx context.Context, pod *v1.Pod, opts metav1.CreateOptions) (*v1.Pod, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(podsResource, c.ClusterPath, c.Namespace, pod), &v1.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}

func (c *podsClient) Update(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(podsResource, c.ClusterPath, c.Namespace, pod), &v1.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}

func (c *podsClient) UpdateStatus(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (*v1.Pod, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(podsResource, c.ClusterPath, "status", c.Namespace, pod), &v1.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}

func (c *podsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(podsResource, c.ClusterPath, c.Namespace, name, opts), &v1.Pod{})

	return err
}

func (c *podsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(podsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PodList{})
	return err
}

func (c *podsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.Pod, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(podsResource, c.ClusterPath, c.Namespace, name), &v1.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}

func (c *podsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1.PodList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(podsResource, podsKind, c.ClusterPath, c.Namespace, opts), &v1.PodList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PodList{ListMeta: obj.(*v1.PodList).ListMeta}
	for _, item := range obj.(*v1.PodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *podsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(podsResource, c.ClusterPath, c.Namespace, opts))

}

func (c *podsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.Pod, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}

func (c *podsClient) Apply(ctx context.Context, applyConfiguration *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*v1.Pod, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1.Pod{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Pod), err
}

func (c *podsClient) ApplyStatus(ctx context.Context, applyConfiguration *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (*v1.Pod, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(podsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &v1.Pod{})

	return obj.(*v1.Pod), err
}
