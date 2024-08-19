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
	"github.com/kcp-dev/logicalcluster/v3"
	v1 "k8s.io/api/authorization/v1"
	gentype "k8s.io/client-go/gentype"
	authorizationv1client "k8s.io/client-go/kuberentes/typed/authorization/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// SelfSubjectAccessReviewsClusterGetter has a method to return a SelfSubjectAccessReviewClusterInterface.
// A group's client should implement this interface.
type SelfSubjectAccessReviewsClusterGetter interface {
	SelfSubjectAccessReviews() SelfSubjectAccessReviewClusterInterface
}

// SelfSubjectAccessReviewInterface has methods to work with SelfSubjectAccessReview resources.
type SelfSubjectAccessReviewClusterInterface interface {
	Cluster(logicalcluster.Path) authorizationv1client.SelfSubjectAccessReviewInterface

	SelfSubjectAccessReviewExpansion
}

// selfSubjectAccessReviews implements SelfSubjectAccessReviewInterface
type selfSubjectAccessReviews struct {
	*gentype.Client[*v1.SelfSubjectAccessReview]
}

// newSelfSubjectAccessReviews returns a SelfSubjectAccessReviews
func newSelfSubjectAccessReviews(c *AuthorizationV1Client) *selfSubjectAccessReviews {
	return &selfSubjectAccessReviews{
		gentype.NewClient[*v1.SelfSubjectAccessReview](
			"selfsubjectaccessreviews",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.SelfSubjectAccessReview { return &v1.SelfSubjectAccessReview{} }),
	}
}
