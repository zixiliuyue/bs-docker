

package v4http

import (
	"github.com/Tencent/bk-bcs/bcs-common/common"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	bhttp "github.com/Tencent/bk-bcs/bcs-common/common/http"
	bcstype "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

// ScaleApplication scale application transparent forwarding
func (s *Scheduler) ScaleApplication(ns, name, instances string, kind bcstype.BcsDataType) (string, error) {
	blog.Info("scale application (%s.%s) instances (%s)", ns, name, instances)

	if s.GetHost() == "" {
		blog.Error("no scheduler is connected by driver")
		err := bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+"scheduler not exist")
		return err.Error(), err
	}

	url := s.GetHost() + "/v1/apps/" + ns + "/" + name + "/scale?instances=" + instances + "&kind=" + string(kind)
	blog.Info("post a request to url(%s)", url)

	reply, err := s.client.POST(url, nil, nil)
	if err != nil {
		blog.Error("post request to url(%s) failed! err(%s)", url, err.Error())
		err = bhttp.InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr+err.Error())
		return err.Error(), err
	}

	return string(reply), nil
}
