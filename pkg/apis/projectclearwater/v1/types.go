package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SproutClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SproutCluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SproutCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SproutClusterSpec   `json:"spec"`
	Status            SproutClusterStatus `json:"status,omitempty"`
}

type SproutClusterSpec struct {
	Shards int32 `json:"shards"`
	Scale int32 `json:"scale"`
}

type SproutClusterStatus struct {
	ShardNodes []string `json:"shardNodes"`
	BonoNodes []string `json:"bonoNodes"`
}
