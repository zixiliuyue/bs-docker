

package gin

// Version is the current release version of the gin instrumentation.
func Version() string {
	return "0.37.0"
	// This string is updated by the pre_release.sh script during release
}

// SemVersion is the semantic version to be supplied to tracer/meter creation.
func SemVersion() string {
	return "semver:" + Version()
}
