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
	v1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	storagev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	upstreamstoragev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/testing"
)

var csidriversResource = v1beta1.SchemeGroupVersion.WithResource("csidrivers")

var csidriversKind = v1beta1.SchemeGroupVersion.WithKind("CSIDriver")

// cSIDriversClusterClient implements cSIDriverInterface
type cSIDriversClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cSIDriversClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamstoragev1beta1client.CSIDriverInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cSIDriversClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors.
func (c *cSIDriversClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CSIDriverList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(csidriversResource, csidriversKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.CSIDriverList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CSIDriverList{ListMeta: obj.(*v1beta1.CSIDriverList).ListMeta}
	for _, item := range obj.(*v1beta1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIDrivers across all clusters.
func (c *cSIDriversClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(csidriversResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type cSIDriversClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *cSIDriversClient) Create(ctx context.Context, cSIDriver *v1beta1.CSIDriver, opts metav1.CreateOptions) (*v1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(csidriversResource, c.ClusterPath, cSIDriver), &v1beta1.CSIDriver{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

func (c *cSIDriversClient) Update(ctx context.Context, cSIDriver *v1beta1.CSIDriver, opts metav1.UpdateOptions) (*v1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(csidriversResource, c.ClusterPath, cSIDriver), &v1beta1.CSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

func (c *cSIDriversClient) UpdateStatus(ctx context.Context, cSIDriver *v1beta1.CSIDriver, opts metav1.UpdateOptions) (*v1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(csidriversResource, c.ClusterPath, "status", cSIDriver), &v1beta1.CSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

func (c *cSIDriversClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(csidriversResource, c.ClusterPath, name, opts), &v1beta1.CSIDriver{})

	return err
}

func (c *cSIDriversClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(csidriversResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CSIDriverList{})
	return err
}

func (c *cSIDriversClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(csidriversResource, c.ClusterPath, name), &v1beta1.CSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

func (c *cSIDriversClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.CSIDriverList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(csidriversResource, csidriversKind, c.ClusterPath, opts), &v1beta1.CSIDriverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CSIDriverList{ListMeta: obj.(*v1beta1.CSIDriverList).ListMeta}
	for _, item := range obj.(*v1beta1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cSIDriversClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(csidriversResource, c.ClusterPath, opts))

}

func (c *cSIDriversClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.CSIDriver, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(csidriversResource, c.ClusterPath, name, pt, data, subresources...), &v1beta1.CSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

func (c *cSIDriversClient) Apply(ctx context.Context, applyConfiguration *storagev1beta1.CSIDriverApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.CSIDriver, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(csidriversResource, c.ClusterPath, *name, types.ApplyPatchType, data), &v1beta1.CSIDriver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CSIDriver), err
}

func (c *cSIDriversClient) ApplyStatus(ctx context.Context, applyConfiguration *storagev1beta1.CSIDriverApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.CSIDriver, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(csidriversResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &v1beta1.CSIDriver{})

	return obj.(*v1beta1.CSIDriver), err
}
