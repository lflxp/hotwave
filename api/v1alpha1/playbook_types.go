package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PlaybookSpec defines the desired state of Playbook
type PlaybookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Playbook. Edit Playbook_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// PlaybookStatus defines the observed state of Playbook
type PlaybookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Playbook is the Schema for the playbooks API
type Playbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlaybookSpec   `json:"spec,omitempty"`
	Status PlaybookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlaybookList contains a list of Playbook
type PlaybookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Playbook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Playbook{}, &PlaybookList{})
}
