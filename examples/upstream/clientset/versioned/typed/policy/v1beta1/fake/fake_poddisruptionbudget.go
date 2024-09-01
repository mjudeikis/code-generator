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
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	v1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	policyv1beta1 "k8s.io/client-go/applyconfigurations/policy/v1beta1"
	upstreampolicyv1beta1client "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/policy/v1beta1"
)

var poddisruptionbudgetsResource = v1beta1.SchemeGroupVersion.WithResource("poddisruptionbudgets")

var poddisruptionbudgetsKind = v1beta1.SchemeGroupVersion.WithKind("PodDisruptionBudget")

// podDisruptionBudgetsClusterClient implements podDisruptionBudgetInterface
type podDisruptionBudgetsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *podDisruptionBudgetsClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.PodDisruptionBudgetNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &podDisruptionBudgetsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PodDisruptionBudgets that match those selectors.
func (c *podDisruptionBudgetsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PodDisruptionBudgetList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(poddisruptionbudgetsResource, poddisruptionbudgetsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.PodDisruptionBudgetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PodDisruptionBudgetList{ListMeta: obj.(*v1beta1.PodDisruptionBudgetList).ListMeta}
	for _, item := range obj.(*v1beta1.PodDisruptionBudgetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podDisruptionBudgets across all clusters.
func (c *podDisruptionBudgetsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(poddisruptionbudgetsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type podDisruptionBudgetsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *podDisruptionBudgetsNamespacer) Namespace(namespace string) upstreampolicyv1beta1client.PodDisruptionBudgetInterface {
	return &podDisruptionBudgetsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type podDisruptionBudgetsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *podDisruptionBudgetsClient) Create(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts metav1.CreateOptions) (*v1beta1.PodDisruptionBudget, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, podDisruptionBudget), &v1beta1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodDisruptionBudget), err
}

func (c *podDisruptionBudgetsClient) Update(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts metav1.UpdateOptions) (*v1beta1.PodDisruptionBudget, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, podDisruptionBudget), &v1beta1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodDisruptionBudget), err
}

func (c *podDisruptionBudgetsClient) UpdateStatus(ctx context.Context, podDisruptionBudget *v1beta1.PodDisruptionBudget, opts metav1.UpdateOptions) (*v1beta1.PodDisruptionBudget, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(poddisruptionbudgetsResource, c.ClusterPath, "status", c.Namespace, podDisruptionBudget), &v1beta1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodDisruptionBudget), err
}

func (c *podDisruptionBudgetsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, name, opts), &v1beta1.PodDisruptionBudget{})

	return err
}

func (c *podDisruptionBudgetsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PodDisruptionBudgetList{})
	return err
}

func (c *podDisruptionBudgetsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.PodDisruptionBudget, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, name), &v1beta1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodDisruptionBudget), err
}

func (c *podDisruptionBudgetsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.PodDisruptionBudgetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(poddisruptionbudgetsResource, poddisruptionbudgetsKind, c.ClusterPath, c.Namespace, opts), &v1beta1.PodDisruptionBudgetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PodDisruptionBudgetList{ListMeta: obj.(*v1beta1.PodDisruptionBudgetList).ListMeta}
	for _, item := range obj.(*v1beta1.PodDisruptionBudgetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *podDisruptionBudgetsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, opts))

}

func (c *podDisruptionBudgetsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.PodDisruptionBudget, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodDisruptionBudget), err
}

func (c *podDisruptionBudgetsClient) Apply(ctx context.Context, applyConfiguration *policyv1beta1.PodDisruptionBudgetApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.PodDisruptionBudget, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta1.PodDisruptionBudget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PodDisruptionBudget), err
}

func (c *podDisruptionBudgetsClient) ApplyStatus(ctx context.Context, applyConfiguration *policyv1beta1.PodDisruptionBudgetApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.PodDisruptionBudget, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(poddisruptionbudgetsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &v1beta1.PodDisruptionBudget{})

	return obj.(*v1beta1.PodDisruptionBudget), err
}
