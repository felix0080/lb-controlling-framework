/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "git.code.oa.com/k8s/lb-controlling-framework/pkg/apis/lbcf.tke.cloud.tencent.com/v1beta1"
	scheme "git.code.oa.com/k8s/lb-controlling-framework/pkg/client-go/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackendGroupsGetter has a method to return a BackendGroupInterface.
// A group's client should implement this interface.
type BackendGroupsGetter interface {
	BackendGroups(namespace string) BackendGroupInterface
}

// BackendGroupInterface has methods to work with BackendGroup resources.
type BackendGroupInterface interface {
	Create(*v1beta1.BackendGroup) (*v1beta1.BackendGroup, error)
	Update(*v1beta1.BackendGroup) (*v1beta1.BackendGroup, error)
	UpdateStatus(*v1beta1.BackendGroup) (*v1beta1.BackendGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.BackendGroup, error)
	List(opts v1.ListOptions) (*v1beta1.BackendGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackendGroup, err error)
	BackendGroupExpansion
}

// backendGroups implements BackendGroupInterface
type backendGroups struct {
	client rest.Interface
	ns     string
}

// newBackendGroups returns a BackendGroups
func newBackendGroups(c *LbcfV1beta1Client, namespace string) *backendGroups {
	return &backendGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backendGroup, and returns the corresponding backendGroup object, and an error if there is any.
func (c *backendGroups) Get(name string, options v1.GetOptions) (result *v1beta1.BackendGroup, err error) {
	result = &v1beta1.BackendGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackendGroups that match those selectors.
func (c *backendGroups) List(opts v1.ListOptions) (result *v1beta1.BackendGroupList, err error) {
	result = &v1beta1.BackendGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backendgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backendGroups.
func (c *backendGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backendgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a backendGroup and creates it.  Returns the server's representation of the backendGroup, and an error, if there is any.
func (c *backendGroups) Create(backendGroup *v1beta1.BackendGroup) (result *v1beta1.BackendGroup, err error) {
	result = &v1beta1.BackendGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backendgroups").
		Body(backendGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backendGroup and updates it. Returns the server's representation of the backendGroup, and an error, if there is any.
func (c *backendGroups) Update(backendGroup *v1beta1.BackendGroup) (result *v1beta1.BackendGroup, err error) {
	result = &v1beta1.BackendGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendgroups").
		Name(backendGroup.Name).
		Body(backendGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *backendGroups) UpdateStatus(backendGroup *v1beta1.BackendGroup) (result *v1beta1.BackendGroup, err error) {
	result = &v1beta1.BackendGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backendgroups").
		Name(backendGroup.Name).
		SubResource("status").
		Body(backendGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the backendGroup and deletes it. Returns an error if one occurs.
func (c *backendGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backendGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backendgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backendGroup.
func (c *backendGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackendGroup, err error) {
	result = &v1beta1.BackendGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backendgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
