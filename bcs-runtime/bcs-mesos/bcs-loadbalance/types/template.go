

package types

// TemplateData data holder for haproxy.cfg.template
type TemplateData struct {
	HTTP    HTTPServiceInfoList      // HTTP service info
	HTTPS   HTTPServiceInfoList      // HTTPS service info
	TCP     FourLayerServiceInfoList // TCP service info
	UDP     FourLayerServiceInfoList // UDP service info
	LogFlag bool                     // log flag, true will open log writer
	SSLCert string                   // SSL certificate path, true will listen https
}
