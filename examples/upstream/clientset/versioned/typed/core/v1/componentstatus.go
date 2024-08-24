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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamcorev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ComponentStatusesClusterGetter has a method to return a ComponentStatusClusterInterface.
// A group's client should implement this interface.
type ComponentStatusesClusterGetter interface {
	ComponentStatuses() ComponentStatusClusterInterface
}

// ComponentStatusClusterInterface has methods to work with ComponentStatus resources.
type ComponentStatusClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*corev1.ComponentStatusList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamcorev1client.ComponentStatusInterface
	ComponentStatusExpansion
}

type componentStatusesClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamcorev1client.CoreV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *componentStatusesClusterInterface) Cluster(clusterPath logicalcluster.Path) upstreamcorev1client.ComponentStatusInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ComponentStatuses()
}

// List returns the entire collection of all ComponentStatuses that are available in all clusters.
func (c *componentStatusesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.ComponentStatusList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ComponentStatuses().List(ctx, opts)
}

// Watch begins to watch all ComponentStatuses across all clusters.
func (c *componentStatusesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ComponentStatuses().Watch(ctx, opts)
}
