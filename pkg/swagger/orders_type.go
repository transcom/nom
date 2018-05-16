/*
 * MyMove Orders Gateway
 *
 * API to submit, amend, and cancel orders for move.mil.
 *
 * API version: 0.0.4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger
// OrdersType : The common types fit into the acronym ASTRO-UB.   * **A**ccession - Joining the military   * **S**eparation / Retirement - Leaving the military   * **T**raining   * **R**otational   * **O**perational   * **U**nit Move - When an entire unit is reassigned to another installation, often as a deployment   * **B**RAC - Base Realignment and Closure. As of this writing, the most recent iteration of BRAC has ended, but Congress may start another one in the future.  Consequences of this field include   * Separation and retirement moves currently require the member to go through in-person counseling at the TMO / PPPO. 
type OrdersType string

// List of OrdersType
const (
	ACCESSION OrdersType = "accession"
	BETWEEN_DUTY_STATIONS OrdersType = "between-duty-stations"
	BRAC OrdersType = "brac"
	COT OrdersType = "cot"
	EMERGENCY_EVAC OrdersType = "emergency-evac"
	IPCOT OrdersType = "ipcot"
	LOW_COST_TRAVEL OrdersType = "low-cost-travel"
	OPERATIONAL OrdersType = "operational"
	OTEIP OrdersType = "oteip"
	ROTATIONAL OrdersType = "rotational"
	SEPARATION OrdersType = "separation"
	SPECIAL_PURPOSE OrdersType = "special-purpose"
	TRAINING OrdersType = "training"
	UNIT_MOVE OrdersType = "unit-move"
)
