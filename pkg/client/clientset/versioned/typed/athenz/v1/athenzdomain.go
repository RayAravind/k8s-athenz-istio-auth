// Copyright 2019, Verizon Media Inc.
// Licensed under the terms of the 3-Clause BSD license. See LICENSE file in
// github.com/yahoo/k8s-athenz-istio-auth for terms.
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/yahoo/k8s-athenz-istio-auth/pkg/apis/athenz/v1"
	scheme "github.com/yahoo/k8s-athenz-istio-auth/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AthenzDomainsGetter has a method to return a AthenzDomainInterface.
// A group's client should implement this interface.
type AthenzDomainsGetter interface {
	AthenzDomains(namespace string) AthenzDomainInterface
}

// AthenzDomainInterface has methods to work with AthenzDomain resources.
type AthenzDomainInterface interface {
	Create(*v1.AthenzDomain) (*v1.AthenzDomain, error)
	Update(*v1.AthenzDomain) (*v1.AthenzDomain, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AthenzDomain, error)
	List(opts meta_v1.ListOptions) (*v1.AthenzDomainList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AthenzDomain, err error)
	AthenzDomainExpansion
}

// athenzDomains implements AthenzDomainInterface
type athenzDomains struct {
	client rest.Interface
	ns     string
}

// newAthenzDomains returns a AthenzDomains
func newAthenzDomains(c *AthenzV1Client, namespace string) *athenzDomains {
	return &athenzDomains{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the athenzDomain, and returns the corresponding athenzDomain object, and an error if there is any.
func (c *athenzDomains) Get(name string, options meta_v1.GetOptions) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("athenzdomains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AthenzDomains that match those selectors.
func (c *athenzDomains) List(opts meta_v1.ListOptions) (result *v1.AthenzDomainList, err error) {
	result = &v1.AthenzDomainList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("athenzdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested athenzDomains.
func (c *athenzDomains) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("athenzdomains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a athenzDomain and creates it.  Returns the server's representation of the athenzDomain, and an error, if there is any.
func (c *athenzDomains) Create(athenzDomain *v1.AthenzDomain) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("athenzdomains").
		Body(athenzDomain).
		Do().
		Into(result)
	return
}

// Update takes the representation of a athenzDomain and updates it. Returns the server's representation of the athenzDomain, and an error, if there is any.
func (c *athenzDomains) Update(athenzDomain *v1.AthenzDomain) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("athenzdomains").
		Name(athenzDomain.Name).
		Body(athenzDomain).
		Do().
		Into(result)
	return
}

// Delete takes name of the athenzDomain and deletes it. Returns an error if one occurs.
func (c *athenzDomains) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("athenzdomains").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *athenzDomains) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("athenzdomains").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched athenzDomain.
func (c *athenzDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AthenzDomain, err error) {
	result = &v1.AthenzDomain{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("athenzdomains").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
