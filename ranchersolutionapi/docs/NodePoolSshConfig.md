# NodePoolSshConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallDefaultKeys** | Pointer to **bool** | Define whether public keys marked as default should be installed on this node. These are public keys that were already recorded on this system. Use &lt;a href&#x3D;&#39;https://developers.phoenixnap.com/docs/bmc/1/routes/ssh-keys/get&#39; target&#x3D;&#39;_blank&#39;&gt;GET /ssh-keys&lt;/a&gt; to retrieve a list of possible values. | [optional] [default to true]
**Keys** | Pointer to **[]string** | List of public SSH keys. | [optional] 
**KeyIds** | Pointer to **[]string** | List of public SSH key identifiers. These are public keys that were already recorded on this system. Use &lt;a href&#x3D;&#39;https://developers.phoenixnap.com/docs/bmc/1/routes/ssh-keys/get&#39; target&#x3D;&#39;_blank&#39;&gt;GET /ssh-keys&lt;/a&gt; to retrieve a list of possible values. | [optional] 

## Methods

### NewNodePoolSshConfig

`func NewNodePoolSshConfig() *NodePoolSshConfig`

NewNodePoolSshConfig instantiates a new NodePoolSshConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolSshConfigWithDefaults

`func NewNodePoolSshConfigWithDefaults() *NodePoolSshConfig`

NewNodePoolSshConfigWithDefaults instantiates a new NodePoolSshConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallDefaultKeys

`func (o *NodePoolSshConfig) GetInstallDefaultKeys() bool`

GetInstallDefaultKeys returns the InstallDefaultKeys field if non-nil, zero value otherwise.

### GetInstallDefaultKeysOk

`func (o *NodePoolSshConfig) GetInstallDefaultKeysOk() (*bool, bool)`

GetInstallDefaultKeysOk returns a tuple with the InstallDefaultKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallDefaultKeys

`func (o *NodePoolSshConfig) SetInstallDefaultKeys(v bool)`

SetInstallDefaultKeys sets InstallDefaultKeys field to given value.

### HasInstallDefaultKeys

`func (o *NodePoolSshConfig) HasInstallDefaultKeys() bool`

HasInstallDefaultKeys returns a boolean if a field has been set.

### GetKeys

`func (o *NodePoolSshConfig) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *NodePoolSshConfig) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *NodePoolSshConfig) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *NodePoolSshConfig) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetKeyIds

`func (o *NodePoolSshConfig) GetKeyIds() []string`

GetKeyIds returns the KeyIds field if non-nil, zero value otherwise.

### GetKeyIdsOk

`func (o *NodePoolSshConfig) GetKeyIdsOk() (*[]string, bool)`

GetKeyIdsOk returns a tuple with the KeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyIds

`func (o *NodePoolSshConfig) SetKeyIds(v []string)`

SetKeyIds sets KeyIds field to given value.

### HasKeyIds

`func (o *NodePoolSshConfig) HasKeyIds() bool`

HasKeyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


