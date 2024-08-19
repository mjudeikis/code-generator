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

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "k8s.io/code-generator/examples/upstream/applyconfiguration/core/v1"
)

// FakeReplicationControllers implements ReplicationControllerInterface
type FakeReplicationControllers struct {
	Fake *FakeCoreV1
	ns   string
}

var replicationcontrollersResource = v1.SchemeGroupVersion.WithResource("replicationcontrollers")

var replicationcontrollersKind = v1.SchemeGroupVersion.WithKind("ReplicationController")

// Get takes name of the replicationController, and returns the corresponding replicationController object, and an error if there is any.
func (c *FakeReplicationControllers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ReplicationController, err error) {
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(replicationcontrollersResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// List takes label and field selectors, and returns the list of ReplicationControllers that match those selectors.
func (c *FakeReplicationControllers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ReplicationControllerList, err error) {
	emptyResult := &v1.ReplicationControllerList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(replicationcontrollersResource, replicationcontrollersKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ReplicationControllerList{ListMeta: obj.(*v1.ReplicationControllerList).ListMeta}
	for _, item := range obj.(*v1.ReplicationControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicationControllers.
func (c *FakeReplicationControllers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(replicationcontrollersResource, c.ns, opts))

}

// Create takes the representation of a replicationController and creates it.  Returns the server's representation of the replicationController, and an error, if there is any.
func (c *FakeReplicationControllers) Create(ctx context.Context, replicationController *v1.ReplicationController, opts metav1.CreateOptions) (result *v1.ReplicationController, err error) {
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(replicationcontrollersResource, c.ns, replicationController, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// Update takes the representation of a replicationController and updates it. Returns the server's representation of the replicationController, and an error, if there is any.
func (c *FakeReplicationControllers) Update(ctx context.Context, replicationController *v1.ReplicationController, opts metav1.UpdateOptions) (result *v1.ReplicationController, err error) {
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(replicationcontrollersResource, c.ns, replicationController, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicationControllers) UpdateStatus(ctx context.Context, replicationController *v1.ReplicationController, opts metav1.UpdateOptions) (result *v1.ReplicationController, err error) {
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(replicationcontrollersResource, "status", c.ns, replicationController, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// Delete takes name of the replicationController and deletes it. Returns an error if one occurs.
func (c *FakeReplicationControllers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(replicationcontrollersResource, c.ns, name, opts), &v1.ReplicationController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicationControllers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(replicationcontrollersResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ReplicationControllerList{})
	return err
}

// Patch applies the patch and returns the patched replicationController.
func (c *FakeReplicationControllers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ReplicationController, err error) {
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicationcontrollersResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied replicationController.
func (c *FakeReplicationControllers) Apply(ctx context.Context, replicationController *corev1.ReplicationControllerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicationController, err error) {
	if replicationController == nil {
		return nil, fmt.Errorf("replicationController provided to Apply must not be nil")
	}
	data, err := json.Marshal(replicationController)
	if err != nil {
		return nil, err
	}
	name := replicationController.Name
	if name == nil {
		return nil, fmt.Errorf("replicationController.Name must be provided to Apply")
	}
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicationcontrollersResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeReplicationControllers) ApplyStatus(ctx context.Context, replicationController *corev1.ReplicationControllerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicationController, err error) {
	if replicationController == nil {
		return nil, fmt.Errorf("replicationController provided to Apply must not be nil")
	}
	data, err := json.Marshal(replicationController)
	if err != nil {
		return nil, err
	}
	name := replicationController.Name
	if name == nil {
		return nil, fmt.Errorf("replicationController.Name must be provided to Apply")
	}
	emptyResult := &v1.ReplicationController{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(replicationcontrollersResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ReplicationController), err
}

// GetScale takes name of the replicationController, and returns the corresponding scale object, and an error if there is any.
func (c *FakeReplicationControllers) GetScale(ctx context.Context, replicationControllerName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(replicationcontrollersResource, c.ns, "scale", replicationControllerName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeReplicationControllers) UpdateScale(ctx context.Context, replicationControllerName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	emptyResult := &autoscalingv1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(replicationcontrollersResource, "scale", c.ns, scale, opts), &autoscalingv1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*autoscalingv1.Scale), err
}
