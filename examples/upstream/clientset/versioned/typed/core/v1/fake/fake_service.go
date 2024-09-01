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

var servicesResource = v1.SchemeGroupVersion.WithResource("services")

var servicesKind = v1.SchemeGroupVersion.WithKind("Service")

// servicesClusterClient implements serviceInterface
type servicesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *servicesClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.ServiceNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &servicesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Services that match those selectors.
func (c *servicesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ServiceList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(servicesResource, servicesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.ServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ServiceList{ListMeta: obj.(*v1.ServiceList).ListMeta}
	for _, item := range obj.(*v1.ServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested services across all clusters.
func (c *servicesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(servicesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type servicesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *servicesNamespacer) Namespace(namespace string) upstreamcorev1client.ServiceInterface {
	return &servicesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type servicesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *servicesClient) Create(ctx context.Context, service *v1.Service, opts metav1.CreateOptions) (*v1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(servicesResource, c.ClusterPath, c.Namespace, service), &v1.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Service), err
}

func (c *servicesClient) Update(ctx context.Context, service *v1.Service, opts metav1.UpdateOptions) (*v1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(servicesResource, c.ClusterPath, c.Namespace, service), &v1.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Service), err
}

func (c *servicesClient) UpdateStatus(ctx context.Context, service *v1.Service, opts metav1.UpdateOptions) (*v1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(servicesResource, c.ClusterPath, "status", c.Namespace, service), &v1.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Service), err
}

func (c *servicesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(servicesResource, c.ClusterPath, c.Namespace, name, opts), &v1.Service{})

	return err
}

func (c *servicesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(servicesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ServiceList{})
	return err
}

func (c *servicesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(servicesResource, c.ClusterPath, c.Namespace, name), &v1.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Service), err
}

func (c *servicesClient) List(ctx context.Context, opts metav1.ListOptions) (*v1.ServiceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(servicesResource, servicesKind, c.ClusterPath, c.Namespace, opts), &v1.ServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ServiceList{ListMeta: obj.(*v1.ServiceList).ListMeta}
	for _, item := range obj.(*v1.ServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *servicesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(servicesResource, c.ClusterPath, c.Namespace, opts))

}

func (c *servicesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.Service, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(servicesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Service), err
}

func (c *servicesClient) Apply(ctx context.Context, applyConfiguration *corev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (*v1.Service, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(servicesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1.Service{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Service), err
}

func (c *servicesClient) ApplyStatus(ctx context.Context, applyConfiguration *corev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (*v1.Service, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(servicesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &v1.Service{})

	return obj.(*v1.Service), err
}
