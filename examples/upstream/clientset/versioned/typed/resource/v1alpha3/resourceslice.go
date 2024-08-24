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

package v1alpha3

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha3 "k8s.io/api/resource/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamresourcev1alpha3client "k8s.io/client-go/kubernetes/typed/resource/v1alpha3"
)

// ResourceSlicesClusterGetter has a method to return a ResourceSliceClusterInterface.
// A group's client should implement this interface.
type ResourceSlicesClusterGetter interface {
	ResourceSlices() ResourceSliceClusterInterface
}

// ResourceSliceClusterInterface has methods to work with ResourceSlice resources.
type ResourceSliceClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.ResourceSliceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamresourcev1alpha3client.ResourceSliceInterface
	ResourceSliceExpansion
}

type resourceSlicesClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamresourcev1alpha3client.ResourceV1alpha3Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *resourceSlicesClusterInterface) Cluster(clusterPath logicalcluster.Path) upstreamresourcev1alpha3client.ResourceSliceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ResourceSlices()
}

// List returns the entire collection of all ResourceSlices that are available in all clusters.
func (c *resourceSlicesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1alpha3.ResourceSliceList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ResourceSlices().List(ctx, opts)
}

// Watch begins to watch all ResourceSlices across all clusters.
func (c *resourceSlicesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ResourceSlices().Watch(ctx, opts)
}
