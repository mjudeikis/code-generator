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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/kcp-dev/logicalcluster/v3"
)

// ControllerRevisionListerExpansion allows custom methods to be added to
// ControllerRevisionLister.
type ControllerRevisionListerExpansion interface{}

// ControllerRevisionClusterListerExpansion allows custom methods to be added to
// ControllerRevisionLister.
type ControllerRevisionClusterListerExpansion interface {
	// Cluster returns a lister that can list and get ControllerRevision in one workspace.
	Cluster(clusterName logicalcluster.Name) ControllerRevisionLister
}

// ControllerRevisionNamespaceListerExpansion allows custom methods to be added to
// ControllerRevisionNamespaceLister.
type ControllerRevisionNamespaceListerExpansion interface{}

// DeploymentListerExpansion allows custom methods to be added to
// DeploymentLister.
type DeploymentListerExpansion interface{}

// DeploymentClusterListerExpansion allows custom methods to be added to
// DeploymentLister.
type DeploymentClusterListerExpansion interface {
	// Cluster returns a lister that can list and get Deployment in one workspace.
	Cluster(clusterName logicalcluster.Name) DeploymentLister
}

// DeploymentNamespaceListerExpansion allows custom methods to be added to
// DeploymentNamespaceLister.
type DeploymentNamespaceListerExpansion interface{}

// StatefulSetListerExpansion allows custom methods to be added to
// StatefulSetLister.
type StatefulSetListerExpansion interface{}

// StatefulSetClusterListerExpansion allows custom methods to be added to
// StatefulSetLister.
type StatefulSetClusterListerExpansion interface {
	// Cluster returns a lister that can list and get StatefulSet in one workspace.
	Cluster(clusterName logicalcluster.Name) StatefulSetLister
}

// StatefulSetNamespaceListerExpansion allows custom methods to be added to
// StatefulSetNamespaceLister.
type StatefulSetNamespaceListerExpansion interface{}
