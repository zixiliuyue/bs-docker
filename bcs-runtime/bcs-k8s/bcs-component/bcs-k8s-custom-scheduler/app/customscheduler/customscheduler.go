

package customscheduler

import (
	"fmt"

	"github.com/Tencent/bk-bcs/bcs-common/common/http/httpserver"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/config"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/pkg/actions"
)

// CustomScheduler xxx
type CustomScheduler struct {
	config   *config.CustomSchedulerConfig
	httpServ *httpserver.HttpServer
}

// NewCustomScheduler creates an CustomScheduler object
func NewCustomScheduler(conf *config.CustomSchedulerConfig) *CustomScheduler {
	customSched := &CustomScheduler{
		config:   conf,
		httpServ: httpserver.NewHttpServer(conf.Port, conf.Address, conf.Sock),
	}

	if conf.ServCert.IsSSL {
		customSched.httpServ.SetSsl(
			conf.ServCert.CAFile,
			conf.ServCert.CertFile,
			conf.ServCert.KeyFile,
			conf.ServCert.CertPasswd)
	}

	customSched.httpServ.SetInsecureServer(conf.InsecureAddress, conf.InsecurePort)

	return customSched
}

// Start xxx
func (p *CustomScheduler) Start() error {

	p.httpServ.RegisterWebServer("", nil, actions.GetApiAction())
	router := p.httpServ.GetRouter()
	webContainer := p.httpServ.GetWebContainer()
	router.Handle("/{sub_path:.*}", webContainer)
	if err := p.httpServ.ListenAndServeMux(p.config.VerifyClientTLS); err != nil {
		return fmt.Errorf("http ListenAndServe error %s", err.Error())
	}

	return nil
}
