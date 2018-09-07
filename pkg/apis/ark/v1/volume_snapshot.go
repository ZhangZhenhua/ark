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

// VolumeSnapshot is... TODO
type VolumeSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   VolumeSnapshotSpec   `json:"spec"`
	Status VolumeSnapshotStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VolumeSnapshotList is a list of VolumeSnapshots.
type VolumeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []VolumeSnapshot `json:"items"`
}

// VolumeSnapshotSpec defines the specification for an Ark VolumeSnapshot.
type VolumeSnapshotSpec struct {
	// Type is the type of the disk/volume in the cloud provider
	// API.
	Type string `json:"type"`

	// AvailabilityZone is the where the volume is provisioned
	// in the cloud provider.
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// Iops is the optional value of provisioned IOPS for the
	// disk/volume in the cloud provider API.
	Iops *int64 `json:"iops,omitempty"`

	// Backup is a string containing the name of name of the Ark backup this snapshot is associated with.
	Backup string `json:"backup"`

	// Location is a string containing the name of a VolumeSnapshotLocation.
	Location string `json:"volumeLocation"`
}

// VolumeSnapshotStatus captures the current status of an Ark VolumeSnapshot.
type VolumeSnapshotStatus struct {
	// SnapshotID is the ID of the snapshot taken in the cloud
	// provider API of this volume.
	SnapshotID string `json:"snapshotID"`

	// Phase is the current state of the VolumeSnapshot.
	Phase VolumeSnapshotPhase `json:"phase,omitempty"`

	// ValidationErrors is a slice of all validation errors (if
	// applicable).
	ValidationErrors []string `json:"validationErrors"`

	// StartTimestamp records the time a backup was started.
	// Separate from CreationTimestamp, since that value changes
	// on restores.
	// The server's time is used for StartTimestamps
	StartTimestamp metav1.Time `json:"startTimestamp"`

	// CompletionTimestamp records the time a backup was completed.
	// Completion time is recorded even on failed backups.
	// Completion time is recorded before uploading the backup object.
	// The server's time is used for CompletionTimestamps
	CompletionTimestamp metav1.Time `json:"completionTimestamp"`
}

// VolumeSnapshotPhase is the lifecyle phase of an Ark VolumeSnapshot.
type VolumeSnapshotPhase string

const (
	// VolumeSnapshotPhaseNew means the volume snapshot has been created but not
	// yet processed by the VolumeSnapshotController.
	VolumeSnapshotPhaseNew VolumeSnapshotPhase = "New"

	// VolumeSnapshotPhaseCompleted means the location is unavailable to read and write from.
	VolumeSnapshotPhaseCompleted VolumeSnapshotPhase = "Completed"

	// VolumeSnapshotPhaseFailed means the volume snapshot was unable to execute.
	VolumeSnapshotPhaseFailed VolumeSnapshotPhase = "Failed"
)
