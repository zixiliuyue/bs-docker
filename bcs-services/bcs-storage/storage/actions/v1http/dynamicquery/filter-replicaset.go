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
 *
 */

package dynamicquery

import (
	"github.com/Tencent/bk-bcs/bcs-common/pkg/odm/operator"
)

// ReplicaSetFilter filter of replica set
type ReplicaSetFilter struct {
	ClusterId         string `json:"clusterId" filter:"clusterId"`
	Name              string `json:"name,omitempty" filter:"resourceName"`
	Namespace         string `json:"namespace,omitempty" filter:"namespace"`
	Replicas          string `json:"replicas,omitempty" filter:"data.status.replicas,int"`
	AvailableReplicas string `json:"availableReplicas,omitempty" filter:"data.status.availableReplicas,int"`
	ReadyReplicas     string `json:"readyReplicas,omitempty" filter:"data.status.readyReplicas,int"`
	CreateTimeBegin   string `json:"createTimeBegin,omitempty" filter:"data.metadata.creationTimestamp,timeL"`
	CreateTimeEnd     string `json:"createTimeEnd,omitempty" filter:"data.metadata.creationTimestamp,timeR"`
}

const rsNestedTimeLayout = NestedTimeLayout

// GetCondition get condition
func (t ReplicaSetFilter) GetCondition() *operator.Condition {
	return qGenerate(t, rsNestedTimeLayout)
}