/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/multiclusterdns/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceDNSRecords implements ServiceDNSRecordInterface
type FakeServiceDNSRecords struct {
	Fake *FakeMulticlusterdnsV1alpha1
	ns   string
}

var servicednsrecordsResource = schema.GroupVersionResource{Group: "multiclusterdns.federation.k8s.io", Version: "v1alpha1", Resource: "servicednsrecords"}

var servicednsrecordsKind = schema.GroupVersionKind{Group: "multiclusterdns.federation.k8s.io", Version: "v1alpha1", Kind: "ServiceDNSRecord"}

// Get takes name of the serviceDNSRecord, and returns the corresponding serviceDNSRecord object, and an error if there is any.
func (c *FakeServiceDNSRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicednsrecordsResource, c.ns, name), &v1alpha1.ServiceDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDNSRecord), err
}

// List takes label and field selectors, and returns the list of ServiceDNSRecords that match those selectors.
func (c *FakeServiceDNSRecords) List(opts v1.ListOptions) (result *v1alpha1.ServiceDNSRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicednsrecordsResource, servicednsrecordsKind, c.ns, opts), &v1alpha1.ServiceDNSRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceDNSRecordList{ListMeta: obj.(*v1alpha1.ServiceDNSRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceDNSRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceDNSRecords.
func (c *FakeServiceDNSRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicednsrecordsResource, c.ns, opts))

}

// Create takes the representation of a serviceDNSRecord and creates it.  Returns the server's representation of the serviceDNSRecord, and an error, if there is any.
func (c *FakeServiceDNSRecords) Create(serviceDNSRecord *v1alpha1.ServiceDNSRecord) (result *v1alpha1.ServiceDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicednsrecordsResource, c.ns, serviceDNSRecord), &v1alpha1.ServiceDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDNSRecord), err
}

// Update takes the representation of a serviceDNSRecord and updates it. Returns the server's representation of the serviceDNSRecord, and an error, if there is any.
func (c *FakeServiceDNSRecords) Update(serviceDNSRecord *v1alpha1.ServiceDNSRecord) (result *v1alpha1.ServiceDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicednsrecordsResource, c.ns, serviceDNSRecord), &v1alpha1.ServiceDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDNSRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceDNSRecords) UpdateStatus(serviceDNSRecord *v1alpha1.ServiceDNSRecord) (*v1alpha1.ServiceDNSRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicednsrecordsResource, "status", c.ns, serviceDNSRecord), &v1alpha1.ServiceDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDNSRecord), err
}

// Delete takes name of the serviceDNSRecord and deletes it. Returns an error if one occurs.
func (c *FakeServiceDNSRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicednsrecordsResource, c.ns, name), &v1alpha1.ServiceDNSRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceDNSRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicednsrecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceDNSRecordList{})
	return err
}

// Patch applies the patch and returns the patched serviceDNSRecord.
func (c *FakeServiceDNSRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDNSRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicednsrecordsResource, c.ns, name, data, subresources...), &v1alpha1.ServiceDNSRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDNSRecord), err
}