

package v4http

import (
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
)

// CreateAdmissionWebhook xxx
func (s *Scheduler) CreateAdmissionWebhook(body []byte) (string, error) {

	blog.Info("create admissionwebhook data(%s)", string(body))

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/admissionwebhook"
	blog.Info("post a request to url(%s), request:%s", url, string(body))

	reply, err := s.client.POST(url, nil, body)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// UpdateAdmissionWebhook xxx
func (s *Scheduler) UpdateAdmissionWebhook(body []byte) (string, error) {

	blog.Info("update admissionwebhook data(%s)", string(body))

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/admissionwebhook"
	blog.Info("put a request to url(%s), request:%s", url, string(body))

	reply, err := s.client.PUT(url, nil, body)
	if err != nil {
		blog.Error("put request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// DeleteAdmissionWebhook xxx
func (s *Scheduler) DeleteAdmissionWebhook(ns string, name string) (string, error) {
	blog.Info("delete admissionwebhook(%s, %s)", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/admissionwebhook/" + ns + "/" + name
	blog.Info("delete url(%s)", url)

	reply, err := s.client.DELETE(url, nil, nil)
	if err != nil {
		blog.Error("delete url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// FetchAdmissionWebhook xxx
func (s *Scheduler) FetchAdmissionWebhook(ns string, name string) (string, error) {
	blog.Info("fetch admissionwebhook(%s, %s)", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/admissionwebhook/" + ns + "/" + name
	blog.Info("fetch url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("fetch url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// FetchAllAdmissionWebhooks xxx
func (s *Scheduler) FetchAllAdmissionWebhooks() (string, error) {
	blog.Info("fetch all admissionwebhooks")

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/admissionwebhooks"
	blog.Info("fetch url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("fetch url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
