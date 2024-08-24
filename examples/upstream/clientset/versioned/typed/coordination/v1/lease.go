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

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	coordinationv1 "k8s.io/api/coordination/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamcoordinationv1client "k8s.io/client-go/kubernetes/typed/coordination/v1"
)

// LeasesClusterGetter has a method to return a LeaseClusterInterface.
// A group's client should implement this interface.
type LeasesClusterGetter interface {
	Leases() LeaseClusterInterface
}

// LeaseClusterInterface has methods to work with Lease resources.
type LeaseClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*coordinationv1.LeaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) LeaseNamespacer
	LeaseExpansion
}

type leasesClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamcoordinationv1client.CoordinationV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *leasesClusterInterface) Cluster(clusterPath logicalcluster.Path) LeaseNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &leasesNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all Leases that are available in all clusters.
func (c *leasesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*coordinationv1.LeaseList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Leases(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all Leases across all clusters.
func (c *leasesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).Leases(metav1.NamespaceAll).Watch(ctx, opts)
}

// LeaseNamespacer can scope to objects within a namespace, returning a LeaseInterface.
type LeaseNamespacer interface {
	Namespace(name string) upstreamcoordinationv1client.LeaseInterface
}

type leasesNamespacer struct {
	clientCache kcpclient.Cache[*upstreamcoordinationv1client.CoordinationV1Client]
	clusterPath logicalcluster.Path
}

func (n *leasesNamespacer) Namespace(namespace string) upstreamcoordinationv1client.LeaseInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).Leases(namespace)
}
