

package zk

import (
	"encoding/json"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

func getSecretRootPath() string {
	return "/" + bcsRootNode + "/" + secretNode
}

// SaveSecret xxx
func (store *managerStore) SaveSecret(secret *commtypes.BcsSecret) error {

	data, err := json.Marshal(secret)
	if err != nil {
		return err
	}

	path := getSecretRootPath() + "/" + secret.ObjectMeta.NameSpace + "/" + secret.ObjectMeta.Name

	return store.Db.Insert(path, string(data))
}

// FetchSecret xxx
func (store *managerStore) FetchSecret(ns, name string) (*commtypes.BcsSecret, error) {

	path := getSecretRootPath() + "/" + ns + "/" + name

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	secret := &commtypes.BcsSecret{}
	if err := json.Unmarshal(data, secret); err != nil {
		blog.Error("fail to unmarshal secret(%s). err:%s", string(data), err.Error())
		return nil, err
	}

	return secret, nil
}

// DeleteSecret xxx
func (store *managerStore) DeleteSecret(ns, name string) error {

	path := getSecretRootPath() + "/" + ns + "/" + name
	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete secret(%s) err:%s", path, err.Error())
		return err
	}

	return nil
}

// ListSecrets xxx
func (store *managerStore) ListSecrets(runAs string) ([]*commtypes.BcsSecret, error) {

	path := getSecretRootPath() + "/" + runAs // defaultRunAs

	IDs, err := store.Db.List(path)
	if err != nil {
		blog.Error("fail to list secret ids, err:%s", err.Error())
		return nil, err
	}

	if nil == IDs {
		blog.Error("no secret in (%s)", runAs)
		return nil, nil
	}

	var objs []*commtypes.BcsSecret

	for _, ID := range IDs {
		obj, err := store.FetchSecret(runAs, ID)
		if err != nil {
			blog.Error("fail to fetch secret by ID(%s)", ID)
			continue
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

// ListAllSecrets xxx
func (store *managerStore) ListAllSecrets() ([]*commtypes.BcsSecret, error) {
	nss, err := store.ListObjectNamespaces(secretNode)
	if err != nil {
		return nil, err
	}

	var objs []*commtypes.BcsSecret
	for _, ns := range nss {
		obj, err := store.ListSecrets(ns)
		if err != nil {
			blog.Error("fail to fetch secret by ns(%s)", ns)
			continue
		}

		objs = append(objs, obj...)
	}

	return objs, nil
}
