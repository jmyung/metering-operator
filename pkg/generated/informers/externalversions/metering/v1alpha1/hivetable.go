// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	meteringv1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	versioned "github.com/operator-framework/operator-metering/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/operator-framework/operator-metering/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/generated/listers/metering/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HiveTableInformer provides access to a shared informer and lister for
// HiveTables.
type HiveTableInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HiveTableLister
}

type hiveTableInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHiveTableInformer constructs a new informer for HiveTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHiveTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHiveTableInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHiveTableInformer constructs a new informer for HiveTable type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHiveTableInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1alpha1().HiveTables(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeteringV1alpha1().HiveTables(namespace).Watch(options)
			},
		},
		&meteringv1alpha1.HiveTable{},
		resyncPeriod,
		indexers,
	)
}

func (f *hiveTableInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHiveTableInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hiveTableInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meteringv1alpha1.HiveTable{}, f.defaultInformer)
}

func (f *hiveTableInformer) Lister() v1alpha1.HiveTableLister {
	return v1alpha1.NewHiveTableLister(f.Informer().GetIndexer())
}