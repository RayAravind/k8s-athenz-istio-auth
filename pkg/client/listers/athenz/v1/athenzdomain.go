// Copyright 2019, VerizonMedia Inc.
// Licensed under the terms of the 3-Clause BSD license. See LICENSE file in github.com/yahoo/k8s-athenz-istio-auth
// for terms.
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/yahoo/k8s-athenz-istio-auth/pkg/apis/athenz/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AthenzDomainLister helps list AthenzDomains.
type AthenzDomainLister interface {
	// List lists all AthenzDomains in the indexer.
	List(selector labels.Selector) (ret []*v1.AthenzDomain, err error)
	// AthenzDomains returns an object that can list and get AthenzDomains.
	AthenzDomains(namespace string) AthenzDomainNamespaceLister
	AthenzDomainListerExpansion
}

// athenzDomainLister implements the AthenzDomainLister interface.
type athenzDomainLister struct {
	indexer cache.Indexer
}

// NewAthenzDomainLister returns a new AthenzDomainLister.
func NewAthenzDomainLister(indexer cache.Indexer) AthenzDomainLister {
	return &athenzDomainLister{indexer: indexer}
}

// List lists all AthenzDomains in the indexer.
func (s *athenzDomainLister) List(selector labels.Selector) (ret []*v1.AthenzDomain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AthenzDomain))
	})
	return ret, err
}

// AthenzDomains returns an object that can list and get AthenzDomains.
func (s *athenzDomainLister) AthenzDomains(namespace string) AthenzDomainNamespaceLister {
	return athenzDomainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AthenzDomainNamespaceLister helps list and get AthenzDomains.
type AthenzDomainNamespaceLister interface {
	// List lists all AthenzDomains in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AthenzDomain, err error)
	// Get retrieves the AthenzDomain from the indexer for a given namespace and name.
	Get(name string) (*v1.AthenzDomain, error)
	AthenzDomainNamespaceListerExpansion
}

// athenzDomainNamespaceLister implements the AthenzDomainNamespaceLister
// interface.
type athenzDomainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AthenzDomains in the indexer for a given namespace.
func (s athenzDomainNamespaceLister) List(selector labels.Selector) (ret []*v1.AthenzDomain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AthenzDomain))
	})
	return ret, err
}

// Get retrieves the AthenzDomain from the indexer for a given namespace and name.
func (s athenzDomainNamespaceLister) Get(name string) (*v1.AthenzDomain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("athenzdomain"), name)
	}
	return obj.(*v1.AthenzDomain), nil
}
