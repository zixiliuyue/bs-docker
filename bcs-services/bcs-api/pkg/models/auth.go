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

package models

import "time"

const (
	// PermBackendTypeDefault ..
	PermBackendTypeDefault = iota + 1
	// PermBackendTypeSyncOnce ..
	PermBackendTypeSyncOnce

	// ClusterPermNameView ..
	ClusterPermNameView = "view"
	// ClusterPermNameEdit ..
	ClusterPermNameEdit = "edit"
)

// UserClusterPermission stores the user cluster permission
type UserClusterPermission struct {
	ID uint `gorm:"primary_key"`
	// Backend is where this permission comes from:
	//   - default: this permission was generated by bke itself
	//   - sync_once: this permission was synced from other sources, one cluster at a time
	//
	// The same user+cluster pair can have multiple perm records with different backends, the final permission check
	// result will be caculated by the priority of different backends.
	Backend   int    `gorm:"unique_index:idx_backend_user_cluster_id_name;default:1;"`
	UserID    uint   `gorm:"unique_index:idx_backend_user_cluster_id_name"`
	ClusterID string `gorm:"unique_index:idx_backend_user_cluster_id_name"`
	Name      string `gorm:"unique_index:idx_backend_user_cluster_id_name;size:16"`
	// IsActive means if this permission is valid, the user does not have this permission when it's false,
	// this field's purpose is to record the permission synchronize procedure even if the result is false
	IsActive  bool
	UpdatedAt time.Time
	CreatedAt time.Time
}
