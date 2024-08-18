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
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	appsv1 "acme.corp/pkg/apis/apps/v1"
	applyconfigurationsappsv1 "acme.corp/pkg/generated/applyconfigurations/apps/v1"
	applyconfigurationsautoscalingv1 "acme.corp/pkg/generated/applyconfigurations/autoscaling/v1"
	appsv1client "acme.corp/pkg/generated/clientset/versioned/typed/apps/v1"
	kcpappsv1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/apps/v1"
)

var deploymentsResource = schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"}
var deploymentsKind = schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}

type deploymentsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *deploymentsClusterClient) Cluster(clusterPath logicalcluster.Path) kcpappsv1.DeploymentsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &deploymentsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors across all clusters.
func (c *deploymentsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(deploymentsResource, deploymentsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &appsv1.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.DeploymentList{ListMeta: obj.(*appsv1.DeploymentList).ListMeta}
	for _, item := range obj.(*appsv1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Deployments across all clusters.
func (c *deploymentsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(deploymentsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type deploymentsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *deploymentsNamespacer) Namespace(namespace string) appsv1client.DeploymentInterface {
	return &deploymentsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type deploymentsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *deploymentsClient) Create(ctx context.Context, deployment *appsv1.Deployment, opts metav1.CreateOptions) (*appsv1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(deploymentsResource, c.ClusterPath, c.Namespace, deployment), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

func (c *deploymentsClient) Update(ctx context.Context, deployment *appsv1.Deployment, opts metav1.UpdateOptions) (*appsv1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(deploymentsResource, c.ClusterPath, c.Namespace, deployment), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

func (c *deploymentsClient) UpdateStatus(ctx context.Context, deployment *appsv1.Deployment, opts metav1.UpdateOptions) (*appsv1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(deploymentsResource, c.ClusterPath, "status", c.Namespace, deployment), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

func (c *deploymentsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(deploymentsResource, c.ClusterPath, c.Namespace, name, opts), &appsv1.Deployment{})
	return err
}

func (c *deploymentsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(deploymentsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1.DeploymentList{})
	return err
}

func (c *deploymentsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*appsv1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(deploymentsResource, c.ClusterPath, c.Namespace, name), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deploymentsClient) List(ctx context.Context, opts metav1.ListOptions) (*appsv1.DeploymentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(deploymentsResource, deploymentsKind, c.ClusterPath, c.Namespace, opts), &appsv1.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.DeploymentList{ListMeta: obj.(*appsv1.DeploymentList).ListMeta}
	for _, item := range obj.(*appsv1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *deploymentsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(deploymentsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *deploymentsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*appsv1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

func (c *deploymentsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsappsv1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.Deployment, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

func (c *deploymentsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsappsv1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*appsv1.Deployment, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &appsv1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.Deployment), err
}

func (c *deploymentsClient) GetScale(ctx context.Context, deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetSubresourceAction(deploymentsResource, c.ClusterPath, "scale", c.Namespace, deploymentName), &autoscalingv1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}

func (c *deploymentsClient) UpdateScale(ctx context.Context, deploymentName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(deploymentsResource, c.ClusterPath, "scale", c.Namespace, scale), &autoscalingv1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}

func (c *deploymentsClient) ApplyScale(ctx context.Context, deploymentName string, applyConfiguration *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.Scale, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &autoscalingv1.Scale{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
