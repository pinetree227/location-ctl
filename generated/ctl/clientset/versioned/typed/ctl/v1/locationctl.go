/*
Copyright 2023.

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
	"context"
	"time"

	v1 "github.com/pinetree227/location-ctl/api/ctl/v1"
	scheme "github.com/pinetree227/location-ctl/generated/ctl/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LocationCtlsGetter has a method to return a LocationCtlInterface.
// A group's client should implement this interface.
type LocationCtlsGetter interface {
	LocationCtls(namespace string) LocationCtlInterface
}

// LocationCtlInterface has methods to work with LocationCtl resources.
type LocationCtlInterface interface {
	Create(ctx context.Context, locationCtl *v1.LocationCtl, opts metav1.CreateOptions) (*v1.LocationCtl, error)
	Update(ctx context.Context, locationCtl *v1.LocationCtl, opts metav1.UpdateOptions) (*v1.LocationCtl, error)
	UpdateStatus(ctx context.Context, locationCtl *v1.LocationCtl, opts metav1.UpdateOptions) (*v1.LocationCtl, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LocationCtl, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LocationCtlList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocationCtl, err error)
	LocationCtlExpansion
}

// locationCtls implements LocationCtlInterface
type locationCtls struct {
	client rest.Interface
	ns     string
}

// newLocationCtls returns a LocationCtls
func newLocationCtls(c *CtlV1Client, namespace string) *locationCtls {
	return &locationCtls{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the locationCtl, and returns the corresponding locationCtl object, and an error if there is any.
func (c *locationCtls) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LocationCtl, err error) {
	result = &v1.LocationCtl{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locationctls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocationCtls that match those selectors.
func (c *locationCtls) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LocationCtlList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LocationCtlList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locationctls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested locationCtls.
func (c *locationCtls) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("locationctls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a locationCtl and creates it.  Returns the server's representation of the locationCtl, and an error, if there is any.
func (c *locationCtls) Create(ctx context.Context, locationCtl *v1.LocationCtl, opts metav1.CreateOptions) (result *v1.LocationCtl, err error) {
	result = &v1.LocationCtl{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("locationctls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(locationCtl).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a locationCtl and updates it. Returns the server's representation of the locationCtl, and an error, if there is any.
func (c *locationCtls) Update(ctx context.Context, locationCtl *v1.LocationCtl, opts metav1.UpdateOptions) (result *v1.LocationCtl, err error) {
	result = &v1.LocationCtl{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locationctls").
		Name(locationCtl.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(locationCtl).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *locationCtls) UpdateStatus(ctx context.Context, locationCtl *v1.LocationCtl, opts metav1.UpdateOptions) (result *v1.LocationCtl, err error) {
	result = &v1.LocationCtl{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locationctls").
		Name(locationCtl.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(locationCtl).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the locationCtl and deletes it. Returns an error if one occurs.
func (c *locationCtls) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locationctls").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *locationCtls) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locationctls").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched locationCtl.
func (c *locationCtls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocationCtl, err error) {
	result = &v1.LocationCtl{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("locationctls").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
