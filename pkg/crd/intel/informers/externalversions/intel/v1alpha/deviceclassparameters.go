/*
 * Copyright (c) 2022, Intel Corporation.  All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha

import (
	"context"
	time "time"

	versioned "github.com/intel/intel-resource-drivers-for-kubernetes/pkg/crd/intel/clientset/versioned"
	internalinterfaces "github.com/intel/intel-resource-drivers-for-kubernetes/pkg/crd/intel/informers/externalversions/internalinterfaces"
	v1alpha "github.com/intel/intel-resource-drivers-for-kubernetes/pkg/crd/intel/listers/intel/v1alpha"
	intelv1alpha "github.com/intel/intel-resource-drivers-for-kubernetes/pkg/crd/intel/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DeviceClassParametersInformer provides access to a shared informer and lister for
// DeviceClassParameters.
type DeviceClassParametersInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha.DeviceClassParametersLister
}

type deviceClassParametersInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDeviceClassParametersInformer constructs a new informer for DeviceClassParameters type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeviceClassParametersInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeviceClassParametersInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDeviceClassParametersInformer constructs a new informer for DeviceClassParameters type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeviceClassParametersInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GpuV1alpha().DeviceClassParameters().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GpuV1alpha().DeviceClassParameters().Watch(context.TODO(), options)
			},
		},
		&intelv1alpha.DeviceClassParameters{},
		resyncPeriod,
		indexers,
	)
}

func (f *deviceClassParametersInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeviceClassParametersInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deviceClassParametersInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&intelv1alpha.DeviceClassParameters{}, f.defaultInformer)
}

func (f *deviceClassParametersInformer) Lister() v1alpha.DeviceClassParametersLister {
	return v1alpha.NewDeviceClassParametersLister(f.Informer().GetIndexer())
}