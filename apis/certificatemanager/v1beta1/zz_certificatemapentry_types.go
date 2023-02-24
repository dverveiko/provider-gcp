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

type CertificateMapEntryObservation struct {

	// Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A serving state of this Certificate Map Entry.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateMapEntryParameters struct {

	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects//locations//certificates/*.
	// +kubebuilder:validation:Required
	Certificates []*string `json:"certificates" tf:"certificates,omitempty"`

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A map entry that is inputted into the cetrificate map
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.CertificateMap
	// +kubebuilder:validation:Optional
	Map *string `json:"map,omitempty" tf:"map,omitempty"`

	// Reference to a CertificateMap in certificatemanager to populate map.
	// +kubebuilder:validation:Optional
	MapRef *v1.Reference `json:"mapRef,omitempty" tf:"-"`

	// Selector for a CertificateMap in certificatemanager to populate map.
	// +kubebuilder:validation:Optional
	MapSelector *v1.Selector `json:"mapSelector,omitempty" tf:"-"`

	// A predefined matcher for particular cases, other than SNI selection
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// CertificateMapEntrySpec defines the desired state of CertificateMapEntry
type CertificateMapEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateMapEntryParameters `json:"forProvider"`
}

// CertificateMapEntryStatus defines the observed state of CertificateMapEntry.
type CertificateMapEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateMapEntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateMapEntry is the Schema for the CertificateMapEntrys API. CertificateMapEntry is a list of certificate configurations, that have been issued for a particular hostname
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CertificateMapEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateMapEntrySpec   `json:"spec"`
	Status            CertificateMapEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateMapEntryList contains a list of CertificateMapEntrys
type CertificateMapEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateMapEntry `json:"items"`
}

// Repository type metadata.
var (
	CertificateMapEntry_Kind             = "CertificateMapEntry"
	CertificateMapEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateMapEntry_Kind}.String()
	CertificateMapEntry_KindAPIVersion   = CertificateMapEntry_Kind + "." + CRDGroupVersion.String()
	CertificateMapEntry_GroupVersionKind = CRDGroupVersion.WithKind(CertificateMapEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateMapEntry{}, &CertificateMapEntryList{})
}
