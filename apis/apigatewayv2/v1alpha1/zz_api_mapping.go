/*
Copyright 2019 The Crossplane Authors.

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

// APIMappingParameters defines the desired state of APIMapping
type APIMappingParameters struct {
	// Region is which region the APIMapping will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region,omitempty"`

	// +kubebuilder:validation:Required
	APIID *string `json:"apiID"`

	APIMappingKey *string `json:"apiMappingKey,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName"`

	// +kubebuilder:validation:Required
	Stage                      *string `json:"stage"`
	CustomAPIMappingParameters `json:",inline"`
}

// APIMappingSpec defines the desired state of APIMapping
type APIMappingSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  APIMappingParameters `json:"forProvider"`
}

// APIMappingObservation defines the observed state of APIMapping
type APIMappingObservation struct {
	APIMappingID *string `json:"apiMappingID,omitempty"`
}

// APIMappingStatus defines the observed state of APIMapping.
type APIMappingStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     APIMappingObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// APIMapping is the Schema for the APIMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type APIMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIMappingSpec   `json:"spec,omitempty"`
	Status            APIMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIMappingList contains a list of APIMappings
type APIMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIMapping `json:"items"`
}

// Repository type metadata.
var (
	APIMappingKind             = "APIMapping"
	APIMappingGroupKind        = schema.GroupKind{Group: Group, Kind: APIMappingKind}.String()
	APIMappingKindAPIVersion   = APIMappingKind + "." + GroupVersion.String()
	APIMappingGroupVersionKind = GroupVersion.WithKind(APIMappingKind)
)

func init() {
	SchemeBuilder.Register(&APIMapping{}, &APIMappingList{})
}
