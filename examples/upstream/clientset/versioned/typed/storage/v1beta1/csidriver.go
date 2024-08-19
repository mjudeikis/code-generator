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
	v1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	storagev1beta1client "k8s.io/client-go/kuberentes/typed/storage/v1beta1"
	storagev1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/storage/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// CSIDriversClusterGetter has a method to return a CSIDriverClusterInterface.
// A group's client should implement this interface.
type CSIDriversClusterGetter interface {
	CSIDrivers() CSIDriverClusterInterface
}

// CSIDriverInterface has methods to work with CSIDriver resources.
type CSIDriverClusterInterface interface {
	Cluster(logicalcluster.Path) storagev1beta1client.CSIDriverInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CSIDriverList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	CSIDriverExpansion
}

// cSIDrivers implements CSIDriverInterface
type cSIDrivers struct {
	*gentype.ClientWithListAndApply[*v1beta1.CSIDriver, *v1beta1.CSIDriverList, *storagev1beta1.CSIDriverApplyConfiguration]
}

// newCSIDrivers returns a CSIDrivers
func newCSIDrivers(c *StorageV1beta1Client) *cSIDrivers {
	return &cSIDrivers{
		gentype.NewClientWithListAndApply[*v1beta1.CSIDriver, *v1beta1.CSIDriverList, *storagev1beta1.CSIDriverApplyConfiguration](
			"csidrivers",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta1.CSIDriver { return &v1beta1.CSIDriver{} },
			func() *v1beta1.CSIDriverList { return &v1beta1.CSIDriverList{} }),
	}
}
