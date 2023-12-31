

// Package version xxx
package version

import "fmt"

var (
	version   = "unknown"
	gitCommit = "$Format:%H$"          // sha1 from git, output of $(git rev-parse HEAD)
	buildDate = "1970-01-01T00:00:00Z" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// SimpleVersion xxx
type SimpleVersion struct {
	version   string
	gitCommit string
	buildDate string
}

// ForDisplay xxx
func (v *SimpleVersion) ForDisplay() string {
	return fmt.Sprintf(
		`Version: "%s", GitCommit: "%s", BuildDate: "%s"`,
		v.version,
		v.gitCommit,
		v.buildDate)
}

// Get xxx
func Get() *SimpleVersion {
	return &SimpleVersion{
		version:   version,
		gitCommit: gitCommit,
		buildDate: buildDate,
	}
}
