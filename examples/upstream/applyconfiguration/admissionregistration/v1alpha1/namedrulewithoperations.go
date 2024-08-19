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

package v1alpha1

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	v1 "k8s.io/code-generator/examples/upstream/applyconfiguration/admissionregistration/v1"
)

// NamedRuleWithOperationsApplyConfiguration represents a declarative configuration of the NamedRuleWithOperations type for use
// with apply.
type NamedRuleWithOperationsApplyConfiguration struct {
	ResourceNames                           []string `json:"resourceNames,omitempty"`
	v1.RuleWithOperationsApplyConfiguration `json:",inline"`
}

// NamedRuleWithOperationsApplyConfiguration constructs a declarative configuration of the NamedRuleWithOperations type for use with
// apply.
func NamedRuleWithOperations() *NamedRuleWithOperationsApplyConfiguration {
	return &NamedRuleWithOperationsApplyConfiguration{}
}

// WithResourceNames adds the given value to the ResourceNames field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResourceNames field.
func (b *NamedRuleWithOperationsApplyConfiguration) WithResourceNames(values ...string) *NamedRuleWithOperationsApplyConfiguration {
	for i := range values {
		b.ResourceNames = append(b.ResourceNames, values[i])
	}
	return b
}

// WithOperations adds the given value to the Operations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Operations field.
func (b *NamedRuleWithOperationsApplyConfiguration) WithOperations(values ...admissionregistrationv1.OperationType) *NamedRuleWithOperationsApplyConfiguration {
	for i := range values {
		b.Operations = append(b.Operations, values[i])
	}
	return b
}

// WithAPIGroups adds the given value to the APIGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIGroups field.
func (b *NamedRuleWithOperationsApplyConfiguration) WithAPIGroups(values ...string) *NamedRuleWithOperationsApplyConfiguration {
	for i := range values {
		b.APIGroups = append(b.APIGroups, values[i])
	}
	return b
}

// WithAPIVersions adds the given value to the APIVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIVersions field.
func (b *NamedRuleWithOperationsApplyConfiguration) WithAPIVersions(values ...string) *NamedRuleWithOperationsApplyConfiguration {
	for i := range values {
		b.APIVersions = append(b.APIVersions, values[i])
	}
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *NamedRuleWithOperationsApplyConfiguration) WithResources(values ...string) *NamedRuleWithOperationsApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *NamedRuleWithOperationsApplyConfiguration) WithScope(value admissionregistrationv1.ScopeType) *NamedRuleWithOperationsApplyConfiguration {
	b.Scope = &value
	return b
}
