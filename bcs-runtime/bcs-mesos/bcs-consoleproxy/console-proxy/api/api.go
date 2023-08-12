

// Package api xxx
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/common/ssl"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/manager"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-consoleproxy/console-proxy/types"
)

// Router is api router
type Router struct {
	sync.RWMutex
	conf    *config.ConsoleProxyConfig
	backend manager.Manager
}

// CreateExecReq is createExec request struct
type CreateExecReq struct {
	ContainerID string   `json:"container_id,omitempty"`
	Cmd         []string `json:"cmd,omitempty"`
	User        string   `json:"user,omitempty"`
}

// ResizeExecReq is resizeExec request struct
type ResizeExecReq struct {
	ExecID string `json:"exec_id,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

// NewRouter return api router
func NewRouter(b manager.Manager, conf *config.ConsoleProxyConfig) *Router {
	r := &Router{
		backend: b,
		conf:    conf,
	}

	r.initRoutes()
	return r
}

func (r *Router) initRoutes() {

	// handler container web console
	mux := http.NewServeMux()
	mux.HandleFunc("/bcsapi/v1/consoleproxy/create_exec", r.createExec)
	mux.HandleFunc("/bcsapi/v1/consoleproxy/start_exec", r.startExec)
	mux.HandleFunc("/bcsapi/v1/consoleproxy/resize_exec", r.resizeExec)
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", r.conf.Address, r.conf.Port),
		Handler: mux,
	}
	if r.conf.ServCert.IsSSL {
		tlsConf, err := ssl.ServerTslConf(r.conf.ServCert.CAFile, r.conf.ServCert.CertFile, r.conf.ServCert.KeyFile,
			r.conf.ServCert.CertPasswd)
		if err != nil {
			blog.Error("fail to load certfile, err:%s", err.Error())
			return
		}
		s.TLSConfig = tlsConf
		blog.Info("Start https service on(%s:%d)", r.conf.Address, r.conf.Port)
		go func() {
			err := s.ListenAndServeTLS("", "")
			fmt.Printf("tls server failed: %v\n", err)
		}()
	} else {
		blog.Info("Start http service on(%s:%d)", r.conf.Address, r.conf.Port)
		go func() {
			err := s.ListenAndServe()
			fmt.Printf("insecure server failed: %v\n", err)
		}()
	}
}

func (r *Router) createExec(w http.ResponseWriter, req *http.Request) {

	var createExecReq CreateExecReq
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createExecReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if createExecReq.ContainerID == "" {
		http.Error(w, "container_id must be provided", http.StatusBadRequest)
		return
	}

	if createExecReq.User == "" {
		createExecReq.User = "root"
	}
	if createExecReq.Cmd == nil {
		createExecReq.Cmd = r.conf.Cmd
	}
	webconsole := &types.WebSocketConfig{
		ContainerID: createExecReq.ContainerID,
		User:        createExecReq.User,
		Cmd:         createExecReq.Cmd,
	}

	r.backend.CreateExec(w, req, webconsole)
}

func (r *Router) startExec(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	execID := req.FormValue("exec_id")
	containerID := req.FormValue("container_id")

	webconsole := &types.WebSocketConfig{
		ExecID:      execID,
		ContainerID: containerID,
		Origin:      req.Header.Get("Origin"),
	}

	// handler container web console
	r.backend.StartExec(w, req, webconsole)
}

func (r *Router) resizeExec(w http.ResponseWriter, req *http.Request) {

	var resizeExecReq ResizeExecReq
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&resizeExecReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if resizeExecReq.ExecID == "" {
		http.Error(w, "exec_id must be provided", http.StatusBadRequest)
		return
	}

	webconsole := &types.WebSocketConfig{
		ExecID: resizeExecReq.ExecID,
		Height: resizeExecReq.Height,
		Width:  resizeExecReq.Width,
	}

	r.backend.ResizeExec(w, req, webconsole)
}
