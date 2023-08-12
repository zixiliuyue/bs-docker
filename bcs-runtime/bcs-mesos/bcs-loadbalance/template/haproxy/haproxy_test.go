
package haproxy

import (
	"bytes"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/bcs-loadbalance/types"
	"html/template"
	"io/ioutil"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	haproxyManager := &Manager{
		envConfig: loadEnvConfig(),
	}
	testData := []struct {
		tmpData   *types.TemplateData
		referFile string
	}{
		{
			tmpData: &types.TemplateData{
				HTTP: types.HTTPServiceInfoList{
					{
						ServiceInfo: types.ServiceInfo{
							Name:        "bcs_frontend_1_8089",
							ServicePort: 8089,
							Balance:     "roundrobin",
						},
						BCSVHost: "www.test.com",
						Backends: []types.HTTPBackend{
							{
								Path:         "/test1",
								UpstreamName: "bcs_frontend_1_8089_test1",
								BackendList: []types.Backend{
									{
										IP:     "127.0.10.1",
										Port:   8888,
										Weight: 10,
									},
									{
										IP:     "127.0.10.2",
										Port:   8888,
										Weight: 10,
									},
									{
										IP:     "127.0.10.3",
										Port:   8888,
										Weight: 10,
									},
								},
							},
						},
					},
				},
				TCP: types.FourLayerServiceInfoList{
					types.FourLayerServiceInfo{
						ServiceInfo: types.ServiceInfo{
							Name:        "bcs_frontend_3000",
							ServicePort: 3000,
							Balance:     "roundrobin",
						},
						Backends: []types.Backend{
							{
								IP:     "127.0.10.7",
								Port:   8888,
								Weight: 10,
							},
							{
								IP:     "127.0.10.8",
								Port:   8888,
								Weight: 10,
							},
							{
								IP:     "127.0.10.9",
								Port:   8888,
								Weight: 10,
							},
						},
					},
				},
			},
			referFile: "./test/haproxy.test1.cfg",
		},
	}

	for _, test := range testData {
		config := haproxyManager.convertData(test.tmpData)
		config.generateServerName()
		config.generateRenderData()
		var err error
		var tmpl *template.Template
		// loading template file
		tmpl, err = template.ParseFiles("../../image/config/haproxy.cfg.template")
		if err != nil {
			t.Error(err.Error())
		}
		var b bytes.Buffer
		tmpl.Execute(&b, config)
		referData, err := ioutil.ReadFile(test.referFile)
		if err != nil {
			t.Errorf("read refer file %s failed, err %s", test.referFile, err.Error())
		}
		if b.String() != string(referData) {
			t.Errorf("expect %s but get %s", string(referData), b.String())
		}
	}
}
