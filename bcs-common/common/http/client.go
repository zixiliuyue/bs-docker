

package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Tencent/bk-bcs/bcs-common/common"
)

// HeaderSet http header set
type HeaderSet struct {
	Key   string
	Value string
}

// Request http request encapsulation
func Request(url, method string, header http.Header, body io.Reader) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		reply := InternalError(common.BcsErrCommHttpNewRequest, common.BcsErrCommHttpNewRequestStr)
		return reply.Error(), fmt.Errorf("fail to new a http request. err:%s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req.Close = true

	// header
	if header != nil {
		req.Header = header
	}

	client := &http.Client{}

	rsp, err := client.Do(req)
	if err != nil {
		reply := InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr)
		return reply.Error(), fmt.Errorf("fail to do http request. err:%s", err.Error())
	}

	defer rsp.Body.Close()

	replyData, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		reply := InternalError(common.BcsErrCommHttpReadRsp, common.BcsErrCommHttpReadRspStr)
		return reply.Error(), fmt.Errorf("read response failed. err:%s", err.Error())
	}

	return string(replyData), nil
}

// RequestV2 v2 version for http request encapsulation
func RequestV2(url, method string, headSet []*HeaderSet, body io.Reader) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		reply := InternalError(common.BcsErrCommHttpNewRequest, common.BcsErrCommHttpNewRequestStr)
		return reply.Error(), fmt.Errorf("fail to new a http request. err:%s", err.Error())
	}

	for _, header := range headSet {
		req.Header.Set(header.Key, header.Value)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req.Close = true

	client := &http.Client{}

	rsp, err := client.Do(req)
	if err != nil {
		reply := InternalError(common.BcsErrCommHttpDo, common.BcsErrCommHttpDoStr)
		return reply.Error(), fmt.Errorf("fail to do http request. err:%s", err.Error())
	}

	defer rsp.Body.Close()

	replyData, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		reply := InternalError(common.BcsErrCommHttpReadRsp, common.BcsErrCommHttpReadRspStr)
		return reply.Error(), fmt.Errorf("read response failed. err:%s", err.Error())
	}

	return string(replyData), nil
}
