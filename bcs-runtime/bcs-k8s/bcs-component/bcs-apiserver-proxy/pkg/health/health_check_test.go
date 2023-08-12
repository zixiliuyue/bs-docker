

package health

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var (
	addr        = "127.0.0.1"
	port uint32 = 8888
)

func TestHealthConfig_IsHTTPAPIHealth(t *testing.T) {
	server := fmt.Sprintf("%s:%d", addr, port)
	startHTTPServer(server)

	time.Sleep(2 * time.Second)

	health, err := NewHealthConfig("http", "/health")
	if err != nil {
		t.Fatalf("NewHealthConfig failed: %v", err)
		return
	}

	ok := health.IsHTTPAPIHealth(addr, port)
	if !ok {
		t.Fatalf("IsHTTPAPIHealth failed")
		return
	}

	t.Log("IsHTTPAPIHealth successful")
}

func startHTTPServer(addr string) {
	http.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("ok"))
	})
	go http.ListenAndServe(addr, nil)
}
