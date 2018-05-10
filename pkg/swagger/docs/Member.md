# Member

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **string** | In languages that use Western order, like English, this is the first name. | [default to null]
**FamilyName** | **string** | In languages that use Western order, like English, this is the last name. NOTE: some services lump suffixes (like Jr. or Sr.) in with the family name! | [default to null]
**MiddleName** | **string** | Middle name or middle initial | [optional] [default to null]
**Suffix** | **string** | Jr., Sr., III, etc. | [optional] [default to null]
**Affiliation** | [***Affiliation**](Affiliation.md) |  | [default to null]
**Rank** | [***Rank**](Rank.md) |  | [default to null]
**Title** | **string** | If supplied, this is the preferred form of address or preferred human-readable rank. This is especially useful when a rank has multiple possible titles. For example, in the Army, an E-4 can be either a Specialist or a Corporal. In the Marine Corps, an E-8 can be either a Master Sergeant or a First Sergeant, and they do care about the distinction.  If omitted, use the default name for the rank for the provided affiliation.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


