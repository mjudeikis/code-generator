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

package v1beta1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/rbac/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	rbacv1beta1client "k8s.io/client-go/kuberentes/typed/rbac/v1beta1"
	rbacv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/rbac/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ClusterRoleBindingsClusterGetter has a method to return a ClusterRoleBindingClusterInterface.
// A group's client should implement this interface.
type ClusterRoleBindingsClusterGetter interface {
	ClusterRoleBindings() ClusterRoleBindingClusterInterface
}

// ClusterRoleBindingInterface has methods to work with ClusterRoleBinding resources.
type ClusterRoleBindingClusterInterface interface {
	Cluster(logicalcluster.Path) rbacv1beta1client.ClusterRoleBindingInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterRoleBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ClusterRoleBindingExpansion
}

// clusterRoleBindings implements ClusterRoleBindingInterface
type clusterRoleBindings struct {
	*gentype.ClientWithListAndApply[*v1beta1.ClusterRoleBinding, *v1beta1.ClusterRoleBindingList, *rbacv1beta1.ClusterRoleBindingApplyConfiguration]
}

// newClusterRoleBindings returns a ClusterRoleBindings
func newClusterRoleBindings(c *RbacV1beta1Client) *clusterRoleBindings {
	return &clusterRoleBindings{
		gentype.NewClientWithListAndApply[*v1beta1.ClusterRoleBinding, *v1beta1.ClusterRoleBindingList, *rbacv1beta1.ClusterRoleBindingApplyConfiguration](
			"clusterrolebindings",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.ClusterRoleBinding { return &v1beta1.ClusterRoleBinding{} },
			func() *v1beta1.ClusterRoleBindingList { return &v1beta1.ClusterRoleBindingList{} }),
	}
}
