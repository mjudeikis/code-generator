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
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/authorization/v1beta1"
	gentype "k8s.io/client-go/gentype"
	authorizationv1beta1client "k8s.io/client-go/kuberentes/typed/authorization/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// SelfSubjectRulesReviewsClusterGetter has a method to return a SelfSubjectRulesReviewClusterInterface.
// A group's client should implement this interface.
type SelfSubjectRulesReviewsClusterGetter interface {
	SelfSubjectRulesReviews() SelfSubjectRulesReviewClusterInterface
}

// SelfSubjectRulesReviewInterface has methods to work with SelfSubjectRulesReview resources.
type SelfSubjectRulesReviewClusterInterface interface {
	Cluster(logicalcluster.Path) authorizationv1beta1client.SelfSubjectRulesReviewInterface

	SelfSubjectRulesReviewExpansion
}

// selfSubjectRulesReviews implements SelfSubjectRulesReviewInterface
type selfSubjectRulesReviews struct {
	*gentype.Client[*v1beta1.SelfSubjectRulesReview]
}

// newSelfSubjectRulesReviews returns a SelfSubjectRulesReviews
func newSelfSubjectRulesReviews(c *AuthorizationV1beta1Client) *selfSubjectRulesReviews {
	return &selfSubjectRulesReviews{
		gentype.NewClient[*v1beta1.SelfSubjectRulesReview](
			"selfsubjectrulesreviews",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.SelfSubjectRulesReview { return &v1beta1.SelfSubjectRulesReview{} }),
	}
}
