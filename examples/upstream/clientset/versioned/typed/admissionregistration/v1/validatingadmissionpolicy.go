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
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	admissionregistrationv1client "k8s.io/client-go/kuberentes/typed/admissionregistration/v1"
	applyconfigurationadmissionregistrationv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/admissionregistration/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ValidatingAdmissionPoliciesClusterGetter has a method to return a ValidatingAdmissionPolicyClusterInterface.
// A group's client should implement this interface.
type ValidatingAdmissionPoliciesClusterGetter interface {
	ValidatingAdmissionPolicies() ValidatingAdmissionPolicyClusterInterface
}

// ValidatingAdmissionPolicyInterface has methods to work with ValidatingAdmissionPolicy resources.
type ValidatingAdmissionPolicyClusterInterface interface {
	Cluster(logicalcluster.Path) admissionregistrationv1client.ValidatingAdmissionPolicyInterface

	List(ctx context.Context, opts v1.ListOptions) (*admissionregistrationv1.ValidatingAdmissionPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ValidatingAdmissionPolicyExpansion
}

// validatingAdmissionPolicies implements ValidatingAdmissionPolicyInterface
type validatingAdmissionPolicies struct {
	*gentype.ClientWithListAndApply[*admissionregistrationv1.ValidatingAdmissionPolicy, *admissionregistrationv1.ValidatingAdmissionPolicyList, *applyconfigurationadmissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration]
}

// newValidatingAdmissionPolicies returns a ValidatingAdmissionPolicies
func newValidatingAdmissionPolicies(c *AdmissionregistrationV1Client) *validatingAdmissionPolicies {
	return &validatingAdmissionPolicies{
		gentype.NewClientWithListAndApply[*admissionregistrationv1.ValidatingAdmissionPolicy, *admissionregistrationv1.ValidatingAdmissionPolicyList, *applyconfigurationadmissionregistrationv1.ValidatingAdmissionPolicyApplyConfiguration](
			"validatingadmissionpolicies",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *admissionregistrationv1.ValidatingAdmissionPolicy {
				return &admissionregistrationv1.ValidatingAdmissionPolicy{}
			},
			func() *admissionregistrationv1.ValidatingAdmissionPolicyList {
				return &admissionregistrationv1.ValidatingAdmissionPolicyList{}
			}),
	}
}
