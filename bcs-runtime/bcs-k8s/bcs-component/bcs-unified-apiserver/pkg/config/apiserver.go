

package config

const (
	// BcsStorageMode bcs storage mode
	BcsStorageMode = "bcs-storage"
	// KubernetesMode kubernetes mode
	KubernetesMode = "kubernetes"
)

// APIServer 当前选择的集群ID
type APIServer struct {
	ClusterId string `yaml:"cluster_id"`
	StoreMode string `yaml:"store_mode"`
}
