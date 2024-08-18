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

	rbacv1beta1 "acme.corp/pkg/apis/rbac/v1beta1"
	applyconfigurationsrbacv1beta1 "acme.corp/pkg/generated/applyconfigurations/rbac/v1beta1"
	rbacv1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/rbac/v1beta1"
)

var clusterRolesResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1beta1", Resource: "clusterroles"}
var clusterRolesKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1beta1", Kind: "ClusterRole"}

type clusterRolesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterRolesClusterClient) Cluster(clusterPath logicalcluster.Path) rbacv1beta1client.ClusterRoleInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &clusterRolesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ClusterRoles that match those selectors across all clusters.
func (c *clusterRolesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1beta1.ClusterRoleList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterRolesResource, clusterRolesKind, logicalcluster.Wildcard, opts), &rbacv1beta1.ClusterRoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1beta1.ClusterRoleList{ListMeta: obj.(*rbacv1beta1.ClusterRoleList).ListMeta}
	for _, item := range obj.(*rbacv1beta1.ClusterRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ClusterRoles across all clusters.
func (c *clusterRolesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterRolesResource, logicalcluster.Wildcard, opts))
}

type clusterRolesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *clusterRolesClient) Create(ctx context.Context, clusterRole *rbacv1beta1.ClusterRole, opts metav1.CreateOptions) (*rbacv1beta1.ClusterRole, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(clusterRolesResource, c.ClusterPath, clusterRole), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}

func (c *clusterRolesClient) Update(ctx context.Context, clusterRole *rbacv1beta1.ClusterRole, opts metav1.UpdateOptions) (*rbacv1beta1.ClusterRole, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(clusterRolesResource, c.ClusterPath, clusterRole), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}

func (c *clusterRolesClient) UpdateStatus(ctx context.Context, clusterRole *rbacv1beta1.ClusterRole, opts metav1.UpdateOptions) (*rbacv1beta1.ClusterRole, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(clusterRolesResource, c.ClusterPath, "status", clusterRole), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}

func (c *clusterRolesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(clusterRolesResource, c.ClusterPath, name, opts), &rbacv1beta1.ClusterRole{})
	return err
}

func (c *clusterRolesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(clusterRolesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &rbacv1beta1.ClusterRoleList{})
	return err
}

func (c *clusterRolesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*rbacv1beta1.ClusterRole, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(clusterRolesResource, c.ClusterPath, name), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}

// List takes label and field selectors, and returns the list of ClusterRoles that match those selectors.
func (c *clusterRolesClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1beta1.ClusterRoleList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(clusterRolesResource, clusterRolesKind, c.ClusterPath, opts), &rbacv1beta1.ClusterRoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1beta1.ClusterRoleList{ListMeta: obj.(*rbacv1beta1.ClusterRoleList).ListMeta}
	for _, item := range obj.(*rbacv1beta1.ClusterRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *clusterRolesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(clusterRolesResource, c.ClusterPath, opts))
}

func (c *clusterRolesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1beta1.ClusterRole, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterRolesResource, c.ClusterPath, name, pt, data, subresources...), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}

func (c *clusterRolesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1beta1.ClusterRoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1beta1.ClusterRole, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterRolesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}

func (c *clusterRolesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1beta1.ClusterRoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1beta1.ClusterRole, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(clusterRolesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &rbacv1beta1.ClusterRole{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1beta1.ClusterRole), err
}
