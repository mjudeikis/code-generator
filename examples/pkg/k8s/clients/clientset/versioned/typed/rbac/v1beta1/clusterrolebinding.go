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

package v1beta1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	rbacv1beta1 "acme.corp/pkg/apis/rbac/v1beta1"
	rbacv1beta1client "acme.corp/pkg/generated/clientset/versioned/typed/rbac/v1beta1"
)

// ClusterRoleBindingsClusterGetter has a method to return a ClusterRoleBindingClusterInterface.
// A group's cluster client should implement this interface.
type ClusterRoleBindingsClusterGetter interface {
	ClusterRoleBindings() ClusterRoleBindingClusterInterface
}

// ClusterRoleBindingClusterInterface can operate on ClusterRoleBindings across all clusters,
// or scope down to one cluster and return a rbacv1beta1client.ClusterRoleBindingInterface.
type ClusterRoleBindingClusterInterface interface {
	Cluster(logicalcluster.Path) rbacv1beta1client.ClusterRoleBindingInterface
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1beta1.ClusterRoleBindingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type clusterRoleBindingsClusterInterface struct {
	clientCache kcpclient.Cache[*rbacv1beta1client.RbacV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterRoleBindingsClusterInterface) Cluster(clusterPath logicalcluster.Path) rbacv1beta1client.ClusterRoleBindingInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ClusterRoleBindings()
}

// List returns the entire collection of all ClusterRoleBindings across all clusters.
func (c *clusterRoleBindingsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1beta1.ClusterRoleBindingList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterRoleBindings().List(ctx, opts)
}

// Watch begins to watch all ClusterRoleBindings across all clusters.
func (c *clusterRoleBindingsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterRoleBindings().Watch(ctx, opts)
}
