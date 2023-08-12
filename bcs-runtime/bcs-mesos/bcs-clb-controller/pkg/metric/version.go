

package metric

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/version"

	restful "github.com/emicklei/go-restful"
)

// VersionMetricResp xxx
type VersionMetricResp struct {
	Version   string `json:"version"`
	Tag       string `json:"tag"`
	BuildTime string `json:"buildtime"`
	GitHash   string `json:"githash"`
}

// VersionMetric xxx
type VersionMetric struct {
}

// NewVersionMetric xxx
func NewVersionMetric() Resource {
	return &VersionMetric{}
}

// Register xxx
func (vm *VersionMetric) Register(container *restful.Container) {
	// 创建webservice
	ws := new(restful.WebService)
	// 指定路径以及支持的媒体类型
	ws.Path("/version").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.GET("/").To(vm.getVersion))
	// 创建container 不创建则使用restful.Add添加到DefaultContainer
	container.Add(ws)
}

func (vm *VersionMetric) getVersion(req *restful.Request, resp *restful.Response) {
	newResp := VersionMetricResp{
		Version:   version.BcsVersion,
		Tag:       version.BcsTag,
		BuildTime: version.BcsBuildTime,
		GitHash:   version.BcsGitHash,
	}
	resp.WriteEntity(newResp)
}
