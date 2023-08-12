

// Package config xxx
package config

// HostConf host conf
type HostConf struct {
	BKIAMHost  string `yaml:"bk_iam_host"`  // 权限中心
	BSCPAPIURL string `yaml:"bscp_api_url"` // bscp api地址
}

// FrontendConf docs and host conf
type FrontendConf struct {
	Docs map[string]string `yaml:"docs"`
	Host *HostConf         `yaml:"hosts"`
}

// defaultFrontendConf 默认配置
func defaultUIConf() *FrontendConf {
	c := &FrontendConf{
		Docs: map[string]string{},
		Host: &HostConf{},
	}
	return c
}
