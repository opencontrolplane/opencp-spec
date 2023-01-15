package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DomainRecords is the Schema for the DomainRecords API
type DomainRecords struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Type     string `json:"type"`
	Priority *int32 `json:"priority,omitempty"`
	TTL      *int32 `json:"ttl"`
}

// DomainSpec defines the desired state of DNS
type DomainSpec struct {
	Records []DomainRecords `json:"records"`
}

// DomainStatus defines the observed state of DNS
type DomainStatus struct {
	State string `json:"state,omitempty"`
}

// Domain is the Schema for the DNS API
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *DomainSpec   `json:"spec,omitempty"`
	Status            *DomainStatus `json:"status,omitempty"`
}

// DNSList contains a list of DNS
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}
