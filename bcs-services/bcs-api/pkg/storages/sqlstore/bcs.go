

package sqlstore

import (
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
)

// CreateBCSClusterInfo xxx
func CreateBCSClusterInfo(external *m.BCSClusterInfo) error {
	err := GCoreDB.Create(external).Error
	return err
}

// QueryBCSClusterInfo xxx
// Query BCSClusterInfo search a BCSClusterInfo object using given conditions
func QueryBCSClusterInfo(info *m.BCSClusterInfo) *m.BCSClusterInfo {
	result := m.BCSClusterInfo{}
	GCoreDB.Where(info).First(&result)
	if result.ID != 0 {
		return &result
	}
	return nil

}

// GetClusterByBCSInfo query for the cluster by given clusterId
func GetClusterByBCSInfo(sourceProjectID string, sourceClusterID string) *m.Cluster {
	var externalClusterInfo *m.BCSClusterInfo
	if sourceProjectID != "" {
		externalClusterInfo = QueryBCSClusterInfo(&m.BCSClusterInfo{
			SourceProjectId: sourceProjectID,
			SourceClusterId: sourceClusterID,
		})
	} else {
		externalClusterInfo = QueryBCSClusterInfo(&m.BCSClusterInfo{
			SourceClusterId: sourceClusterID,
		})
	}

	if externalClusterInfo == nil {
		return nil
	}

	cluster := m.Cluster{}
	GCoreDB.Where(&m.Cluster{ID: externalClusterInfo.ClusterId}).First(&cluster)
	if cluster.ID != "" {
		return &cluster
	}
	return nil
}
