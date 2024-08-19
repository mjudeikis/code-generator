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

	v1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	resourcev1alpha3 "k8s.io/code-generator/examples/upstream/applyconfiguration/resource/v1alpha3"
)

// FakeResourceClaimTemplates implements ResourceClaimTemplateInterface
type FakeResourceClaimTemplates struct {
	Fake *FakeResourceV1alpha3
	ns   string
}

var resourceclaimtemplatesResource = v1alpha3.SchemeGroupVersion.WithResource("resourceclaimtemplates")

var resourceclaimtemplatesKind = v1alpha3.SchemeGroupVersion.WithKind("ResourceClaimTemplate")

// Get takes name of the resourceClaimTemplate, and returns the corresponding resourceClaimTemplate object, and an error if there is any.
func (c *FakeResourceClaimTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.ResourceClaimTemplate, err error) {
	emptyResult := &v1alpha3.ResourceClaimTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(resourceclaimtemplatesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.ResourceClaimTemplate), err
}

// List takes label and field selectors, and returns the list of ResourceClaimTemplates that match those selectors.
func (c *FakeResourceClaimTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.ResourceClaimTemplateList, err error) {
	emptyResult := &v1alpha3.ResourceClaimTemplateList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(resourceclaimtemplatesResource, resourceclaimtemplatesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.ResourceClaimTemplateList{ListMeta: obj.(*v1alpha3.ResourceClaimTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha3.ResourceClaimTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceClaimTemplates.
func (c *FakeResourceClaimTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(resourceclaimtemplatesResource, c.ns, opts))

}

// Create takes the representation of a resourceClaimTemplate and creates it.  Returns the server's representation of the resourceClaimTemplate, and an error, if there is any.
func (c *FakeResourceClaimTemplates) Create(ctx context.Context, resourceClaimTemplate *v1alpha3.ResourceClaimTemplate, opts v1.CreateOptions) (result *v1alpha3.ResourceClaimTemplate, err error) {
	emptyResult := &v1alpha3.ResourceClaimTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(resourceclaimtemplatesResource, c.ns, resourceClaimTemplate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.ResourceClaimTemplate), err
}

// Update takes the representation of a resourceClaimTemplate and updates it. Returns the server's representation of the resourceClaimTemplate, and an error, if there is any.
func (c *FakeResourceClaimTemplates) Update(ctx context.Context, resourceClaimTemplate *v1alpha3.ResourceClaimTemplate, opts v1.UpdateOptions) (result *v1alpha3.ResourceClaimTemplate, err error) {
	emptyResult := &v1alpha3.ResourceClaimTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(resourceclaimtemplatesResource, c.ns, resourceClaimTemplate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.ResourceClaimTemplate), err
}

// Delete takes name of the resourceClaimTemplate and deletes it. Returns an error if one occurs.
func (c *FakeResourceClaimTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourceclaimtemplatesResource, c.ns, name, opts), &v1alpha3.ResourceClaimTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceClaimTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(resourceclaimtemplatesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.ResourceClaimTemplateList{})
	return err
}

// Patch applies the patch and returns the patched resourceClaimTemplate.
func (c *FakeResourceClaimTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.ResourceClaimTemplate, err error) {
	emptyResult := &v1alpha3.ResourceClaimTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(resourceclaimtemplatesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.ResourceClaimTemplate), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied resourceClaimTemplate.
func (c *FakeResourceClaimTemplates) Apply(ctx context.Context, resourceClaimTemplate *resourcev1alpha3.ResourceClaimTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.ResourceClaimTemplate, err error) {
	if resourceClaimTemplate == nil {
		return nil, fmt.Errorf("resourceClaimTemplate provided to Apply must not be nil")
	}
	data, err := json.Marshal(resourceClaimTemplate)
	if err != nil {
		return nil, err
	}
	name := resourceClaimTemplate.Name
	if name == nil {
		return nil, fmt.Errorf("resourceClaimTemplate.Name must be provided to Apply")
	}
	emptyResult := &v1alpha3.ResourceClaimTemplate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(resourceclaimtemplatesResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.ResourceClaimTemplate), err
}
