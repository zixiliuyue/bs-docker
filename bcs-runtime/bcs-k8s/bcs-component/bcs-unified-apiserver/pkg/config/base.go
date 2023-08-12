

package config

import (
	"time"
)

const (
	// DevEnv xxx
	DevEnv = "dev"
	// StagEnv xxx
	StagEnv = "stag"
	// ProdEnv xxx
	ProdEnv = "prod"
)

// BaseConf xxx
type BaseConf struct {
	AppCode      string              `yaml:"app_code"`
	AppSecret    string              `yaml:"app_secret"`
	TimeZone     string              `yaml:"time_zone"`
	LanguageCode string              `yaml:"language_code"`
	Managers     []string            `yaml:"managers"`
	ManagerMap   map[string]struct{} `yaml:"-"`
	Debug        bool                `yaml:"debug"`
	RunEnv       string              `yaml:"run_env"`
	Location     *time.Location      `yaml:"-"`
}

// Init xxx
func (c *BaseConf) Init() error {
	var err error
	c.AppCode = ""
	c.AppSecret = ""
	c.TimeZone = "Asia/Shanghai"
	c.LanguageCode = ""
	c.Managers = []string{}
	c.ManagerMap = map[string]struct{}{}
	c.Debug = false
	c.RunEnv = DevEnv
	c.Location, err = time.LoadLocation(c.TimeZone)
	if err != nil {
		return err
	}
	return nil
}

// InitManagers xxx
func (c *BaseConf) InitManagers() error {
	for _, manager := range c.Managers {
		c.ManagerMap[manager] = struct{}{}
	}
	return nil
}

// IsManager xxx
func (c *BaseConf) IsManager(username string) bool {
	_, ok := c.ManagerMap[username]
	return ok
}
