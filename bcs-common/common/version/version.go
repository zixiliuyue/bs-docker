

// Package version xxx
package version

import (
	"fmt"
	"runtime"
)

// BcsVersion describes version
// BcsTag show the git tag for this version
// BcsBuildTime show the compile time
var (
	BcsVersion   = "17.03.28"
	BcsTag       = "2017-03-28 Release"
	BcsBuildTime = "2017-03-28 19:50:00"
	BcsGitHash   = "unknown"
	GoVersion    = runtime.Version()
)

// ShowVersion is the default handler which match the --version flag
func ShowVersion() {
	fmt.Printf("%s", GetVersion())
}

// GetVersion xxx
func GetVersion() string {
	version := fmt.Sprintf("Version  :  %s\nTag      :  %s\nBuildTime:  %s\nGitHash  :  %s\nGoVersion:  %s\n",
		BcsVersion, BcsTag, BcsBuildTime, BcsGitHash, GoVersion)
	return version
}
