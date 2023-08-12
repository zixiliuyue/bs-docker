

// Package sqlstore is main SQL database storage
package sqlstore

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/config"
	"github.com/jinzhu/gorm"
	// import empty mysql package
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// GCoreDB xxx
var GCoreDB *gorm.DB

// InitCoreDatabase initialize the GLOBAL database object
func InitCoreDatabase(conf *config.ApiServConfig) error {
	if conf == nil {
		return fmt.Errorf("core_database config not init")
	}

	dsn := conf.BKE.DSN
	if dsn == "" {
		return fmt.Errorf("core_database dsn not configured")
	}
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.DB().SetConnMaxLifetime(60 * time.Second)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(20)

	GCoreDB = db
	return nil
}
