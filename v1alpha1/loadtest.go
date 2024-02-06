package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// LoadTestSpec defines the desired state of LoadTest
type LoadTestSpec struct {
	Script    string `json:"script"`
	TargetURL string `json:"targetURL"`
	Users     int    `json:"users"`
	Duration  int    `json:"duration"`
}

// LoadTestStatus defines the observed state of LoadTest
type LoadTestStatus struct {
	// Add your status fields here
}

// +kubebuilder:object:root=true

// LoadTest is the Schema for the loadtests API
type LoadTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadTestSpec   `json:"spec,omitempty"`
	Status LoadTestStatus `json:"status,omitempty"`
}

// DeepCopy creates a deep copy of the LoadTest struct.
func (lt *LoadTest) DeepCopy() *LoadTest {
	if lt == nil {
		return nil
	}

	copied := &LoadTest{
		TypeMeta:   lt.TypeMeta,
		ObjectMeta: lt.ObjectMeta,
		Spec:       lt.Spec,
		Status:     lt.Status,
	}
	return copied
}

// DeepCopyObject implements client.Object.
func (lt *LoadTest) DeepCopyObject() runtime.Object {
	return lt.DeepCopy()
}

// +kubebuilder:object:root=true

// LoadTestList contains a list of LoadTest
type LoadTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadTest `json:"items"`
}
