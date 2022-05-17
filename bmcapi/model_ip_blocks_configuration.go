/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API.  Deprovision servers, get or edit SSH key details, assign public IPs, assign servers to networks and a lot more.  Manage your infrastructure more efficiently using just a few simple API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b> 

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// IpBlocksConfiguration The IP blocks to assign to this server. <b>This is an exclusive allocation</b>, i.e. the IP blocks cannot be shared with other servers. If IpBlocksConfiguration is not defined, the purchase of a new IP block is determined by the networkType field.
type IpBlocksConfiguration struct {
	// Determines the approach for configuring IP blocks for the server being provisioned. If PURCHASE_NEW is selected, the smallest supported range, depending on the operating system, is allocated to the server.
	ConfigurationType *string `json:"configurationType,omitempty"`
	// Used to specify the previously purchased IP blocks to assign to this server upon provisioning. Used alongside the USER_DEFINED configurationType.
	IpBlocks *[]ServerIpBlock `json:"ipBlocks,omitempty"`
}

// NewIpBlocksConfiguration instantiates a new IpBlocksConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlocksConfiguration() *IpBlocksConfiguration {
	this := IpBlocksConfiguration{}
	var configurationType string = "PURCHASE_NEW"
	this.ConfigurationType = &configurationType
	return &this
}

// NewIpBlocksConfigurationWithDefaults instantiates a new IpBlocksConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlocksConfigurationWithDefaults() *IpBlocksConfiguration {
	this := IpBlocksConfiguration{}
	var configurationType string = "PURCHASE_NEW"
	this.ConfigurationType = &configurationType
	return &this
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *IpBlocksConfiguration) GetConfigurationType() string {
	if o == nil || o.ConfigurationType == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlocksConfiguration) GetConfigurationTypeOk() (*string, bool) {
	if o == nil || o.ConfigurationType == nil {
		return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *IpBlocksConfiguration) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType != nil {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given string and assigns it to the ConfigurationType field.
func (o *IpBlocksConfiguration) SetConfigurationType(v string) {
	o.ConfigurationType = &v
}

// GetIpBlocks returns the IpBlocks field value if set, zero value otherwise.
func (o *IpBlocksConfiguration) GetIpBlocks() []ServerIpBlock {
	if o == nil || o.IpBlocks == nil {
		var ret []ServerIpBlock
		return ret
	}
	return *o.IpBlocks
}

// GetIpBlocksOk returns a tuple with the IpBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlocksConfiguration) GetIpBlocksOk() (*[]ServerIpBlock, bool) {
	if o == nil || o.IpBlocks == nil {
		return nil, false
	}
	return o.IpBlocks, true
}

// HasIpBlocks returns a boolean if a field has been set.
func (o *IpBlocksConfiguration) HasIpBlocks() bool {
	if o != nil && o.IpBlocks != nil {
		return true
	}

	return false
}

// SetIpBlocks gets a reference to the given []ServerIpBlock and assigns it to the IpBlocks field.
func (o *IpBlocksConfiguration) SetIpBlocks(v []ServerIpBlock) {
	o.IpBlocks = &v
}

func (o IpBlocksConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigurationType != nil {
		toSerialize["configurationType"] = o.ConfigurationType
	}
	if o.IpBlocks != nil {
		toSerialize["ipBlocks"] = o.IpBlocks
	}
	return json.Marshal(toSerialize)
}

type NullableIpBlocksConfiguration struct {
	value *IpBlocksConfiguration
	isSet bool
}

func (v NullableIpBlocksConfiguration) Get() *IpBlocksConfiguration {
	return v.value
}

func (v *NullableIpBlocksConfiguration) Set(val *IpBlocksConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlocksConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlocksConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlocksConfiguration(val *IpBlocksConfiguration) *NullableIpBlocksConfiguration {
	return &NullableIpBlocksConfiguration{value: val, isSet: true}
}

func (v NullableIpBlocksConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlocksConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


