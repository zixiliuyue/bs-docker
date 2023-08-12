

package bcsegress

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// init all flags
	viper.SetEnvPrefix("BCS")
	pflag.String("namespace", "bcs-system", "namespace that BCSEgress controller is running in")
	viper.BindEnv("namespace")
	pflag.String("name", "egress-controller", "name that BCSEgress controller is running in with")
	viper.BindEnv("name")
	pflag.String("tamplate", "./template/nginx-template.conf", "tamplate use for proxy configuration generation")
	viper.BindEnv("tamplate")
	pflag.String("generate_dir", "./generate/", "directory for configuration generating")
	viper.BindEnv("generate_dir")
	viper.BindPFlags(pflag.CommandLine)
}

// NewOptionFromFlagAndEnv create option from env or command line
func NewOptionFromFlagAndEnv() *EgressOption {
	egress := &EgressOption{
		Namespace:    viper.GetString("namespace"),
		Name:         viper.GetString("name"),
		TemplateFile: viper.GetString("template"),
		GenerateDir:  viper.GetString("generate_dir"),
	}
	return egress
}

// EgressOption all options that required for BCSEgressController
type EgressOption struct {
	Namespace       string
	Name            string
	TemplateFile    string
	GenerateDir     string
	ProxyExecutable string
	ProxyConfig     string
}
