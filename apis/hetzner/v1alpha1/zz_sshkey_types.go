// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SSHKeyInitParameters struct {

	// User-defined labels (key-value pairs) should be created with.
	// User-defined [labels](https://docs.hetzner.cloud/#labels) (key-value pairs) for the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the SSH key.
	// Name of the SSH key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The public key. If this is a file, it can be read using the file interpolation function
	// Public key of the SSH key pair. If this is a file, it can be read using the `file` interpolation function.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SSHKeyObservation struct {

	// (string) The fingerprint of the SSH key
	// Fingerprint of the SSH public key.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// (int) The unique ID of the key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined labels (key-value pairs) should be created with.
	// User-defined [labels](https://docs.hetzner.cloud/#labels) (key-value pairs) for the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the SSH key.
	// Name of the SSH key.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The public key. If this is a file, it can be read using the file interpolation function
	// Public key of the SSH key pair. If this is a file, it can be read using the `file` interpolation function.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SSHKeyParameters struct {

	// User-defined labels (key-value pairs) should be created with.
	// User-defined [labels](https://docs.hetzner.cloud/#labels) (key-value pairs) for the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the SSH key.
	// Name of the SSH key.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The public key. If this is a file, it can be read using the file interpolation function
	// Public key of the SSH key pair. If this is a file, it can be read using the `file` interpolation function.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

// SSHKeySpec defines the desired state of SSHKey
type SSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SSHKeyInitParameters `json:"initProvider,omitempty"`
}

// SSHKeyStatus defines the observed state of SSHKey.
type SSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SSHKey is the Schema for the SSHKeys API. Provides a Hetzner Cloud SSH key resource to manage SSH keys for server access.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hcloud}
type SSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   SSHKeySpec   `json:"spec"`
	Status SSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKeyList contains a list of SSHKeys
type SSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHKey `json:"items"`
}

// Repository type metadata.
var (
	SSHKey_Kind             = "SSHKey"
	SSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHKey_Kind}.String()
	SSHKey_KindAPIVersion   = SSHKey_Kind + "." + CRDGroupVersion.String()
	SSHKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
}
