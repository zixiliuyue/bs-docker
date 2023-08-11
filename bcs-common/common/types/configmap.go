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

package types

// BcsConfigMapSourceType xxx
type BcsConfigMapSourceType string

const (
	// BcsConfigMapSourceType_FILE xxx
	BcsConfigMapSourceType_FILE BcsConfigMapSourceType = "file"
	// BcsConfigMapSourceType_HTTP xxx
	BcsConfigMapSourceType_HTTP BcsConfigMapSourceType = "http"
	// BcsConfigMapSourceType_HTTPS xxx
	BcsConfigMapSourceType_HTTPS BcsConfigMapSourceType = "https"
	// BcsConfigMapSourceType_FTP xxx
	BcsConfigMapSourceType_FTP BcsConfigMapSourceType = "ftp"
	// BcsConfigMapSourceType_FTPS xxx
	BcsConfigMapSourceType_FTPS BcsConfigMapSourceType = "ftps"
)

// BcsConfigMapItem item hold configmap
type BcsConfigMapItem struct {
	Type         BcsConfigMapSourceType `json:"type"`
	Content      string                 `json:"content"`
	RemoteUser   string                 `josn:"remoteUser,omitempty"`
	RemotePasswd string                 `json:"remotePasswd,omitempty"`
}

// BcsConfigMap definition
type BcsConfigMap struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata"`
	Data       map[string]BcsConfigMapItem `json:"datas"` // map key for index, map value is content
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsConfigMap) DeepCopyInto(out *BcsConfigMap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]BcsConfigMapItem, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BcsConfigMapSpec.
func (in *BcsConfigMap) DeepCopy() *BcsConfigMap {
	if in == nil {
		return nil
	}
	out := new(BcsConfigMap)
	in.DeepCopyInto(out)
	return out
}