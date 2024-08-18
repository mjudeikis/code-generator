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

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	rbacv1 "acme.corp/pkg/apis/rbac/v1"
	rbacv1client "acme.corp/pkg/generated/clientset/versioned/typed/rbac/v1"
)

// RoleBindingsClusterGetter has a method to return a RoleBindingClusterInterface.
// A group's cluster client should implement this interface.
type RoleBindingsClusterGetter interface {
	RoleBindings() RoleBindingClusterInterface
}

// RoleBindingClusterInterface can operate on RoleBindings across all clusters,
// or scope down to one cluster and return a RoleBindingsNamespacer.
type RoleBindingClusterInterface interface {
	Cluster(logicalcluster.Path) RoleBindingsNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleBindingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type roleBindingsClusterInterface struct {
	clientCache kcpclient.Cache[*rbacv1client.RbacV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *roleBindingsClusterInterface) Cluster(clusterPath logicalcluster.Path) RoleBindingsNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &roleBindingsNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all RoleBindings across all clusters.
func (c *roleBindingsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*rbacv1.RoleBindingList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).RoleBindings(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all RoleBindings across all clusters.
func (c *roleBindingsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).RoleBindings(metav1.NamespaceAll).Watch(ctx, opts)
}

// RoleBindingsNamespacer can scope to objects within a namespace, returning a rbacv1client.RoleBindingInterface.
type RoleBindingsNamespacer interface {
	Namespace(string) rbacv1client.RoleBindingInterface
}

type roleBindingsNamespacer struct {
	clientCache kcpclient.Cache[*rbacv1client.RbacV1Client]
	clusterPath logicalcluster.Path
}

func (n *roleBindingsNamespacer) Namespace(namespace string) rbacv1client.RoleBindingInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).RoleBindings(namespace)
}
