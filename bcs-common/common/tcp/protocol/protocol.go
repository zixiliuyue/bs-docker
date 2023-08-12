

// Package protocol xxx
package protocol

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// ProtocolMagicNum 防篡改魔法数字
const ProtocolMagicNum uint64 = 0x75062D90E03C798A

// MsgHead 通讯协议头协议头定义
type MsgHead struct {
	Magic     uint64
	Type      uint32
	Length    uint32
	ExtID     int64
	Timestamp int64
}

// HeadLength 获取协议头长度
func HeadLength() int {
	return int(unsafe.Sizeof(MsgHead{})) // nolint
}

// ConvertToMsgHead 将字节数组 转换为协议头
func ConvertToMsgHead(buffer []byte) (*MsgHead, error) {

	tmpBuff := bytes.NewBuffer(buffer)

	msgHead := &MsgHead{}
	err := binary.Read(tmpBuff, binary.BigEndian, msgHead)

	if nil != err {
		return nil, err
	}

	return msgHead, nil
}

// Package 打包数据
func Package(msgHead *MsgHead, data []byte) ([]byte, error) {

	buffer := new(bytes.Buffer)

	// package the head
	if err := binary.Write(buffer, binary.BigEndian, msgHead); nil != err {
		blog.Error("can not package the head, error %s", err.Error())
		return nil, err
	}

	// package the data after the head
	head := buffer.Bytes()
	if nil != data && len(data) > 0 {
		head = append(head, data...)
	}
	return head, nil
}
