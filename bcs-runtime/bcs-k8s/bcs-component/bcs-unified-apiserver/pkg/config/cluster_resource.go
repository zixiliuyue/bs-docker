

package config

// ClusterKind 集群类型
type ClusterKind string

var (
	// IsolatedCLuster isolated cluster
	IsolatedCLuster ClusterKind = "isolated"
	// SharedCluster shared cluster
	SharedCluster ClusterKind = "shared"
	// FederatedCluter federated cluster
	FederatedCluter ClusterKind = "federated"
)

// ClusterResource 集群配置
type ClusterResource struct {
	Kind      string   `yaml:"kind"`
	Member    string   `yaml:"member"`
	Members   []string `yaml:"members"`
	Master    string   `yaml:"master"`
	ClusterId string   `yaml:"cluster_id"`
}
