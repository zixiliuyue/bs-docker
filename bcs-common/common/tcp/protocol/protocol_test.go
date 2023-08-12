

package protocol

import (
	"bytes"
	"encoding/binary"
	"testing"
	"time"
)

func TestHeadLength(t *testing.T) {
	t.Logf("Protocol Head Length:%d", HeadLength())
}

func TestConvertToMsgHead(t *testing.T) {

	buffer := new(bytes.Buffer)

	msgHead := &MsgHead{}
	msgHead.Magic = ProtocolMagicNum
	msgHead.ExtID = 0x01
	msgHead.Length = 1
	msgHead.Timestamp = time.Now().UnixNano()
	msgHead.Type = 1

	// package the head
	if err := binary.Write(buffer, binary.BigEndian, msgHead); nil != err {
		t.Errorf("can not package the head, error %s", err.Error())
	}

	// package the data after the head
	head := buffer.Bytes()
	t.Logf("MsgHead bytes:%+#v", head)

	tmpHead, tmpErr := ConvertToMsgHead(head)
	if nil != tmpErr {
		t.Errorf("convert to msghead failed, error %s", tmpErr.Error())
	} else {
		t.Logf("MsgHead: %+#v", tmpHead)
	}
}

func TestPackage(t *testing.T) {

	msgHead := &MsgHead{}
	msgHead.ExtID = 0x01
	msgHead.Length = 1
	msgHead.Timestamp = time.Now().UnixNano()
	msgHead.Type = 1

	resp, respErr := Package(msgHead, []byte("hello world"))
	if nil != respErr {
		t.Errorf("convert to msghead failed, error %s", respErr.Error())
	} else {
		t.Logf("data: %+v", resp)
	}

}
