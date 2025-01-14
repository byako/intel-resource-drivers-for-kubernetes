/*
 * Copyright (c) 2023, Intel Corporation.  All Rights Reserved.
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

// GpuAllocationStateInformer provides access to a shared informer and lister for
// GpuAllocationStates.
type GpuAllocationStateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha.GpuAllocationStateLister
}

type gpuAllocationStateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGpuAllocationStateInformer constructs a new informer for GpuAllocationState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGpuAllocationStateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGpuAllocationStateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGpuAllocationStateInformer constructs a new informer for GpuAllocationState type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGpuAllocationStateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GpuV1alpha().GpuAllocationStates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GpuV1alpha().GpuAllocationStates(namespace).Watch(context.TODO(), options)
			},
		},
		&intelv1alpha.GpuAllocationState{},
		resyncPeriod,
		indexers,
	)
}

func (f *gpuAllocationStateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGpuAllocationStateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gpuAllocationStateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&intelv1alpha.GpuAllocationState{}, f.defaultInformer)
}

func (f *gpuAllocationStateInformer) Lister() v1alpha.GpuAllocationStateLister {
	return v1alpha.NewGpuAllocationStateLister(f.Informer().GetIndexer())
}
