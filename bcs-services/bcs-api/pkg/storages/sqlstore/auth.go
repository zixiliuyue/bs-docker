

package sqlstore

import (
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"time"
)

// GetUserClusterPerm xxx
func GetUserClusterPerm(user *m.User, cluster *m.Cluster, name string) *m.UserClusterPermission {
	var result m.UserClusterPermission
	GCoreDB.Where(m.UserClusterPermission{UserID: user.ID, ClusterID: cluster.ID, Name: name}).First(&result)
	if result.ID == 0 {
		return nil
	}
	return &result
}

// SaveUserClusterPerm saves a user cluster permission to database, it will update the `UpdatedAt` filed to
// current timestamp.
func SaveUserClusterPerm(backend int, user *m.User, cluster *m.Cluster, name string, isActive bool) error {
	var result m.UserClusterPermission
	// Create or update, source: https://github.com/jinzhu/gorm/issues/1307
	dbScoped := GCoreDB.Where(m.UserClusterPermission{
		Backend:   backend,
		UserID:    user.ID,
		ClusterID: cluster.ID,
		Name:      name,
	}).Assign(
		m.UserClusterPermission{
			IsActive:  isActive,
			UpdatedAt: time.Now(),
		},
	).FirstOrCreate(&result)
	return dbScoped.Error
}
