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

package v1alpha1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	rbacv1alpha1client "k8s.io/client-go/kuberentes/typed/rbac/v1alpha1"
	rbacv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/rbac/v1alpha1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// RolesClusterGetter has a method to return a RoleClusterInterface.
// A group's client should implement this interface.
type RolesClusterGetter interface {
	Roles(namespace string) RoleClusterInterface
}

// RoleInterface has methods to work with Role resources.
type RoleClusterInterface interface {
	Cluster(logicalcluster.Path) rbacv1alpha1client.RoleInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	RoleExpansion
}

// roles implements RoleInterface
type roles struct {
	*gentype.ClientWithListAndApply[*v1alpha1.Role, *v1alpha1.RoleList, *rbacv1alpha1.RoleApplyConfiguration]
}

// newRoles returns a Roles
func newRoles(c *RbacV1alpha1Client, namespace string) *roles {
	return &roles{
		gentype.NewClientWithListAndApply[*v1alpha1.Role, *v1alpha1.RoleList, *rbacv1alpha1.RoleApplyConfiguration](
			"roles",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Role { return &v1alpha1.Role{} },
			func() *v1alpha1.RoleList { return &v1alpha1.RoleList{} }),
	}
}
