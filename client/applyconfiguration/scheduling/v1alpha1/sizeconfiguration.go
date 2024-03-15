/*


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

// SizeConfigurationApplyConfiguration represents an declarative configuration of the SizeConfiguration type for use
// with apply.
type SizeConfigurationApplyConfiguration struct {
	Name     *string                              `json:"name,omitempty"`
	Criteria *NodeCountCriteriaApplyConfiguration `json:"criteria,omitempty"`
	Effects  *EffectsApplyConfiguration           `json:"effects,omitempty"`
}

// SizeConfigurationApplyConfiguration constructs an declarative configuration of the SizeConfiguration type for use with
// apply.
func SizeConfiguration() *SizeConfigurationApplyConfiguration {
	return &SizeConfigurationApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SizeConfigurationApplyConfiguration) WithName(value string) *SizeConfigurationApplyConfiguration {
	b.Name = &value
	return b
}

// WithCriteria sets the Criteria field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Criteria field is set to the value of the last call.
func (b *SizeConfigurationApplyConfiguration) WithCriteria(value *NodeCountCriteriaApplyConfiguration) *SizeConfigurationApplyConfiguration {
	b.Criteria = value
	return b
}

// WithEffects sets the Effects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Effects field is set to the value of the last call.
func (b *SizeConfigurationApplyConfiguration) WithEffects(value *EffectsApplyConfiguration) *SizeConfigurationApplyConfiguration {
	b.Effects = value
	return b
}