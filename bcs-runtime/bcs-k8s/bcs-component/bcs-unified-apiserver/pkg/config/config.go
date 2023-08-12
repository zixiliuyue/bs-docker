

// Package config xxx
package config

import (
	"sync"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Configurations : manage all configurations
type Configurations struct {
	mtx                sync.Mutex
	Base               *BaseConf                   `yaml:"base_conf"`
	Logging            *LogConf                    `yaml:"logging"`
	BCS                *BCSConf                    `yaml:"bcs_conf"`
	APIServer          *APIServer                  `yaml:"apiserver"`
	BCSEnvMap          map[BCSClusterEnv]*BCSConf  `yaml:"-"`
	BCSEnvConf         []*BCSConf                  `yaml:"bcs_env_conf"`
	ClusterResources   []*ClusterResource          `yaml:"cluster_resources"`
	ClusterResourceMap map[string]*ClusterResource `yaml:"-"`
}

// Init xxx
// ReadFrom : read from file
func (c *Configurations) Init() error {
	c.Base = &BaseConf{}
	c.Base.Init()

	// logging
	c.Logging = &LogConf{}
	c.Logging.Init()

	// BCS Config
	c.BCS = &BCSConf{}
	c.BCS.Init()

	c.BCSEnvConf = []*BCSConf{}
	c.BCSEnvMap = map[BCSClusterEnv]*BCSConf{}

	return nil
}

// IsDevMode 是否本地开发模式
func (c *Configurations) IsDevMode() bool {
	return c.Base.RunEnv == DevEnv
}

// G : Global Configurations
var G = &Configurations{}

// init 初始化
func init() {
	G.Init()
}

// ReadFrom : read from file
func (c *Configurations) ReadFrom(content []byte) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if len(content) == 0 {
		return errors.New("conf content is empty, will use default values")
	}

	err := yaml.Unmarshal(content, &G)
	if err != nil {
		return err
	}
	c.Logging.InitBlog()
	c.Base.InitManagers()
	c.InitClusterResources()

	// 把列表类型转换为map，方便检索
	for _, conf := range c.BCSEnvConf {
		c.BCSEnvMap[conf.ClusterEnv] = conf
	}

	if err := c.BCS.InitJWTPubKey(); err != nil {
		return err
	}

	return nil
}

// InitClusterResources xxx
func (c *Configurations) InitClusterResources() error {
	c.ClusterResourceMap = make(map[string]*ClusterResource)
	for _, v := range c.ClusterResources {
		c.ClusterResourceMap[v.ClusterId] = v
	}
	return nil
}

// GetMember xxx
func (c *Configurations) GetMember(clusterId string) (*ClusterResource, bool) {
	resource, ok := c.ClusterResourceMap[clusterId]
	if !ok {
		return nil, false
	}
	return resource, true
}
