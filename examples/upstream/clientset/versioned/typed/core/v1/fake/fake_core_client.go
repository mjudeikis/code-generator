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

package fake

import (
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	upstreamcorev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	rest "k8s.io/client-go/rest"
	v1 "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/core/v1"
)

type CoreV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *CoreV1ClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamcorev1client.CoreV1Interface {

	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &CoreV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *CoreV1ClusterClient) ComponentStatuses() v1.ComponentStatusClusterInterface {
	return &componentStatusesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ConfigMaps(namespace string) v1.ConfigMapClusterInterface {
	return &configMapsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Endpoints(namespace string) v1.EndpointsClusterInterface {
	return &endpointsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Events(namespace string) v1.EventClusterInterface {
	return &eventsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) LimitRanges(namespace string) v1.LimitRangeClusterInterface {
	return &limitRangesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Namespaces() v1.NamespaceClusterInterface {
	return &namespacesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Nodes() v1.NodeClusterInterface {
	return &nodesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) PersistentVolumes() v1.PersistentVolumeClusterInterface {
	return &persistentVolumesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) PersistentVolumeClaims(namespace string) v1.PersistentVolumeClaimClusterInterface {
	return &persistentVolumeClaimsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Pods(namespace string) v1.PodClusterInterface {
	return &podsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) PodTemplates(namespace string) v1.PodTemplateClusterInterface {
	return &podTemplatesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ReplicationControllers(namespace string) v1.ReplicationControllerClusterInterface {
	return &replicationControllersClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ResourceQuotas(namespace string) v1.ResourceQuotaClusterInterface {
	return &resourceQuotasClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Secrets(namespace string) v1.SecretClusterInterface {
	return &secretsClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) Services(namespace string) v1.ServiceClusterInterface {
	return &servicesClusterClient{Fake: c.Fake}
}

func (c *CoreV1ClusterClient) ServiceAccounts(namespace string) v1.ServiceAccountClusterInterface {
	return &serviceAccountsClusterClient{Fake: c.Fake}
}

type CoreV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *CoreV1Client) ComponentStatuses() upstreamcorev1client.ComponentStatusInterface {

	return &componentStatusesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *CoreV1Client) ConfigMaps(namespace string) upstreamcorev1client.ConfigMapInterface {

	return &configMapsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) Endpoints(namespace string) upstreamcorev1client.EndpointsInterface {

	return &endpointsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) Events(namespace string) upstreamcorev1client.EventInterface {

	return &eventsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) LimitRanges(namespace string) upstreamcorev1client.LimitRangeInterface {

	return &limitRangesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) Namespaces() upstreamcorev1client.NamespaceInterface {

	return &namespacesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *CoreV1Client) Nodes() upstreamcorev1client.NodeInterface {

	return &nodesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *CoreV1Client) PersistentVolumes() upstreamcorev1client.PersistentVolumeInterface {

	return &persistentVolumesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *CoreV1Client) PersistentVolumeClaims(namespace string) upstreamcorev1client.PersistentVolumeClaimInterface {

	return &persistentVolumeClaimsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) Pods(namespace string) upstreamcorev1client.PodInterface {

	return &podsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) PodTemplates(namespace string) upstreamcorev1client.PodTemplateInterface {

	return &podTemplatesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) ReplicationControllers(namespace string) upstreamcorev1client.ReplicationControllerInterface {

	return &replicationControllersClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) ResourceQuotas(namespace string) upstreamcorev1client.ResourceQuotaInterface {

	return &resourceQuotasClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) Secrets(namespace string) upstreamcorev1client.SecretInterface {

	return &secretsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) Services(namespace string) upstreamcorev1client.ServiceInterface {

	return &servicesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *CoreV1Client) ServiceAccounts(namespace string) upstreamcorev1client.ServiceAccountInterface {

	return &serviceAccountsClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}
