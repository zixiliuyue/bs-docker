

package zk

import (
	"encoding/json"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

func getEndpointRootPath() string {
	return "/" + bcsRootNode + "/" + endpointNode
}

// SaveEndpoint xxx
func (store *managerStore) SaveEndpoint(endpoint *commtypes.BcsEndpoint) error {

	data, err := json.Marshal(endpoint)
	if err != nil {
		return err
	}

	path := getEndpointRootPath() + "/" + endpoint.ObjectMeta.NameSpace + "/" + endpoint.ObjectMeta.Name

	return store.Db.Insert(path, string(data))
}

// FetchEndpoint xxx
func (store *managerStore) FetchEndpoint(ns, name string) (*commtypes.BcsEndpoint, error) {

	path := getEndpointRootPath() + "/" + ns + "/" + name

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	endpoint := &commtypes.BcsEndpoint{}
	if err := json.Unmarshal(data, endpoint); err != nil {
		blog.Error("fail to unmarshal endpoint(%s). err:%s", string(data), err.Error())
		return nil, err
	}

	return endpoint, nil
}

// DeleteEndpoint xxx
func (store *managerStore) DeleteEndpoint(ns, name string) error {

	path := getEndpointRootPath() + "/" + ns + "/" + name
	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete endpoint(%s) err:%s", path, err.Error())
		return err
	}

	return nil
}
