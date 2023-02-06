package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Kubeconfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Clusters       []*Cluster `json:"clusters,omitempty"`
	Users          []*User    `json:"users,omitempty"`
	Contexts       []*Context `json:"contexts,omitempty"`
	CurrentContext string     `json:"currentContext,omitempty"`
}

type Cluster struct {
	Server                   string `json:"server,omitempty"`
	CertificateAuthorityData string `json:"certificateAuthorityData,omitempty"`
}

type User struct {
	Name string    `json:"name,omitempty"`
	User *UserInfo `json:"user,omitempty"`
}

type UserInfo struct {
	ClientCertificateData string `json:"clientCertificateData,omitempty"`
	ClientKeyData         string `json:"clientKeyData,omitempty"`
}

type Context struct {
	Cluster string `json:"cluster,omitempty"`
	User    string `json:"user,omitempty"`
}