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

// ServerPublicNetwork Public network details of bare metal server.
type ServerPublicNetwork struct {
	// The network identifier.
	Id string `json:"id"`
	// IPs to configure/configured on the server. IPs must be within the network's range.
	Ips []string `json:"ips"`
	// The status of the assignment to the network.
	StatusDescription *string `json:"statusDescription,omitempty"`
}

// NewServerPublicNetwork instantiates a new ServerPublicNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerPublicNetwork(id string, ips []string) *ServerPublicNetwork {
	this := ServerPublicNetwork{}
	this.Id = id
	this.Ips = ips
	return &this
}

// NewServerPublicNetworkWithDefaults instantiates a new ServerPublicNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerPublicNetworkWithDefaults() *ServerPublicNetwork {
	this := ServerPublicNetwork{}
	return &this
}

// GetId returns the Id field value
func (o *ServerPublicNetwork) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServerPublicNetwork) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServerPublicNetwork) SetId(v string) {
	o.Id = v
}

// GetIps returns the Ips field value
func (o *ServerPublicNetwork) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *ServerPublicNetwork) GetIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ips, true
}

// SetIps sets field value
func (o *ServerPublicNetwork) SetIps(v []string) {
	o.Ips = v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *ServerPublicNetwork) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerPublicNetwork) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *ServerPublicNetwork) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *ServerPublicNetwork) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

func (o ServerPublicNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ips"] = o.Ips
	}
	if o.StatusDescription != nil {
		toSerialize["statusDescription"] = o.StatusDescription
	}
	return json.Marshal(toSerialize)
}

type NullableServerPublicNetwork struct {
	value *ServerPublicNetwork
	isSet bool
}

func (v NullableServerPublicNetwork) Get() *ServerPublicNetwork {
	return v.value
}

func (v *NullableServerPublicNetwork) Set(val *ServerPublicNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableServerPublicNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableServerPublicNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerPublicNetwork(val *ServerPublicNetwork) *NullableServerPublicNetwork {
	return &NullableServerPublicNetwork{value: val, isSet: true}
}

func (v NullableServerPublicNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerPublicNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}