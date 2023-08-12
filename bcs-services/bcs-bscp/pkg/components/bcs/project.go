

package bcs

import (
	"context"
	"fmt"
	"strings"
	"time"

	"bscp.io/pkg/components"
	"bscp.io/pkg/config"
)

// Project 项目信息
type Project struct {
	Name          string `json:"name"`
	ProjectId     string `json:"projectID"`
	Code          string `json:"projectCode"`
	CcBizID       string `json:"businessID"`
	Creator       string `json:"creator"`
	Kind          string `json:"kind"`
	RawCreateTime string `json:"createTime"`
}

// String :
func (p *Project) String() string {
	var displayCode string
	if p.Code == "" {
		displayCode = "-"
	} else {
		displayCode = p.Code
	}
	return fmt.Sprintf("project<%s, %s|%s|%s>", p.Name, displayCode, p.ProjectId, p.CcBizID)
}

// CreateTime xxx
func (p *Project) CreateTime() (time.Time, error) {
	return time.ParseInLocation("2006-01-02T15:04:05Z", p.RawCreateTime, config.G.Base.Location)
}

// ListAuthorizedProjects 通过 用户 获取项目信息
func ListAuthorizedProjects(ctx context.Context, username string) ([]*Project, error) {
	url := fmt.Sprintf("%s/bcsapi/v4/bcsproject/v1/authorized_projects", "")
	resp, err := components.GetClient().R().
		SetContext(ctx).
		SetHeader("X-Bcs-Username", username).
		SetAuthToken("").
		Get(url)

	if err != nil {
		return nil, err
	}

	projects := []*Project{}
	if err := components.UnmarshalBKResult(resp, projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// ListProjects 按项目 Code 查询
func ListProjects(ctx context.Context, projectCodeList []string) ([]*Project, error) {
	url := fmt.Sprintf("%s/bcsapi/v4/bcsproject/v1/projects", "")
	resp, err := components.GetClient().R().
		SetContext(ctx).
		SetHeader("X-Bcs-Username", "").
		SetQueryParam("projectCode", strings.Join(projectCodeList, ",")).
		SetAuthToken("").
		Get(url)

	if err != nil {
		return nil, err
	}

	projects := []*Project{}
	if err := components.UnmarshalBKResult(resp, projects); err != nil {
		return nil, err
	}

	return projects, nil
}
