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

	appsv1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/apps/v1beta1"
	kcpappsv1beta1 "acme.corp/pkg/k8s/clients/clientset/versioned/typed/apps/v1beta1"
)

var _ kcpappsv1beta1.AppsV1beta1ClusterInterface = (*AppsV1beta1ClusterClient)(nil)

type AppsV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *AppsV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) appsv1beta1.AppsV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &AppsV1beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *AppsV1beta1ClusterClient) StatefulSets() kcpappsv1beta1.StatefulSetClusterInterface {
	return &statefulSetsClusterClient{Fake: c.Fake}
}

func (c *AppsV1beta1ClusterClient) Deployments() kcpappsv1beta1.DeploymentClusterInterface {
	return &deploymentsClusterClient{Fake: c.Fake}
}

func (c *AppsV1beta1ClusterClient) ControllerRevisions() kcpappsv1beta1.ControllerRevisionClusterInterface {
	return &controllerRevisionsClusterClient{Fake: c.Fake}
}

var _ appsv1beta1.AppsV1beta1Interface = (*AppsV1beta1Client)(nil)

type AppsV1beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *AppsV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *AppsV1beta1Client) StatefulSets(namespace string) appsv1beta1.StatefulSetInterface {
	return &statefulSetsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1beta1Client) Deployments(namespace string) appsv1beta1.DeploymentInterface {
	return &deploymentsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *AppsV1beta1Client) ControllerRevisions(namespace string) appsv1beta1.ControllerRevisionInterface {
	return &controllerRevisionsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
