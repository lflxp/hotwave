package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PlaybookRunSpec defines the desired state of PlaybookRun
type PlaybookRunSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of PlaybookRun. Edit PlaybookRun_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// PlaybookRunStatus defines the observed state of PlaybookRun
type PlaybookRunStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// PlaybookRun is the Schema for the playbookruns API
type PlaybookRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlaybookRunSpec   `json:"spec,omitempty"`
	Status PlaybookRunStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlaybookRunList contains a list of PlaybookRun
type PlaybookRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlaybookRun `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PlaybookRun{}, &PlaybookRunList{})
}
