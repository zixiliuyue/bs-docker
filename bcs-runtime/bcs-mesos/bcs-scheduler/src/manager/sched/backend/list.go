

package backend

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// ListApplications xxx
func (b *backend) ListApplications(runAs string) ([]*types.Application, error) {
	return b.store.ListApplications(runAs)
}

// ListApplicationTaskGroups xxx
func (b *backend) ListApplicationTaskGroups(runAs, appId string) ([]*types.TaskGroup, error) {
	b.store.LockApplication(runAs + "." + appId)
	defer b.store.UnLockApplication(runAs + "." + appId)

	return b.store.ListTaskGroups(runAs, appId)
}

// ListApplicationVersions is used to list all versions for application from db specified by application id.
func (b *backend) ListApplicationVersions(runAs, appId string) ([]string, error) {
	return b.store.ListVersions(runAs, appId)
}
