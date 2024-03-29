/*
 * Certificate Manager Service API
 *
 * Using the Certificate Manager Service, you can conveniently provision and manage SSL certificates with IONOS services and your internal connected resources. For the [Application Load Balancer](https://api.ionos.com/docs/cloud/v6/#Application-Load-Balancers-get-datacenters-datacenterId-applicationloadbalancers), you usually need a certificate to encrypt your HTTPS traffic.  The service provides the basic functions of uploading and deleting your certificates for this purpose.
 *
 * API version: 1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ApiInfoDto The information about the API.
type ApiInfoDto struct {
	// The API entry point.
	Href *string `json:"href,omitempty"`
	// The API name.
	Name *string `json:"name,omitempty"`
	// The API version.
	Version *string `json:"version,omitempty"`
}

// NewApiInfoDto instantiates a new ApiInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInfoDto() *ApiInfoDto {
	this := ApiInfoDto{}

	return &this
}

// NewApiInfoDtoWithDefaults instantiates a new ApiInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInfoDtoWithDefaults() *ApiInfoDto {
	this := ApiInfoDto{}
	return &this
}

// GetHref returns the Href field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiInfoDto) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiInfoDto) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *ApiInfoDto) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *ApiInfoDto) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiInfoDto) GetName() *string {
	if o == nil {
		return nil
	}

	return o.Name

}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiInfoDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Name, true
}

// SetName sets field value
func (o *ApiInfoDto) SetName(v string) {

	o.Name = &v

}

// HasName returns a boolean if a field has been set.
func (o *ApiInfoDto) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ApiInfoDto) GetVersion() *string {
	if o == nil {
		return nil
	}

	return o.Version

}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiInfoDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Version, true
}

// SetVersion sets field value
func (o *ApiInfoDto) SetVersion(v string) {

	o.Version = &v

}

// HasVersion returns a boolean if a field has been set.
func (o *ApiInfoDto) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

func (o ApiInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	return json.Marshal(toSerialize)
}

type NullableApiInfoDto struct {
	value *ApiInfoDto
	isSet bool
}

func (v NullableApiInfoDto) Get() *ApiInfoDto {
	return v.value
}

func (v *NullableApiInfoDto) Set(val *ApiInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInfoDto(val *ApiInfoDto) *NullableApiInfoDto {
	return &NullableApiInfoDto{value: val, isSet: true}
}

func (v NullableApiInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
