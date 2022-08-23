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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// FakeTCPRoutes implements TCPRouteInterface
type FakeTCPRoutes struct {
	Fake *FakeGatewayV1alpha2
	ns   string
}

var tcproutesResource = schema.GroupVersionResource{Group: "gateway.networking.k8s.io", Version: "v1alpha2", Resource: "tcproutes"}

var tcproutesKind = schema.GroupVersionKind{Group: "gateway.networking.k8s.io", Version: "v1alpha2", Kind: "TCPRoute"}

// Get takes name of the tCPRoute, and returns the corresponding tCPRoute object, and an error if there is any.
func (c *FakeTCPRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.TCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tcproutesResource, c.ns, name), &v1alpha2.TCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TCPRoute), err
}

// List takes label and field selectors, and returns the list of TCPRoutes that match those selectors.
func (c *FakeTCPRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.TCPRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tcproutesResource, tcproutesKind, c.ns, opts), &v1alpha2.TCPRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.TCPRouteList{ListMeta: obj.(*v1alpha2.TCPRouteList).ListMeta}
	for _, item := range obj.(*v1alpha2.TCPRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tCPRoutes.
func (c *FakeTCPRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tcproutesResource, c.ns, opts))

}

// Create takes the representation of a tCPRoute and creates it.  Returns the server's representation of the tCPRoute, and an error, if there is any.
func (c *FakeTCPRoutes) Create(ctx context.Context, tCPRoute *v1alpha2.TCPRoute, opts v1.CreateOptions) (result *v1alpha2.TCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tcproutesResource, c.ns, tCPRoute), &v1alpha2.TCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TCPRoute), err
}

// Update takes the representation of a tCPRoute and updates it. Returns the server's representation of the tCPRoute, and an error, if there is any.
func (c *FakeTCPRoutes) Update(ctx context.Context, tCPRoute *v1alpha2.TCPRoute, opts v1.UpdateOptions) (result *v1alpha2.TCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tcproutesResource, c.ns, tCPRoute), &v1alpha2.TCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TCPRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTCPRoutes) UpdateStatus(ctx context.Context, tCPRoute *v1alpha2.TCPRoute, opts v1.UpdateOptions) (*v1alpha2.TCPRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tcproutesResource, "status", c.ns, tCPRoute), &v1alpha2.TCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TCPRoute), err
}

// Delete takes name of the tCPRoute and deletes it. Returns an error if one occurs.
func (c *FakeTCPRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tcproutesResource, c.ns, name), &v1alpha2.TCPRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTCPRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tcproutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.TCPRouteList{})
	return err
}

// Patch applies the patch and returns the patched tCPRoute.
func (c *FakeTCPRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.TCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tcproutesResource, c.ns, name, pt, data, subresources...), &v1alpha2.TCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.TCPRoute), err
}
