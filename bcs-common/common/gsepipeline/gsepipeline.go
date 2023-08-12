

// Package gsepipeline xxx
package gsepipeline

import (
	"encoding/json"
	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	gseclient "github.com/Tencent/bk-bcs/bcs-common/common/gsepipeline/client"
)

type gseStorage struct {
	client gseclient.AsyncProducer
}

// NewGseStorage create a new gseclient
func NewGseStorage(endpoint string) (Storage, error) {

	blog.Info("endpoint:%v", endpoint)

	client, err := gseclient.New(endpoint)
	if nil != err {
		return nil, err
	}

	if nil == err {
		gseStorageClient := &gseStorage{client: client}

		return gseStorageClient, nil
	}

	return nil, err

}

// AddStats pushing Stats info
func (gse *gseStorage) AddStats(dmsg LogMsg) error {

	b, err := json.Marshal(dmsg)
	if err != nil {
		return err
	}

	gse.client.Input(&gseclient.ProducerMessage{DataID: uint32(dmsg.DataID), Value: b})

	return nil
}

// Close close pipeline
func (gse *gseStorage) Close() error {

	gse.client.Close()
	return nil
}
