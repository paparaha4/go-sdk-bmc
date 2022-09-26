/*
Network Storage API

Create, list, edit, and delete storage networks with the Network Storage API. Use storage networks to expand storage capacity on a private network. <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='' target='_blank'>here</a> </span> <br> <b>All URLs are relative to (https://api.phoenixnap.com/network-storage/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkstorageapi

import (
	"encoding/json"
)

// Permissions Permissions for a volume.
type Permissions struct {
	Nfs *NfsPermissions `json:"nfs,omitempty"`
}

// NewPermissions instantiates a new Permissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissions() *Permissions {
	this := Permissions{}
	return &this
}

// NewPermissionsWithDefaults instantiates a new Permissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsWithDefaults() *Permissions {
	this := Permissions{}
	return &this
}

// GetNfs returns the Nfs field value if set, zero value otherwise.
func (o *Permissions) GetNfs() NfsPermissions {
	if o == nil || o.Nfs == nil {
		var ret NfsPermissions
		return ret
	}
	return *o.Nfs
}

// GetNfsOk returns a tuple with the Nfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permissions) GetNfsOk() (*NfsPermissions, bool) {
	if o == nil || o.Nfs == nil {
		return nil, false
	}
	return o.Nfs, true
}

// HasNfs returns a boolean if a field has been set.
func (o *Permissions) HasNfs() bool {
	if o != nil && o.Nfs != nil {
		return true
	}

	return false
}

// SetNfs gets a reference to the given NfsPermissions and assigns it to the Nfs field.
func (o *Permissions) SetNfs(v NfsPermissions) {
	o.Nfs = &v
}

func (o Permissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nfs != nil {
		toSerialize["nfs"] = o.Nfs
	}
	return json.Marshal(toSerialize)
}

type NullablePermissions struct {
	value *Permissions
	isSet bool
}

func (v NullablePermissions) Get() *Permissions {
	return v.value
}

func (v *NullablePermissions) Set(val *Permissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissions(val *Permissions) *NullablePermissions {
	return &NullablePermissions{value: val, isSet: true}
}

func (v NullablePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}