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

package v1alpha1

import (
	"time"
	"context"

	v1alpha1 "github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	scheme "github.com/litmuschaos/chaos-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ChaosResultsGetter has a method to return a ChaosResultInterface.
// A group's client should implement this interface.
type ChaosResultsGetter interface {
	ChaosResults(namespace string) ChaosResultInterface
}

// ChaosResultInterface has methods to work with ChaosResult resources.
type ChaosResultInterface interface {
	Create(ctx context.Context, chaosResult *v1alpha1.ChaosResult) (*v1alpha1.ChaosResult, error)
	Update(ctx context.Context, chaosResult *v1alpha1.ChaosResult) (*v1alpha1.ChaosResult, error)
	UpdateStatus(ctx context.Context, chaosResult *v1alpha1.ChaosResult) (*v1alpha1.ChaosResult, error)
	Delete(ctx context.Context,name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context,options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context,name string, options v1.GetOptions) (*v1alpha1.ChaosResult, error)
	List(ctx context.Context,opts v1.ListOptions) (*v1alpha1.ChaosResultList, error)
	Watch(ctx context.Context,opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context,name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChaosResult, err error)
	ChaosResultExpansion
}

// chaosResults implements ChaosResultInterface
type chaosResults struct {
	client rest.Interface
	ns     string
}

// newChaosResults returns a ChaosResults
func newChaosResults(c *LitmuschaosV1alpha1Client, namespace string) *chaosResults {
	return &chaosResults{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chaosResult, and returns the corresponding chaosResult object, and an error if there is any.
func (c *chaosResults) Get(ctx context.Context,name string, options v1.GetOptions) (result *v1alpha1.ChaosResult, err error) {
	result = &v1alpha1.ChaosResult{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chaosresults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChaosResults that match those selectors.
func (c *chaosResults) List(ctx context.Context,opts v1.ListOptions) (result *v1alpha1.ChaosResultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ChaosResultList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chaosresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chaosResults.
func (c *chaosResults) Watch(ctx context.Context,opts v1.ListOptions) (result watch.Interface,err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	result,err= c.client.Get().
		Namespace(c.ns).
		Resource("chaosresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
	return

}

// Create takes the representation of a chaosResult and creates it.  Returns the server's representation of the chaosResult, and an error, if there is any.
func (c *chaosResults) Create(ctx context.Context,chaosResult *v1alpha1.ChaosResult) (result *v1alpha1.ChaosResult, err error) {
	result = &v1alpha1.ChaosResult{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chaosresults").
		Body(chaosResult).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a chaosResult and updates it. Returns the server's representation of the chaosResult, and an error, if there is any.
func (c *chaosResults) Update(ctx context.Context,chaosResult *v1alpha1.ChaosResult) (result *v1alpha1.ChaosResult, err error) {
	result = &v1alpha1.ChaosResult{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chaosresults").
		Name(chaosResult.Name).
		Body(chaosResult).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *chaosResults) UpdateStatus(ctx context.Context,chaosResult *v1alpha1.ChaosResult) (result *v1alpha1.ChaosResult, err error) {
	result = &v1alpha1.ChaosResult{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chaosresults").
		Name(chaosResult.Name).
		SubResource("status").
		Body(chaosResult).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the chaosResult and deletes it. Returns an error if one occurs.
func (c *chaosResults) Delete(ctx context.Context,name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chaosresults").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chaosResults) DeleteCollection(ctx context.Context,options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chaosresults").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched chaosResult.
func (c *chaosResults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChaosResult, err error) {
	result = &v1alpha1.ChaosResult{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chaosresults").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
