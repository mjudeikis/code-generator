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

	flowcontrolv1beta2 "acme.corp/pkg/generated/clientset/versioned/typed/flowcontrol/v1beta2"
	kcpflowcontrolv1beta2 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/flowcontrol/v1beta2"
)

var _ kcpflowcontrolv1beta2.FlowcontrolV1beta2ClusterInterface = (*FlowcontrolV1beta2ClusterClient)(nil)

type FlowcontrolV1beta2ClusterClient struct {
	*kcptesting.Fake
}

func (c *FlowcontrolV1beta2ClusterClient) Cluster(clusterPath logicalcluster.Path) flowcontrolv1beta2.FlowcontrolV1beta2Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &FlowcontrolV1beta2Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *FlowcontrolV1beta2ClusterClient) FlowSchemas() kcpflowcontrolv1beta2.FlowSchemaClusterInterface {
	return &flowSchemasClusterClient{Fake: c.Fake}
}

func (c *FlowcontrolV1beta2ClusterClient) PriorityLevelConfigurations() kcpflowcontrolv1beta2.PriorityLevelConfigurationClusterInterface {
	return &priorityLevelConfigurationsClusterClient{Fake: c.Fake}
}

var _ flowcontrolv1beta2.FlowcontrolV1beta2Interface = (*FlowcontrolV1beta2Client)(nil)

type FlowcontrolV1beta2Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *FlowcontrolV1beta2Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *FlowcontrolV1beta2Client) FlowSchemas() flowcontrolv1beta2.FlowSchemaInterface {
	return &flowSchemasClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *FlowcontrolV1beta2Client) PriorityLevelConfigurations() flowcontrolv1beta2.PriorityLevelConfigurationInterface {
	return &priorityLevelConfigurationsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
