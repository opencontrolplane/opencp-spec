package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ObjectStorageCredentialSpec defines the desired state of ObjectStorageCredential
type ObjectStorageCredentialSpec struct {
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
}

// ObjectStorage defines the desired status of ObjectStorage
type ObjectStorageCredentialStatus struct {
	State string `json:"state,omitempty"`
}

// ObjectStorage is the Schema for the ObjectStorage API
type ObjectStorageCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectStorageCredentialSpec   `json:"spec,omitempty"`
	Status ObjectStorageCredentialStatus `json:"status,omitempty"`
}

// ObjectStorage contains a list of ObjectStorage
type ObjectStorageCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectStorageCredential `json:"items"`
}
