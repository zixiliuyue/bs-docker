

package zk

import (
	"encoding/json"
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
)

func getTaskRootPath() string {
	return "/" + bcsRootNode + "/" + applicationNode + "/"
}

func createTaskPath(taskID string) (string, error) {
	runAs, appID := types.GetRunAsAndAppIDbyTaskID(taskID)
	if "" == appID {
		err := fmt.Errorf("fail to get appid from taskID(%s)", taskID)
		return "", err
	}

	if "" == runAs {
		err := fmt.Errorf("fail to get runAs from taskID(%s)", taskID)
		return "", err
	}

	taskGroupID := types.GetTaskGroupID(taskID)
	if "" == taskGroupID {
		err := fmt.Errorf("fail to get taskgroup id from taskID(%s)", taskID)
		return "", err
	}

	path := getTaskRootPath() + runAs + "/" + appID + "/" + taskGroupID + "/" + taskID

	return path, nil
}

// SaveTask xxx
func (store *managerStore) SaveTask(task *types.Task) error {

	blog.V(3).Infof("save task(id:%s)", task.ID)

	data, err := json.Marshal(task)
	if err != nil {
		blog.Error("fail to encode object task(ID:%s) by json. err:%s", task.ID, err.Error())
		return err
	}

	path, err := createTaskPath(task.ID)
	if err != nil {
		blog.Error("fail to create task path. err(%s)", err.Error())
		return err
	}

	if err := store.Db.Insert(path, string(data)); err != nil {
		return err
	}

	saveCacheTask(task)

	return nil
}

// ListTasks xxx
func (store *managerStore) ListTasks(runAs, appID string) ([]*types.Task, error) {

	path := getTaskRootPath() + runAs + "/" + appID
	taskGroupIds, err := store.Db.List(path)
	if err != nil {
		blog.Error("fail to get taskgroup ids by app path(%s), err:%s", path, err.Error())
		return nil, err
	}

	var tasksList []*types.Task

	for _, taskGroupID := range taskGroupIds {
		taskGroupPath := path + "/" + taskGroupID
		taskIds, err := store.Db.List(taskGroupPath)
		if err != nil {
			blog.Error("fail to get task ids by taskgroupPath(%s), err:%s", taskGroupPath, err.Error())
			return nil, err
		}

		for _, taskId := range taskIds {
			task, err := store.FetchTask(taskId)
			if err != nil {
				blog.Error("fail to get task by taskId(%s), err:%s", taskId, err.Error())
				return nil, err
			}

			tasksList = append(tasksList, task)
		}
	}

	return tasksList, nil
}

// DeleteTask xxx
func (store *managerStore) DeleteTask(taskId string) error {
	path, err := createTaskPath(taskId)
	if err != nil {
		blog.Error("fail to create task path. err(%s)", err.Error())
		return err
	}

	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete task(id:%s), err:%s", taskId, err.Error())
		return err
	}

	deleteCacheTask(taskId)

	return nil
}

// FetchTask xxx
func (store *managerStore) FetchTask(taskId string) (*types.Task, error) {

	cacheTask, cacheErr := fetchCacheTask(taskId)
	if cacheErr == nil && cacheTask != nil {
		return cacheTask, nil
	}

	path, err := createTaskPath(taskId)
	if err != nil {
		blog.Error("fail to create task path. err(%s)", err.Error())
		return nil, err
	}

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	var task types.Task
	if err := json.Unmarshal(data, &task); err != nil {
		blog.Error("fail to unmarshal task. task(%s), err:%s", string(data), err.Error())
		return nil, err
	}

	return &task, nil
}

// FetchDBTask xxx
func (store *managerStore) FetchDBTask(taskId string) (*types.Task, error) {

	path, err := createTaskPath(taskId)
	if err != nil {
		blog.Error("fail to create task path. err(%s)", err.Error())
		return nil, err
	}

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	var task types.Task
	if err := json.Unmarshal(data, &task); err != nil {
		blog.Error("fail to unmarshal task. task(%s), err:%s", string(data), err.Error())
		return nil, err
	}

	return &task, nil
}
