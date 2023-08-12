

package config

import (
	"time"
	_ "time/tzdata" // tzdata ..
)

const (
	// DevEnv ..
	DevEnv = "dev"
	// StagEnv ..
	StagEnv = "stag"
	// ProdEnv ..
	ProdEnv = "prod"
)

// BaseConf :
type BaseConf struct {
	AppCode      string         `yaml:"app_code"`
	AppSecret    string         `yaml:"app_secret"`
	BKPaaSHost   string         `yaml:"bk_paas_host"` // esb 调用地址
	TimeZone     string         `yaml:"time_zone"`
	LanguageCode string         `yaml:"language_code"`
	RunEnv       string         `yaml:"run_env"`
	Location     *time.Location `yaml:"-"`
}

// Init :
func (c *BaseConf) Init() error {
	var err error
	c.TimeZone = "Asia/Shanghai"
	c.LanguageCode = "en-us"
	c.RunEnv = DevEnv
	c.Location, err = time.LoadLocation(c.TimeZone)
	if err != nil {
		return err
	}
	return nil
}

// EtcdConf defines etcd related runtime
type EtcdConf struct {
	Endpoints string `yaml:"endpoints"`
	CA        string `yaml:"ca"`
	Cert      string `yaml:"cert"`
	Key       string `yaml:"key"`
}

// defaultEtcdConf
func defaultEtcdConf() *EtcdConf {
	return &EtcdConf{
		Endpoints: "127.0.0.1:2379",
		CA:        "",
		Cert:      "",
		Key:       "",
	}
}
