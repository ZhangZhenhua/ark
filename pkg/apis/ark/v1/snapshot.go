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

// Snapshot is... TODO
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   SnapshotSpec   `json:"spec"`
	Status SnapshotStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SnapshotList is a list of Snapshots.
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Snapshot `json:"items"`
}

// SnapshotSpec defines the specification for an Ark Snapshot.
type SnapshotSpec struct {
	// Provider is the provider of the backup storage.
	Provider string `json:"provider"`

	// Config is for provider-specific configuration fields.
	Config map[string]string `json:"config"`

	StorageType `json:",inline"`
}

// SnapshotPhase is the lifecyle phase of an Ark Snapshot.
type SnapshotPhase string

const (
	// SnapshotPhaseAvailable means the location is available to read and write from.
	SnapshotPhaseAvailable SnapshotPhase = "Available"

	// SnapshotPhaseUnavailable means the location is unavailable to read and write from.
	SnapshotPhaseUnavailable SnapshotPhase = "Unavailable"
)

// SnapshotAccessMode represents the permissions for a Snapshot.
type SnapshotAccessMode string

const (
	// SnapshotAccessModeReadOnly represents read-only access to a Snapshot.
	SnapshotAccessModeReadOnly SnapshotAccessMode = "ReadOnly"

	// SnapshotAccessModeReadWrite represents read and write access to a Snapshot.
	SnapshotAccessModeReadWrite SnapshotAccessMode = "ReadWrite"
)

// SnapshotStatus describes the current status of an Ark Snapshot.
type SnapshotStatus struct {
	Phase      SnapshotPhase      `json:"phase,omitempty"`
	AccessMode SnapshotAccessMode `json:"accessMode,omitempty"`
}
