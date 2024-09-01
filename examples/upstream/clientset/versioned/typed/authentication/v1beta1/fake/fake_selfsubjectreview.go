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
	v1beta1 "k8s.io/api/authentication/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	authenticationv1beta1 "k8s.io/client-go/applyconfigurations/authentication/v1beta1"
	upstreamauthenticationv1beta1client "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	"k8s.io/client-go/testing"
)

var selfsubjectreviewsResource = v1beta1.SchemeGroupVersion.WithResource("selfsubjectreviews")

var selfsubjectreviewsKind = v1beta1.SchemeGroupVersion.WithKind("SelfSubjectReview")

// selfSubjectReviewsClusterClient implements selfSubjectReviewInterface
type selfSubjectReviewsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *selfSubjectReviewsClusterClient) Cluster(clusterPath logicalcluster.Path) upstreamauthenticationv1beta1client.SelfSubjectReviewInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &selfSubjectReviewsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of SelfSubjectReviews that match those selectors.
func (c *selfSubjectReviewsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SelfSubjectReviewList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(selfsubjectreviewsResource, selfsubjectreviewsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta1.SelfSubjectReviewList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SelfSubjectReviewList{ListMeta: obj.(*v1beta1.SelfSubjectReviewList).ListMeta}
	for _, item := range obj.(*v1beta1.SelfSubjectReviewList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested selfSubjectReviews across all clusters.
func (c *selfSubjectReviewsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(selfsubjectreviewsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type selfSubjectReviewsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *selfSubjectReviewsClient) Create(ctx context.Context, selfSubjectReview *v1beta1.SelfSubjectReview, opts metav1.CreateOptions) (*v1beta1.SelfSubjectReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(selfsubjectreviewsResource, c.ClusterPath, selfSubjectReview), &v1beta1.SelfSubjectReview{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectReview), err
}

func (c *selfSubjectReviewsClient) Update(ctx context.Context, selfSubjectReview *v1beta1.SelfSubjectReview, opts metav1.UpdateOptions) (*v1beta1.SelfSubjectReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(selfsubjectreviewsResource, c.ClusterPath, selfSubjectReview), &v1beta1.SelfSubjectReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectReview), err
}

func (c *selfSubjectReviewsClient) UpdateStatus(ctx context.Context, selfSubjectReview *v1beta1.SelfSubjectReview, opts metav1.UpdateOptions) (*v1beta1.SelfSubjectReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(selfsubjectreviewsResource, c.ClusterPath, "status", selfSubjectReview), &v1beta1.SelfSubjectReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectReview), err
}

func (c *selfSubjectReviewsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(selfsubjectreviewsResource, c.ClusterPath, name, opts), &v1beta1.SelfSubjectReview{})

	return err
}

func (c *selfSubjectReviewsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(selfsubjectreviewsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SelfSubjectReviewList{})
	return err
}

func (c *selfSubjectReviewsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta1.SelfSubjectReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(selfsubjectreviewsResource, c.ClusterPath, name), &v1beta1.SelfSubjectReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectReview), err
}

func (c *selfSubjectReviewsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta1.SelfSubjectReviewList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(selfsubjectreviewsResource, selfsubjectreviewsKind, c.ClusterPath, opts), &v1beta1.SelfSubjectReviewList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SelfSubjectReviewList{ListMeta: obj.(*v1beta1.SelfSubjectReviewList).ListMeta}
	for _, item := range obj.(*v1beta1.SelfSubjectReviewList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *selfSubjectReviewsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(selfsubjectreviewsResource, c.ClusterPath, opts))

}

func (c *selfSubjectReviewsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta1.SelfSubjectReview, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(selfsubjectreviewsResource, c.ClusterPath, name, pt, data, subresources...), &v1beta1.SelfSubjectReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectReview), err
}

func (c *selfSubjectReviewsClient) Apply(ctx context.Context, applyConfiguration *authenticationv1beta1.SelfSubjectReviewApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.SelfSubjectReview, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(selfsubjectreviewsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &v1beta1.SelfSubjectReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SelfSubjectReview), err
}

func (c *selfSubjectReviewsClient) ApplyStatus(ctx context.Context, applyConfiguration *authenticationv1beta1.SelfSubjectReviewApplyConfiguration, opts metav1.ApplyOptions) (*v1beta1.SelfSubjectReview, error) {
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

	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(selfsubjectreviewsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &v1beta1.SelfSubjectReview{})

	return obj.(*v1beta1.SelfSubjectReview), err
}
