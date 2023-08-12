

package sqlstore

import (
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
)

// GetCredentials query for clusterCredentials by clusterId
func GetCredentials(clusterId string) *m.ClusterCredentials {
	credentials := m.ClusterCredentials{}
	GCoreDB.Where(&m.ClusterCredentials{ClusterId: clusterId}).First(&credentials)
	if credentials.ID != 0 {
		return &credentials
	}
	return nil
}

// SaveCredentials saves the current cluster credentials
func SaveCredentials(clusterId, serverAddresses, caCertData, userToken, clusterDomain string) error {
	var credentials m.ClusterCredentials
	// Create or update, source: https://github.com/jinzhu/gorm/issues/1307
	dbScoped := GCoreDB.Where(m.ClusterCredentials{ClusterId: clusterId}).Assign(
		m.ClusterCredentials{
			ServerAddresses: serverAddresses,
			CaCertData:      caCertData,
			UserToken:       userToken,
			ClusterDomain:   clusterDomain,
		},
	).FirstOrCreate(&credentials)
	return dbScoped.Error
}

// SaveWsCredentials saves the credentials of cluster registered by websocket
func SaveWsCredentials(serverKey, clientModule, serverAddress, caCertData, userToken string) error {
	var credentials m.WsClusterCredentials
	// Create or update, source: https://github.com/jinzhu/gorm/issues/1307
	dbScoped := GCoreDB.Where(m.WsClusterCredentials{ServerKey: serverKey}).Assign(
		m.WsClusterCredentials{
			ClientModule:  clientModule,
			ServerAddress: serverAddress,
			CaCertData:    caCertData,
			UserToken:     userToken,
		},
	).FirstOrCreate(&credentials)
	return dbScoped.Error
}

// GetWsCredentials query for clusterCredentials of cluster registered by websocket
func GetWsCredentials(serverKey string) *m.WsClusterCredentials {
	credentials := m.WsClusterCredentials{}
	GCoreDB.Where(&m.WsClusterCredentials{ServerKey: serverKey}).First(&credentials)
	if credentials.ID != 0 {
		return &credentials
	}
	return nil
}

// DelWsCredentials xxx
func DelWsCredentials(serverKey string) {
	credentials := m.WsClusterCredentials{}
	GCoreDB.Where(&m.WsClusterCredentials{ServerKey: serverKey}).Delete(&credentials)
}

// GetWsCredentialsByClusterId xxx
func GetWsCredentialsByClusterId(clusterId string) []*m.WsClusterCredentials {
	var credentials []*m.WsClusterCredentials
	query := clusterId + "-%"
	GCoreDB.Where("server_key LIKE ?", query).Find(&credentials)
	return credentials
}
