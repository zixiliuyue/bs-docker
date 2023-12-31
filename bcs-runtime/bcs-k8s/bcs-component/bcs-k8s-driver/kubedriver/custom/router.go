

package custom

import (
	"regexp"
	"strings"
)

// APIUrl url for
type APIUrl string

// UrlHandlerMap key can't be same with K8S URI
// Warning: It can not be start with "api" or "apis"
var UrlHandlerMap = map[string]APIHandler{
	"cluster/resources": &ClusterResourceAPIHandler{},
	"bcsclient/.+":      &BcsClientAPIHandler{},
}

// APIRouterInterface interface for http request route
type APIRouterInterface interface {
	Route(subPath string) *APIHandler
}

// APIRouter implementation
type APIRouter struct {
	UrlHandlerMap map[string]APIHandler
}

// NewRouter implementation creator
func NewRouter() (ar *APIRouter) {
	ar = &APIRouter{}
	ar.InitRegisteredUrls()
	return
}

// InitRegisteredUrls init URL
func (ar *APIRouter) InitRegisteredUrls() {
	ar.UrlHandlerMap = UrlHandlerMap
}

// Route http url request route
func (ar *APIRouter) Route(subPath string) (handler APIHandler) {
	subPathWithoutQuery := strings.Split(subPath, "?")[0]
	for url, handler := range ar.UrlHandlerMap {
		if m, _ := regexp.MatchString(url, subPathWithoutQuery); m {
			return handler
		}
	}
	return nil
}
