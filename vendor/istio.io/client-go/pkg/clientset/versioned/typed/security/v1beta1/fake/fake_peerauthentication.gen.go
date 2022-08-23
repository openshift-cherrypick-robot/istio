// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	securityv1beta1 "istio.io/client-go/pkg/applyconfiguration/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePeerAuthentications implements PeerAuthenticationInterface
type FakePeerAuthentications struct {
	Fake *FakeSecurityV1beta1
	ns   string
}

var peerauthenticationsResource = schema.GroupVersionResource{Group: "security.istio.io", Version: "v1beta1", Resource: "peerauthentications"}

var peerauthenticationsKind = schema.GroupVersionKind{Group: "security.istio.io", Version: "v1beta1", Kind: "PeerAuthentication"}

// Get takes name of the peerAuthentication, and returns the corresponding peerAuthentication object, and an error if there is any.
func (c *FakePeerAuthentications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.PeerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(peerauthenticationsResource, c.ns, name), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}

// List takes label and field selectors, and returns the list of PeerAuthentications that match those selectors.
func (c *FakePeerAuthentications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PeerAuthenticationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(peerauthenticationsResource, peerauthenticationsKind, c.ns, opts), &v1beta1.PeerAuthenticationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PeerAuthenticationList{ListMeta: obj.(*v1beta1.PeerAuthenticationList).ListMeta}
	for _, item := range obj.(*v1beta1.PeerAuthenticationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested peerAuthentications.
func (c *FakePeerAuthentications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(peerauthenticationsResource, c.ns, opts))

}

// Create takes the representation of a peerAuthentication and creates it.  Returns the server's representation of the peerAuthentication, and an error, if there is any.
func (c *FakePeerAuthentications) Create(ctx context.Context, peerAuthentication *v1beta1.PeerAuthentication, opts v1.CreateOptions) (result *v1beta1.PeerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(peerauthenticationsResource, c.ns, peerAuthentication), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}

// Update takes the representation of a peerAuthentication and updates it. Returns the server's representation of the peerAuthentication, and an error, if there is any.
func (c *FakePeerAuthentications) Update(ctx context.Context, peerAuthentication *v1beta1.PeerAuthentication, opts v1.UpdateOptions) (result *v1beta1.PeerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(peerauthenticationsResource, c.ns, peerAuthentication), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePeerAuthentications) UpdateStatus(ctx context.Context, peerAuthentication *v1beta1.PeerAuthentication, opts v1.UpdateOptions) (*v1beta1.PeerAuthentication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(peerauthenticationsResource, "status", c.ns, peerAuthentication), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}

// Delete takes name of the peerAuthentication and deletes it. Returns an error if one occurs.
func (c *FakePeerAuthentications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(peerauthenticationsResource, c.ns, name, opts), &v1beta1.PeerAuthentication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePeerAuthentications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(peerauthenticationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PeerAuthenticationList{})
	return err
}

// Patch applies the patch and returns the patched peerAuthentication.
func (c *FakePeerAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.PeerAuthentication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peerauthenticationsResource, c.ns, name, pt, data, subresources...), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied peerAuthentication.
func (c *FakePeerAuthentications) Apply(ctx context.Context, peerAuthentication *securityv1beta1.PeerAuthenticationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.PeerAuthentication, err error) {
	if peerAuthentication == nil {
		return nil, fmt.Errorf("peerAuthentication provided to Apply must not be nil")
	}
	data, err := json.Marshal(peerAuthentication)
	if err != nil {
		return nil, err
	}
	name := peerAuthentication.Name
	if name == nil {
		return nil, fmt.Errorf("peerAuthentication.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peerauthenticationsResource, c.ns, *name, types.ApplyPatchType, data), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePeerAuthentications) ApplyStatus(ctx context.Context, peerAuthentication *securityv1beta1.PeerAuthenticationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.PeerAuthentication, err error) {
	if peerAuthentication == nil {
		return nil, fmt.Errorf("peerAuthentication provided to Apply must not be nil")
	}
	data, err := json.Marshal(peerAuthentication)
	if err != nil {
		return nil, err
	}
	name := peerAuthentication.Name
	if name == nil {
		return nil, fmt.Errorf("peerAuthentication.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peerauthenticationsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1beta1.PeerAuthentication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.PeerAuthentication), err
}
