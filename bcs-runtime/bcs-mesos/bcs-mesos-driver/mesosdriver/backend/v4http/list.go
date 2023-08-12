

package v4http

import (
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
	bcstype "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// ListApplications list all apps
func (s *Scheduler) ListApplications(ns string, kind bcstype.BcsDataType) (string, error) {
	blog.V(3).Infof("list namespace (%s) applications", ns)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/" + ns + "/apps" + "?kind=" + string(kind)
	blog.V(3).Infof("post a request to url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// ListApplicationTasks xxx
func (s *Scheduler) ListApplicationTasks(ns, name string) (string, error) {
	blog.V(3).Infof("list application (%s.%s) tasks", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "/tasks"
	blog.V(3).Infof("post a request to url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// ListApplicationTaskGroups list all taskgroup under application
func (s *Scheduler) ListApplicationTaskGroups(ns, name string) (string, error) {
	blog.V(3).Infof("list application (%s.%s) taskgroups", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "/taskgroups"
	blog.V(3).Infof("post a request to url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// ListApplicationVersions list application inner definition
func (s *Scheduler) ListApplicationVersions(ns, name string) (string, error) {
	blog.V(3).Infof("list application (%s.%s) versions", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "/versions"
	blog.V(3).Infof("post a request to url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
