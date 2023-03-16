package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualMachineSpec defines the desired state of VirtualMachine
type VirtualMachineSpec struct {
	Size         string             `json:"size,omitempty"`
	Firewall     string             `json:"firewall,omitempty"`
	Ipv4         bool               `json:"ipv4,omitempty"`
	Ipv6         bool               `json:"ipv6,omitempty"`
	StaticIP     string             `json:"static_ip,omitempty"`
	Storageclass string             `json:"storageClass,omitempty"`
	Image        string             `json:"image,omitempty"`
	Auth         VirtualMachineAuth `json:"auth,omitempty"`
	Tags         []string           `json:"tags,omitempty"`
	Cloudinit    string             `json:"cloudInit,omitempty"`
	UserScript   string             `json:"userScript,omitempty"`
}

// VirtualMachineAuth defines the auth for the VM
type VirtualMachineAuth struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	SSHKey   string `json:"sshKey,omitempty"`
}

// VirtualMachineStatus defines the observed state of VirtualMachine
type VirtualMachineStatus struct {
	PrivateIP string `json:"privateIP,omitempty"`
	PublicIP  string `json:"publicIP,omitempty"`
	State     string `json:"state,omitempty"`
}

// VirtualMachine is the Schema for the VirtualMachine API
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:"spec,omitempty"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// VirtualMachineList contains a list of VirtualMachine
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}
