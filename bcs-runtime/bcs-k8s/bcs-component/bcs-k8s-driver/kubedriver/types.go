

package kubedriver

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// KubeVersion xxx
type KubeVersion struct {
	Major      string
	Minor      string
	GitVersion string
}

// String 用于打印
func (v KubeVersion) String() string {
	return fmt.Sprintf("git:%s/major:%s/minor:%s", v.GitVersion, v.Major, v.Minor)
}

// IsValid xxx
func (v KubeVersion) IsValid() bool {
	// Minikube may not return minor and major
	return v.GitVersion != ""
}

// KubeAPIPrefer xxx
type KubeAPIPrefer struct {
	Groups []struct {
		Name     string `json:"name"`
		Versions []struct {
			GroupVersion string `json:"groupVersion"`
			Version      string `json:"version"`
		} `json:"versions"`
		PreferredVersion struct {
			GroupVersion string `json:"groupVersion"`
			Version      string `json:"version"`
		} `json:"preferredVersion"`
	} `json:"groups"`
}

// Map xxx
func (p KubeAPIPrefer) Map() map[string]string {
	kubePrefer := map[string]string{}
	for _, group := range p.Groups {
		kubePrefer[group.Name] = group.PreferredVersion.Version
	}
	return kubePrefer
}
