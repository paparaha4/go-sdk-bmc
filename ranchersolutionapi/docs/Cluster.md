# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | (Read-only) The Cluster identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Cluster name. This field is autogenerated if not provided. | [optional] 
**Description** | Pointer to **string** | Cluster description. | [optional] 
**Location** | **string** | Deployment location. | 
**InitialClusterVersion** | Pointer to **string** | (Read-only) The Rancher version that was installed on the cluster during the first creation process. | [optional] [readonly] 
**NodePools** | Pointer to [**[]NodePool**](NodePool.md) | The node pools associated with the cluster. | [optional] 
**Configuration** | Pointer to [**RancherClusterConfig**](RancherClusterConfig.md) |  | [optional] 
**Metadata** | Pointer to [**RancherServerMetadata**](RancherServerMetadata.md) |  | [optional] 
**StatusDescription** | Pointer to **string** | The cluster status | [optional] [readonly] 

## Methods

### NewCluster

`func NewCluster(location string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Cluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *Cluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Cluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Cluster) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetInitialClusterVersion

`func (o *Cluster) GetInitialClusterVersion() string`

GetInitialClusterVersion returns the InitialClusterVersion field if non-nil, zero value otherwise.

### GetInitialClusterVersionOk

`func (o *Cluster) GetInitialClusterVersionOk() (*string, bool)`

GetInitialClusterVersionOk returns a tuple with the InitialClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialClusterVersion

`func (o *Cluster) SetInitialClusterVersion(v string)`

SetInitialClusterVersion sets InitialClusterVersion field to given value.

### HasInitialClusterVersion

`func (o *Cluster) HasInitialClusterVersion() bool`

HasInitialClusterVersion returns a boolean if a field has been set.

### GetNodePools

`func (o *Cluster) GetNodePools() []NodePool`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *Cluster) GetNodePoolsOk() (*[]NodePool, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *Cluster) SetNodePools(v []NodePool)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *Cluster) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetConfiguration

`func (o *Cluster) GetConfiguration() RancherClusterConfig`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Cluster) GetConfigurationOk() (*RancherClusterConfig, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Cluster) SetConfiguration(v RancherClusterConfig)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Cluster) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetMetadata

`func (o *Cluster) GetMetadata() RancherServerMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Cluster) GetMetadataOk() (*RancherServerMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Cluster) SetMetadata(v RancherServerMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Cluster) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatusDescription

`func (o *Cluster) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Cluster) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Cluster) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *Cluster) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

