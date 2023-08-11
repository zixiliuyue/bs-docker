/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-services/bcs-argocd-manager/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ArgocdInstances returns a ArgocdInstanceInformer.
	ArgocdInstances() ArgocdInstanceInformer
	// ArgocdPlugins returns a ArgocdPluginInformer.
	ArgocdPlugins() ArgocdPluginInformer
	// ArgocdProjects returns a ArgocdProjectInformer.
	ArgocdProjects() ArgocdProjectInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ArgocdInstances returns a ArgocdInstanceInformer.
func (v *version) ArgocdInstances() ArgocdInstanceInformer {
	return &argocdInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ArgocdPlugins returns a ArgocdPluginInformer.
func (v *version) ArgocdPlugins() ArgocdPluginInformer {
	return &argocdPluginInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ArgocdProjects returns a ArgocdProjectInformer.
func (v *version) ArgocdProjects() ArgocdProjectInformer {
	return &argocdProjectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
