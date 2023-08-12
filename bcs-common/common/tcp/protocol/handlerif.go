

package protocol

// HandlerIf 用于接收服务端收到的数据
type HandlerIf interface {

	// Write 转发接收到的数据和协议头数据给逻辑层
	Write(head *MsgHead, data []byte) (int, error)
}
