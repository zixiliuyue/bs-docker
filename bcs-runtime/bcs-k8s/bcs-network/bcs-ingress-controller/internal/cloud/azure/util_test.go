

package azure

import (
	"fmt"
	"testing"

	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
)

// 健康检查状态码（仅适用于HTTP/HTTPS转发规则、TCP监听器的HTTP健康检查方式）。可选值：1~31，默认 31。1 表示探测后返回值 1xx 代表健康，
// 2 表示返回 2xx 代表健康，4 表示返回 3xx 代表健康，8 表示返回 4xx 代表健康，16 表示返回 5xx 代表健康。
// 若希望多种返回码都可代表健康，则将相应的值相加。注意：TCP监听器的HTTP健康检查方式，只支持指定一种健康检查状态码。
// TestTransProbeMatch test translate probe match
func TestTransProbeMatch(t *testing.T) {
	for i := 1; i < 32; i++ {
		healthCheck := &networkextensionv1.ListenerHealthCheck{
			HTTPCode: i,
		}
		match := transAgProbeMatch(healthCheck)
		fmt.Printf("%d ", i)
		for _, code := range match.StatusCodes {
			fmt.Printf(" %s", *code)
		}
		fmt.Println()
	}
}
