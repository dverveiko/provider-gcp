/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TopicIAMBindingConditionObservation struct {
}

type TopicIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type TopicIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TopicIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TopicIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Topic
	// +crossplane:generate:reference:refFieldName=TopicRef
	// +crossplane:generate:reference:selectorFieldName=TopicSelector
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

// TopicIAMBindingSpec defines the desired state of TopicIAMBinding
type TopicIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicIAMBindingParameters `json:"forProvider"`
}

// TopicIAMBindingStatus defines the observed state of TopicIAMBinding.
type TopicIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicIAMBinding is the Schema for the TopicIAMBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TopicIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicIAMBindingSpec   `json:"spec"`
	Status            TopicIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicIAMBindingList contains a list of TopicIAMBindings
type TopicIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	TopicIAMBinding_Kind             = "TopicIAMBinding"
	TopicIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicIAMBinding_Kind}.String()
	TopicIAMBinding_KindAPIVersion   = TopicIAMBinding_Kind + "." + CRDGroupVersion.String()
	TopicIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(TopicIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicIAMBinding{}, &TopicIAMBindingList{})
}
