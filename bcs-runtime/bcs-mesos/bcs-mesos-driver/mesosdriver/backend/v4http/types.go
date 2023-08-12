

package v4http

// OperateItem xxx
type OperateItem struct {
	Name     string `json:"name"`
	RunAs    string `json:"namespace"`
	SetId    string `json:"setid"`
	ModuleId string `json:"moduleid"`
}

// ScaleOpeParam xxx
type ScaleOpeParam struct {
	OperateItem
	Instance uint64 `json:"instance"`
}

// DeleteAppOpeParam xxx
type DeleteAppOpeParam struct {
	OperateItem
}

// DeleteTaskGroupsOpeParam xxx
type DeleteTaskGroupsOpeParam struct {
	OperateItem
}

// DeleteTaskGroupOpeParam xxx
type DeleteTaskGroupOpeParam struct {
	OperateItem
	TaskGroupId string `json:"taskgroupid"`
}

// RollbackOpeParam xxx
type RollbackOpeParam struct {
	OperateItem
}

// FetchVersionOpeParam xxx
type FetchVersionOpeParam struct {
	OperateItem
	VersionId string `json:"versionid"`
}

// SendMsgOpeParam xxx
type SendMsgOpeParam struct {
	OperateItem
	MsgType     string      `json:"msgtype"`
	MsgData     interface{} `json:"msgdata"`
	TaskGroupId string      `json:"taskgroupid"`
}
