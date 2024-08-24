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
	v1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamextensionsv1beta1client "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

// DaemonSetsClusterGetter has a method to return a DaemonSetClusterInterface.
// A group's client should implement this interface.
type DaemonSetsClusterGetter interface {
	DaemonSets() DaemonSetClusterInterface
}

// DaemonSetClusterInterface has methods to work with DaemonSet resources.
type DaemonSetClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DaemonSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) DaemonSetNamespacer
	DaemonSetExpansion
}

type daemonSetsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamextensionsv1beta1client.ExtensionsV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *daemonSetsClusterInterface) Cluster(clusterPath logicalcluster.Path) DaemonSetNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &daemonSetsNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all DaemonSets that are available in all clusters.
func (c *daemonSetsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.DaemonSetList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).DaemonSets(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all DaemonSets across all clusters.
func (c *daemonSetsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).DaemonSets(metav1.NamespaceAll).Watch(ctx, opts)
}

// DaemonSetNamespacer can scope to objects within a namespace, returning a DaemonSetInterface.
type DaemonSetNamespacer interface {
	Namespace(name string) upstreamextensionsv1beta1client.DaemonSetInterface
}

type daemonSetsNamespacer struct {
	clientCache kcpclient.Cache[*upstreamextensionsv1beta1client.ExtensionsV1beta1Client]
	clusterPath logicalcluster.Path
}

func (n *daemonSetsNamespacer) Namespace(namespace string) upstreamextensionsv1beta1client.DaemonSetInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).DaemonSets(namespace)
}
