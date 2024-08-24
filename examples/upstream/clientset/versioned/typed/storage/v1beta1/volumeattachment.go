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
	v1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamstoragev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

// VolumeAttachmentsClusterGetter has a method to return a VolumeAttachmentClusterInterface.
// A group's client should implement this interface.
type VolumeAttachmentsClusterGetter interface {
	VolumeAttachments() VolumeAttachmentClusterInterface
}

// VolumeAttachmentClusterInterface has methods to work with VolumeAttachment resources.
type VolumeAttachmentClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeAttachmentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamstoragev1beta1client.VolumeAttachmentInterface
	VolumeAttachmentExpansion
}

type volumeAttachmentsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamstoragev1beta1client.StorageV1beta1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *volumeAttachmentsClusterInterface) Cluster(clusterPath logicalcluster.Path) upstreamstoragev1beta1client.VolumeAttachmentInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).VolumeAttachments()
}

// List returns the entire collection of all VolumeAttachments that are available in all clusters.
func (c *volumeAttachmentsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.VolumeAttachmentList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).VolumeAttachments().List(ctx, opts)
}

// Watch begins to watch all VolumeAttachments across all clusters.
func (c *volumeAttachmentsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).VolumeAttachments().Watch(ctx, opts)
}
