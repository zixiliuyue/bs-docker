

// Package listenerclient xxx
package listenerclient

import (
	"context"
	"fmt"

	cloudListenerType "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/network/v1"
	listenerClientV1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/generated/clientset/versioned/typed/network/v1"
	listerV1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/generated/listers/network/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

// ListenerClient xxx
type ListenerClient struct {
	client  listenerClientV1.NetworkV1Interface
	lister  listerV1.CloudListenerLister
	clbName string
}

// NewListenerClient xxx
func NewListenerClient(clbname string, client listenerClientV1.NetworkV1Interface,
	lister listerV1.CloudListenerLister) (Interface, error) {

	return &ListenerClient{
		client:  client,
		lister:  lister,
		clbName: clbname,
	}, nil
}

// ListListeners xxx
func (lc *ListenerClient) ListListeners() ([]*cloudListenerType.CloudListener, error) {
	selector := labels.NewSelector()
	requirement, err := labels.NewRequirement("bmsf.tencent.com/clbname", selection.Equals, []string{lc.clbName})
	if err != nil {
		return nil, fmt.Errorf("create requirement failed, err %s", err.Error())
	}
	selector = selector.Add(*requirement)
	return lc.lister.List(selector)
}

// Create xxx
func (lc *ListenerClient) Create(listener *cloudListenerType.CloudListener) error {
	_, err := lc.client.CloudListeners(listener.GetNamespace()).Create(
		context.Background(), listener, metav1.CreateOptions{})
	return err
}

// Update xxx
func (lc *ListenerClient) Update(listener *cloudListenerType.CloudListener) error {
	old, err := lc.lister.CloudListeners(listener.GetNamespace()).Get(listener.GetName())
	if err != nil {
		_, createErr := lc.client.CloudListeners(listener.GetNamespace()).Create(
			context.Background(), listener, metav1.CreateOptions{})
		return createErr
	}
	listener.SetResourceVersion(old.GetResourceVersion())
	_, err = lc.client.CloudListeners(listener.GetNamespace()).Update(
		context.Background(), listener, metav1.UpdateOptions{})
	return err
}

// Delete xxx
func (lc *ListenerClient) Delete(listener *cloudListenerType.CloudListener) error {
	return lc.client.CloudListeners(listener.GetNamespace()).Delete(
		context.Background(), listener.GetName(), metav1.DeleteOptions{})
}
