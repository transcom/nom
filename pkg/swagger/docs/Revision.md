# Revision

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SeqNum** | **int32** | Sequence number. As Orders are amended, the Revision with the highest sequence number is considered the current, authoritative version of the Orders, even if its dateIssued is earlier.  The sequence number is NOT required to increase monotonically or sequentially; in other words, if a set of orders is modified twice, the sequence numbers could be 1, 5858300, and 30.  | [default to null]
**Member** | [***Member**](Member.md) |  | [default to null]
**Status** | **string** | Indicates whether these Orders are authorized or canceled. If omitted, then these Orders are assumed to be authorized. | [optional] [default to null]
**DateIssued** | [**time.Time**](time.Time.md) | The date and time that these orders were cut. If omitted, the current date and time will be used. | [optional] [default to null]
**PcaWithoutPcs** | **bool** | Permanent Change of Assignment without Permanent Change of Station. A PCA without PCS happens when a member is assigned to a new unit at the same duty station, or to a new duty station geographically close to the current duty station.  If true, then these orders do not authorize any move expenses. If omitted or false, then these orders are a PCS and should authorize move expenses.  It is not unheard of for the initial revision of orders to have this set to false and then later to be amended to true and vice-versa.  | [optional] [default to null]
**TdyEnRoute** | **bool** | TDY (Temporary Duty Yonder) en-route. If omitted, assume false. | [optional] [default to null]
**TourType** | [***TourType**](TourType.md) |  | [optional] [default to null]
**OrdersType** | [***OrdersType**](OrdersType.md) |  | [default to null]
**HasDependents** | **bool** | True if the service member has any dependents (e.g., spouse, children, caring for an elderly parent, etc.), False otherwise.  When the member has dependents, it usually raises their weight entitlement.  | [default to null]
**LosingUnit** | [***Unit**](Unit.md) |  | [default to null]
**GainingUnit** | [***Unit**](Unit.md) |  | [default to null]
**ReportNoEarlierThan** | **string** | Earliest date that the service member is allowed to report for duty at the new duty station. If omitted, the member is allowed to report as early as desired. | [optional] [default to null]
**ReportNoLaterThan** | **string** | Latest date that the service member is allowed to report for duty at the new duty station. Should be included for most Orders types, but can be missing for Separation / Retirement Orders. | [optional] [default to null]
**PcsAccounting** | [***Accounting**](Accounting.md) |  | [optional] [default to null]
**NtsAccounting** | [***Accounting**](Accounting.md) |  | [optional] [default to null]
**PovShipmentAccounting** | [***Accounting**](Accounting.md) |  | [optional] [default to null]
**PovStorageAccounting** | [***Accounting**](Accounting.md) |  | [optional] [default to null]
**UbAccounting** | [***Accounting**](Accounting.md) |  | [optional] [default to null]
**Comments** | **string** | Free-form text that may or may not contain information relevant to moving. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


