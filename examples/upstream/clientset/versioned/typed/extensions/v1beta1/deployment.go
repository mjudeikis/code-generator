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
	json "encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	extensionsv1beta1client "k8s.io/client-go/kuberentes/typed/extensions/v1beta1"
	extensionsv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/extensions/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// DeploymentsClusterGetter has a method to return a DeploymentClusterInterface.
// A group's client should implement this interface.
type DeploymentsClusterGetter interface {
	Deployments(namespace string) DeploymentClusterInterface
}

// DeploymentInterface has methods to work with Deployment resources.
type DeploymentClusterInterface interface {
	Cluster(logicalcluster.Path) extensionsv1beta1client.DeploymentInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DeploymentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)

	UpdateScale(ctx context.Context, deploymentName string, scale *v1beta1.Scale, opts v1.UpdateOptions) (*v1beta1.Scale, error)

	DeploymentExpansion
}

// deployments implements DeploymentInterface
type deployments struct {
	*gentype.ClientWithListAndApply[*v1beta1.Deployment, *v1beta1.DeploymentList, *extensionsv1beta1.DeploymentApplyConfiguration]
}

// newDeployments returns a Deployments
func newDeployments(c *ExtensionsV1beta1Client, namespace string) *deployments {
	return &deployments{
		gentype.NewClientWithListAndApply[*v1beta1.Deployment, *v1beta1.DeploymentList, *extensionsv1beta1.DeploymentApplyConfiguration](
			"deployments",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.Deployment { return &v1beta1.Deployment{} },
			func() *v1beta1.DeploymentList { return &v1beta1.DeploymentList{} }),
	}
}

// GetScale takes name of the deployment, and returns the corresponding v1beta1.Scale object, and an error if there is any.
func (c *deployments) GetScale(ctx context.Context, deploymentName string, options v1.GetOptions) (result *v1beta1.Scale, err error) {
	result = &v1beta1.Scale{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *deployments) UpdateScale(ctx context.Context, deploymentName string, scale *v1beta1.Scale, opts v1.UpdateOptions) (result *v1beta1.Scale, err error) {
	result = &v1beta1.Scale{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *deployments) ApplyScale(ctx context.Context, deploymentName string, scale *extensionsv1beta1.ScaleApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}

	result = &v1beta1.Scale{}
	err = c.GetClient().Patch(types.ApplyPatchType).
		Namespace(c.GetNamespace()).
		Resource("deployments").
		Name(deploymentName).
		SubResource("scale").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
