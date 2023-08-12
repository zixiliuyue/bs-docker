

package models

import "time"

// BCSClusterInfo xxx
type BCSClusterInfo struct {
	ID uint `gorm:"primary_key"`
	// ClusterId is a "one-to-one" field which connects this clsuterInfo to a bke cluster object
	ClusterId        string
	SourceProjectId  string `gorm:"size:32;unique_index:idx_source_cluster_project"`
	SourceClusterId  string `gorm:"size:100;unique_index:idx_source_cluster_project"`
	ClusterType      uint   `gorm:""`
	TkeClusterId     string
	TkeClusterRegion string
	CreatedAt        time.Time
}
