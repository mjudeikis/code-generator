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

	"github.com/kcp-dev/logicalcluster/v3"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	batchv1client "k8s.io/client-go/kuberentes/typed/batch/v1"
	applyconfigurationbatchv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/batch/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// CronJobsClusterGetter has a method to return a CronJobClusterInterface.
// A group's client should implement this interface.
type CronJobsClusterGetter interface {
	CronJobs(namespace string) CronJobClusterInterface
}

// CronJobInterface has methods to work with CronJob resources.
type CronJobClusterInterface interface {
	Cluster(logicalcluster.Path) batchv1client.CronJobInterface

	List(ctx context.Context, opts v1.ListOptions) (*batchv1.CronJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	CronJobExpansion
}

// cronJobs implements CronJobInterface
type cronJobs struct {
	*gentype.ClientWithListAndApply[*batchv1.CronJob, *batchv1.CronJobList, *applyconfigurationbatchv1.CronJobApplyConfiguration]
}

// newCronJobs returns a CronJobs
func newCronJobs(c *BatchV1Client, namespace string) *cronJobs {
	return &cronJobs{
		gentype.NewClientWithListAndApply[*batchv1.CronJob, *batchv1.CronJobList, *applyconfigurationbatchv1.CronJobApplyConfiguration](
			"cronjobs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *batchv1.CronJob { return &batchv1.CronJob{} },
			func() *batchv1.CronJobList { return &batchv1.CronJobList{} }),
	}
}
