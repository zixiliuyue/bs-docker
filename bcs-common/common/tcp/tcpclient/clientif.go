

package tcpclient

// ClientIf 客户端操作接口
type ClientIf interface {
	IsAlive() bool
	Connect() error
	Send(extID int, data []byte) (int, error)
	DisConnect() error
}
