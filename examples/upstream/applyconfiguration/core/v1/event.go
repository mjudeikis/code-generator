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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// EventApplyConfiguration represents a declarative configuration of the Event type for use
// with apply.
type EventApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	InvolvedObject                   *ObjectReferenceApplyConfiguration `json:"involvedObject,omitempty"`
	Reason                           *string                            `json:"reason,omitempty"`
	Message                          *string                            `json:"message,omitempty"`
	Source                           *EventSourceApplyConfiguration     `json:"source,omitempty"`
	FirstTimestamp                   *metav1.Time                       `json:"firstTimestamp,omitempty"`
	LastTimestamp                    *metav1.Time                       `json:"lastTimestamp,omitempty"`
	Count                            *int32                             `json:"count,omitempty"`
	Type                             *string                            `json:"type,omitempty"`
	EventTime                        *metav1.MicroTime                  `json:"eventTime,omitempty"`
	Series                           *EventSeriesApplyConfiguration     `json:"series,omitempty"`
	Action                           *string                            `json:"action,omitempty"`
	Related                          *ObjectReferenceApplyConfiguration `json:"related,omitempty"`
	ReportingController              *string                            `json:"reportingComponent,omitempty"`
	ReportingInstance                *string                            `json:"reportingInstance,omitempty"`
}

// Event constructs a declarative configuration of the Event type for use with
// apply.
func Event(name, namespace string) *EventApplyConfiguration {
	b := &EventApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Event")
	b.WithAPIVersion("v1")
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *EventApplyConfiguration) WithKind(value string) *EventApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *EventApplyConfiguration) WithAPIVersion(value string) *EventApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EventApplyConfiguration) WithName(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *EventApplyConfiguration) WithGenerateName(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *EventApplyConfiguration) WithNamespace(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *EventApplyConfiguration) WithUID(value types.UID) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *EventApplyConfiguration) WithResourceVersion(value string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *EventApplyConfiguration) WithGeneration(value int64) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithCreationTimestamp(value metav1.Time) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *EventApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *EventApplyConfiguration) WithLabels(entries map[string]string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *EventApplyConfiguration) WithAnnotations(entries map[string]string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *EventApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *EventApplyConfiguration) WithFinalizers(values ...string) *EventApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *EventApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithInvolvedObject sets the InvolvedObject field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InvolvedObject field is set to the value of the last call.
func (b *EventApplyConfiguration) WithInvolvedObject(value *ObjectReferenceApplyConfiguration) *EventApplyConfiguration {
	b.InvolvedObject = value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *EventApplyConfiguration) WithReason(value string) *EventApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *EventApplyConfiguration) WithMessage(value string) *EventApplyConfiguration {
	b.Message = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *EventApplyConfiguration) WithSource(value *EventSourceApplyConfiguration) *EventApplyConfiguration {
	b.Source = value
	return b
}

// WithFirstTimestamp sets the FirstTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FirstTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithFirstTimestamp(value metav1.Time) *EventApplyConfiguration {
	b.FirstTimestamp = &value
	return b
}

// WithLastTimestamp sets the LastTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTimestamp field is set to the value of the last call.
func (b *EventApplyConfiguration) WithLastTimestamp(value metav1.Time) *EventApplyConfiguration {
	b.LastTimestamp = &value
	return b
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *EventApplyConfiguration) WithCount(value int32) *EventApplyConfiguration {
	b.Count = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *EventApplyConfiguration) WithType(value string) *EventApplyConfiguration {
	b.Type = &value
	return b
}

// WithEventTime sets the EventTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EventTime field is set to the value of the last call.
func (b *EventApplyConfiguration) WithEventTime(value metav1.MicroTime) *EventApplyConfiguration {
	b.EventTime = &value
	return b
}

// WithSeries sets the Series field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Series field is set to the value of the last call.
func (b *EventApplyConfiguration) WithSeries(value *EventSeriesApplyConfiguration) *EventApplyConfiguration {
	b.Series = value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *EventApplyConfiguration) WithAction(value string) *EventApplyConfiguration {
	b.Action = &value
	return b
}

// WithRelated sets the Related field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Related field is set to the value of the last call.
func (b *EventApplyConfiguration) WithRelated(value *ObjectReferenceApplyConfiguration) *EventApplyConfiguration {
	b.Related = value
	return b
}

// WithReportingController sets the ReportingController field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReportingController field is set to the value of the last call.
func (b *EventApplyConfiguration) WithReportingController(value string) *EventApplyConfiguration {
	b.ReportingController = &value
	return b
}

// WithReportingInstance sets the ReportingInstance field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReportingInstance field is set to the value of the last call.
func (b *EventApplyConfiguration) WithReportingInstance(value string) *EventApplyConfiguration {
	b.ReportingInstance = &value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *EventApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
