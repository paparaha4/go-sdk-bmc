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

// PrivateNetworkConfiguration Private network details of bare metal server.
type PrivateNetworkConfiguration struct {
	// Deprecated in favour of a common gateway address across all networks available under NetworkConfiguration.<br> The address of the gateway assigned / to assign to the server.<br> When used as part of request body, IP address has to be part of private network assigned to this server.<br> Gateway address also has to be assigned on an already deployed resource unless the `force` query parameter is true.
	// Deprecated
	GatewayAddress *string `json:"gatewayAddress,omitempty"`
	// (Write-only) Determines the approach for configuring private network(s) for the server being provisioned. Currently this field should be set to `USE_OR_CREATE_DEFAULT` or `USER_DEFINED`.
	ConfigurationType *string `json:"configurationType,omitempty"`
	// The list of private networks this server is member of. When this field is part of request body, it'll be used to specify the private networks to assign to this server upon provisioning. Used alongside the `USER_DEFINED` configurationType.
	PrivateNetworks []ServerPrivateNetwork `json:"privateNetworks,omitempty"`
}

// NewPrivateNetworkConfiguration instantiates a new PrivateNetworkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateNetworkConfiguration() *PrivateNetworkConfiguration {
	this := PrivateNetworkConfiguration{}
	var configurationType string = "USE_OR_CREATE_DEFAULT"
	this.ConfigurationType = &configurationType
	return &this
}

// NewPrivateNetworkConfigurationWithDefaults instantiates a new PrivateNetworkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateNetworkConfigurationWithDefaults() *PrivateNetworkConfiguration {
	this := PrivateNetworkConfiguration{}
	var configurationType string = "USE_OR_CREATE_DEFAULT"
	this.ConfigurationType = &configurationType
	return &this
}

// GetGatewayAddress returns the GatewayAddress field value if set, zero value otherwise.
// Deprecated
func (o *PrivateNetworkConfiguration) GetGatewayAddress() string {
	if o == nil || o.GatewayAddress == nil {
		var ret string
		return ret
	}
	return *o.GatewayAddress
}

// GetGatewayAddressOk returns a tuple with the GatewayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PrivateNetworkConfiguration) GetGatewayAddressOk() (*string, bool) {
	if o == nil || o.GatewayAddress == nil {
		return nil, false
	}
	return o.GatewayAddress, true
}

// HasGatewayAddress returns a boolean if a field has been set.
func (o *PrivateNetworkConfiguration) HasGatewayAddress() bool {
	if o != nil && o.GatewayAddress != nil {
		return true
	}

	return false
}

// SetGatewayAddress gets a reference to the given string and assigns it to the GatewayAddress field.
// Deprecated
func (o *PrivateNetworkConfiguration) SetGatewayAddress(v string) {
	o.GatewayAddress = &v
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *PrivateNetworkConfiguration) GetConfigurationType() string {
	if o == nil || o.ConfigurationType == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkConfiguration) GetConfigurationTypeOk() (*string, bool) {
	if o == nil || o.ConfigurationType == nil {
		return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *PrivateNetworkConfiguration) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType != nil {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given string and assigns it to the ConfigurationType field.
func (o *PrivateNetworkConfiguration) SetConfigurationType(v string) {
	o.ConfigurationType = &v
}

// GetPrivateNetworks returns the PrivateNetworks field value if set, zero value otherwise.
func (o *PrivateNetworkConfiguration) GetPrivateNetworks() []ServerPrivateNetwork {
	if o == nil || o.PrivateNetworks == nil {
		var ret []ServerPrivateNetwork
		return ret
	}
	return o.PrivateNetworks
}

// GetPrivateNetworksOk returns a tuple with the PrivateNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateNetworkConfiguration) GetPrivateNetworksOk() ([]ServerPrivateNetwork, bool) {
	if o == nil || o.PrivateNetworks == nil {
		return nil, false
	}
	return o.PrivateNetworks, true
}

// HasPrivateNetworks returns a boolean if a field has been set.
func (o *PrivateNetworkConfiguration) HasPrivateNetworks() bool {
	if o != nil && o.PrivateNetworks != nil {
		return true
	}

	return false
}

// SetPrivateNetworks gets a reference to the given []ServerPrivateNetwork and assigns it to the PrivateNetworks field.
func (o *PrivateNetworkConfiguration) SetPrivateNetworks(v []ServerPrivateNetwork) {
	o.PrivateNetworks = v
}

func (o PrivateNetworkConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GatewayAddress != nil {
		toSerialize["gatewayAddress"] = o.GatewayAddress
	}
	if o.ConfigurationType != nil {
		toSerialize["configurationType"] = o.ConfigurationType
	}
	if o.PrivateNetworks != nil {
		toSerialize["privateNetworks"] = o.PrivateNetworks
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateNetworkConfiguration struct {
	value *PrivateNetworkConfiguration
	isSet bool
}

func (v NullablePrivateNetworkConfiguration) Get() *PrivateNetworkConfiguration {
	return v.value
}

func (v *NullablePrivateNetworkConfiguration) Set(val *PrivateNetworkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateNetworkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateNetworkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateNetworkConfiguration(val *PrivateNetworkConfiguration) *NullablePrivateNetworkConfiguration {
	return &NullablePrivateNetworkConfiguration{value: val, isSet: true}
}

func (v NullablePrivateNetworkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateNetworkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
