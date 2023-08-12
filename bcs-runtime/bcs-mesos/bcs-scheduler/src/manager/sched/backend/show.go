

package backend

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// FetchApplication xxx
func (b *backend) FetchApplication(runAs, appId string) (*types.Application, error) {
	return b.store.FetchApplication(runAs, appId)
}

// FetchApplicationVersion is used to fetch specified version from db by version id and application id.
func (b *backend) FetchApplicationVersion(runAs, appId, versionId string) (*types.Version, error) {
	return b.store.FetchVersion(runAs, appId, versionId)
}
