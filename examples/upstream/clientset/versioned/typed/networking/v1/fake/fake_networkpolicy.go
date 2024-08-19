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

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	networkingv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/networking/v1"
)

// FakeNetworkPolicies implements NetworkPolicyInterface
type FakeNetworkPolicies struct {
	Fake *FakeNetworkingV1
	ns   string
}

var networkpoliciesResource = v1.SchemeGroupVersion.WithResource("networkpolicies")

var networkpoliciesKind = v1.SchemeGroupVersion.WithKind("NetworkPolicy")

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *FakeNetworkPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkPolicy, err error) {
	emptyResult := &v1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(networkpoliciesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *FakeNetworkPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkPolicyList, err error) {
	emptyResult := &v1.NetworkPolicyList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(networkpoliciesResource, networkpoliciesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.NetworkPolicyList{ListMeta: obj.(*v1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*v1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *FakeNetworkPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(networkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkPolicy and creates it.  Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Create(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.CreateOptions) (result *v1.NetworkPolicy, err error) {
	emptyResult := &v1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(networkpoliciesResource, c.ns, networkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// Update takes the representation of a networkPolicy and updates it. Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *FakeNetworkPolicies) Update(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.UpdateOptions) (result *v1.NetworkPolicy, err error) {
	emptyResult := &v1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(networkpoliciesResource, c.ns, networkPolicy, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// Delete takes name of the networkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkpoliciesResource, c.ns, name, opts), &v1.NetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(networkpoliciesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.NetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkPolicy.
func (c *FakeNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicy, err error) {
	emptyResult := &v1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkpoliciesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NetworkPolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied networkPolicy.
func (c *FakeNetworkPolicies) Apply(ctx context.Context, networkPolicy *networkingv1.NetworkPolicyApplyConfiguration, opts metav1.ApplyOptions) (result *v1.NetworkPolicy, err error) {
	if networkPolicy == nil {
		return nil, fmt.Errorf("networkPolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(networkPolicy)
	if err != nil {
		return nil, err
	}
	name := networkPolicy.Name
	if name == nil {
		return nil, fmt.Errorf("networkPolicy.Name must be provided to Apply")
	}
	emptyResult := &v1.NetworkPolicy{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(networkpoliciesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.NetworkPolicy), err
}
