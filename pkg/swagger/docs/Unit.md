# Unit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable name of the Unit. | [optional] [default to null]
**Uic** | **string** | Unit Identification Code - a six character alphanumeric code that uniquely identifies each United States Department of Defense entity. Used in Army, Air Force, and Navy orders.  Note that the Navy has the habit of omitting the leading character, which is always \&quot;N\&quot; for them.  | [optional] [default to null]
**City** | **string** | May be FPO or APO for OCONUS commands. | [optional] [default to null]
**Locality** | **string** | State (US). OCONUS units may not have the equivalent information available. | [optional] [default to null]
**Country** | **string** | ISO 3166-1 alpha-2 country code. If blank, but city and locality or postalCode are not blank, assume US | [optional] [default to null]
**PostalCode** | **string** | In the USA, this is the ZIP Code. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


