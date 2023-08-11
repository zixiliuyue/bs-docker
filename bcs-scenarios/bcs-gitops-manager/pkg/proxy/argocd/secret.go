/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package argocd

import (
	"context"
	"net/http"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-common/pkg/auth/iam"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// SecretPlugin for internal project authorization
type SecretPlugin struct {
	*mux.Router
	middleware MiddlewareInterface
	session    *Session
}

// all argocd secret URL:
// * required Project Edit projectPermission
// NOTE 因为密钥是敏感数据，所以get权限需要使用项目可编辑权限，后续可能需要可配置
// POST：   /api/v1/secrets/{project}/{path}，创建
// PUT：    /api/v1/secrets/{project}/{path}，指定更新(创建新版本)
// DELETE： /api/v1/secrets/{project}/{path}，完全销毁secret路径
// GET：/api/v1/secrets/{project}/{path}?version={version}，获取版本密钥secret信息
// GET：/api/v1/secrets/{project}/{path}/metadata，返回secret的metadata，包括版本信息以及创建修改时间等
// GET：/api/v1/secrets/{project}/{path}/version，返回secret的版本信息
// POST：/api/v1/secrets/{project}/{path}/rollback，secret版本回滚，会根据某个版本创建新版本

// * required Project View projectPermission
// GET：/api/v1/secrets/list?project={project}&path={path}，返回列表
//

// Init all project sub path handler
// project plugin is a subRouter, all path registered is relative
func (plugin *SecretPlugin) Init() error {
	// Create or Update(create new version) with preifx /api/v1/secrets/{project}/{path}
	plugin.Path("/{project}/{path}").Methods("POST", "PUT").
		Handler(plugin.middleware.HttpWrapper(plugin.createPutSecretHandler))
	// Delete with preifx /api/v1/secrets/{project}/{path}
	plugin.Path("/{project}/{path}").Methods("DELETE").
		Handler(plugin.middleware.HttpWrapper(plugin.deleteSecretHandler))
	// Get with preifx /api/v1/secrets/{project}/{path}?version={version}
	plugin.Path("/{project}/{path}").Methods("GET").
		Handler(plugin.middleware.HttpWrapper(plugin.getSecretHandler)).Queries("version", "{version}")

	// force check query, GET /api/v1/secrets/{project}/{path}/list
	// 这里因为path可能为空，无法在url中定义空字段，所以用parameter来做path
	plugin.Path("/{project}/list").Methods("GET").Queries("path", "{path}").
		Handler(plugin.middleware.HttpWrapper(plugin.listSecretHandler))
	// force check query, GET /api/v1/secrets/{project}/{path}/metadata
	plugin.Path("/{project}/{path}/metadata").Methods("GET").
		Handler(plugin.middleware.HttpWrapper(plugin.getMetadataHandler))
	// force check query, GET /api/v1/secrets/{project}/{path}/version
	plugin.Path("/{project}/{path}/version").Methods("GET").
		Handler(plugin.middleware.HttpWrapper(plugin.getVersionHandler))
	// force check query, POST /api/v1/secrets/{project}/{path}/rollback?version={version}
	plugin.Path("/{project}/{path}/rollback").Methods("POST").Queries("version", "{version}").
		Handler(plugin.middleware.HttpWrapper(plugin.rollbackHandler))

	blog.Infof("secret plugin init successfully")
	return nil
}

// POST,PUT /api/v1/secrets, create new secret && update(create new version) new secret
// validate project detail from request
func (plugin *SecretPlugin) createPutSecretHandler(ctx context.Context, r *http.Request) *httpResponse {
	project := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, project, iam.ProjectEdit)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project '%s' edit permission failed", project),
		}
	}

	resp := plugin.middleware.ProxySecretRequest(r)
	return &httpResponse{
		obj: resp,
	}
}

// Delete with preifx /api/v1/secrets/{project}/{path}
func (plugin *SecretPlugin) deleteSecretHandler(ctx context.Context, r *http.Request) *httpResponse {
	projectName := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, projectName, iam.ProjectEdit)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project permission failed"),
		}
	}

	resp := plugin.middleware.ProxySecretRequest(r)
	return &httpResponse{
		obj: resp,
	}
}

// Get with preifx /api/v1/secrets/{project}/{path}?version={version}
func (plugin *SecretPlugin) getSecretHandler(ctx context.Context, r *http.Request) *httpResponse {
	projectName := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, projectName, iam.ProjectEdit)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project permission failed"),
		}
	}

	resp := plugin.middleware.ProxySecretRequest(r)
	if resp.Data == nil {
		resp.Data = map[string]interface{}{}
	}
	return &httpResponse{
		obj: resp,
	}
}

// GET /api/v1/secrets/{project}/list？path=${path}
func (plugin *SecretPlugin) listSecretHandler(ctx context.Context, r *http.Request) *httpResponse {
	projectName := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, projectName, iam.ProjectView)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project permission failed"),
		}
	}

	resp := plugin.middleware.ProxySecretRequest(r)
	if resp.Data == nil {
		resp.Data = []string{}
	}
	return &httpResponse{
		obj: resp,
	}
}

// GET /api/v1/secrets/{project}/{path}/metadata
func (plugin *SecretPlugin) getMetadataHandler(ctx context.Context, r *http.Request) *httpResponse {
	projectName := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, projectName, iam.ProjectEdit)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project permission failed"),
		}
	}
	resp := plugin.middleware.ProxySecretRequest(r)
	return &httpResponse{
		obj: resp,
	}
}

// GET /api/v1/secrets/{project}/{path}/version
func (plugin *SecretPlugin) getVersionHandler(ctx context.Context, r *http.Request) *httpResponse {
	projectName := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, projectName, iam.ProjectEdit)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project permission failed"),
		}
	}

	resp := plugin.middleware.ProxySecretRequest(r)
	return &httpResponse{
		obj: resp,
	}
}

// POST /api/v1/secrets/{project}/{path}/rollback
func (plugin *SecretPlugin) rollbackHandler(ctx context.Context, r *http.Request) *httpResponse {
	projectName := mux.Vars(r)["project"]

	_, statusCode, err := plugin.middleware.CheckProjectPermission(ctx, projectName, iam.ProjectEdit)
	if statusCode != http.StatusOK {
		return &httpResponse{
			statusCode: statusCode,
			err:        errors.Wrapf(err, "check project permission failed"),
		}
	}

	resp := plugin.middleware.ProxySecretRequest(r)
	return &httpResponse{
		obj: resp,
	}
}
