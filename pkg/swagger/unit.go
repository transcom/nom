/*
 * MyMove Orders Gateway
 *
 * API to submit, amend, and cancel orders for move.mil.
 *
 * API version: 0.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Unit struct {

	// Human-readable name of the Unit.
	Name string `json:"name,omitempty"`

	// Unit Identification Code - a six character alphanumeric code that uniquely identifies each United States Department of Defense entity. Used in Army, Air Force, and Navy orders.  Note that the Navy has the habit of omitting the leading character, which is always \"N\" for them. 
	Uic string `json:"uic,omitempty"`

	// May be FPO or APO for OCONUS commands.
	City string `json:"city"`

	// State (US). OCONUS units may not have the equivalent information available.
	Locality string `json:"locality,omitempty"`

	// ISO 3166-1 alpha-2 country code. If blank, assume US
	Country string `json:"country,omitempty"`

	// In the USA, this is the ZIP Code.
	PostalCode string `json:"postal-code"`
}
