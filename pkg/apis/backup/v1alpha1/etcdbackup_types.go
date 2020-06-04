package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EtcdBackupSpec defines the desired state of EtcdBackup
type EtcdBackupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Resources    *v1.ResourceRequirements `json:"resources"`
	NodeSelector map[string]string        `json:"nodeSelector,omitempty"`
	Tolerations  []v1.Toleration          `json:"tolerations,omitempty"`
	Schedule     string                   `json:"schedule"`
}

// EtcdBackupStatus defines the observed state of EtcdBackup
type EtcdBackupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	CronJob   string `json:"cronJobs"`
	Schedule  string `json:"schedules"`
	Suspended bool   `json:"suspended"`
	//Conditions map[string]ClusterConditions `json:"clusterCondition,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdBackup is the Schema for the etcdbackups API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=etcdbackups,scope=Namespaced
type EtcdBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdBackupSpec   `json:"spec,omitempty"`
	Status EtcdBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdBackupList contains a list of EtcdBackup
type EtcdBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EtcdBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EtcdBackup{}, &EtcdBackupList{})
}
