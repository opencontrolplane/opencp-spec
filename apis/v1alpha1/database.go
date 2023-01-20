package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DomainSpec defines the desired state of DNS
type DatabaseSpec struct {
	Nodes         int    `json:"nodes,omitempty"`
	Size          string `json:"size,omitempty"`
	Firewall      string `json:"firewall,omitempty"`
	Engine        string `json:"engine,omitempty"`
	EngineVersion string `json:"engineVersion,omitempty"`
}

// DatabaseStatus defines the observed state of Database
type DatabaseStatus struct {
	Username string `json:"username,omitempty"`
	// Password string `json:"password,omitempty"`
	Port     int    `json:"port,omitempty"`
	PublicIP string `json:"publicIP,omitempty"`
	State    string `json:"state,omitempty"`
}

// Database is the Schema for the databases API
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec,omitempty"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// DatabaseList contains a list of Database
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}
