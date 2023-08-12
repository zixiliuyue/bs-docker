

package v4http

import (
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
	// bcstype "github.com/Tencent/bk-bcs/bcs-common/common/types"
	// "github.com/Tencent/bk-bcs/bcs-common/pkg/scheduler/schetypes"
	// "encoding/json"
)

// CreateService create service request forwarding
func (s *Scheduler) CreateService(body []byte) (string, error) {

	blog.Info("create service data(%s)", string(body))

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/service"
	blog.Info("post a request to url(%s), request:%s", url, string(body))

	reply, err := s.client.POST(url, nil, body)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// UpdateService update service request forwarding
func (s *Scheduler) UpdateService(body []byte) (string, error) {

	blog.Info("update service data(%s)", string(body))

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/service"
	blog.Info("put a request to url(%s), request:%s", url, string(body))

	reply, err := s.client.PUT(url, nil, body)
	if err != nil {
		blog.Error("put request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// DeleteService delete service request forwarding
func (s *Scheduler) DeleteService(ns string, name string) (string, error) {
	blog.Info("delete service(%s, %s)", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/service/" + ns + "/" + name
	blog.Info("delete url(%s)", url)

	reply, err := s.client.DELETE(url, nil, nil)
	if err != nil {
		blog.Error("delete url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
