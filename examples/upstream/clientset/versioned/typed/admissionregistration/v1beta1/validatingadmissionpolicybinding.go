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

	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	admissionregistrationv1beta1client "k8s.io/client-go/kuberentes/typed/admissionregistration/v1beta1"
	admissionregistrationv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/admissionregistration/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ValidatingAdmissionPolicyBindingsClusterGetter has a method to return a ValidatingAdmissionPolicyBindingClusterInterface.
// A group's client should implement this interface.
type ValidatingAdmissionPolicyBindingsClusterGetter interface {
	ValidatingAdmissionPolicyBindings() ValidatingAdmissionPolicyBindingClusterInterface
}

// ValidatingAdmissionPolicyBindingInterface has methods to work with ValidatingAdmissionPolicyBinding resources.
type ValidatingAdmissionPolicyBindingClusterInterface interface {
	Cluster(logicalcluster.Path) admissionregistrationv1beta1client.ValidatingAdmissionPolicyBindingInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ValidatingAdmissionPolicyBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ValidatingAdmissionPolicyBindingExpansion
}

// validatingAdmissionPolicyBindings implements ValidatingAdmissionPolicyBindingInterface
type validatingAdmissionPolicyBindings struct {
	*gentype.ClientWithListAndApply[*v1beta1.ValidatingAdmissionPolicyBinding, *v1beta1.ValidatingAdmissionPolicyBindingList, *admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingApplyConfiguration]
}

// newValidatingAdmissionPolicyBindings returns a ValidatingAdmissionPolicyBindings
func newValidatingAdmissionPolicyBindings(c *AdmissionregistrationV1beta1Client) *validatingAdmissionPolicyBindings {
	return &validatingAdmissionPolicyBindings{
		gentype.NewClientWithListAndApply[*v1beta1.ValidatingAdmissionPolicyBinding, *v1beta1.ValidatingAdmissionPolicyBindingList, *admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingApplyConfiguration](
			"validatingadmissionpolicybindings",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.ValidatingAdmissionPolicyBinding { return &v1beta1.ValidatingAdmissionPolicyBinding{} },
			func() *v1beta1.ValidatingAdmissionPolicyBindingList {
				return &v1beta1.ValidatingAdmissionPolicyBindingList{}
			}),
	}
}
