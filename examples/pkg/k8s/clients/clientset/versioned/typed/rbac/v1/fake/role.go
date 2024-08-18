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

	rbacv1 "acme.corp/pkg/apis/rbac/v1"
	applyconfigurationsrbacv1 "acme.corp/pkg/generated/applyconfigurations/rbac/v1"
	rbacv1client "acme.corp/pkg/generated/clientset/versioned/typed/rbac/v1"
	kcprbacv1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/rbac/v1"
)

var rolesResource = schema.GroupVersionResource{Group: "rbac.authorization.k8s.io", Version: "v1", Resource: "roles"}
var rolesKind = schema.GroupVersionKind{Group: "rbac.authorization.k8s.io", Version: "v1", Kind: "Role"}

type rolesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *rolesClusterClient) Cluster(clusterPath logicalcluster.Path) kcprbacv1.RolesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &rolesNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Roles that match those selectors across all clusters.
func (c *rolesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(rolesResource, rolesKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &rbacv1.RoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1.RoleList{ListMeta: obj.(*rbacv1.RoleList).ListMeta}
	for _, item := range obj.(*rbacv1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Roles across all clusters.
func (c *rolesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(rolesResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type rolesNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *rolesNamespacer) Namespace(namespace string) rbacv1client.RoleInterface {
	return &rolesClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type rolesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *rolesClient) Create(ctx context.Context, role *rbacv1.Role, opts metav1.CreateOptions) (*rbacv1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(rolesResource, c.ClusterPath, c.Namespace, role), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}

func (c *rolesClient) Update(ctx context.Context, role *rbacv1.Role, opts metav1.UpdateOptions) (*rbacv1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(rolesResource, c.ClusterPath, c.Namespace, role), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}

func (c *rolesClient) UpdateStatus(ctx context.Context, role *rbacv1.Role, opts metav1.UpdateOptions) (*rbacv1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(rolesResource, c.ClusterPath, "status", c.Namespace, role), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}

func (c *rolesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(rolesResource, c.ClusterPath, c.Namespace, name, opts), &rbacv1.Role{})
	return err
}

func (c *rolesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(rolesResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &rbacv1.RoleList{})
	return err
}

func (c *rolesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*rbacv1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(rolesResource, c.ClusterPath, c.Namespace, name), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}

// List takes label and field selectors, and returns the list of Roles that match those selectors.
func (c *rolesClient) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(rolesResource, rolesKind, c.ClusterPath, c.Namespace, opts), &rbacv1.RoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rbacv1.RoleList{ListMeta: obj.(*rbacv1.RoleList).ListMeta}
	for _, item := range obj.(*rbacv1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *rolesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(rolesResource, c.ClusterPath, c.Namespace, opts))
}

func (c *rolesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*rbacv1.Role, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolesResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}

func (c *rolesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1.RoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1.Role, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}

func (c *rolesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsrbacv1.RoleApplyConfiguration, opts metav1.ApplyOptions) (*rbacv1.Role, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(rolesResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &rbacv1.Role{})
	if obj == nil {
		return nil, err
	}
	return obj.(*rbacv1.Role), err
}
