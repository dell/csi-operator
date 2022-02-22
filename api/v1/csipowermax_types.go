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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CSIPowerMaxSpec defines the desired state of CSIPowerMax
type CSIPowerMaxSpec struct {
	// Driver is the specification for the CSI PowerMax Driver
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Driver"
	Driver Driver `json:"driver" yaml:"driver"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=csipowermaxes,scope=Namespaced

// CSIPowerMax is the Schema for the csipowermaxes API
// +operator-sdk:csv:customresourcedefinitions:displayName="CSI PowerMax",resources={{Deployment,v1,powermax-controller},{DameonSet,v1,powermax-node}}
type CSIPowerMax struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CSIPowerMaxSpec `json:"spec,omitempty"`
	Status DriverStatus    `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CSIPowerMaxList contains a list of CSIPowerMax
type CSIPowerMaxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CSIPowerMax `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CSIPowerMax{}, &CSIPowerMaxList{})
}
