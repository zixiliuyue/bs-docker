

package config

import (
	"path/filepath"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// LogConf : config for logging
type LogConf struct {
	Level    string `yaml:"level"`
	File     string `yaml:"file"`
	Stderr   bool   `yaml:"stderr"`
	CmdFile  string `yaml:"-"`
	CmdLevel string `yaml:"-"`
}

// Init : init default logging config
func (c *LogConf) Init() error {
	// only for development
	c.Level = "info"
	c.File = ""
	c.Stderr = true
	c.CmdFile = ""
	c.CmdLevel = "info"

	return nil
}

// InitBlog 初始化 blog 模块, 注意只能初始化一次
func (c *LogConf) InitBlog() error {
	var blogLevel int32
	blogDir := ""

	// blog 只有 0 和 3 个等级
	switch c.Level {
	case "debug":
		blogLevel = 3
	default:
		blogLevel = 0
	}

	// 不会自动创建目录, 需要管理员手动创建
	if c.File != "" {
		logFile, err := filepath.Abs(c.File)
		if err != nil {
			return err
		}

		blogDir = filepath.Dir(logFile)
	}

	blogConf := conf.LogConfig{
		Verbosity:    blogLevel,
		AlsoToStdErr: c.Stderr,
		LogDir:       blogDir,
		LogMaxSize:   500,
		LogMaxNum:    10,
	}

	blog.InitLogs(blogConf)
	return nil
}
