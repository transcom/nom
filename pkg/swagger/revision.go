/*
 * my.move.mil Orders Gateway
 *
 * API to submit, amend, and cancel orders for my.move.mil.
 *
 * API version: 0.0.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Revision struct {

	// Sequence number. When an order is amended, it should get a higher sequence number than all previous revisions of the same orders. The orders with the highest sequence number is considered the current, authoritative version, even if its dateIssued is earlier.  The sequence number is NOT required to increase monotonically or sequentially; in other words, if a set of orders is modified twice, the sequence numbers could be 1, 5858300, and 30. 
	SeqNum int32 `json:"seqNum"`

	Member *Member `json:"member"`

	// Indicates whether these Orders are authorized or canceled.
	Status string `json:"status"`

	// The date and time that these orders were cut.
	DateIssued time.Time `json:"dateIssued,omitempty"`

	// Permanent Change of Assignment without Permanent Change of Station. A PCA without PCS happens when a member is assigned to a new unit at the same duty station, or to a new duty station geographically close to the current duty station.  If true, then these orders do not authorize any move expenses. If omitted or false, then these orders are a PCS and should authorize move expenses.  It is not unheard of for the initial revision of orders to have this set to false and then later to be amended to true and vice-versa. 
	PcaWithoutPcs bool `json:"pcaWithoutPcs,omitempty"`

	// TDY (Temporary Duty Yonder) en-route. If omitted, assume false. 
	TdyEnRoute bool `json:"tdyEnRoute,omitempty"`

	TourType *TourType `json:"tourType,omitempty"`

	OrdersType *OrdersType `json:"ordersType"`

	// True if the service member has dependents (e.g., spouse, children, caring for an elderly parent, etc.), False otherwise.  When the member has dependents, it usually raises their weight entitlement. 
	HasDependents bool `json:"hasDependents"`

	LosingUnit *Unit `json:"losingUnit"`

	GainingUnit *Unit `json:"gainingUnit"`

	// Earliest date that the service member is allowed to report for duty at the new duty station.
	ReportNoEarlierThan string `json:"reportNoEarlierThan,omitempty"`

	// Latest date that the service member is allowed to report for duty at the new duty station.
	ReportNoLaterThan string `json:"reportNoLaterThan"`

	PcsAccounting *Accounting `json:"pcsAccounting,omitempty"`

	NtsAccounting *Accounting `json:"ntsAccounting,omitempty"`

	PovShipmentAccounting *Accounting `json:"povShipmentAccounting,omitempty"`

	PovStorageAccounting *Accounting `json:"povStorageAccounting,omitempty"`

	UbAccounting *Accounting `json:"ubAccounting,omitempty"`

	// Free-form text that may or may not contain information relevant to moving.
	Comments string `json:"comments,omitempty"`
}
