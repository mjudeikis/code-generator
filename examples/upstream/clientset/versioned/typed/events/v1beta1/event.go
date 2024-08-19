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
	v1beta1 "k8s.io/api/events/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	eventsv1beta1client "k8s.io/client-go/kuberentes/typed/events/v1beta1"
	eventsv1beta1 "k8s.io/code-generator/examples/upstream/applyconfiguration/events/v1beta1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// EventsClusterGetter has a method to return a EventClusterInterface.
// A group's client should implement this interface.
type EventsClusterGetter interface {
	Events(namespace string) EventClusterInterface
}

// EventInterface has methods to work with Event resources.
type EventClusterInterface interface {
	Cluster(logicalcluster.Path) eventsv1beta1client.EventInterface

	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.EventList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	EventExpansion
}

// events implements EventInterface
type events struct {
	*gentype.ClientWithListAndApply[*v1beta1.Event, *v1beta1.EventList, *eventsv1beta1.EventApplyConfiguration]
}

// newEvents returns a Events
func newEvents(c *EventsV1beta1Client, namespace string) *events {
	return &events{
		gentype.NewClientWithListAndApply[*v1beta1.Event, *v1beta1.EventList, *eventsv1beta1.EventApplyConfiguration](
			"events",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.Event { return &v1beta1.Event{} },
			func() *v1beta1.EventList { return &v1beta1.EventList{} }),
	}
}
