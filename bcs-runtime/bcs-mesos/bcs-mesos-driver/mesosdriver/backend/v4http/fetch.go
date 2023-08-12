

package v4http

import (
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
	bcstype "github.com/Tencent/bk-bcs/bcs-common/common/types"
	"github.com/bitly/go-simplejson"
	"sort"
)

// FetchApplication xxx
func (s *Scheduler) FetchApplication(ns, name string, kind bcstype.BcsDataType) (string, error) {
	blog.V(3).Infof("fetch application (%s.%s)", ns, name)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "?kind=" + string(kind)
	blog.V(3).Infof("post a request to url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

// FetchApplicationVersion xxx
func (s *Scheduler) FetchApplicationVersion(ns, name, versionID string) (string, error) {
	blog.V(3).Infof("fetch application (%s.%s) version (%s)", ns, name, versionID)

	if "" == versionID {
		verID, err := s.getLatestVersionId(ns, name)
		if err != nil {
			return err.Error(), err
		}

		versionID = verID
	}

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "/versions/" + versionID
	blog.V(3).Infof("post a request to url(%s)", url)

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}

func (s *Scheduler) getLatestVersionId(runAs, appId string) (string, error) {
	url := s.GetHost() + "/v1/apps/" + runAs + "/" + appId + "/versions"

	reply, err := s.client.GET(url, nil, nil)
	if err != nil {
		blog.Error("get request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	// {"result":true,"code":0,"message":"success","data":["1486469930525129787"]}
	rpyJson, err := simplejson.NewJson(reply)
	if err != nil {
		blog.Error("parse respone from listversions failed! err:%s", err.Error())
		err = bhttp.InternalError(common.BcsErrCommJsonDecode, common.BcsErrCommJsonDecodeStr+err.Error())
		return err.Error(), err
	}

	verIDs, _ := rpyJson.Get("data").Array()
	versionIDs := []string{}
	for _, id := range verIDs {
		switch id.(type) {
		case string:
			versionIDs = append(versionIDs, id.(string))
		}
	}

	if len(versionIDs) <= 0 {
		return "", bhttp.InternalError(common.BcsErrMesosDriverNoVersionId, common.BcsErrMesosDriverNoVersionIdStr)
	}

	sort.Strings(versionIDs)

	return versionIDs[len(versionIDs)-1], nil
}
