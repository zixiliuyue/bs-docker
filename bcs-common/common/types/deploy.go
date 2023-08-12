

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ConfigSet xxx
type ConfigSet struct {
	Common map[string]interface{}            `json:"common"`
	Conf   map[string]map[string]interface{} `json:"conf"`
}

// ParseConfigSet xxx
func ParseConfigSet(data interface{}) (c *ConfigSet, err error) {
	c = new(ConfigSet)
	var tmp []byte
	if tmp, err = json.Marshal(data); err != nil {
		return c, err
	}
	err = json.Unmarshal(tmp, c)
	return c, err
}

// ClusterSet xxx
type ClusterSet struct {
	ClusterId     string    `json:"clusterId"`
	ClusterConfig ConfigSet `json:"clusterConfig"`
}

// DeployConfig xxx
type DeployConfig struct {
	Service       string       `json:"service"`
	ServiceConfig ConfigSet    `json:"serviceConfig"`
	Clusters      []ClusterSet `json:"clusters"`
	StableVersion string       `json:"stableVersion"`
}

// RenderConfig xxx
type RenderConfig struct {
	MesosZk          string `render:"mesosZkHost"`
	MesosZkSpace     string `render:"mesosZkHostSpace"`
	MesosZkSemicolon string `render:"mesosZkHostSemicolon"`
	MesosZkRaw       string `render:"mesosZkRaw"`
	MesosMaster      string `render:"mesosMaster"`
	MesosQuorum      string `render:"mesosQuorum"`
	Dns              string `render:"dnsUpStream"`
	ClusterId        string `render:"clusterId"`
	ClusterIdNum     string `render:"clusterIdNumber"`
	City             string `render:"city"`
	JfrogUrl         string `render:"jfrogUrl"`
	NeedNat          string `render:"needNat"`
}

var tagFormat = "${%s}"

// Render xxx
func (rc RenderConfig) Render(s string) (r string) {
	r = s

	typeOf := reflect.TypeOf(rc)
	n := typeOf.NumField()
	i := 0
	for i < n {
		field := typeOf.Field(i)
		tag := field.Tag.Get("render")
		value := reflect.ValueOf(rc).FieldByName(field.Name).String()
		r = strings.Replace(r, fmt.Sprintf(tagFormat, tag), value, -1)
		i++
	}
	return r
}
