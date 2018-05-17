# Orders

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Universally Unique IDentifier. Generated internally. | [optional] [default to null]
**OrdersNum** | **string** | Orders number. Supposed to be unique, but in practice uniqueness is not guaranteed for all branches of service.  # Army Typically found in the upper-left hand corner of printed orders.  # Navy Corresponds to the CT (Commercial Travel) SDN. On printed orders, this is found on the SDN line in the &#x60;------- ACCOUNTING DATA -------&#x60; section in the &#x60;PCS ACCOUNTING DATA&#x60; paragraph.  The BUPERS Orders number is not suitable, because it includes the sailor&#39;s full SSN, and the included four digit date code could repeat for a sailor if he or she gets orders exactly 10 years apart.  No-cost moves do not have a CT SDN, because they involve no travel. Without a CT SDN, USN Orders have nothing to use for the Orders number. Such Orders won&#39;t authorize any PCS expenses either, so they do not need to be submitted to this API.  # Marine Corps Corresponds to the CT (Commercial Travel) SDN. On Web Orders, the CT SDN is found in the table at the bottom, in the last column of the row that begins with \&quot;Travel\&quot;.  No-cost moves do not have a CT SDN, because they involve no travel. Without a CT SDN, USMC Orders have nothing to use for the Orders number. Such Orders won&#39;t authorize any PCS expenses either, so they do not need to be submitted to this API.  # Air Force Corresponds to the SPECIAL ORDER NO, found in box 27 on AF Form 899.  # Coast Guard Corresponds to the Travel Order No.  # Civilian Corresponds to the Travel Authorization Number.  | [default to null]
**Edipi** | **string** | Electronic Data Interchange Personal Identifier, AKA the 10 digit DoD ID Number of the member | [default to null]
**Issuer** | **string** | Military Department or Civilian Agency that authorized these orders; e.g., Department of the Army, Department of the Navy, Defense Information Systems Agency (DISA), etc. The issuer is inferred by the system when the Orders are first created, based on the authenticated client certificate. | [default to null]
**Revisions** | [**[]Revision**](Revision.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


