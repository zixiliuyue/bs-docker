

// Package server xxx
package server

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/config"
	m "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/models"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/rbac"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/rbac/template"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/storages/sqlstore"
	sqlutils "github.com/Tencent/bk-bcs/bcs-services/bcs-api/pkg/storages/sqlstore/utils"
)

// Setup do initialize jobs for server
func Setup(conf *config.ApiServConfig) {

	if err := sqlstore.InitCoreDatabase(conf); err != nil {
		blog.Fatalf("unable to connect to core database: %s", err.Error())
	}
	// Migrate db schemas
	sqlstore.GCoreDB.AutoMigrate(
		// Auth
		&m.User{},
		&m.UserToken{},
		&m.UserClusterPermission{},
		&m.ExternalUserRecord{},
		// Cluster
		&m.Cluster{},
		&m.ClusterCredentials{},
		&m.RegisterToken{},
		&m.WsClusterCredentials{},

		// BCS
		&m.BCSClusterInfo{},

		// Network
		&m.TkeLbSubnet{},
		&m.TkeCidr{},
	)

	if conf != nil {
		// Migrate db data
		sqlutils.CreateBootstrapUsers(conf.BKE.BootStrapUsers)
	}
}

// StartRbacSync xxx
// sync rbac data to k8s clusters
func StartRbacSync(conf *config.ApiServConfig) {
	template.InitRbacTemplates()

	if conf == nil {
		return
	}

	// 检查是否从配置文件读取 rbac数据
	if conf.BKE.TurnOnConf {
		blog.Info("read rbac data from conf")
		go rbac.SyncRbacFromConf(conf.BKE.RbacDatas)
	}

	if conf.BKE.TurnOnAuth {
		go rbac.SyncRbacFromAuth()
	}
}
