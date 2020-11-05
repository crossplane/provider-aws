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

// IntegrationResponseParameters defines the desired state of IntegrationResponse
type IntegrationResponseParameters struct {
	// Region is which region the IntegrationResponse will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region,omitempty"`

	// +kubebuilder:validation:Required
	APIID *string `json:"apiID"`

	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationID *string `json:"integrationID"`

	// +kubebuilder:validation:Required
	IntegrationResponseKey *string `json:"integrationResponseKey"`

	ResponseParameters map[string]*string `json:"responseParameters,omitempty"`

	ResponseTemplates map[string]*string `json:"responseTemplates,omitempty"`

	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`
}

// IntegrationResponseSpec defines the desired state of IntegrationResponse
type IntegrationResponseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IntegrationResponseParameters `json:"forProvider"`
}

// IntegrationResponseObservation defines the observed state of IntegrationResponse
type IntegrationResponseObservation struct {
	IntegrationResponseID *string `json:"integrationResponseID,omitempty"`
}

// IntegrationResponseStatus defines the observed state of IntegrationResponse.
type IntegrationResponseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IntegrationResponseObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// IntegrationResponse is the Schema for the IntegrationResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IntegrationResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationResponseSpec   `json:"spec,omitempty"`
	Status            IntegrationResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationResponseList contains a list of IntegrationResponses
type IntegrationResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IntegrationResponse `json:"items"`
}

// Repository type metadata.
var (
	IntegrationResponseKind             = "IntegrationResponse"
	IntegrationResponseGroupKind        = schema.GroupKind{Group: Group, Kind: IntegrationResponseKind}.String()
	IntegrationResponseKindAPIVersion   = IntegrationResponseKind + "." + GroupVersion.String()
	IntegrationResponseGroupVersionKind = GroupVersion.WithKind(IntegrationResponseKind)
)

func init() {
	SchemeBuilder.Register(&IntegrationResponse{}, &IntegrationResponseList{})
}
