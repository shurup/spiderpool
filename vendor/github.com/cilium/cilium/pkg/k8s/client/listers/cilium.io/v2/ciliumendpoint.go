// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumEndpointLister helps list CiliumEndpoints.
// All objects returned here must be treated as read-only.
type CiliumEndpointLister interface {
	// List lists all CiliumEndpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumEndpoint, err error)
	// CiliumEndpoints returns an object that can list and get CiliumEndpoints.
	CiliumEndpoints(namespace string) CiliumEndpointNamespaceLister
	CiliumEndpointListerExpansion
}

// ciliumEndpointLister implements the CiliumEndpointLister interface.
type ciliumEndpointLister struct {
	indexer cache.Indexer
}

// NewCiliumEndpointLister returns a new CiliumEndpointLister.
func NewCiliumEndpointLister(indexer cache.Indexer) CiliumEndpointLister {
	return &ciliumEndpointLister{indexer: indexer}
}

// List lists all CiliumEndpoints in the indexer.
func (s *ciliumEndpointLister) List(selector labels.Selector) (ret []*v2.CiliumEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumEndpoint))
	})
	return ret, err
}

// CiliumEndpoints returns an object that can list and get CiliumEndpoints.
func (s *ciliumEndpointLister) CiliumEndpoints(namespace string) CiliumEndpointNamespaceLister {
	return ciliumEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CiliumEndpointNamespaceLister helps list and get CiliumEndpoints.
// All objects returned here must be treated as read-only.
type CiliumEndpointNamespaceLister interface {
	// List lists all CiliumEndpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumEndpoint, err error)
	// Get retrieves the CiliumEndpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumEndpoint, error)
	CiliumEndpointNamespaceListerExpansion
}

// ciliumEndpointNamespaceLister implements the CiliumEndpointNamespaceLister
// interface.
type ciliumEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CiliumEndpoints in the indexer for a given namespace.
func (s ciliumEndpointNamespaceLister) List(selector labels.Selector) (ret []*v2.CiliumEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumEndpoint))
	})
	return ret, err
}

// Get retrieves the CiliumEndpoint from the indexer for a given namespace and name.
func (s ciliumEndpointNamespaceLister) Get(name string) (*v2.CiliumEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumendpoint"), name)
	}
	return obj.(*v2.CiliumEndpoint), nil
}