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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	upstreamadmissionregistrationv1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

// ValidatingWebhookConfigurationsClusterGetter has a method to return a ValidatingWebhookConfigurationClusterInterface.
// A group's client should implement this interface.
type ValidatingWebhookConfigurationsClusterGetter interface {
	ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClusterInterface
}

// ValidatingWebhookConfigurationClusterInterface has methods to work with ValidatingWebhookConfiguration resources.
type ValidatingWebhookConfigurationClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*admissionregistrationv1.ValidatingWebhookConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) upstreamadmissionregistrationv1client.ValidatingWebhookConfigurationInterface
	ValidatingWebhookConfigurationExpansion
}

type validatingWebhookConfigurationsClusterInterface struct {
	clientCache kcpclient.Cache[*upstreamadmissionregistrationv1client.AdmissionregistrationV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *validatingWebhookConfigurationsClusterInterface) Cluster(clusterPath logicalcluster.Path) upstreamadmissionregistrationv1client.ValidatingWebhookConfigurationInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).ValidatingWebhookConfigurations()
}

// List returns the entire collection of all ValidatingWebhookConfigurations that are available in all clusters.
func (c *validatingWebhookConfigurationsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1.ValidatingWebhookConfigurationList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ValidatingWebhookConfigurations().List(ctx, opts)
}

// Watch begins to watch all ValidatingWebhookConfigurations across all clusters.
func (c *validatingWebhookConfigurationsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).ValidatingWebhookConfigurations().Watch(ctx, opts)
}
