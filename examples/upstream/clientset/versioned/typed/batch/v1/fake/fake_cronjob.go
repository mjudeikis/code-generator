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

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	batchv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/batch/v1"
)

// FakeCronJobs implements CronJobInterface
type FakeCronJobs struct {
	Fake *FakeBatchV1
	ns   string
}

var cronjobsResource = v1.SchemeGroupVersion.WithResource("cronjobs")

var cronjobsKind = v1.SchemeGroupVersion.WithKind("CronJob")

// Get takes name of the cronJob, and returns the corresponding cronJob object, and an error if there is any.
func (c *FakeCronJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CronJob, err error) {
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(cronjobsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}

// List takes label and field selectors, and returns the list of CronJobs that match those selectors.
func (c *FakeCronJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CronJobList, err error) {
	emptyResult := &v1.CronJobList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(cronjobsResource, cronjobsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CronJobList{ListMeta: obj.(*v1.CronJobList).ListMeta}
	for _, item := range obj.(*v1.CronJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cronJobs.
func (c *FakeCronJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(cronjobsResource, c.ns, opts))

}

// Create takes the representation of a cronJob and creates it.  Returns the server's representation of the cronJob, and an error, if there is any.
func (c *FakeCronJobs) Create(ctx context.Context, cronJob *v1.CronJob, opts metav1.CreateOptions) (result *v1.CronJob, err error) {
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(cronjobsResource, c.ns, cronJob, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}

// Update takes the representation of a cronJob and updates it. Returns the server's representation of the cronJob, and an error, if there is any.
func (c *FakeCronJobs) Update(ctx context.Context, cronJob *v1.CronJob, opts metav1.UpdateOptions) (result *v1.CronJob, err error) {
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(cronjobsResource, c.ns, cronJob, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCronJobs) UpdateStatus(ctx context.Context, cronJob *v1.CronJob, opts metav1.UpdateOptions) (result *v1.CronJob, err error) {
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(cronjobsResource, "status", c.ns, cronJob, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}

// Delete takes name of the cronJob and deletes it. Returns an error if one occurs.
func (c *FakeCronJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cronjobsResource, c.ns, name, opts), &v1.CronJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCronJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(cronjobsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CronJobList{})
	return err
}

// Patch applies the patch and returns the patched cronJob.
func (c *FakeCronJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CronJob, err error) {
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(cronjobsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cronJob.
func (c *FakeCronJobs) Apply(ctx context.Context, cronJob *batchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CronJob, err error) {
	if cronJob == nil {
		return nil, fmt.Errorf("cronJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(cronJob)
	if err != nil {
		return nil, err
	}
	name := cronJob.Name
	if name == nil {
		return nil, fmt.Errorf("cronJob.Name must be provided to Apply")
	}
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(cronjobsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeCronJobs) ApplyStatus(ctx context.Context, cronJob *batchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CronJob, err error) {
	if cronJob == nil {
		return nil, fmt.Errorf("cronJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(cronJob)
	if err != nil {
		return nil, err
	}
	name := cronJob.Name
	if name == nil {
		return nil, fmt.Errorf("cronJob.Name must be provided to Apply")
	}
	emptyResult := &v1.CronJob{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(cronjobsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.CronJob), err
}
