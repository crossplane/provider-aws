/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// BackupParameters defines the desired state of Backup
type BackupParameters struct {
	// Region is which region the Backup will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The name of the table.
	// +kubebuilder:validation:Required
	TableName              *string `json:"tableName"`
	CustomBackupParameters `json:",inline"`
}

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BackupParameters `json:"forProvider"`
}

// BackupObservation defines the observed state of Backup
type BackupObservation struct {
	// ARN associated with the backup.
	BackupARN *string `json:"backupARN,omitempty"`
	// Time at which the backup was created. This is the request time of the backup.
	BackupCreationDateTime *metav1.Time `json:"backupCreationDateTime,omitempty"`
	// Time at which the automatic on-demand backup created by DynamoDB will expire.
	// This SYSTEM on-demand backup expires automatically 35 days after its creation.
	BackupExpiryDateTime *metav1.Time `json:"backupExpiryDateTime,omitempty"`
	// Name of the requested backup.
	BackupName *string `json:"backupName,omitempty"`
	// Size of the backup in bytes.
	BackupSizeBytes *int64 `json:"backupSizeBytes,omitempty"`
	// Backup can be in one of the following states: CREATING, ACTIVE, DELETED.
	BackupStatus *string `json:"backupStatus,omitempty"`
	// BackupType:
	//
	//    * USER - You create and manage these using the on-demand backup feature.
	//
	//    * SYSTEM - If you delete a table with point-in-time recovery enabled,
	//    a SYSTEM backup is automatically created and is retained for 35 days (at
	//    no additional cost). System backups allow you to restore the deleted table
	//    to the state it was in just before the point of deletion.
	//
	//    * AWS_BACKUP - On-demand backup created by you from AWS Backup service.
	BackupType *string `json:"backupType,omitempty"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BackupObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// Backup is the Schema for the Backups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSpec   `json:"spec,omitempty"`
	Status            BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backups
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// Repository type metadata.
var (
	BackupKind             = "Backup"
	BackupGroupKind        = schema.GroupKind{Group: Group, Kind: BackupKind}.String()
	BackupKindAPIVersion   = BackupKind + "." + GroupVersion.String()
	BackupGroupVersionKind = GroupVersion.WithKind(BackupKind)
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
