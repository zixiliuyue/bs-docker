

package v4http

import (
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
)

// CreateConfigMap xxx
func (s *Scheduler) CreateConfigMap(body []byte) (string, error) {

	blog.Info("create configmap data(%s)", string(body))

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/configmap"
	blog.Info("post a request to url(%s), request:%s", url, string(body))

	reply, err := s.client.POST(url, nil, body)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// UpdateConfigMap xxx
func (s *Scheduler) UpdateConfigMap(body []byte) (string, error) {

	blog.Info("update configmap data(%s)", string(body))

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/configmap"
	blog.Info("put a request to url(%s), request:%s", url, string(body))

	reply, err := s.client.PUT(url, nil, body)
	if err != nil {
		blog.Error("put request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// DeleteConfigMap xxx
func (s *Scheduler) DeleteConfigMap(ns string, name string) (string, error) {
	blog.Info("delete configmap(%s, %s)", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/configmap/" + ns + "/" + name
	blog.Info("delete url(%s)", url)

	reply, err := s.client.DELETE(url, nil, nil)
	if err != nil {
		blog.Error("delete url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
