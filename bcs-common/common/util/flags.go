

package util

import (
	goflag "flag"
	"os"
	"strings"

	"github.com/Tencent/bk-bcs/bcs-common/common/version"

	"github.com/spf13/pflag"
)

// WordSepNormalizeFunc changes all flags that contain "_" separators
func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		return pflag.NormalizedName(strings.Replace(name, "_", "-", -1))
	}
	return pflag.NormalizedName(name)
}

// WarnWordSepNormalizeFunc changes and warns for flags that contain "_" separators
func WarnWordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	if strings.Contains(name, "_") {
		nname := strings.Replace(name, "_", "-", -1)
		// glog.Warningf("%s is DEPRECATED and will be removed in a future version. Use %s instead.", name, nname)

		return pflag.NormalizedName(nname)
	}
	return pflag.NormalizedName(name)
}

// AddCommonFlags add common flags that is needed by all modules
func AddCommonFlags(cmdline *pflag.FlagSet) *bool {
	version := cmdline.Bool("version", false, "show version information")
	return version
}

// InitFlags normalizes and parses the command line flags
func InitFlags() {
	pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	var help bool
	pflag.CommandLine.BoolVarP(&help, "help", "h", false, "show this help info")

	ver := AddCommonFlags(pflag.CommandLine)
	pflag.Parse()

	if help {
		pflag.PrintDefaults()
		os.Exit(0)
	}

	// add handler if flag include --version/-v
	if *ver {
		version.ShowVersion()
		os.Exit(0)
	}
}
