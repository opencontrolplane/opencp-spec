package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubernetesClusterSpec defines the desired state of KubernetesCluster
type KubernetesClusterSpec struct {
	Pools       []Pool `json:"pools"`
	Version     string `json:"version"`
	Firewall    string `json:"firewall"`
	CNIPlugin   string `json:"cniPlugin,omitempty"`
	ClusterType string `json:"clusterType,omitempty"`
	Kubeconfig  string `json:"kubeconfig,omitempty"`
}

// Pool is a strcut to manage the Cluster
type Pool struct {
	ID          string `json:"id"`
	Size        string `json:"size"`
	Count       int32  `json:"count"`
	Autoscaling bool   `json:"autoscaling,omitempty"`
	MinSize     int64  `json:"minSize,omitempty"`
	MaxSize     int64  `json:"maxSize,omitempty"`
}

// KubernetesClusterStatus defines the observed state of KubernetesCluster
type KubernetesClusterStatus struct {
	State    string `json:"state,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	PublicIP string `json:"publicIP,omitempty"`
}

// KubernetesCluster is the Schema for the Kubernetes Cluster API
type KubernetesCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubernetesClusterSpec   `json:"spec,omitempty"`
	Status KubernetesClusterStatus `json:"status,omitempty"`
}

// KuberenetesClusterList contains a list of KuberenetesCluster
type KuberenetesClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesCluster `json:"items"`
}
