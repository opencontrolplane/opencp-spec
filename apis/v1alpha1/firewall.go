package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	IngressRules []FirewallRules `json:"ingress"`
	EgressRules  []FirewallRules `json:"egress"`
}

// Rules is the struct to define the rules
type FirewallRules struct {
	Source   []string `json:"source"`
	Label    string   `json:"label,omitempty"`
	Ports    string   `json:"ports"`
	Protocol string   `json:"protocol"`
	Action   string   `json:"action"`
}

// FirewallStatus
type FirewallStatus struct {
	State      string `json:"state,omitempty"`
	TotalRules string `json:"totalrules,omitempty"`
}

// Firewall is the Schema for the Firewall API
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirewallSpec   `json:"spec,omitempty"`
	Status FirewallStatus `json:"status,omitempty"`
}

// FirewallList contains a list of Firewall
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}
