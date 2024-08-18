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

	discoveryv1 "acme.corp/pkg/apis/discovery/v1"
	discoveryv1client "acme.corp/pkg/generated/clientset/versioned/typed/discovery/v1"
)

// EndpointSlicesClusterGetter has a method to return a EndpointSliceClusterInterface.
// A group's cluster client should implement this interface.
type EndpointSlicesClusterGetter interface {
	EndpointSlices() EndpointSliceClusterInterface
}

// EndpointSliceClusterInterface can operate on EndpointSlices across all clusters,
// or scope down to one cluster and return a EndpointSlicesNamespacer.
type EndpointSliceClusterInterface interface {
	Cluster(logicalcluster.Path) EndpointSlicesNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*discoveryv1.EndpointSliceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type endpointSlicesClusterInterface struct {
	clientCache kcpclient.Cache[*discoveryv1client.DiscoveryV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *endpointSlicesClusterInterface) Cluster(clusterPath logicalcluster.Path) EndpointSlicesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &endpointSlicesNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all EndpointSlices across all clusters.
func (c *endpointSlicesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*discoveryv1.EndpointSliceList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).EndpointSlices(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all EndpointSlices across all clusters.
func (c *endpointSlicesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).EndpointSlices(metav1.NamespaceAll).Watch(ctx, opts)
}

// EndpointSlicesNamespacer can scope to objects within a namespace, returning a discoveryv1client.EndpointSliceInterface.
type EndpointSlicesNamespacer interface {
	Namespace(string) discoveryv1client.EndpointSliceInterface
}

type endpointSlicesNamespacer struct {
	clientCache kcpclient.Cache[*discoveryv1client.DiscoveryV1Client]
	clusterPath logicalcluster.Path
}

func (n *endpointSlicesNamespacer) Namespace(namespace string) discoveryv1client.EndpointSliceInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).EndpointSlices(namespace)
}
