

package iam

import (
	"fmt"
)

//
const (
	// SysProject resource project
	SysProject TypeID = "project"
)

// ProjectResourcePath build IAMPath for project resource
type ProjectResourcePath struct{}

// BuildIAMPath build IAMPath, related resource project
func (rp ProjectResourcePath) BuildIAMPath() string {
	return ""
}

const (
	// ProjectCreate xxx
	ProjectCreate ActionID = "project_create"
	// ProjectView xxx
	ProjectView ActionID = "project_view"
	// ProjectEdit xxx
	ProjectEdit ActionID = "project_edit"
	// ProjectDelete xxx
	ProjectDelete ActionID = "project_delete"
)

const (
	// SysCluster resource cluster
	SysCluster TypeID = "cluster"
)

// ClusterResourcePath build IAMPath for cluster resource
type ClusterResourcePath struct {
	ProjectID     string
	ClusterCreate bool
}

// BuildIAMPath build IAMPath, related resource project when clusterCreate
func (rp ClusterResourcePath) BuildIAMPath() string {
	if rp.ClusterCreate {
		return ""
	}
	return fmt.Sprintf("/project,%s/", rp.ProjectID)
}

// actions
// cluster resource actions
const (
	// ClusterCreate xxx
	ClusterCreate ActionID = "cluster_create"
	// ClusterView xxx
	ClusterView ActionID = "cluster_view"
	// ClusterManage xxx
	ClusterManage ActionID = "cluster_manage"
	// ClusterDelete xxx
	ClusterDelete ActionID = "cluster_delete"
	// ClusterUse xxx
	ClusterUse ActionID = "cluster_use"
)
