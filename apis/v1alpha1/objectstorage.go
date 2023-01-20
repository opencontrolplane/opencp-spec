package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ObjectStorage defines the desired state of ObjectStorage
type ObjectStorageSpec struct {
	Size              int    `json:"size,omitempty"`
	StorageCredential string `json:"storageCredential,omitempty"`
}

// ObjectStorage defines the desired status of ObjectStorage
type ObjectStorageStatus struct {
	State             string `json:"state,omitempty"`
	EndPoint          string `json:"endpoint,omitempty"`
	StorageCredential string `json:"storageCredential,omitempty"`
}

// ObjectStorage is the Schema for the ObjectStorage API
type ObjectStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObjectStorageSpec   `json:"spec,omitempty"`
	Status ObjectStorageStatus `json:"status,omitempty"`
}

// ObjectStorage contains a list of ObjectStorage
type ObjectStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectStorage `json:"items"`
}
