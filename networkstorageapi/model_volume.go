/*
Network Storage API

Create, list, edit, and delete storage networks with the Network Storage API. Use storage networks to expand storage capacity on a private network. <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bare-metal-cloud-storage' target='_blank'>here</a> </span> <br> <b>All URLs are relative to (https://api-dev.phoenixnap.com/network-storage/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkstorageapi

import (
	"encoding/json"
	"time"
)

// Volume Volume for a storage network.
type Volume struct {
	// Volume ID.
	Id *string `json:"id,omitempty"`
	// Volume friendly name.
	Name *string `json:"name,omitempty"`
	// Volume description.
	Description *string `json:"description,omitempty"`
	// Volume's full path. It is in form of `/{volumeId}/pathSuffix`'.
	Path *string `json:"path,omitempty"`
	// Last part of volume's path.
	PathSuffix *string `json:"pathSuffix,omitempty"`
	// Maximum capacity in GB.
	CapacityInGb *int32 `json:"capacityInGb,omitempty"`
	// File system protocol. Currently this field should be set to `NFS`.
	Protocol    *string      `json:"protocol,omitempty"`
	Status      *Status      `json:"status,omitempty"`
	CreatedOn   *time.Time   `json:"createdOn,omitempty"`
	Permissions *Permissions `json:"permissions,omitempty"`
}

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume() *Volume {
	this := Volume{}
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Volume) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Volume) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Volume) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Volume) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Volume) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Volume) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Volume) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Volume) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Volume) SetDescription(v string) {
	o.Description = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Volume) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Volume) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Volume) SetPath(v string) {
	o.Path = &v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *Volume) GetPathSuffix() string {
	if o == nil || o.PathSuffix == nil {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPathSuffixOk() (*string, bool) {
	if o == nil || o.PathSuffix == nil {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *Volume) HasPathSuffix() bool {
	if o != nil && o.PathSuffix != nil {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *Volume) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetCapacityInGb returns the CapacityInGb field value if set, zero value otherwise.
func (o *Volume) GetCapacityInGb() int32 {
	if o == nil || o.CapacityInGb == nil {
		var ret int32
		return ret
	}
	return *o.CapacityInGb
}

// GetCapacityInGbOk returns a tuple with the CapacityInGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCapacityInGbOk() (*int32, bool) {
	if o == nil || o.CapacityInGb == nil {
		return nil, false
	}
	return o.CapacityInGb, true
}

// HasCapacityInGb returns a boolean if a field has been set.
func (o *Volume) HasCapacityInGb() bool {
	if o != nil && o.CapacityInGb != nil {
		return true
	}

	return false
}

// SetCapacityInGb gets a reference to the given int32 and assigns it to the CapacityInGb field.
func (o *Volume) SetCapacityInGb(v int32) {
	o.CapacityInGb = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Volume) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Volume) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Volume) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Volume) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Volume) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *Volume) SetStatus(v Status) {
	o.Status = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Volume) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Volume) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Volume) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Volume) GetPermissions() Permissions {
	if o == nil || o.Permissions == nil {
		var ret Permissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetPermissionsOk() (*Permissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Volume) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given Permissions and assigns it to the Permissions field.
func (o *Volume) SetPermissions(v Permissions) {
	o.Permissions = &v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.PathSuffix != nil {
		toSerialize["pathSuffix"] = o.PathSuffix
	}
	if o.CapacityInGb != nil {
		toSerialize["capacityInGb"] = o.CapacityInGb
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CreatedOn != nil {
		toSerialize["createdOn"] = o.CreatedOn
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	return json.Marshal(toSerialize)
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
