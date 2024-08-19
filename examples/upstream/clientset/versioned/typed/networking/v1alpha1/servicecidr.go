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
	v1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	networkingv1alpha1client "k8s.io/client-go/kuberentes/typed/networking/v1alpha1"
	networkingv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/networking/v1alpha1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ServiceCIDRsClusterGetter has a method to return a ServiceCIDRClusterInterface.
// A group's client should implement this interface.
type ServiceCIDRsClusterGetter interface {
	ServiceCIDRs() ServiceCIDRClusterInterface
}

// ServiceCIDRInterface has methods to work with ServiceCIDR resources.
type ServiceCIDRClusterInterface interface {
	Cluster(logicalcluster.Path) networkingv1alpha1client.ServiceCIDRInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceCIDRList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ServiceCIDRExpansion
}

// serviceCIDRs implements ServiceCIDRInterface
type serviceCIDRs struct {
	*gentype.ClientWithListAndApply[*v1alpha1.ServiceCIDR, *v1alpha1.ServiceCIDRList, *networkingv1alpha1.ServiceCIDRApplyConfiguration]
}

// newServiceCIDRs returns a ServiceCIDRs
func newServiceCIDRs(c *NetworkingV1alpha1Client) *serviceCIDRs {
	return &serviceCIDRs{
		gentype.NewClientWithListAndApply[*v1alpha1.ServiceCIDR, *v1alpha1.ServiceCIDRList, *networkingv1alpha1.ServiceCIDRApplyConfiguration](
			"servicecidrs",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1alpha1.ServiceCIDR { return &v1alpha1.ServiceCIDR{} },
			func() *v1alpha1.ServiceCIDRList { return &v1alpha1.ServiceCIDRList{} }),
	}
}
