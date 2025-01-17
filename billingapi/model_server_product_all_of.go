/*
Billing API

Automate your infrastructure billing with the Bare Metal Cloud Billing API. Reserve your server instances to ensure guaranteed resource availability for 12, 24, and 36 months. Retrieve your server’s rated usage for a given period and enable or disable auto-renewals.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/phoenixnap-bare-metal-cloud-billing-models' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/billing/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package billingapi

import (
	"encoding/json"
)

// ServerProductAllOf Server product.
type ServerProductAllOf struct {
	Metadata ServerProductMetadata `json:"metadata"`
}

// NewServerProductAllOf instantiates a new ServerProductAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProductAllOf(metadata ServerProductMetadata) *ServerProductAllOf {
	this := ServerProductAllOf{}
	this.Metadata = metadata
	return &this
}

// NewServerProductAllOfWithDefaults instantiates a new ServerProductAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProductAllOfWithDefaults() *ServerProductAllOf {
	this := ServerProductAllOf{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *ServerProductAllOf) GetMetadata() ServerProductMetadata {
	if o == nil {
		var ret ServerProductMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ServerProductAllOf) GetMetadataOk() (*ServerProductMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ServerProductAllOf) SetMetadata(v ServerProductMetadata) {
	o.Metadata = v
}

func (o ServerProductAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableServerProductAllOf struct {
	value *ServerProductAllOf
	isSet bool
}

func (v NullableServerProductAllOf) Get() *ServerProductAllOf {
	return v.value
}

func (v *NullableServerProductAllOf) Set(val *ServerProductAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProductAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProductAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProductAllOf(val *ServerProductAllOf) *NullableServerProductAllOf {
	return &NullableServerProductAllOf{value: val, isSet: true}
}

func (v NullableServerProductAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProductAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
