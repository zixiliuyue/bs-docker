

package etcd

import (
	"context"

	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckEndpointExist check if endpoint exists
func (store *managerStore) CheckEndpointExist(endpoint *commtypes.BcsEndpoint) (string, bool) {
	client := store.BkbcsClient.BcsEndpoints(endpoint.NameSpace)
	v2End, err := client.Get(context.Background(), endpoint.Name, metav1.GetOptions{})
	if err == nil {
		return v2End.ResourceVersion, true
	}

	return "", false
}

// SaveEndpoint save endpoint to db
func (store *managerStore) SaveEndpoint(endpoint *commtypes.BcsEndpoint) error {
	err := store.checkNamespace(endpoint.NameSpace)
	if err != nil {
		return err
	}

	client := store.BkbcsClient.BcsEndpoints(endpoint.NameSpace)
	v2End := &v2.BcsEndpoint{
		TypeMeta: metav1.TypeMeta{
			Kind:       CrdBcsEndpoint,
			APIVersion: ApiversionV2,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        endpoint.Name,
			Namespace:   endpoint.NameSpace,
			Labels:      store.filterSpecialLabels(endpoint.Labels),
			Annotations: endpoint.Annotations,
		},
		Spec: v2.BcsEndpointSpec{
			BcsEndpoint: *endpoint,
		},
	}

	rv, exist := store.CheckEndpointExist(endpoint)
	if exist {
		v2End.ResourceVersion = rv
		_, err = client.Update(context.Background(), v2End, metav1.UpdateOptions{})
	} else {
		_, err = client.Create(context.Background(), v2End, metav1.CreateOptions{})
	}
	return err
}

// FetchEndpoint get endpoint by name and namespace
func (store *managerStore) FetchEndpoint(ns, name string) (*commtypes.BcsEndpoint, error) {
	client := store.BkbcsClient.BcsEndpoints(ns)
	v2End, err := client.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &v2End.Spec.BcsEndpoint, nil
}

// DeleteEndpoint delete deployment by name and namespace
func (store *managerStore) DeleteEndpoint(ns, name string) error {
	client := store.BkbcsClient.BcsEndpoints(ns)
	err := client.Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}
