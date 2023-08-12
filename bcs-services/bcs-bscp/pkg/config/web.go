

package config

import (
	"net/url"
	"path"
)

// WebConf web 相关配置
type WebConf struct {
	Host             string   `yaml:"host"`
	RoutePrefix      string   `yaml:"route_prefix"` // vue路由, 静态资源前缀
	PreferredDomains string   `yaml:"preferred_domains"`
	BaseURL          *url.URL `yaml:"-"`
}

// init 初始化
func (c *WebConf) init() error {
	u, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, c.RoutePrefix)

	c.BaseURL = u
	return nil
}

// defaultWebConf 默认配置
func defaultWebConf() *WebConf {
	c := &WebConf{
		Host:             "",
		RoutePrefix:      "",
		PreferredDomains: "",
	}
	return c
}
