/*
Copyright 2023.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LocationCtlSpec defines the desired state of LocationCtl
type LocationCtlSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of LocationCtl. Edit locationctl_types.go to remove/update
        //+kubebuilder:validation:Required
        PodX string `json:"podx,omitempty"`

        //+kubebuilder:validation:Required
        PodY string `json:"pody,omitempty"`

        // Replicas is the number of viewers.
        // +kubebuilder:default=1
        // +optional
        Replicas int32 `json:"replicas,omitempty"`
	// +kubebuilder:default=0
        // +optional
	Update int32 `json:"update,omitempty"`
        // +optional
	Apptype string `json:"realtime,omitempty"`


}

// LocationCtlStatus defines the observed state of LocationCtl
type LocationCtlStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}
//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// LocationCtl is the Schema for the locationctls API
type LocationCtl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocationCtlSpec   `json:"spec,omitempty"`
	Status LocationCtlStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LocationCtlList contains a list of LocationCtl
type LocationCtlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationCtl `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LocationCtl{}, &LocationCtlList{})
}
