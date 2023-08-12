

package backend

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
	"reflect"
	"sort"
)

// SaveApplication register application in db.
func (b *backend) SaveApplication(application *types.Application) error {
	return b.store.SaveApplication(application)
}

// SaveVersion register application version in db.
func (b *backend) SaveVersion(runAs, appId string, version *types.Version) error {

	blog.Info("save version(%s.%s)", runAs, appId)
	versions, err := b.store.ListVersions(runAs, appId)
	if err != nil {
		blog.Error("list versions(%s.%s) err:%s", runAs, appId, err.Error())
		return err
	}

	versionName := version.Name
	if len(versions) != 0 {
		sort.Strings(versions)
		newestVersion, err := b.store.FetchVersion(runAs, appId, versions[len(versions)-1])
		if err != nil {
			blog.Error("fetch version(%s.%s), versionNo(%s) err:%s", runAs, appId, versions[len(versions)-1], err.Error())
			return err
		}
		version.Name = newestVersion.Name
		if reflect.DeepEqual(version, newestVersion) {
			version.Name = versionName
			return nil
		}
	}

	version.Name = versionName
	return b.store.SaveVersion(version)
}

// GetVersion xxx
func (b *backend) GetVersion(runAs, appId string) (*types.Version, error) {
	return b.store.GetVersion(runAs, appId)
}
