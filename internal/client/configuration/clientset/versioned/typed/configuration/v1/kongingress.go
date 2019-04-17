/*
Copyright 2018 The Kong Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/kong/kubernetes-ingress-controller/internal/apis/configuration/v1"
	scheme "github.com/kong/kubernetes-ingress-controller/internal/client/configuration/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KongIngressesGetter has a method to return a KongIngressInterface.
// A group's client should implement this interface.
type KongIngressesGetter interface {
	KongIngresses(namespace string) KongIngressInterface
}

// KongIngressInterface has methods to work with KongIngress resources.
type KongIngressInterface interface {
	Create(*v1.KongIngress) (*v1.KongIngress, error)
	Update(*v1.KongIngress) (*v1.KongIngress, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KongIngress, error)
	List(opts metav1.ListOptions) (*v1.KongIngressList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KongIngress, err error)
	KongIngressExpansion
}

// kongIngresses implements KongIngressInterface
type kongIngresses struct {
	client rest.Interface
	ns     string
}

// newKongIngresses returns a KongIngresses
func newKongIngresses(c *ConfigurationV1Client, namespace string) *kongIngresses {
	return &kongIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kongIngress, and returns the corresponding kongIngress object, and an error if there is any.
func (c *kongIngresses) Get(name string, options metav1.GetOptions) (result *v1.KongIngress, err error) {
	result = &v1.KongIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KongIngresses that match those selectors.
func (c *kongIngresses) List(opts metav1.ListOptions) (result *v1.KongIngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KongIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kongIngresses.
func (c *kongIngresses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kongingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kongIngress and creates it.  Returns the server's representation of the kongIngress, and an error, if there is any.
func (c *kongIngresses) Create(kongIngress *v1.KongIngress) (result *v1.KongIngress, err error) {
	result = &v1.KongIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kongingresses").
		Body(kongIngress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kongIngress and updates it. Returns the server's representation of the kongIngress, and an error, if there is any.
func (c *kongIngresses) Update(kongIngress *v1.KongIngress) (result *v1.KongIngress, err error) {
	result = &v1.KongIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kongingresses").
		Name(kongIngress.Name).
		Body(kongIngress).
		Do().
		Into(result)
	return
}

// Delete takes name of the kongIngress and deletes it. Returns an error if one occurs.
func (c *kongIngresses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongingresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kongIngresses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongingresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kongIngress.
func (c *kongIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KongIngress, err error) {
	result = &v1.KongIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kongingresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
