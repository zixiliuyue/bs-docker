

package v4http

import (
	"fmt"
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
)

// RescheduleTaskgroup xxx
func (s *Scheduler) RescheduleTaskgroup(taskgroupId, hostRetainTime string) (string, error) {
	blog.Info("rescheduler taskgroup(%s)", taskgroupId)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := fmt.Sprintf("%s/v1/taskgroup/%s/rescheduler?hostRetainTime=%s", s.GetHost(), taskgroupId, hostRetainTime)

	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.PUT(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// RestartTaskGroup xxx
func (s *Scheduler) RestartTaskGroup(taskGroupID string) (string, error) {
	blog.Info("restart taskGroup(%s)", taskGroupID)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := fmt.Sprintf("%s/v1/taskgroup/%s/restart", s.GetHost(), taskGroupID)

	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.POST(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// ReloadTaskGroup xxx
func (s *Scheduler) ReloadTaskGroup(taskGroupID string) (string, error) {
	blog.Info("reload taskGroup(%s)", taskGroupID)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := fmt.Sprintf("%s/v1/taskgroup/%s/reload", s.GetHost(), taskGroupID)

	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.POST(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
