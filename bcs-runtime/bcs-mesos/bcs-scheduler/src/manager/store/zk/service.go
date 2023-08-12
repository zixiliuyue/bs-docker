

package zk

import (
	"encoding/json"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

func getServiceRootPath() string {
	return "/" + bcsRootNode + "/" + serviceNode
}

// SaveService xxx
func (store *managerStore) SaveService(service *commtypes.BcsService) error {

	data, err := json.Marshal(service)
	if err != nil {
		return err
	}

	path := getServiceRootPath() + "/" + service.ObjectMeta.NameSpace + "/" + service.ObjectMeta.Name

	return store.Db.Insert(path, string(data))
}

// FetchService xxx
func (store *managerStore) FetchService(ns, name string) (*commtypes.BcsService, error) {

	path := getServiceRootPath() + "/" + ns + "/" + name

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	service := &commtypes.BcsService{}
	if err := json.Unmarshal(data, service); err != nil {
		blog.Error("fail to unmarshal service(%s). err:%s", string(data), err.Error())
		return nil, err
	}

	return service, nil
}

// DeleteService xxx
func (store *managerStore) DeleteService(ns, name string) error {

	path := getServiceRootPath() + "/" + ns + "/" + name
	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete service(%s) err:%s", path, err.Error())
		return err
	}

	return nil
}

// ListServices xxx
func (store *managerStore) ListServices(runAs string) ([]*commtypes.BcsService, error) {

	path := getServiceRootPath() + "/" + runAs // defaultRunAs
	IDs, err := store.Db.List(path)
	if err != nil {
		blog.Error("fail to list service ids, err:%s", err.Error())
		return nil, err
	}

	if nil == IDs {
		blog.V(3).Infof("no service in (%s)", runAs)
		return nil, nil
	}

	var objs []*commtypes.BcsService

	for _, ID := range IDs {
		obj, err := store.FetchService(runAs, ID)
		if err != nil {
			blog.Error("fail to fetch service by ID(%s)", ID)
			continue
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

// ListAllServices xxx
func (store *managerStore) ListAllServices() ([]*commtypes.BcsService, error) {
	nss, err := store.ListObjectNamespaces(serviceNode)
	if err != nil {
		return nil, err
	}

	var objs []*commtypes.BcsService
	for _, ns := range nss {
		blog.Infof("list all services ns %s", ns)
		obj, err := store.ListServices(ns)
		if err != nil {
			blog.Error("fail to fetch service by ns(%s)", ns)
			continue
		}

		objs = append(objs, obj...)
	}

	return objs, nil
}
