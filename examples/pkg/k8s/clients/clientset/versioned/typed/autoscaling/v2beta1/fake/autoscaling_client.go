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

	autoscalingv2beta1 "acme.corp/pkg/generated/clientset/versioned/typed/autoscaling/v2beta1"
	kcpautoscalingv2beta1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/autoscaling/v2beta1"
)

var _ kcpautoscalingv2beta1.AutoscalingV2beta1ClusterInterface = (*AutoscalingV2beta1ClusterClient)(nil)

type AutoscalingV2beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AutoscalingV2beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) autoscalingv2beta1.AutoscalingV2beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AutoscalingV2beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AutoscalingV2beta1ClusterClient) HorizontalPodAutoscalers() kcpautoscalingv2beta1.HorizontalPodAutoscalerClusterInterface {
	return &horizontalPodAutoscalersClusterClient{Fake: c.Fake}
}

var _ autoscalingv2beta1.AutoscalingV2beta1Interface = (*AutoscalingV2beta1Client)(nil)

type AutoscalingV2beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AutoscalingV2beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AutoscalingV2beta1Client) HorizontalPodAutoscalers(namespace string) autoscalingv2beta1.HorizontalPodAutoscalerInterface {
	return &horizontalPodAutoscalersClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
