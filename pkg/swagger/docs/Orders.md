# Orders

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Universally Unique IDentifier. Generated internally. | [optional] [default to null]
**OrdersNum** | **string** | Orders number. Supposed to be unique, but in practice uniqueness is not guaranteed for all branches of service.  ## Army Typically found in the upper-left hand corner of printed orders.  ## Navy Corresponds to a transformation (see below) of the BUPERS Orders number, which is originally formatted as \&quot;{Julian Day of the year (from 1 to 366)}{last digit of the year} {xxx-xx-xxxx SSN}\&quot;. It would be unique (because of the SSN), except that it&#39;s possible for a set of orders to be cut on the same day 10 years later for the same sailor, resulting in a collision.  On printed orders, this is typically the first line after \&quot;RMKS/\&quot;.  ### Transformation Because the BUPERS Orders Number contains PII (the SSN) and could potentially not be unique (because it only allows a single digit for the year), this API expects the Orders number to be reformatted as \&quot;{Julian day}{4 digit year} {EDIPI}\&quot;. The full 4 digit year can be found in the dateIssued field of the first revision, while the EDIPI can be copied from the edipi field.  Transforming this back to the BUPERS format can be achieved by removing the first three digits of the year field and replacing the EDIPI with the member&#39;s SSN. The SSN can be fetched using DMDC&#39;s Identity Web Services Real-time Broker Service.  ## Marine Corps Corresponds to the CT (Commercial Travel) SDN. On Web Orders, the CT SDN is found in the table at the bottom, in the last column of the row that begins with \&quot;Travel\&quot;.  ## Air Force Corresponds to the Special Order number. On AF Form 899, the \&quot;SPECIAL ORDERS NO\&quot; is found in box 27.  ## Coast Guard Corresponds to the Travel Order No.  ## Civilian Corresponds to the Travel Authorization Number.  | [default to null]
**Edipi** | **string** | Electronic Data Interchange Personal Identifier, AKA the 10 digit DoD ID Number of the member | [default to null]
**IssuingAuthority** | **string** | Military Department or Civilian Agency that authorized these orders; e.g., Department of the Army, Department of the Navy, Defense Information Systems Agency (DISA), etc. | [default to null]
**Revisions** | [**[]Revision**](Revision.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


