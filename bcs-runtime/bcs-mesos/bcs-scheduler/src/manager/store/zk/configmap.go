

package zk

import (
	"encoding/json"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

func getConfigMapRootPath() string {
	return "/" + bcsRootNode + "/" + configMapNode
}

// SaveConfigMap xxx
func (store *managerStore) SaveConfigMap(configmap *commtypes.BcsConfigMap) error {

	data, err := json.Marshal(configmap)
	if err != nil {
		return err
	}

	path := getConfigMapRootPath() + "/" + configmap.ObjectMeta.NameSpace + "/" + configmap.ObjectMeta.Name

	return store.Db.Insert(path, string(data))
}

// FetchConfigMap xxx
func (store *managerStore) FetchConfigMap(ns, name string) (*commtypes.BcsConfigMap, error) {

	path := getConfigMapRootPath() + "/" + ns + "/" + name

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	configmap := &commtypes.BcsConfigMap{}
	if err := json.Unmarshal(data, configmap); err != nil {
		blog.Error("fail to unmarshal configmap(%s). err:%s", string(data), err.Error())
		return nil, err
	}

	return configmap, nil
}

// DeleteConfigMap xxx
func (store *managerStore) DeleteConfigMap(ns, name string) error {

	path := getConfigMapRootPath() + "/" + ns + "/" + name
	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete configmap(%s) err:%s", path, err.Error())
		return err
	}

	return nil
}

// ListConfigmaps xxx
func (store *managerStore) ListConfigmaps(runAs string) ([]*commtypes.BcsConfigMap, error) {

	path := getConfigMapRootPath() + "/" + runAs // defaultRunAs

	IDs, err := store.Db.List(path)
	if err != nil {
		blog.Error("fail to list configmap ids, err:%s", err.Error())
		return nil, err
	}

	if nil == IDs {
		blog.V(3).Infof("no configmap in (%s)", runAs)
		return nil, nil
	}

	var objs []*commtypes.BcsConfigMap

	for _, ID := range IDs {
		obj, err := store.FetchConfigMap(runAs, ID)
		if err != nil {
			blog.Error("fail to fetch configmap by ID(%s)", ID)
			continue
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

// ListAllConfigmaps xxx
func (store *managerStore) ListAllConfigmaps() ([]*commtypes.BcsConfigMap, error) {
	nss, err := store.ListObjectNamespaces(configMapNode)
	if err != nil {
		return nil, err
	}

	var objs []*commtypes.BcsConfigMap
	for _, ns := range nss {
		obj, err := store.ListConfigmaps(ns)
		if err != nil {
			blog.Error("fail to fetch configmap by ns(%s)", ns)
			continue
		}

		objs = append(objs, obj...)
	}

	return objs, nil
}
