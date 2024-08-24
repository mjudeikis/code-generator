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

package v1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	apiv1 "k8s.io/code-generator/examples/single/api/v1"
)

// TestTypesClusterGetter has a method to return a TestTypeClusterInterface.
// A group's client should implement this interface.
type TestTypesClusterGetter interface {
	TestTypes() TestTypeClusterInterface
}

// TestTypeClusterInterface has methods to work with TestType resources.
type TestTypeClusterInterface interface {
	List(ctx context.Context, opts v1.ListOptions) (*apiv1.TestTypeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Cluster(logicalcluster.Path) TestTypeNamespacer
	TestTypeExpansion
}

type testTypesClusterInterface struct {
	clientCache kcpclient.Cache[*ExampleV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *testTypesClusterInterface) Cluster(clusterPath logicalcluster.Path) TestTypeNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &testTypesNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all TestTypes that are available in all clusters.
func (c *testTypesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*apiv1.TestTypeList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).TestTypes(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all TestTypes across all clusters.
func (c *testTypesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).TestTypes(metav1.NamespaceAll).Watch(ctx, opts)
}

// TestTypeNamespacer can scope to objects within a namespace, returning a TestTypeInterface.
type TestTypeNamespacer interface {
	Namespace(name string) TestTypeInterface
}

type testTypesNamespacer struct {
	clientCache kcpclient.Cache[*ExampleV1Client]
	clusterPath logicalcluster.Path
}

func (n *testTypesNamespacer) Namespace(namespace string) TestTypeInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).TestTypes(namespace)
}
