# OsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetrisController** | Pointer to [**OsConfigurationNetrisController**](OsConfigurationNetrisController.md) |  | [optional] 
**NetrisSoftgate** | Pointer to [**OsConfigurationNetrisSoftgate**](OsConfigurationNetrisSoftgate.md) |  | [optional] 
**Windows** | Pointer to [**OsConfigurationWindows**](OsConfigurationWindows.md) |  | [optional] 
**RootPassword** | Pointer to **string** | (Read-only) Auto-generated password set for user &#39;root&#39; on an ESXi or Proxmox server.&lt;br&gt;  The password is not stored and therefore will only be returned in response to provisioning a server. Copy and save it for future reference. | [optional] [readonly] 
**ManagementUiUrl** | Pointer to **string** | (Read-only) The URL of the management UI which will only be returned in response to provisioning a server. | [optional] [readonly] 
**ManagementAccessAllowedIps** | Pointer to **[]string** | List of IPs allowed to access the Management UI. Supported in single IP, CIDR and range format. When undefined, Management UI is disabled. This will only be returned in response to provisioning a server. | [optional] 
**InstallOsToRam** | Pointer to **bool** | If true, OS will be installed to and booted from the server&#39;s RAM. On restart RAM OS will be lost and the server will not be reachable unless a custom bootable OS has been deployed. Follow the &lt;a href&#x3D;&#39;https://phoenixnap.com/kb/bmc-custom-os&#39; target&#x3D;&#39;_blank&#39;&gt;instructions&lt;/a&gt; on how to install custom OS on BMC. Only supported for ubuntu/focal and ubuntu/jammy. | [optional] [default to false]
**CloudInit** | Pointer to [**OsConfigurationCloudInit**](OsConfigurationCloudInit.md) |  | [optional] 

## Methods

### NewOsConfiguration

`func NewOsConfiguration() *OsConfiguration`

NewOsConfiguration instantiates a new OsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsConfigurationWithDefaults

`func NewOsConfigurationWithDefaults() *OsConfiguration`

NewOsConfigurationWithDefaults instantiates a new OsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetrisController

`func (o *OsConfiguration) GetNetrisController() OsConfigurationNetrisController`

GetNetrisController returns the NetrisController field if non-nil, zero value otherwise.

### GetNetrisControllerOk

`func (o *OsConfiguration) GetNetrisControllerOk() (*OsConfigurationNetrisController, bool)`

GetNetrisControllerOk returns a tuple with the NetrisController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetrisController

`func (o *OsConfiguration) SetNetrisController(v OsConfigurationNetrisController)`

SetNetrisController sets NetrisController field to given value.

### HasNetrisController

`func (o *OsConfiguration) HasNetrisController() bool`

HasNetrisController returns a boolean if a field has been set.

### GetNetrisSoftgate

`func (o *OsConfiguration) GetNetrisSoftgate() OsConfigurationNetrisSoftgate`

GetNetrisSoftgate returns the NetrisSoftgate field if non-nil, zero value otherwise.

### GetNetrisSoftgateOk

`func (o *OsConfiguration) GetNetrisSoftgateOk() (*OsConfigurationNetrisSoftgate, bool)`

GetNetrisSoftgateOk returns a tuple with the NetrisSoftgate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetrisSoftgate

`func (o *OsConfiguration) SetNetrisSoftgate(v OsConfigurationNetrisSoftgate)`

SetNetrisSoftgate sets NetrisSoftgate field to given value.

### HasNetrisSoftgate

`func (o *OsConfiguration) HasNetrisSoftgate() bool`

HasNetrisSoftgate returns a boolean if a field has been set.

### GetWindows

`func (o *OsConfiguration) GetWindows() OsConfigurationWindows`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *OsConfiguration) GetWindowsOk() (*OsConfigurationWindows, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *OsConfiguration) SetWindows(v OsConfigurationWindows)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *OsConfiguration) HasWindows() bool`

HasWindows returns a boolean if a field has been set.

### GetRootPassword

`func (o *OsConfiguration) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *OsConfiguration) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *OsConfiguration) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *OsConfiguration) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetManagementUiUrl

`func (o *OsConfiguration) GetManagementUiUrl() string`

GetManagementUiUrl returns the ManagementUiUrl field if non-nil, zero value otherwise.

### GetManagementUiUrlOk

`func (o *OsConfiguration) GetManagementUiUrlOk() (*string, bool)`

GetManagementUiUrlOk returns a tuple with the ManagementUiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUiUrl

`func (o *OsConfiguration) SetManagementUiUrl(v string)`

SetManagementUiUrl sets ManagementUiUrl field to given value.

### HasManagementUiUrl

`func (o *OsConfiguration) HasManagementUiUrl() bool`

HasManagementUiUrl returns a boolean if a field has been set.

### GetManagementAccessAllowedIps

`func (o *OsConfiguration) GetManagementAccessAllowedIps() []string`

GetManagementAccessAllowedIps returns the ManagementAccessAllowedIps field if non-nil, zero value otherwise.

### GetManagementAccessAllowedIpsOk

`func (o *OsConfiguration) GetManagementAccessAllowedIpsOk() (*[]string, bool)`

GetManagementAccessAllowedIpsOk returns a tuple with the ManagementAccessAllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAccessAllowedIps

`func (o *OsConfiguration) SetManagementAccessAllowedIps(v []string)`

SetManagementAccessAllowedIps sets ManagementAccessAllowedIps field to given value.

### HasManagementAccessAllowedIps

`func (o *OsConfiguration) HasManagementAccessAllowedIps() bool`

HasManagementAccessAllowedIps returns a boolean if a field has been set.

### GetInstallOsToRam

`func (o *OsConfiguration) GetInstallOsToRam() bool`

GetInstallOsToRam returns the InstallOsToRam field if non-nil, zero value otherwise.

### GetInstallOsToRamOk

`func (o *OsConfiguration) GetInstallOsToRamOk() (*bool, bool)`

GetInstallOsToRamOk returns a tuple with the InstallOsToRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallOsToRam

`func (o *OsConfiguration) SetInstallOsToRam(v bool)`

SetInstallOsToRam sets InstallOsToRam field to given value.

### HasInstallOsToRam

`func (o *OsConfiguration) HasInstallOsToRam() bool`

HasInstallOsToRam returns a boolean if a field has been set.

### GetCloudInit

`func (o *OsConfiguration) GetCloudInit() OsConfigurationCloudInit`

GetCloudInit returns the CloudInit field if non-nil, zero value otherwise.

### GetCloudInitOk

`func (o *OsConfiguration) GetCloudInitOk() (*OsConfigurationCloudInit, bool)`

GetCloudInitOk returns a tuple with the CloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInit

`func (o *OsConfiguration) SetCloudInit(v OsConfigurationCloudInit)`

SetCloudInit sets CloudInit field to given value.

### HasCloudInit

`func (o *OsConfiguration) HasCloudInit() bool`

HasCloudInit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


