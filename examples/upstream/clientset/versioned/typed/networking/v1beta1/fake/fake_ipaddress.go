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
)

var ipaddressesResource = v1beta1.SchemeGroupVersion.WithResource("ipaddresses")

var ipaddressesKind = v1beta1.SchemeGroupVersion.WithKind("IPAddress")

// iPAddressesClusterClient implements iPAddressInterface
type iPAddressesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *iPAddressesClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamnetworkingv1beta1client.IPAddressInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &iPAddressesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of IPAddresses that match those selectors.
func (c *iPAddressesClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IPAddressList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(ipaddressesResource, ipaddressesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.IPAddressList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IPAddressList{ListMeta: obj.(*v1beta1.IPAddressList).ListMeta}
	for _, item := range obj.(*v1beta1.IPAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAddresss across all clusters.
func (c *iPAddressesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(ipaddressesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type iPAddressesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *iPAddressesClient) Create(ctx context.Context, iPAddress *v1beta1.IPAddress, opts metav1.CreateOptions) (*v1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(ipaddressesResource, c.ClusterPath, iPAddress), &v1beta1.IPAddress{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IPAddress), err
}

func (c *iPAddressesClient) Update(ctx context.Context, iPAddress *v1beta1.IPAddress, opts metav1.UpdateOptions) (*v1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(ipaddressesResource, c.ClusterPath, iPAddress), &v1beta1.IPAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IPAddress), err
}

func (c *iPAddressesClient) UpdateStatus(ctx context.Context, iPAddress *v1beta1.IPAddress, opts metav1.UpdateOptions) (*v1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(ipaddressesResource, c.ClusterPath, "status", iPAddress), &v1beta1.IPAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IPAddress), err
}

func (c *iPAddressesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(ipaddressesResource, c.ClusterPath, name, opts), &v1beta1.IPAddress{})

	return err
}

func (c *iPAddressesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(ipaddressesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IPAddressList{})
	return err
}

func (c *iPAddressesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(ipaddressesResource, c.ClusterPath, name), &v1beta1.IPAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IPAddress), err
}

func (c *iPAddressesClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.IPAddressList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(ipaddressesResource, ipaddressesKind, c.ClusterPath, opts), &v1beta1.IPAddressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IPAddressList{ListMeta: obj.(*v1beta1.IPAddressList).ListMeta}
	for _, item := range obj.(*v1beta1.IPAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *iPAddressesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(ipaddressesResource, c.ClusterPath, opts))

}

func (c *iPAddressesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.IPAddress, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(ipaddressesResource, c.ClusterPath, name, pt, data, subresources...), &v1beta1.IPAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IPAddress), err
}

func (c *iPAddressesClient) Apply(ctx context.Context, applyConfiguration *networkingv1beta1.IPAddressApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.IPAddress, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(ipaddressesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &v1beta1.IPAddress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IPAddress), err
}

func (c *iPAddressesClient) ApplyStatus(ctx context.Context, applyConfiguration *networkingv1beta1.IPAddressApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.IPAddress, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(ipaddressesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &v1beta1.IPAddress{})

	return obj.(*v1beta1.IPAddress), err
}
