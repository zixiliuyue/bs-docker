

package imageacceleration

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
)

// imagePullSecret defines the .dockerconfigjson of secret data
type imagePullSecret struct {
	Auths map[string]*imagePullAuth `json:"auths"`
}

type imagePullAuth struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Auth     string `json:"auth"`
}

// handleImagePullSecret will get the imagePullSecrets from pod. And then for-loop them, get the
// .dockerconfigjson value of them. Comparing the registry prefix that defines in the secret with
// prefix defines at configMapKeyMapping of configMap.
// Secret will be updated if the registry auth not found(or not same, username/password)
func (h *Handler) handleImagePullSecret(configmapMapping map[string]string, pod *corev1.Pod) {
	for _, secretItem := range pod.Spec.ImagePullSecrets {
		secret, err := h.cacheManager.GetSecret(pod.Namespace, secretItem.Name)
		if err != nil {
			blog.Warnf("image acceleration get secret failed: %s", err.Error())
			continue
		}
		if err := h.compareAndUpdateSecret(secret, configmapMapping); err != nil {
			blog.Warnf("image acceleration compare and update secret '%s/%s' failed: %s",
				secret.Namespace, secret.Name, err.Error())
		}
	}
}

func (h *Handler) compareAndUpdateSecret(secret *corev1.Secret, configmapMapping map[string]string) error {
	pullSecret, err := h.parseImagePullSecret(secret)
	if err != nil {
		return errors.Wrapf(err, "parse secret failed")
	}
	changed := false
	for imagePrefix, resultPrefix := range configmapMapping {
		registryAuth, ok := pullSecret.Auths[imagePrefix]
		if !ok {
			continue
		}
		resultAuth, ok := pullSecret.Auths[resultPrefix]
		if ok && (registryAuth.Username == resultAuth.Username && registryAuth.Password == resultAuth.Password &&
			registryAuth.Auth == resultAuth.Auth) {
			continue
		}
		pullSecret.Auths[resultPrefix] = registryAuth
		changed = true
	}
	if !changed {
		return nil
	}
	bs, err := json.Marshal(pullSecret)
	if err != nil {
		return errors.Wrapf(err, "marshal secret failed")
	}
	secret.Data[secretImagePullItem] = bs
	if err := h.cacheManager.UpdateSecret(context.Background(), secret); err != nil {
		return errors.Wrapf(err, "update secret failed")
	}
	blog.Infof("image acceleration update secret '%s/%s' success", secret.Namespace, secret.Name)
	return nil
}

func (h *Handler) parseImagePullSecret(secret *corev1.Secret) (*imagePullSecret, error) {
	bs, ok := secret.Data[secretImagePullItem]
	if !ok {
		return nil, errors.Errorf("secret not have '%s'", secretImagePullItem)
	}
	pullSecret := new(imagePullSecret)
	if err := json.Unmarshal(bs, pullSecret); err != nil {
		return nil, errors.Wrapf(err, "unmarshal secret data failed")
	}
	if len(pullSecret.Auths) == 0 {
		return nil, errors.Errorf("secret dockerconfigjson data is empty")
	}
	return pullSecret, nil
}
