

package sqlstore

import (
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"strings"

	"github.com/dchest/uniuri"
)

// GetCluster query for the cluster by given clusterId
func GetCluster(clusterId string) *m.Cluster {
	cluster := m.Cluster{}
	GCoreDB.Where(&m.Cluster{ID: clusterId}).First(&cluster)
	if cluster.ID != "" {
		return &cluster
	}
	return nil
}

// GetAllCluster query for all clusters
func GetAllCluster() []m.Cluster {
	clusters := []m.Cluster{}
	GCoreDB.Find(&clusters)

	return clusters
}

// GetClusterByFuzzyClusterId xxx
// GetCluster query for the cluster by given fuzzy clusterId
func GetClusterByFuzzyClusterId(clusterId string) *m.Cluster {
	query := "%" + strings.ToLower(clusterId) + "%"
	cluster := m.Cluster{}
	GCoreDB.Where("id LIKE ?", query).First(&cluster)
	if cluster.ID != "" {
		return &cluster
	}
	return nil
}

// GetClusterByIdentifier query for the cluster by given identifier, which is a random string generated when the
// cluster was created.
func GetClusterByIdentifier(clusterIdentifier string) *m.Cluster {
	cluster := m.Cluster{}
	GCoreDB.Where(&m.Cluster{Identifier: clusterIdentifier}).First(&cluster)
	if cluster.ID != "" {
		return &cluster
	}
	return nil
}

// CreateCluster xxx
func CreateCluster(cluster *m.Cluster) error {
	// Generate a random identifier by default, prepend the clusterID to avoid name conflict
	if cluster.Identifier == "" {
		cluster.Identifier = strings.ToLower(cluster.ID) + "-" + uniuri.NewLen(16)
	}
	err := GCoreDB.Create(cluster).Error
	return err
}
