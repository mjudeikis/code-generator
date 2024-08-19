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

// NamespacesClusterGetter has a method to return a NamespaceClusterInterface.
// A group's client should implement this interface.
type NamespacesClusterGetter interface {
	Namespaces() NamespaceClusterInterface
}

// NamespaceInterface has methods to work with Namespace resources.
type NamespaceClusterInterface interface {
	Cluster(logicalcluster.Path) corev1client.NamespaceInterface

	List(ctx context.Context, opts v1.ListOptions) (*corev1.NamespaceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	NamespaceExpansion
}

// namespaces implements NamespaceInterface
type namespaces struct {
	*gentype.ClientWithListAndApply[*corev1.Namespace, *corev1.NamespaceList, *applyconfigurationcorev1.NamespaceApplyConfiguration]
}

// newNamespaces returns a Namespaces
func newNamespaces(c *CoreV1Client) *namespaces {
	return &namespaces{
		gentype.NewClientWithListAndApply[*corev1.Namespace, *corev1.NamespaceList, *applyconfigurationcorev1.NamespaceApplyConfiguration](
			"namespaces",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *corev1.Namespace { return &corev1.Namespace{} },
			func() *corev1.NamespaceList { return &corev1.NamespaceList{} }),
	}
}
