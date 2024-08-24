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

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha1 "k8s.io/api/node/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamnodev1alpha1client "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
)

// RuntimeClassesClusterGetter has a method to return a RuntimeClassClusterInterface.
// A group's client should implement this interface.
type RuntimeClassesClusterGetter interface {
	RuntimeClasses() RuntimeClassClusterInterface
}

// RuntimeClassClusterInterface has methods to work with RuntimeClass resources.
type RuntimeClassClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RuntimeClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamnodev1alpha1client.RuntimeClassInterface
	RuntimeClassExpansion
}

type runtimeClassesClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamnodev1alpha1client.NodeV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *runtimeClassesClusterInterface) Cluster(clusterPath logicalcluster.Path) upstreamnodev1alpha1client.RuntimeClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).RuntimeClasses()
}

// List returns the entire collection of all RuntimeClasses that are available in all clusters.
func (c *runtimeClassesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.RuntimeClassList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).RuntimeClasses().List(ctx, opts)
}

// Watch begins to watch all RuntimeClasses across all clusters.
func (c *runtimeClassesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).RuntimeClasses().Watch(ctx, opts)
}
