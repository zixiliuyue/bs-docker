

package types

// HaProxyRequest xxx
type HaProxyRequest struct {
	User      string `json:"user"`
	SetID     string `json:"set_id"`
	AppID     int    `json:"app_id"`
	Name      string `json:"name"`
	NameSpace string `json:"namespace"`
	Nbproc    int    `json:"nbproc"`
	Replicas  int    `json:"replicas"`
	Services  []struct {
		Mode        string `json:"mode"`
		ServiceName string `json:"k8s_svc_name"`
		ListenPort  string `json:"listen_port"`
		SecretName  string `json:"secret_name"`
		SslCert     string `json:"ssl_cert"`
	} `json:"services"`
}
