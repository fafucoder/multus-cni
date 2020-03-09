/*
Copyright 2019 The Lijiao Authors.

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

	v1 "github.com/fafucoder/sriov-crd/pkg/apis/sriov/v1"
	scheme "github.com/fafucoder/sriov-crd/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SriovVFsGetter has a method to return a SriovVFInterface.
// A group's client should implement this interface.
type SriovVFsGetter interface {
	SriovVFs() SriovVFInterface
}

// SriovVFInterface has methods to work with SriovVF resources.
type SriovVFInterface interface {
	Create(*v1.SriovVF) (*v1.SriovVF, error)
	Update(*v1.SriovVF) (*v1.SriovVF, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.SriovVF, error)
	List(opts metav1.ListOptions) (*v1.SriovVFList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SriovVF, err error)
	SriovVFExpansion
}

// sriovVFs implements SriovVFInterface
type sriovVFs struct {
	client rest.Interface
}

// newSriovVFs returns a SriovVFs
func newSriovVFs(c *SriovV1Client) *sriovVFs {
	return &sriovVFs{
		client: c.RESTClient(),
	}
}

// Get takes name of the sriovVF, and returns the corresponding sriovVF object, and an error if there is any.
func (c *sriovVFs) Get(name string, options metav1.GetOptions) (result *v1.SriovVF, err error) {
	result = &v1.SriovVF{}
	err = c.client.Get().
		Resource("sriovvfs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SriovVFs that match those selectors.
func (c *sriovVFs) List(opts metav1.ListOptions) (result *v1.SriovVFList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SriovVFList{}
	err = c.client.Get().
		Resource("sriovvfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sriovVFs.
func (c *sriovVFs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sriovvfs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sriovVF and creates it.  Returns the server's representation of the sriovVF, and an error, if there is any.
func (c *sriovVFs) Create(sriovVF *v1.SriovVF) (result *v1.SriovVF, err error) {
	result = &v1.SriovVF{}
	err = c.client.Post().
		Resource("sriovvfs").
		Body(sriovVF).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sriovVF and updates it. Returns the server's representation of the sriovVF, and an error, if there is any.
func (c *sriovVFs) Update(sriovVF *v1.SriovVF) (result *v1.SriovVF, err error) {
	result = &v1.SriovVF{}
	err = c.client.Put().
		Resource("sriovvfs").
		Name(sriovVF.Name).
		Body(sriovVF).
		Do().
		Into(result)
	return
}

// Delete takes name of the sriovVF and deletes it. Returns an error if one occurs.
func (c *sriovVFs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sriovvfs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sriovVFs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sriovvfs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sriovVF.
func (c *sriovVFs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SriovVF, err error) {
	result = &v1.SriovVF{}
	err = c.client.Patch(pt).
		Resource("sriovvfs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}