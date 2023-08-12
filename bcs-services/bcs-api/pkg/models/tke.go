

package models

import "time"

// TkeLbSubnet xxx
type TkeLbSubnet struct {
	ID            uint   `gorm:"primary_key"`
	ClusterRegion string `gorm:"unique;not null"`
	SubnetId      string `gorm:"size:256;not null"`
}

// TkeCidr xxx
type TkeCidr struct {
	ID        uint   `gorm:"primary_key"`
	Vpc       string `gorm:"not null"`
	Cidr      string `gorm:"not null"`
	IpNumber  uint   `gorm:"not null"`
	Status    string `gorm:"not null"`
	Cluster   *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
