

package config

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
)

// BCSClusterEnv xxx
type BCSClusterEnv string

const (
	// ProdCluster xxx
	ProdCluster BCSClusterEnv = "prod" // 正式环境
	// DebugCLuster xxx
	DebugCLuster BCSClusterEnv = "debug" // debug 环境
	// UatCluster xxx
	UatCluster BCSClusterEnv = "uat" // uat 环境
)

// BCSConf 配置
type BCSConf struct {
	Host         string         `yaml:"host"`
	Token        string         `yaml:"token"`
	Verify       bool           `yaml:"verify"`
	JWTPubKey    string         `yaml:"jwt_public_key"`
	JWTPubKeyObj *rsa.PublicKey `yaml:"-"`
	ClusterEnv   BCSClusterEnv  `yaml:"cluster_env"`
}

// Init xxx
func (c *BCSConf) Init() error {
	// only for development
	c.Host = ""
	c.Token = ""
	c.JWTPubKey = ""
	c.JWTPubKeyObj = nil
	c.Verify = false
	c.ClusterEnv = ProdCluster
	return nil
}

// InitJWTPubKey xxx
func (c *BCSConf) InitJWTPubKey() error {
	if c.JWTPubKey == "" {
		return nil
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(c.JWTPubKey))
	if err != nil {
		return err
	}

	c.JWTPubKeyObj = pubKey
	return nil
}
