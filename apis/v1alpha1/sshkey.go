package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SSHKeySpec defines the desired state of SSHKey
type SSHKeySpec struct {
	PublicKey string `json:"publickey"`
}

// SSHKeyStatus
type SSHKeyStatus struct {
	State       string `json:"state,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

// SSHKey is the Schema for the SSHKey API
type SSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SSHKeySpec   `json:"spec,omitempty"`
	Status SSHKeyStatus `json:"status,omitempty"`
}

// SSHKeyList contains a list of SSHKey
type SSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHKey `json:"items"`
}
