

// Package types xxx
package types

import (
	bcsv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/apis/bkbcs/v1"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-logbeat-sidecar/metric"
)

// Yaml is the structure viewed of log config file, contains metric info
type Yaml struct {
	Local           []Local                 `yaml:"local"`
	Metric          *metric.LogFileInfoType `yaml:"-"`
	BCSLogConfigKey string                  `yaml:"-"`
}

// Local is a single log collection task with single dataid
type Local struct {
	DataID       int                  `yaml:"dataid"`
	OutputFormat string               `yaml:"output_format"`
	Paths        []string             `yaml:"paths"`
	ToJSON       bool                 `yaml:"to_json"`
	Package      bool                 `yaml:"package"`
	ExtMeta      map[string]string    `yaml:"ext_meta"`
	CloseEOF     *bool                `yaml:"close_eof,omitempty"`
	CloseTimeout string               `yaml:"close_timeout,omitempty"`
	Multiline    *bcsv1.MultilineConf `yaml:"multiline,omitempty"`

	// stdout dataid
	StdoutDataid string `yaml:"-"`
	// nonstandard log dataid
	NonstandardDataid string `yaml:"-"`
	// nonstandard paths
	NonstandardPaths []string `yaml:"-"`
	// host paths
	HostPaths []string `yaml:"-"`
	// log tags
	LogTags map[string]string `yaml:"-"`
}
