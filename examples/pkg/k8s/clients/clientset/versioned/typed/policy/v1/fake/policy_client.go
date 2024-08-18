//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	policyv1 "acme.corp/pkg/generated/clientset/versioned/typed/policy/v1"
	kcppolicyv1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/policy/v1"
)

var _ kcppolicyv1.PolicyV1ClusterInterface = (*PolicyV1ClusterClient)(nil)

type PolicyV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *PolicyV1ClusterClient) Cluster(clusterPath logicalcluster.Path) policyv1.PolicyV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &PolicyV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *PolicyV1ClusterClient) PodDisruptionBudgets() kcppolicyv1.PodDisruptionBudgetClusterInterface {
	return &podDisruptionBudgetsClusterClient{Fake: c.Fake}
}

func (c *PolicyV1ClusterClient) Evictions() kcppolicyv1.EvictionClusterInterface {
	return &evictionsClusterClient{Fake: c.Fake}
}

var _ policyv1.PolicyV1Interface = (*PolicyV1Client)(nil)

type PolicyV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *PolicyV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *PolicyV1Client) PodDisruptionBudgets(namespace string) policyv1.PodDisruptionBudgetInterface {
	return &podDisruptionBudgetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *PolicyV1Client) Evictions(namespace string) policyv1.EvictionInterface {
	return &evictionsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
