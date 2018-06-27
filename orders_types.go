package main

import (
	"github.com/transcom/nom/pkg/gen/ordersapi/models"
)

var enlistedOrdersTypes = map[string]models.OrdersType{
	"1": models.OrdersTypeIpcot,          // IPCOT In-place consecutive overseas travel
	"8": models.OrdersTypeOteip,          // Overseas Tour Extension Incentive Program (OTEIP)
	"9": models.OrdersTypeTraining,       // NAVCAD (Naval Cadet) Training
	"A": models.OrdersTypeAccession,      // Accession Travel Recruits
	"B": models.OrdersTypeAccession,      // Non-recruit Accession Travel
	"C": models.OrdersTypeTraining,       // Training Travel
	"D": models.OrdersTypeOperational,    // Operational Travel
	"E": models.OrdersTypeSeparation,     // Separation Travel
	"F": models.OrdersTypeUnitMove,       // Organized Unit/Homeport Change
	"G": models.OrdersTypeAccession,      // Midshipman Accession Travel
	"H": models.OrdersTypeSpecialPurpose, // Special Purpose Reimbursable
	"I": models.OrdersTypeAccession,      // NAVCAD(Naval Cadet) Accession
	"J": models.OrdersTypeAccession,      // Accession Travel Recruits
	"K": models.OrdersTypeAccession,      // Non-recruit Accession Travel
	"L": models.OrdersTypeTraining,       // Training Travel
	"M": models.OrdersTypeRotational,     // Rotational Travel
	"N": models.OrdersTypeSeparation,     // Separation Travel
	"O": models.OrdersTypeUnitMove,       // Organized Unit/Homeport Change
	"P": models.OrdersTypeSeparation,     // Midshipman Separation Travel
	"R": models.OrdersTypeOperational,    // Misc. Operational Non-member
	"X": models.OrdersTypeEmergencyEvac,  // EMERGENCY NON-MEMBER EVACS
	"Y": models.OrdersTypeRotational,     // Misc. Rotational Non-member
	"Z": models.OrdersTypeSeparation,     // NAVCAD(Naval Cadet) Separation
}

var officerOrdersTypes = map[string]models.OrdersType{
	"0": models.OrdersTypeIpcot,          // IPCOT In-place consecutive overseas travel
	"2": models.OrdersTypeAccession,      // Accession Travel
	"3": models.OrdersTypeTraining,       // Training Travel
	"4": models.OrdersTypeOperational,    // Operational Travel
	"5": models.OrdersTypeSeparation,     // Separation Travel
	"6": models.OrdersTypeUnitMove,       // Organized Unit/Homeport Change
	"7": models.OrdersTypeEmergencyEvac,  // Emergency Non-member Evac
	"H": models.OrdersTypeSpecialPurpose, // Special Purpose Reimbursable
	"Q": models.OrdersTypeRotational,     // Misc. Rotational Non-member
	"S": models.OrdersTypeAccession,      // Accession Travel
	"T": models.OrdersTypeTraining,       // Training Travel
	"U": models.OrdersTypeRotational,     // Rotational Travel
	"V": models.OrdersTypeSeparation,     // Separation Travel
	"W": models.OrdersTypeUnitMove,       // Organized Unit/Homeport Change
	"X": models.OrdersTypeRotational,     // Misc. Rotational Non-member
}
