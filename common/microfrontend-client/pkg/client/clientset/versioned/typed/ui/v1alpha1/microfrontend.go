// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kyma-project/kyma/common/microfrontend-client/pkg/apis/ui/v1alpha1"
	scheme "github.com/kyma-project/kyma/common/microfrontend-client/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MicroFrontendsGetter has a method to return a MicroFrontendInterface.
// A group's client should implement this interface.
type MicroFrontendsGetter interface {
	MicroFrontends(namespace string) MicroFrontendInterface
}

// MicroFrontendInterface has methods to work with MicroFrontend resources.
type MicroFrontendInterface interface {
	Create(ctx context.Context, microFrontend *v1alpha1.MicroFrontend, opts v1.CreateOptions) (*v1alpha1.MicroFrontend, error)
	Update(ctx context.Context, microFrontend *v1alpha1.MicroFrontend, opts v1.UpdateOptions) (*v1alpha1.MicroFrontend, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.MicroFrontend, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.MicroFrontendList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MicroFrontend, err error)
	MicroFrontendExpansion
}

// microFrontends implements MicroFrontendInterface
type microFrontends struct {
	client rest.Interface
	ns     string
}

// newMicroFrontends returns a MicroFrontends
func newMicroFrontends(c *UiV1alpha1Client, namespace string) *microFrontends {
	return &microFrontends{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the microFrontend, and returns the corresponding microFrontend object, and an error if there is any.
func (c *microFrontends) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MicroFrontend, err error) {
	result = &v1alpha1.MicroFrontend{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("microfrontends").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MicroFrontends that match those selectors.
func (c *microFrontends) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MicroFrontendList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MicroFrontendList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("microfrontends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested microFrontends.
func (c *microFrontends) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("microfrontends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a microFrontend and creates it.  Returns the server's representation of the microFrontend, and an error, if there is any.
func (c *microFrontends) Create(ctx context.Context, microFrontend *v1alpha1.MicroFrontend, opts v1.CreateOptions) (result *v1alpha1.MicroFrontend, err error) {
	result = &v1alpha1.MicroFrontend{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("microfrontends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(microFrontend).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a microFrontend and updates it. Returns the server's representation of the microFrontend, and an error, if there is any.
func (c *microFrontends) Update(ctx context.Context, microFrontend *v1alpha1.MicroFrontend, opts v1.UpdateOptions) (result *v1alpha1.MicroFrontend, err error) {
	result = &v1alpha1.MicroFrontend{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("microfrontends").
		Name(microFrontend.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(microFrontend).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the microFrontend and deletes it. Returns an error if one occurs.
func (c *microFrontends) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("microfrontends").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *microFrontends) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("microfrontends").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched microFrontend.
func (c *microFrontends) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MicroFrontend, err error) {
	result = &v1alpha1.MicroFrontend{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("microfrontends").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}