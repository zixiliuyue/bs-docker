

package sqlstore

import (
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
)

const (
	// CidrStatusAvailable xxx
	CidrStatusAvailable = "available"
	// CidrStatusUsed xxx
	CidrStatusUsed = "used"
	// CidrStatusReserved xxx
	CidrStatusReserved = "reserved"
)

// CidrCount xxx
type CidrCount struct {
	Count    int    `json:"count"`
	Vpc      string `json:"vpc"`
	IpNumber uint   `json:"ip_number"`
	Status   string `json:"status"`
}

// SaveTkeLbSubnet xxx
func SaveTkeLbSubnet(clusterRegion, subnetId string) error {
	var tkeLbSubnet m.TkeLbSubnet
	dbScoped := GCoreDB.Where(m.TkeLbSubnet{ClusterRegion: clusterRegion}).Assign(
		m.TkeLbSubnet{
			SubnetId: subnetId,
		},
	).FirstOrCreate(&tkeLbSubnet)

	return dbScoped.Error
}

// GetSubnetByClusterRegion xxx
func GetSubnetByClusterRegion(clusterRegion string) *m.TkeLbSubnet {
	tkeLbSubnet := m.TkeLbSubnet{}
	GCoreDB.Where(&m.TkeLbSubnet{ClusterRegion: clusterRegion}).First(&tkeLbSubnet)
	if tkeLbSubnet.ID != 0 {
		return &tkeLbSubnet
	}
	return nil
}

// QueryTkeCidr xxx
func QueryTkeCidr(tkeCidr *m.TkeCidr) *m.TkeCidr {
	result := m.TkeCidr{}
	GCoreDB.Where(tkeCidr).First(&result)
	if result.ID != 0 {
		return &result
	}
	return nil

}

// SaveTkeCidr xxx
func SaveTkeCidr(vpc, cidr string, ipNumber uint, status, cluster string) error {
	tkeCidr := &m.TkeCidr{
		Vpc:      vpc,
		Cidr:     cidr,
		IpNumber: ipNumber,
		Status:   status,
		Cluster:  &cluster,
	}

	err := GCoreDB.Create(tkeCidr).Error
	return err
}

// UpdateTkeCidr xxx
func UpdateTkeCidr(tkeCidr, updatedTkeCidr *m.TkeCidr) error {
	err := GCoreDB.Model(tkeCidr).Updates(*updatedTkeCidr).Error
	return err
}

// CountTkeCidr xxx
func CountTkeCidr() []CidrCount {
	var cidrCounts []CidrCount
	GCoreDB.Table("tke_cidrs").Select("count(*) as count, vpc, ip_number, status").Group("vpc, ip_number, status").
		Scan(&cidrCounts)
	return cidrCounts
}
