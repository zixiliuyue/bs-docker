

package imageacceleration

import (
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-webhook-server/internal/pluginmanager"
)

const (
	// pluginName defines the name of image acceleration plugin
	pluginName = "imageacceleration"
	// configMapName defines the name of configmap that can enable image acceleration
	configMapName = "bcs-image-acceleration"
	// secretImagePullItem defines the image pull data of secret
	secretImagePullItem = ".dockerconfigjson" // // NOCC:gas/crypto(设计如此)

	// configMapKeyEnabled key of configmap, namespace will enable image acceleration if value is "true"
	configMapKeyEnabled = "enabled"
	// configMapKeyMapping key of configmap, users can use it to define custom registry mapping
	configMapKeyMapping = "mapping"

	// containerImageKey the image key path of pod
	containerImageKey = "/spec/containers/%d/image"

	// patchOpReplace the type of patch op that replace the key value
	patchOpReplace = "replace"
)

func init() {
	h := &Handler{}
	pluginmanager.Register(pluginName, h)
	mh := &Handler{}
	pluginmanager.RegisterMesos(pluginName, mh)
}
