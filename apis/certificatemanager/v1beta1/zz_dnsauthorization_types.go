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

type DNSAuthorizationObservation struct {

	// The structure describing the DNS Resource Record that needs to be added
	// to DNS configuration for the authorization to be usable by
	// certificate.
	// Structure is documented below.
	DNSResourceRecord []DNSResourceRecordObservation `json:"dnsResourceRecord,omitempty" tf:"dns_resource_record,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/dnsAuthorizations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DNSAuthorizationParameters struct {

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A domain which is being authorized. A DnsAuthorization resource covers a
	// single domain and its wildcard, e.g. authorization for "example.com" can
	// be used to issue certificates for "example.com" and "*.example.com".
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Set of label tags associated with the DNS Authorization resource.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DNSResourceRecordObservation struct {

	// Data of the DNS Resource Record.
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// Fully qualified name of the DNS Resource Record.
	// E.g. _acme-challenge.example.com.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the DNS Resource Record.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DNSResourceRecordParameters struct {
}

// DNSAuthorizationSpec defines the desired state of DNSAuthorization
type DNSAuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DNSAuthorizationParameters `json:"forProvider"`
}

// DNSAuthorizationStatus defines the observed state of DNSAuthorization.
type DNSAuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DNSAuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DNSAuthorization is the Schema for the DNSAuthorizations API. DnsAuthorization represents a HTTP-reachable backend for a DnsAuthorization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DNSAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DNSAuthorizationSpec   `json:"spec"`
	Status            DNSAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSAuthorizationList contains a list of DNSAuthorizations
type DNSAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSAuthorization `json:"items"`
}

// Repository type metadata.
var (
	DNSAuthorization_Kind             = "DNSAuthorization"
	DNSAuthorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DNSAuthorization_Kind}.String()
	DNSAuthorization_KindAPIVersion   = DNSAuthorization_Kind + "." + CRDGroupVersion.String()
	DNSAuthorization_GroupVersionKind = CRDGroupVersion.WithKind(DNSAuthorization_Kind)
)

func init() {
	SchemeBuilder.Register(&DNSAuthorization{}, &DNSAuthorizationList{})
}
