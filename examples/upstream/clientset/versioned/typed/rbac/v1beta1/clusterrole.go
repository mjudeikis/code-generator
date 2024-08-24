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

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamrbacv1beta1client "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
)

// ClusterRolesClusterGetter has a method to return a ClusterRoleClusterInterface.
// A group's client should implement this interface.
type ClusterRolesClusterGetter interface {
	ClusterRoles() ClusterRoleClusterInterface
}

// ClusterRoleClusterInterface has methods to work with ClusterRole resources.
type ClusterRoleClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamrbacv1beta1client.ClusterRoleInterface
	ClusterRoleExpansion
}

type clusterRolesClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamrbacv1beta1client.RbacV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *clusterRolesClusterInterface) Cluster(clusterPath logicalcluster.Path) upstreamrbacv1beta1client.ClusterRoleInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ClusterRoles()
}

// List returns the entire collection of all ClusterRoles that are available in all clusters.
func (c *clusterRolesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.ClusterRoleList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterRoles().List(ctx, opts)
}

// Watch begins to watch all ClusterRoles across all clusters.
func (c *clusterRolesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ClusterRoles().Watch(ctx, opts)
}
