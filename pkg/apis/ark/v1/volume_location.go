/*
Copyright 2018 the Heptio Ark contributors.

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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeLocation is a location where Ark stores backup objects.
type VolumeLocation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   VolumeLocationSpec   `json:"spec"`
	Status VolumeLocationStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeLocationList is a list of VolumeLocations.
type VolumeLocationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []VolumeLocation `json:"items"`
}

// VolumeLocationSpec defines the specification for an Ark VolumeLocation.
type VolumeLocationSpec struct {
	// Provider is the provider of the backup storage.
	Provider string `json:"provider"`

	// Config is for provider-specific configuration fields.
	Config map[string]string `json:"config"`

	StorageType `json:",inline"`
}

// VolumeLocationPhase is the lifecyle phase of an Ark VolumeLocation.
type VolumeLocationPhase string

const (
	// VolumeLocationPhaseAvailable means the location is available to read and write from.
	VolumeLocationPhaseAvailable VolumeLocationPhase = "Available"

	// VolumeLocationPhaseUnavailable means the location is unavailable to read and write from.
	VolumeLocationPhaseUnavailable VolumeLocationPhase = "Unavailable"
)

// VolumeLocationAccessMode represents the permissions for a VolumeLocation.
type VolumeLocationAccessMode string

const (
	// VolumeLocationAccessModeReadOnly represents read-only access to a VolumeLocation.
	VolumeLocationAccessModeReadOnly VolumeLocationAccessMode = "ReadOnly"

	// VolumeLocationAccessModeReadWrite represents read and write access to a VolumeLocation.
	VolumeLocationAccessModeReadWrite VolumeLocationAccessMode = "ReadWrite"
)

// VolumeLocationStatus describes the current status of an Ark VolumeLocation.
type VolumeLocationStatus struct {
	Phase      VolumeLocationPhase      `json:"phase,omitempty"`
	AccessMode VolumeLocationAccessMode `json:"accessMode,omitempty"`
}
