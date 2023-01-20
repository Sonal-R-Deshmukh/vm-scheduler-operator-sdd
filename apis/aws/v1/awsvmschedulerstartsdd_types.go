/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AWSVMSchedulerStartsddSpec defines the desired state of AWSVMSchedulerStartsdd
type AWSVMSchedulerStartsddSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	InstanceIds string `json:"instanceIds"`

	// Schedule period for the CronJob.
	// This spec allows you to setup the start schedule for VM
	StartSchedule string `json:"startSchedule"`

	Image string `json:"image"`
}

// AWSVMSchedulerStartsddStatus defines the observed state of AWSVMSchedulerStartsdd
type AWSVMSchedulerStartsddStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	VMStartStatus string `json:"vmStartStatus"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AWSVMSchedulerStartsdd is the Schema for the awsvmschedulerstartsdds API
type AWSVMSchedulerStartsdd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AWSVMSchedulerStartsddSpec   `json:"spec,omitempty"`
	Status AWSVMSchedulerStartsddStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AWSVMSchedulerStartsddList contains a list of AWSVMSchedulerStartsdd
type AWSVMSchedulerStartsddList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AWSVMSchedulerStartsdd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AWSVMSchedulerStartsdd{}, &AWSVMSchedulerStartsddList{})
}
