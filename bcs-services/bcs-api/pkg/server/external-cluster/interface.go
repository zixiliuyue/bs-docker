

package cluster

// ExternalCluster xxx
type ExternalCluster interface {
	BindClusterLb() error
	GetMasterVip() (string, error)
	SyncClusterCredentials() error
}
