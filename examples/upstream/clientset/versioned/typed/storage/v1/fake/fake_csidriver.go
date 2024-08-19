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

	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	storagev1 "k8s.io/code-generator/examples/upstream/applyconfiguration/storage/v1"
)

// FakeCSIDrivers implements CSIDriverInterface
type FakeCSIDrivers struct {
	Fake *FakeStorageV1
}

var csidriversResource = v1.SchemeGroupVersion.WithResource("csidrivers")

var csidriversKind = v1.SchemeGroupVersion.WithKind("CSIDriver")

// Get takes name of the cSIDriver, and returns the corresponding cSIDriver object, and an error if there is any.
func (c *FakeCSIDrivers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CSIDriver, err error) {
	emptyResult := &v1.CSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(csidriversResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CSIDriver), err
}

// List takes label and field selectors, and returns the list of CSIDrivers that match those selectors.
func (c *FakeCSIDrivers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CSIDriverList, err error) {
	emptyResult := &v1.CSIDriverList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(csidriversResource, csidriversKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CSIDriverList{ListMeta: obj.(*v1.CSIDriverList).ListMeta}
	for _, item := range obj.(*v1.CSIDriverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cSIDrivers.
func (c *FakeCSIDrivers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(csidriversResource, opts))
}

// Create takes the representation of a cSIDriver and creates it.  Returns the server's representation of the cSIDriver, and an error, if there is any.
func (c *FakeCSIDrivers) Create(ctx context.Context, cSIDriver *v1.CSIDriver, opts metav1.CreateOptions) (result *v1.CSIDriver, err error) {
	emptyResult := &v1.CSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(csidriversResource, cSIDriver, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CSIDriver), err
}

// Update takes the representation of a cSIDriver and updates it. Returns the server's representation of the cSIDriver, and an error, if there is any.
func (c *FakeCSIDrivers) Update(ctx context.Context, cSIDriver *v1.CSIDriver, opts metav1.UpdateOptions) (result *v1.CSIDriver, err error) {
	emptyResult := &v1.CSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(csidriversResource, cSIDriver, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CSIDriver), err
}

// Delete takes name of the cSIDriver and deletes it. Returns an error if one occurs.
func (c *FakeCSIDrivers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(csidriversResource, name, opts), &v1.CSIDriver{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCSIDrivers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(csidriversResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CSIDriverList{})
	return err
}

// Patch applies the patch and returns the patched cSIDriver.
func (c *FakeCSIDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CSIDriver, err error) {
	emptyResult := &v1.CSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(csidriversResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CSIDriver), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cSIDriver.
func (c *FakeCSIDrivers) Apply(ctx context.Context, cSIDriver *storagev1.CSIDriverApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CSIDriver, err error) {
	if cSIDriver == nil {
		return nil, fmt.Errorf("cSIDriver provided to Apply must not be nil")
	}
	data, err := json.Marshal(cSIDriver)
	if err != nil {
		return nil, err
	}
	name := cSIDriver.Name
	if name == nil {
		return nil, fmt.Errorf("cSIDriver.Name must be provided to Apply")
	}
	emptyResult := &v1.CSIDriver{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(csidriversResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CSIDriver), err
}
