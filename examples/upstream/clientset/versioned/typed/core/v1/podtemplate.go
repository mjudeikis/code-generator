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

	"github.com/kcp-dev/logicalcluster/v3"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	corev1client "k8s.io/client-go/kuberentes/typed/core/v1"
	applyconfigurationcorev1 "k8s.io/code-generator/examples/upstream/applyconfiguration/core/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// PodTemplatesClusterGetter has a method to return a PodTemplateClusterInterface.
// A group's client should implement this interface.
type PodTemplatesClusterGetter interface {
	PodTemplates(namespace string) PodTemplateClusterInterface
}

// PodTemplateInterface has methods to work with PodTemplate resources.
type PodTemplateClusterInterface interface {
	Cluster(logicalcluster.Path) corev1client.PodTemplateInterface

	List(ctx context.Context, opts v1.ListOptions) (*corev1.PodTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	PodTemplateExpansion
}

// podTemplates implements PodTemplateInterface
type podTemplates struct {
	*gentype.ClientWithListAndApply[*corev1.PodTemplate, *corev1.PodTemplateList, *applyconfigurationcorev1.PodTemplateApplyConfiguration]
}

// newPodTemplates returns a PodTemplates
func newPodTemplates(c *CoreV1Client, namespace string) *podTemplates {
	return &podTemplates{
		gentype.NewClientWithListAndApply[*corev1.PodTemplate, *corev1.PodTemplateList, *applyconfigurationcorev1.PodTemplateApplyConfiguration](
			"podtemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.PodTemplate { return &corev1.PodTemplate{} },
			func() *corev1.PodTemplateList { return &corev1.PodTemplateList{} }),
	}
}
