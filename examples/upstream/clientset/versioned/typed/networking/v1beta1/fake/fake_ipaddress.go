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

	v1beta1 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	networkingv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/networking/v1beta1"
)

// FakeIPAddresses implements IPAddressInterface
type FakeIPAddresses struct {
	Fake *FakeNetworkingV1beta1
}

var ipaddressesResource = v1beta1.SchemeGroupVersion.WithResource("ipaddresses")

var ipaddressesKind = v1beta1.SchemeGroupVersion.WithKind("IPAddress")

// Get takes name of the iPAddress, and returns the corresponding iPAddress object, and an error if there is any.
func (c *FakeIPAddresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IPAddress, err error) {
	emptyResult := &v1beta1.IPAddress{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(ipaddressesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.IPAddress), err
}

// List takes label and field selectors, and returns the list of IPAddresses that match those selectors.
func (c *FakeIPAddresses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IPAddressList, err error) {
	emptyResult := &v1beta1.IPAddressList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(ipaddressesResource, ipaddressesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IPAddressList{ListMeta: obj.(*v1beta1.IPAddressList).ListMeta}
	for _, item := range obj.(*v1beta1.IPAddressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAddresses.
func (c *FakeIPAddresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(ipaddressesResource, opts))
}

// Create takes the representation of a iPAddress and creates it.  Returns the server's representation of the iPAddress, and an error, if there is any.
func (c *FakeIPAddresses) Create(ctx context.Context, iPAddress *v1beta1.IPAddress, opts v1.CreateOptions) (result *v1beta1.IPAddress, err error) {
	emptyResult := &v1beta1.IPAddress{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(ipaddressesResource, iPAddress, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.IPAddress), err
}

// Update takes the representation of a iPAddress and updates it. Returns the server's representation of the iPAddress, and an error, if there is any.
func (c *FakeIPAddresses) Update(ctx context.Context, iPAddress *v1beta1.IPAddress, opts v1.UpdateOptions) (result *v1beta1.IPAddress, err error) {
	emptyResult := &v1beta1.IPAddress{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(ipaddressesResource, iPAddress, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.IPAddress), err
}

// Delete takes name of the iPAddress and deletes it. Returns an error if one occurs.
func (c *FakeIPAddresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ipaddressesResource, name, opts), &v1beta1.IPAddress{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPAddresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(ipaddressesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.IPAddressList{})
	return err
}

// Patch applies the patch and returns the patched iPAddress.
func (c *FakeIPAddresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.IPAddress, err error) {
	emptyResult := &v1beta1.IPAddress{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ipaddressesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.IPAddress), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied iPAddress.
func (c *FakeIPAddresses) Apply(ctx context.Context, iPAddress *networkingv1beta1.IPAddressApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.IPAddress, err error) {
	if iPAddress == nil {
		return nil, fmt.Errorf("iPAddress provided to Apply must not be nil")
	}
	data, err := json.Marshal(iPAddress)
	if err != nil {
		return nil, err
	}
	name := iPAddress.Name
	if name == nil {
		return nil, fmt.Errorf("iPAddress.Name must be provided to Apply")
	}
	emptyResult := &v1beta1.IPAddress{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(ipaddressesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.IPAddress), err
}
