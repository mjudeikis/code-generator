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
	"context"
	json "encoding/json"
	"fmt"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	admissionregistrationv1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	upstreamadmissionregistrationv1beta1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	"k8s.io/client-go/testing"
)

var mutatingwebhookconfigurationsResource = v1beta1.SchemeGroupVersion.WithResource("mutatingwebhookconfigurations")

var mutatingwebhookconfigurationsKind = v1beta1.SchemeGroupVersion.WithKind("MutatingWebhookConfiguration")

// mutatingWebhookConfigurationsClusterClient implements mutatingWebhookConfigurationInterface
type mutatingWebhookConfigurationsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *mutatingWebhookConfigurationsClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamadmissionregistrationv1beta1client.MutatingWebhookConfigurationInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &mutatingWebhookConfigurationsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of MutatingWebhookConfigurations that match those selectors.
func (c *mutatingWebhookConfigurationsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MutatingWebhookConfigurationList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(mutatingwebhookconfigurationsResource, mutatingwebhookconfigurationsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.MutatingWebhookConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MutatingWebhookConfigurationList{ListMeta: obj.(*v1beta1.MutatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*v1beta1.MutatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mutatingWebhookConfigurations across all clusters.
func (c *mutatingWebhookConfigurationsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(mutatingwebhookconfigurationsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type mutatingWebhookConfigurationsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *mutatingWebhookConfigurationsClient) Create(ctx context.Context, mutatingWebhookConfiguration *v1beta1.MutatingWebhookConfiguration, opts metav1.CreateOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(mutatingwebhookconfigurationsResource, c.ClusterPath, mutatingWebhookConfiguration), &v1beta1.MutatingWebhookConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}

func (c *mutatingWebhookConfigurationsClient) Update(ctx context.Context, mutatingWebhookConfiguration *v1beta1.MutatingWebhookConfiguration, opts metav1.UpdateOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(mutatingwebhookconfigurationsResource, c.ClusterPath, mutatingWebhookConfiguration), &v1beta1.MutatingWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}

func (c *mutatingWebhookConfigurationsClient) UpdateStatus(ctx context.Context, mutatingWebhookConfiguration *v1beta1.MutatingWebhookConfiguration, opts metav1.UpdateOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(mutatingwebhookconfigurationsResource, c.ClusterPath, "status", mutatingWebhookConfiguration), &v1beta1.MutatingWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}

func (c *mutatingWebhookConfigurationsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(mutatingwebhookconfigurationsResource, c.ClusterPath, name, opts), &v1beta1.MutatingWebhookConfiguration{})

	return err
}

func (c *mutatingWebhookConfigurationsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(mutatingwebhookconfigurationsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MutatingWebhookConfigurationList{})
	return err
}

func (c *mutatingWebhookConfigurationsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(mutatingwebhookconfigurationsResource, c.ClusterPath, name), &v1beta1.MutatingWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}

func (c *mutatingWebhookConfigurationsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.MutatingWebhookConfigurationList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(mutatingwebhookconfigurationsResource, mutatingwebhookconfigurationsKind, c.ClusterPath, opts), &v1beta1.MutatingWebhookConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MutatingWebhookConfigurationList{ListMeta: obj.(*v1beta1.MutatingWebhookConfigurationList).ListMeta}
	for _, item := range obj.(*v1beta1.MutatingWebhookConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *mutatingWebhookConfigurationsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(mutatingwebhookconfigurationsResource, c.ClusterPath, opts))

}

func (c *mutatingWebhookConfigurationsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.MutatingWebhookConfiguration, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, c.ClusterPath, name, pt, data, subresources...), &v1beta1.MutatingWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}

func (c *mutatingWebhookConfigurationsClient) Apply(ctx context.Context, applyConfiguration *admissionregistrationv1beta1.MutatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &v1beta1.MutatingWebhookConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}

func (c *mutatingWebhookConfigurationsClient) ApplyStatus(ctx context.Context, applyConfiguration *admissionregistrationv1beta1.MutatingWebhookConfigurationApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.MutatingWebhookConfiguration, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(mutatingwebhookconfigurationsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &v1beta1.MutatingWebhookConfiguration{})

	return obj.(*v1beta1.MutatingWebhookConfiguration), err
}
