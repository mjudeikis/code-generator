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

package v1beta1

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/client-go/rest"

	networkingv1beta1 "acme.corp/pkg/generated/clientset/versioned/typed/networking/v1beta1"
)

type NetworkingV1beta1ClusterInterface interface {
	NetworkingV1beta1ClusterScoper
	IngressesClusterGetter
	IngressClassesClusterGetter
}

type NetworkingV1beta1ClusterScoper interface {
	Cluster(logicalcluster.Path) networkingv1beta1.NetworkingV1beta1Interface
}

type NetworkingV1beta1ClusterClient struct {
	clientCache kcpclient.Cache[*networkingv1beta1.NetworkingV1beta1Client]
}

func (c *NetworkingV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) networkingv1beta1.NetworkingV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *NetworkingV1beta1ClusterClient) Ingresses() IngressClusterInterface {
	return &ingressesClusterInterface{clientCache: c.clientCache}
}

func (c *NetworkingV1beta1ClusterClient) IngressClasses() IngressClassClusterInterface {
	return &ingressClassesClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new NetworkingV1beta1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*NetworkingV1beta1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new NetworkingV1beta1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*NetworkingV1beta1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*networkingv1beta1.NetworkingV1beta1Client]{
		NewForConfigAndClient: networkingv1beta1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}
	return &NetworkingV1beta1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new NetworkingV1beta1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkingV1beta1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
