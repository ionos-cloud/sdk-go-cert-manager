/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates  with IONOS services and your internal connected resources.   For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic. The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// CertificateReadList struct for CertificateReadList
type CertificateReadList struct {
	// ID of the list of Certificate resources.
	Id *string `json:"id"`
	// The type of the resource.
	Type *string `json:"type"`
	// The URL of the list of Certificate resources.
	Href *string `json:"href"`
	// The list of Certificate resources.
	Items *[]CertificateRead `json:"items,omitempty"`
	// The offset specified in the request (if none was specified, the default offset is 0).
	Offset *int32 `json:"offset"`
	// The limit specified in the request (if none was specified, use the endpoint's default pagination limit).
	Limit *int32 `json:"limit"`
	Links *Links `json:"_links"`
}

// NewCertificateReadList instantiates a new CertificateReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateReadList(id string, type_ string, href string, offset int32, limit int32, links Links) *CertificateReadList {
	this := CertificateReadList{}

	this.Id = &id
	this.Type = &type_
	this.Href = &href
	this.Offset = &offset
	this.Limit = &limit
	this.Links = &links

	return &this
}

// NewCertificateReadListWithDefaults instantiates a new CertificateReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateReadListWithDefaults() *CertificateReadList {
	this := CertificateReadList{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateReadList) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *CertificateReadList) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *CertificateReadList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateReadList) GetType() *string {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *CertificateReadList) SetType(v string) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *CertificateReadList) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CertificateReadList) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *CertificateReadList) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *CertificateReadList) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetItems returns the Items field value
// If the value is explicit nil, the zero value for []CertificateRead will be returned
func (o *CertificateReadList) GetItems() *[]CertificateRead {
	if o == nil {
		return nil
	}

	return o.Items

}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetItemsOk() (*[]CertificateRead, bool) {
	if o == nil {
		return nil, false
	}

	return o.Items, true
}

// SetItems sets field value
func (o *CertificateReadList) SetItems(v []CertificateRead) {

	o.Items = &v

}

// HasItems returns a boolean if a field has been set.
func (o *CertificateReadList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// GetOffset returns the Offset field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CertificateReadList) GetOffset() *int32 {
	if o == nil {
		return nil
	}

	return o.Offset

}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Offset, true
}

// SetOffset sets field value
func (o *CertificateReadList) SetOffset(v int32) {

	o.Offset = &v

}

// HasOffset returns a boolean if a field has been set.
func (o *CertificateReadList) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// GetLimit returns the Limit field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CertificateReadList) GetLimit() *int32 {
	if o == nil {
		return nil
	}

	return o.Limit

}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Limit, true
}

// SetLimit sets field value
func (o *CertificateReadList) SetLimit(v int32) {

	o.Limit = &v

}

// HasLimit returns a boolean if a field has been set.
func (o *CertificateReadList) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// GetLinks returns the Links field value
// If the value is explicit nil, the zero value for Links will be returned
func (o *CertificateReadList) GetLinks() *Links {
	if o == nil {
		return nil
	}

	return o.Links

}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateReadList) GetLinksOk() (*Links, bool) {
	if o == nil {
		return nil, false
	}

	return o.Links, true
}

// SetLinks sets field value
func (o *CertificateReadList) SetLinks(v Links) {

	o.Links = &v

}

// HasLinks returns a boolean if a field has been set.
func (o *CertificateReadList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

func (o CertificateReadList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	return json.Marshal(toSerialize)
}

type NullableCertificateReadList struct {
	value *CertificateReadList
	isSet bool
}

func (v NullableCertificateReadList) Get() *CertificateReadList {
	return v.value
}

func (v *NullableCertificateReadList) Set(val *CertificateReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateReadList(val *CertificateReadList) *NullableCertificateReadList {
	return &NullableCertificateReadList{value: val, isSet: true}
}

func (v NullableCertificateReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
