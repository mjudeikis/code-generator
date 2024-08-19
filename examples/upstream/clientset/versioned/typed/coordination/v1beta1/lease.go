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

package v1beta1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/coordination/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	coordinationv1beta1client "k8s.io/client-go/kuberentes/typed/coordination/v1beta1"
	coordinationv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/coordination/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// LeasesClusterGetter has a method to return a LeaseClusterInterface.
// A group's client should implement this interface.
type LeasesClusterGetter interface {
	Leases(namespace string) LeaseClusterInterface
}

// LeaseInterface has methods to work with Lease resources.
type LeaseClusterInterface interface {
	Cluster(logicalcluster.Path) coordinationv1beta1client.LeaseInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.LeaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	LeaseExpansion
}

// leases implements LeaseInterface
type leases struct {
	*gentype.ClientWithListAndApply[*v1beta1.Lease, *v1beta1.LeaseList, *coordinationv1beta1.LeaseApplyConfiguration]
}

// newLeases returns a Leases
func newLeases(c *CoordinationV1beta1Client, namespace string) *leases {
	return &leases{
		gentype.NewClientWithListAndApply[*v1beta1.Lease, *v1beta1.LeaseList, *coordinationv1beta1.LeaseApplyConfiguration](
			"leases",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.Lease { return &v1beta1.Lease{} },
			func() *v1beta1.LeaseList { return &v1beta1.LeaseList{} }),
	}
}
