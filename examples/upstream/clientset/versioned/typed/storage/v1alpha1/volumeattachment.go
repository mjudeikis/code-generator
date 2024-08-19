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

	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha1 "k8s.io/api/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	storagev1alpha1client "k8s.io/client-go/kuberentes/typed/storage/v1alpha1"
	storagev1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/storage/v1alpha1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// VolumeAttachmentsClusterGetter has a method to return a VolumeAttachmentClusterInterface.
// A group's client should implement this interface.
type VolumeAttachmentsClusterGetter interface {
	VolumeAttachments() VolumeAttachmentClusterInterface
}

// VolumeAttachmentInterface has methods to work with VolumeAttachment resources.
type VolumeAttachmentClusterInterface interface {
	Cluster(logicalcluster.Path) storagev1alpha1client.VolumeAttachmentInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VolumeAttachmentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	VolumeAttachmentExpansion
}

// volumeAttachments implements VolumeAttachmentInterface
type volumeAttachments struct {
	*gentype.ClientWithListAndApply[*v1alpha1.VolumeAttachment, *v1alpha1.VolumeAttachmentList, *storagev1alpha1.VolumeAttachmentApplyConfiguration]
}

// newVolumeAttachments returns a VolumeAttachments
func newVolumeAttachments(c *StorageV1alpha1Client) *volumeAttachments {
	return &volumeAttachments{
		gentype.NewClientWithListAndApply[*v1alpha1.VolumeAttachment, *v1alpha1.VolumeAttachmentList, *storagev1alpha1.VolumeAttachmentApplyConfiguration](
			"volumeattachments",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.VolumeAttachment { return &v1alpha1.VolumeAttachment{} },
			func() *v1alpha1.VolumeAttachmentList { return &v1alpha1.VolumeAttachmentList{} }),
	}
}
