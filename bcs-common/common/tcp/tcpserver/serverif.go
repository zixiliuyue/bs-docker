

package tcpserver

// ServerIf 服务接口
type ServerIf interface {
	Start() error
	Stop() error
}
