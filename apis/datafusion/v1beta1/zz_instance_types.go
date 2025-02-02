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

type InstanceObservation struct {

	// The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Cloud Storage bucket generated by Data Fusion in the customer project.
	GcsBucket *string `json:"gcsBucket,omitempty" tf:"gcs_bucket,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/instances/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Endpoint on which the Data Fusion UI and REST APIs are accessible.
	ServiceEndpoint *string `json:"serviceEndpoint,omitempty" tf:"service_endpoint,omitempty"`

	// The current state of this Data Fusion instance.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage *string `json:"stateMessage,omitempty" tf:"state_message,omitempty"`

	// The name of the tenant project.
	TenantProjectID *string `json:"tenantProjectId,omitempty" tf:"tenant_project_id,omitempty"`

	// The time the instance was last updated in RFC3339 UTC "Zulu" format, accurate to nanoseconds.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type InstanceParameters struct {

	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines.
	// +kubebuilder:validation:Optional
	DataprocServiceAccount *string `json:"dataprocServiceAccount,omitempty" tf:"dataproc_service_account,omitempty"`

	// An optional description of the instance.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Option to enable Stackdriver Logging.
	// +kubebuilder:validation:Optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging,omitempty"`

	// Option to enable Stackdriver Monitoring.
	// +kubebuilder:validation:Optional
	EnableStackdriverMonitoring *bool `json:"enableStackdriverMonitoring,omitempty" tf:"enable_stackdriver_monitoring,omitempty"`

	// The resource labels for instance to use to annotate any related underlying resources,
	// such as Compute Engine VMs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkConfig []NetworkConfigParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// Map of additional options used to configure the behavior of Data Fusion instance.
	// +kubebuilder:validation:Optional
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// Specifies whether the Data Fusion instance should be private. If set to
	// true, all Data Fusion nodes will have private IP addresses and will not be
	// able to access the public internet.
	// +kubebuilder:validation:Optional
	PrivateInstance *bool `json:"privateInstance,omitempty" tf:"private_instance,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the Data Fusion instance.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Represents the type of Data Fusion instance. Each type is configured with
	// the default settings for processing and memory.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Current version of the Data Fusion.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NetworkConfigObservation struct {
}

type NetworkConfigParameters struct {

	// The IP range in CIDR notation to use for the managed Data Fusion instance
	// nodes. This range must not overlap with any other ranges used in the Data Fusion instance network.
	// +kubebuilder:validation:Required
	IPAllocation *string `json:"ipAllocation" tf:"ip_allocation,omitempty"`

	// Name of the network in the project with which the tenant project
	// will be peered for executing pipelines. In case of shared VPC where the network resides in another host
	// project the network should specified in the form of projects/{host-project-id}/global/networks/{network}
	// +kubebuilder:validation:Required
	Network *string `json:"network" tf:"network,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. Represents a Data Fusion instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
