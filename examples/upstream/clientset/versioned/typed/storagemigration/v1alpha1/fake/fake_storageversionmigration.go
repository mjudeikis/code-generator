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

	v1alpha1 "k8s.io/api/storagemigration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	storagemigrationv1alpha1 "k8s.io/code-generator/examples/upstream/applyconfiguration/storagemigration/v1alpha1"
)

// FakeStorageVersionMigrations implements StorageVersionMigrationInterface
type FakeStorageVersionMigrations struct {
	Fake *FakeStoragemigrationV1alpha1
}

var storageversionmigrationsResource = v1alpha1.SchemeGroupVersion.WithResource("storageversionmigrations")

var storageversionmigrationsKind = v1alpha1.SchemeGroupVersion.WithKind("StorageVersionMigration")

// Get takes name of the storageVersionMigration, and returns the corresponding storageVersionMigration object, and an error if there is any.
func (c *FakeStorageVersionMigrations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.StorageVersionMigration, err error) {
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(storageversionmigrationsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}

// List takes label and field selectors, and returns the list of StorageVersionMigrations that match those selectors.
func (c *FakeStorageVersionMigrations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StorageVersionMigrationList, err error) {
	emptyResult := &v1alpha1.StorageVersionMigrationList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(storageversionmigrationsResource, storageversionmigrationsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageVersionMigrationList{ListMeta: obj.(*v1alpha1.StorageVersionMigrationList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageVersionMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageVersionMigrations.
func (c *FakeStorageVersionMigrations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(storageversionmigrationsResource, opts))
}

// Create takes the representation of a storageVersionMigration and creates it.  Returns the server's representation of the storageVersionMigration, and an error, if there is any.
func (c *FakeStorageVersionMigrations) Create(ctx context.Context, storageVersionMigration *v1alpha1.StorageVersionMigration, opts v1.CreateOptions) (result *v1alpha1.StorageVersionMigration, err error) {
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(storageversionmigrationsResource, storageVersionMigration, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}

// Update takes the representation of a storageVersionMigration and updates it. Returns the server's representation of the storageVersionMigration, and an error, if there is any.
func (c *FakeStorageVersionMigrations) Update(ctx context.Context, storageVersionMigration *v1alpha1.StorageVersionMigration, opts v1.UpdateOptions) (result *v1alpha1.StorageVersionMigration, err error) {
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(storageversionmigrationsResource, storageVersionMigration, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageVersionMigrations) UpdateStatus(ctx context.Context, storageVersionMigration *v1alpha1.StorageVersionMigration, opts v1.UpdateOptions) (result *v1alpha1.StorageVersionMigration, err error) {
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(storageversionmigrationsResource, "status", storageVersionMigration, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}

// Delete takes name of the storageVersionMigration and deletes it. Returns an error if one occurs.
func (c *FakeStorageVersionMigrations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(storageversionmigrationsResource, name, opts), &v1alpha1.StorageVersionMigration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageVersionMigrations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(storageversionmigrationsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageVersionMigrationList{})
	return err
}

// Patch applies the patch and returns the patched storageVersionMigration.
func (c *FakeStorageVersionMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.StorageVersionMigration, err error) {
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(storageversionmigrationsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied storageVersionMigration.
func (c *FakeStorageVersionMigrations) Apply(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigrationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.StorageVersionMigration, err error) {
	if storageVersionMigration == nil {
		return nil, fmt.Errorf("storageVersionMigration provided to Apply must not be nil")
	}
	data, err := json.Marshal(storageVersionMigration)
	if err != nil {
		return nil, err
	}
	name := storageVersionMigration.Name
	if name == nil {
		return nil, fmt.Errorf("storageVersionMigration.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(storageversionmigrationsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeStorageVersionMigrations) ApplyStatus(ctx context.Context, storageVersionMigration *storagemigrationv1alpha1.StorageVersionMigrationApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.StorageVersionMigration, err error) {
	if storageVersionMigration == nil {
		return nil, fmt.Errorf("storageVersionMigration provided to Apply must not be nil")
	}
	data, err := json.Marshal(storageVersionMigration)
	if err != nil {
		return nil, err
	}
	name := storageVersionMigration.Name
	if name == nil {
		return nil, fmt.Errorf("storageVersionMigration.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.StorageVersionMigration{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(storageversionmigrationsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.StorageVersionMigration), err
}
