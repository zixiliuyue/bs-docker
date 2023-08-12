

package etcd

import (
	"context"
	"fmt"

	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CheckCustomResourceRegisterExist check if crd is registered
func (store *managerStore) CheckCustomResourceRegisterExist(crr *commtypes.Crr) (string, bool) {
	client := store.BkbcsClient.Crrs(DefaultNamespace)
	obj, err := client.Get(context.Background(), crr.Spec.Names.Kind, metav1.GetOptions{})
	if err == nil {
		return obj.ResourceVersion, true
	}

	return "", false
}

// SaveCustomResourceRegister save custom resource register
func (store *managerStore) SaveCustomResourceRegister(crr *commtypes.Crr) error {
	client := store.BkbcsClient.Crrs(DefaultNamespace)
	v2Crr := &v2.Crr{
		TypeMeta: metav1.TypeMeta{
			Kind:       CrdCrr,
			APIVersion: ApiversionV2,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      crr.Spec.Names.Kind,
			Namespace: DefaultNamespace,
		},
		Spec: v2.CrrSpec{
			Crr: *crr,
		},
	}

	var err error
	rv, exist := store.CheckCustomResourceRegisterExist(crr)
	if exist {
		v2Crr.ResourceVersion = rv
		_, err = client.Update(context.Background(), v2Crr, metav1.UpdateOptions{})
	} else {
		_, err = client.Create(context.Background(), v2Crr, metav1.CreateOptions{})
	}
	return err
}

// DeleteCustomResourceRegister delete custom resource register
func (store *managerStore) DeleteCustomResourceRegister(name string) error {
	client := store.BkbcsClient.Crrs(DefaultNamespace)
	err := client.Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

// ListCustomResourceRegister list custom resource register
func (store *managerStore) ListCustomResourceRegister() ([]*commtypes.Crr, error) {
	client := store.BkbcsClient.Crrs(DefaultNamespace)
	v2Crrs, err := client.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	crrs := make([]*commtypes.Crr, 0, len(v2Crrs.Items))
	for _, crr := range v2Crrs.Items {
		obj := crr.Spec.Crr
		crrs = append(crrs, &obj)
	}

	return crrs, nil
}

// getCrdNamespace xxx
// crd namespace = crd.kind=crd.namespace
func getCrdNamespace(kind, ns string) string {
	return fmt.Sprintf("%s-%s", kind, ns)
}

// CheckCustomResourceDefinitionExist check if custom resource definition exists
func (store *managerStore) CheckCustomResourceDefinitionExist(crd *commtypes.Crd) (string, bool) {
	client := store.BkbcsClient.Crds(getCrdNamespace(string(crd.Kind), crd.NameSpace))
	v2Crd, err := client.Get(context.Background(), crd.Name, metav1.GetOptions{})
	if err == nil {
		return v2Crd.ResourceVersion, true
	}

	return "", false
}

// SaveCustomResourceDefinition save custom resource definition into db
func (store *managerStore) SaveCustomResourceDefinition(crd *commtypes.Crd) error {
	// crd namespace = crd.kind=crd.namespace
	realNs := getCrdNamespace(string(crd.Kind), crd.NameSpace)
	err := store.checkNamespace(realNs)
	if err != nil {
		return err
	}

	client := store.BkbcsClient.Crds(realNs)
	v2Crd := &v2.Crd{
		TypeMeta: metav1.TypeMeta{
			Kind:       CrdCrd,
			APIVersion: ApiversionV2,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      crd.Name,
			Namespace: realNs,
		},
		Spec: v2.CrdSpec{
			Crd: *crd,
		},
	}

	rv, exist := store.CheckCustomResourceDefinitionExist(crd)
	if exist {
		v2Crd.ResourceVersion = rv
		_, err = client.Update(context.Background(), v2Crd, metav1.UpdateOptions{})
	} else {
		_, err = client.Create(context.Background(), v2Crd, metav1.CreateOptions{})
	}
	return err
}

// DeleteCustomResourceDefinition delete custom resource definition
func (store *managerStore) DeleteCustomResourceDefinition(kind, ns, name string) error {
	client := store.BkbcsClient.Crds(getCrdNamespace(kind, ns))
	err := client.Delete(context.Background(), name, metav1.DeleteOptions{})
	return err
}

// ListAllCrds list all crds
func (store *managerStore) ListAllCrds(kind string) ([]*commtypes.Crd, error) {
	client := store.BkbcsClient.Crds("")
	v2Crds, err := client.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	crds := make([]*commtypes.Crd, 0, len(v2Crds.Items))
	for _, crd := range v2Crds.Items {
		if strings.Contains(crd.Namespace, kind) {
			obj := crd.Spec.Crd
			crds = append(crds, &obj)
		}
	}

	return crds, nil
}

// ListCustomResourceDefinition list crds by kind and namespace
func (store *managerStore) ListCustomResourceDefinition(kind, ns string) ([]*commtypes.Crd, error) {
	client := store.BkbcsClient.Crds(getCrdNamespace(kind, ns))
	v2Crds, err := client.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	crds := make([]*commtypes.Crd, 0, len(v2Crds.Items))
	for _, crd := range v2Crds.Items {
		obj := crd.Spec.Crd
		crds = append(crds, &obj)
	}

	return crds, nil
}

// FetchCustomResourceDefinition get custom resource definition
func (store *managerStore) FetchCustomResourceDefinition(kind, ns, name string) (*commtypes.Crd, error) {
	client := store.BkbcsClient.Crds(getCrdNamespace(kind, ns))
	v2Crd, err := client.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &v2Crd.Spec.Crd, nil
}
