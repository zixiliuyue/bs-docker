

package zk

import (
	"encoding/json"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

// SaveFrameworkID xxx
func (store *managerStore) SaveFrameworkID(frameworkId string) error {

	framework := &types.Framework{ID: frameworkId}
	data, err := json.Marshal(framework)
	if err != nil {
		blog.Error("fail to encode object framework by json. err:%s", err.Error())
		return err
	}

	path := "/" + bcsRootNode + "/" + frameWorkNode

	return store.Db.Insert(path, string(data))
}

// FetchFrameworkID xxx
func (store *managerStore) FetchFrameworkID() (string, error) {

	path := "/" + bcsRootNode + "/" + frameWorkNode

	data, err := store.Db.Fetch(path)
	if err != nil {
		blog.Error("fail to get framework id, err(%s)", err.Error())
		return "", err
	}

	framework := &types.Framework{}
	if err := json.Unmarshal(data, framework); err != nil {
		blog.Error("fail to unmarshal framework(%s), err(%s)", string(data), err.Error())
		return "", err
	}

	return framework.ID, nil
}

// HasFrameworkID xxx
func (store *managerStore) HasFrameworkID() (bool, error) {
	_, err := store.FetchFrameworkID()
	if err != nil {
		return false, err
	}

	return true, nil
}
