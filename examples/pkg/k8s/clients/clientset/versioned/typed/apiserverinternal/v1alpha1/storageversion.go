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

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	internalv1alpha1 "acme.corp/pkg/apis/apiserverinternal/v1alpha1"
	internalv1alpha1client "acme.corp/pkg/generated/clientset/versioned/typed/apiserverinternal/v1alpha1"
)

// StorageVersionsClusterGetter has a method to return a StorageVersionClusterInterface.
// A group's cluster client should implement this interface.
type StorageVersionsClusterGetter interface {
	StorageVersions() StorageVersionClusterInterface
}

// StorageVersionClusterInterface can operate on StorageVersions across all clusters,
// or scope down to one cluster and return a internalv1alpha1client.StorageVersionInterface.
type StorageVersionClusterInterface interface {
	Cluster(logicalcluster.Path) internalv1alpha1client.StorageVersionInterface
	List(ctx context.Context, opts metav1.ListOptions) (*internalv1alpha1.StorageVersionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type storageVersionsClusterInterface struct {
	clientCache kcpclient.Cache[*internalv1alpha1client.InternalV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *storageVersionsClusterInterface) Cluster(clusterPath logicalcluster.Path) internalv1alpha1client.StorageVersionInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).StorageVersions()
}

// List returns the entire collection of all StorageVersions across all clusters.
func (c *storageVersionsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*internalv1alpha1.StorageVersionList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).StorageVersions().List(ctx, opts)
}

// Watch begins to watch all StorageVersions across all clusters.
func (c *storageVersionsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).StorageVersions().Watch(ctx, opts)
}
