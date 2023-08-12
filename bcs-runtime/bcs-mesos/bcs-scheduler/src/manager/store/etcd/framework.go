

package etcd

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// FrameworkNode node name of mesos framework
	FrameworkNode = "frameworknode"
)

// CheckFrameworkIDExist check if framework id exists
func (store *managerStore) CheckFrameworkIDExist() (string, bool) {
	client := store.BkbcsClient.Frameworks(DefaultNamespace)
	v2Fw, err := client.Get(context.Background(), FrameworkNode, metav1.GetOptions{})
	if err == nil {
		return v2Fw.ResourceVersion, true
	}

	return "", false
}

// SaveFrameworkID save framework id
func (store *managerStore) SaveFrameworkID(frameworkId string) error {
	client := store.BkbcsClient.Frameworks(DefaultNamespace)
	v2Fw := &v2.Framework{
		TypeMeta: metav1.TypeMeta{
			Kind:       CrdFramework,
			APIVersion: ApiversionV2,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      FrameworkNode,
			Namespace: DefaultNamespace,
		},
		Spec: v2.FrameworkSpec{
			FrameworkId: frameworkId,
		},
	}

	var err error
	rv, exist := store.CheckFrameworkIDExist()
	if exist {
		v2Fw.ResourceVersion = rv
		_, err = client.Update(context.Background(), v2Fw, metav1.UpdateOptions{})
	} else {
		_, err = client.Create(context.Background(), v2Fw, metav1.CreateOptions{})
	}
	return err
}

// FetchFrameworkID get framework id
func (store *managerStore) FetchFrameworkID() (string, error) {
	client := store.BkbcsClient.Frameworks(DefaultNamespace)
	v2Fw, err := client.Get(context.Background(), FrameworkNode, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	return v2Fw.Spec.FrameworkId, nil
}

// HasFrameworkID check if there is framework id
func (store *managerStore) HasFrameworkID() (bool, error) {
	_, err := store.FetchFrameworkID()
	if err != nil {
		return false, err
	}

	return true, nil
}
