

package driver

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	protoExec "github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-process-executor/process-executor/protobuf/executor"

	"github.com/golang/protobuf/proto"
)

// HttpConnection xxx
type HttpConnection struct {
	endpoint string // remote http endpoint info
	uri      string // remote http endpoint uri
	streamID string // http header Mesos-Stream-Id

	client *http.Client
}

// NewHttpConnection xxx
func NewHttpConnection(endpoint, uri string) *HttpConnection {
	httpTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		ResponseHeaderTimeout: 5 * time.Second,
	}

	return &HttpConnection{
		endpoint: endpoint,
		uri:      uri,
		client: &http.Client{
			Transport: httpTransport,
		},
	}
}

// Send xxx
func (conn *HttpConnection) Send(call *protoExec.Call, keepAlive bool) (*http.Response, error) {
	// create targetURL
	targetURL := fmt.Sprintf("%s%s", conn.endpoint, conn.uri)
	// proto serialization
	payLoad, err := proto.Marshal(call)
	if err != nil {
		blog.Errorf("proto.Marshal call error %s", err.Error())
		return nil, err
	}

	// create http request
	request, err := http.NewRequest("POST", targetURL, bytes.NewReader(payLoad))
	if err != nil {
		blog.Errorf("NewRequest url %s error %s", targetURL, err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-protobuf")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "bcs-container-executor/1.0")
	request.Header.Set("Connection", "Keep-Alive")

	response, err := conn.client.Do(request)
	if err != nil {
		blog.Errorf("conn request uri %s error %s", targetURL, err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusAccepted && response.StatusCode != http.StatusOK {
		reply, _ := ioutil.ReadAll(response.Body)
		err = fmt.Errorf("request url %s response statuscode %d body %s", targetURL, response.StatusCode, string(reply))
		blog.Errorf(err.Error())
		return nil, err
	}

	if keepAlive {
		return response, nil
	}

	response.Body.Close()
	return nil, nil
}
