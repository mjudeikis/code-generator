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
	v1 "k8s.io/api/policy/v1"
	gentype "k8s.io/client-go/gentype"
	policyv1client "k8s.io/client-go/kuberentes/typed/policy/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// EvictionsClusterGetter has a method to return a EvictionClusterInterface.
// A group's client should implement this interface.
type EvictionsClusterGetter interface {
	Evictions(namespace string) EvictionClusterInterface
}

// EvictionInterface has methods to work with Eviction resources.
type EvictionClusterInterface interface {
	Cluster(logicalcluster.Path) policyv1client.EvictionInterface

	EvictionExpansion
}

// evictions implements EvictionInterface
type evictions struct {
	*gentype.Client[*v1.Eviction]
}

// newEvictions returns a Evictions
func newEvictions(c *PolicyV1Client, namespace string) *evictions {
	return &evictions{
		gentype.NewClient[*v1.Eviction](
			"evictions",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.Eviction { return &v1.Eviction{} }),
	}
}
