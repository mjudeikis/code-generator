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

	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	resourcev1alpha3client "k8s.io/client-go/kuberentes/typed/resource/v1alpha3"
	resourcev1alpha3 "k8s.io/code-generator/examples/upstream/applyconfiguration/resource/v1alpha3"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ResourceClaimsClusterGetter has a method to return a ResourceClaimClusterInterface.
// A group's client should implement this interface.
type ResourceClaimsClusterGetter interface {
	ResourceClaims(namespace string) ResourceClaimClusterInterface
}

// ResourceClaimInterface has methods to work with ResourceClaim resources.
type ResourceClaimClusterInterface interface {
	Cluster(logicalcluster.Path) resourcev1alpha3client.ResourceClaimInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.ResourceClaimList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	ResourceClaimExpansion
}

// resourceClaims implements ResourceClaimInterface
type resourceClaims struct {
	*gentype.ClientWithListAndApply[*v1alpha3.ResourceClaim, *v1alpha3.ResourceClaimList, *resourcev1alpha3.ResourceClaimApplyConfiguration]
}

// newResourceClaims returns a ResourceClaims
func newResourceClaims(c *ResourceV1alpha3Client, namespace string) *resourceClaims {
	return &resourceClaims{
		gentype.NewClientWithListAndApply[*v1alpha3.ResourceClaim, *v1alpha3.ResourceClaimList, *resourcev1alpha3.ResourceClaimApplyConfiguration](
			"resourceclaims",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha3.ResourceClaim { return &v1alpha3.ResourceClaim{} },
			func() *v1alpha3.ResourceClaimList { return &v1alpha3.ResourceClaimList{} }),
	}
}
