

package check

import (
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networkextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/networkextension/v1"
)

func TestGetPortConflictMap(t *testing.T) {
	pChecker := &PortBindChecker{}
	portbindingList := &networkextensionv1.PortBindingList{
		Items: []networkextensionv1.PortBinding{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-1",
					Namespace: "default",
				},
				Spec: networkextensionv1.PortBindingSpec{
					PortBindingList: []*networkextensionv1.PortBindingItem{
						{
							PoolName:      "pool1",
							PoolNamespace: "default",
							PoolItemName:  "item1",
							StartPort:     1230,
							EndPort:       1250,
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-2",
					Namespace: "default",
				},
				Spec: networkextensionv1.PortBindingSpec{
					PortBindingList: []*networkextensionv1.PortBindingItem{
						{
							PoolName:      "pool1",
							PoolNamespace: "default",
							PoolItemName:  "item1",
							StartPort:     1250,
							EndPort:       1270,
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-3",
					Namespace: "default",
				},
				Spec: networkextensionv1.PortBindingSpec{
					PortBindingList: []*networkextensionv1.PortBindingItem{
						{
							PoolName:      "pool1",
							PoolNamespace: "default",
							PoolItemName:  "item1",
							StartPort:     1230,
							EndPort:       1250,
						},
					},
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "pod-4",
					Namespace: "default",
				},
				Spec: networkextensionv1.PortBindingSpec{
					PortBindingList: []*networkextensionv1.PortBindingItem{
						{
							PoolName:      "pool1",
							PoolNamespace: "default",
							PoolItemName:  "item1",
							StartPort:     1250,
							EndPort:       1270,
						},
					},
				},
			},
		},
	}
	conflictMap := pChecker.getPortConflictMap(portbindingList)
	fmt.Println(conflictMap)
}
