

package v4http

import (
	"encoding/json"
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
	bcstype "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// DeleteApplication xxx
func (s *Scheduler) DeleteApplication(ns, name, enforce string, kind bcstype.BcsDataType) (string, error) {
	blog.Info("delete application (%s.%s) enforce(%s)", ns, name, enforce)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "?enforce=" + enforce + "&kind=" + string(kind)
	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.DELETE(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// DeleteApplicationTaskGroups xxx
func (s *Scheduler) DeleteApplicationTaskGroups(body []byte) (string, error) {
	blog.Info("delete application taskgroups, param(%s)", string(body))

	// encodings the parameter of deleting taskgroups operation
	var param DeleteTaskGroupsOpeParam
	if err := json.Unmarshal(body, &param); err != nil {
		blog.Error("parse deleting application taskgroups operation parameters failed, err(%s)", err.Error())
		err = bhttp.InternalError(common.BcsErrCommJsonDecode, common.BcsErrCommJsonDecodeStr+err.Error())
		return err.Error(), err
	}

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + param.RunAs + "/" + param.Name + "/taskgroups"
	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.DELETE(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// DeleteApplicationTaskGroup xxx
func (s *Scheduler) DeleteApplicationTaskGroup(body []byte) (string, error) {
	blog.Info("delete application taskgroup, param(%s)", string(body))

	// encoding the parameter of deleting taskgroup operation
	var param DeleteTaskGroupOpeParam
	if err := json.Unmarshal(body, &param); err != nil {
		blog.Error("parse deleting application taskgroup operation parameters failed, err(%s)", err.Error())
		err = bhttp.InternalError(common.BcsErrCommJsonDecode, common.BcsErrCommJsonDecodeStr+err.Error())
		return err.Error(), err
	}

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + param.RunAs + "/" + param.Name + "/taskgroups/" + param.TaskGroupId
	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.DELETE(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
