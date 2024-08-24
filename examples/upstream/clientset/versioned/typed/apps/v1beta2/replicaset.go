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

package v1beta2

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta2 "k8s.io/api/apps/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamappsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
)

// ReplicaSetsClusterGetter has a method to return a ReplicaSetClusterInterface.
// A group's client should implement this interface.
type ReplicaSetsClusterGetter interface {
	ReplicaSets() ReplicaSetClusterInterface
}

// ReplicaSetClusterInterface has methods to work with ReplicaSet resources.
type ReplicaSetClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.ReplicaSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) ReplicaSetNamespacer
	ReplicaSetExpansion
}

type replicaSetsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamappsv1beta2client.AppsV1beta2Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *replicaSetsClusterInterface) Cluster(clusterPath logicalcluster.Path) ReplicaSetNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &replicaSetsNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all ReplicaSets that are available in all clusters.
func (c *replicaSetsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1beta2.ReplicaSetList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ReplicaSets(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all ReplicaSets across all clusters.
func (c *replicaSetsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ReplicaSets(metav1.NamespaceAll).Watch(ctx, opts)
}

// ReplicaSetNamespacer can scope to objects within a namespace, returning a ReplicaSetInterface.
type ReplicaSetNamespacer interface {
	Namespace(name string) upstreamappsv1beta2client.ReplicaSetInterface
}

type replicaSetsNamespacer struct {
	clientCache kcpclient.Cache[*upstreamappsv1beta2client.AppsV1beta2Client]
	clusterPath logicalcluster.Path
}

func (n *replicaSetsNamespacer) Namespace(namespace string) upstreamappsv1beta2client.ReplicaSetInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).ReplicaSets(namespace)
}
