

package health

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

var (
	// ErrSchemeInValid invalid scheme
	ErrSchemeInValid = errors.New("scheme invalid, should be http/https")
	// ErrHealthConfigNotInited show HealthConfig not inited
	ErrHealthConfigNotInited = errors.New("healthConfig not inited")
)

// HealthCheck is interface for check addr:port health
type HealthCheck interface {
	IsHTTPAPIHealth(addr string, port uint32) bool
}

const (
	schemeHTTPS = "https"
	schemeHTTP  = "http"
)

func validateScheme(scheme string) error {
	if scheme != schemeHTTPS && scheme != schemeHTTP {
		return ErrSchemeInValid
	}

	return nil
}

// NewHealthConfig init HealthConfig
func NewHealthConfig(scheme string, path string) (HealthCheck, error) {
	err := validateScheme(scheme)
	if err != nil {
		return nil, err
	}

	return &HealthConfig{
		Shem: scheme,
		Path: path,
	}, nil
}

// HealthConfig conf immutable schem/path
type HealthConfig struct {
	Shem string
	Path string
}

// IsHTTPAPIHealth for check schem://addr:port/Path health
func (hc *HealthConfig) IsHTTPAPIHealth(addr string, port uint32) bool {
	if hc == nil {
		blog.Errorf("IsHTTPAPIHealth empty: %v", ErrHealthConfigNotInited)
		return false
	}

	url := fmt.Sprintf("%s://%s:%d%s", hc.Shem, addr, port, hc.Path)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// NOCC:gas/tls(设计如此)
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		blog.Errorf("http health check failed, %v", err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		blog.Errorf("IsHTTPAPIHealth[%s] error: %v", url, err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}
