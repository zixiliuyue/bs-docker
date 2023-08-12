
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	tkexv1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/apis/tkex/v1alpha1"
	versioned "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/generated/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubebkbcs/generated/listers/tkex/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GameStatefulSetInformer provides access to a shared informer and lister for
// GameStatefulSets.
type GameStatefulSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GameStatefulSetLister
}

type gameStatefulSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGameStatefulSetInformer constructs a new informer for GameStatefulSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGameStatefulSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGameStatefulSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGameStatefulSetInformer constructs a new informer for GameStatefulSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGameStatefulSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TkexV1alpha1().GameStatefulSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TkexV1alpha1().GameStatefulSets(namespace).Watch(context.TODO(), options)
			},
		},
		&tkexv1alpha1.GameStatefulSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *gameStatefulSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGameStatefulSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gameStatefulSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tkexv1alpha1.GameStatefulSet{}, f.defaultInformer)
}

func (f *gameStatefulSetInformer) Lister() v1alpha1.GameStatefulSetLister {
	return v1alpha1.NewGameStatefulSetLister(f.Informer().GetIndexer())
}
