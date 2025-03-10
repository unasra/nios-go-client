# RecordA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRte53RecordInfo** | Pointer to **string** | Aws Route 53 record information. | [optional] 
**CloudInfo** | Pointer to **string** | Structure containing all cloud API related information for this object. | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **string** | The time of the record creation in Epoch seconds format. | [optional] 
**Creator** | Pointer to **string** | The record creator. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DiscoveredData** | Pointer to **string** | The discovered data for this A record. | [optional] 
**DnsName** | Pointer to **string** | The name for an A record in punycode format. | [optional] 
**Extattrs** | Pointer to **string** | Extensible attributes associated with the object. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the record. | [optional] 
**LastQueried** | Pointer to **string** | The time of the last DNS query in Epoch seconds format. | [optional] 
**MsAdUserData** | Pointer to **string** | The Microsoft Active Directory user related information. | [optional] 
**Name** | Pointer to **string** | Name for A record in FQDN format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] 
**RemoveAssociatedPtr** | Pointer to **bool** | Whether to remove associated PTR records while deleting the A record. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The shared record group this record belongs to. | [optional] 
**Ttl** | Pointer to **int32** | Time-to-live value of the record, in seconds. | [optional] 
**UseTtl** | Pointer to **bool** | Flag to indicate whether the TTL value should be used for the A record. | [optional] 
**View** | Pointer to **string** | View that this record is part of. | [optional] 
**Zone** | Pointer to **string** | The zone in which the record resides. | [optional] 

## Methods

### NewRecordA

`func NewRecordA() *RecordA`

NewRecordA instantiates a new RecordA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordAWithDefaults

`func NewRecordAWithDefaults() *RecordA`

NewRecordAWithDefaults instantiates a new RecordA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRte53RecordInfo

`func (o *RecordA) GetAwsRte53RecordInfo() string`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordA) GetAwsRte53RecordInfoOk() (*string, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordA) SetAwsRte53RecordInfo(v string)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordA) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordA) GetCloudInfo() string`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordA) GetCloudInfoOk() (*string, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordA) SetCloudInfo(v string)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordA) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordA) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordA) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordA) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordA) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordA) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordA) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordA) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordA) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordA) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordA) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordA) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordA) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordA) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordA) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordA) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordA) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordA) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordA) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordA) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordA) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordA) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordA) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordA) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordA) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *RecordA) GetDiscoveredData() string`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *RecordA) GetDiscoveredDataOk() (*string, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *RecordA) SetDiscoveredData(v string)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *RecordA) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordA) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordA) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordA) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordA) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtattrs

`func (o *RecordA) GetExtattrs() string`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *RecordA) GetExtattrsOk() (*string, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *RecordA) SetExtattrs(v string)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *RecordA) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordA) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordA) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordA) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordA) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv4addr

`func (o *RecordA) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *RecordA) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *RecordA) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *RecordA) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordA) GetLastQueried() string`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordA) GetLastQueriedOk() (*string, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordA) SetLastQueried(v string)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordA) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *RecordA) GetMsAdUserData() string`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *RecordA) GetMsAdUserDataOk() (*string, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *RecordA) SetMsAdUserData(v string)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *RecordA) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *RecordA) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordA) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordA) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordA) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordA) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordA) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordA) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordA) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetRemoveAssociatedPtr

`func (o *RecordA) GetRemoveAssociatedPtr() bool`

GetRemoveAssociatedPtr returns the RemoveAssociatedPtr field if non-nil, zero value otherwise.

### GetRemoveAssociatedPtrOk

`func (o *RecordA) GetRemoveAssociatedPtrOk() (*bool, bool)`

GetRemoveAssociatedPtrOk returns a tuple with the RemoveAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAssociatedPtr

`func (o *RecordA) SetRemoveAssociatedPtr(v bool)`

SetRemoveAssociatedPtr sets RemoveAssociatedPtr field to given value.

### HasRemoveAssociatedPtr

`func (o *RecordA) HasRemoveAssociatedPtr() bool`

HasRemoveAssociatedPtr returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *RecordA) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *RecordA) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *RecordA) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *RecordA) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *RecordA) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordA) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordA) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordA) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordA) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordA) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordA) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordA) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordA) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordA) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordA) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordA) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordA) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordA) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordA) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordA) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


