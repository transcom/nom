package main

import (
	"strings"

	"github.com/transcom/nom/pkg/swagger"
)

// Rank correlates a title with its DoD paygrade and indicates whether the rank belongs to a commissioned or warrant officer
type RankTitlePaygrade struct {
	officer  bool
	title    string
	paygrade swagger.Rank
}

type rate struct {
	prefix   string
	suffix   string
	paygrade swagger.Rank
}

var officerRanks = map[string]RankTitlePaygrade{
	"OC":   RankTitlePaygrade{officer: true, title: "Officer Candidate", paygrade: swagger.E_5},
	"ENS":  RankTitlePaygrade{officer: true, title: "Ensign", paygrade: swagger.O_1},
	"LTJG": RankTitlePaygrade{officer: true, title: "Lieutenant Junior Grade", paygrade: swagger.O_2},
	"LT":   RankTitlePaygrade{officer: true, title: "Lieutenant", paygrade: swagger.O_3},
	"LCDR": RankTitlePaygrade{officer: true, title: "Lieutenant Commander", paygrade: swagger.O_4},
	"CDR":  RankTitlePaygrade{officer: true, title: "Commander", paygrade: swagger.O_5},
	"CAPT": RankTitlePaygrade{officer: true, title: "Captain", paygrade: swagger.O_6},
	"RDML": RankTitlePaygrade{officer: true, title: "Rear Admiral (lower half)", paygrade: swagger.O_7},
	"RDMU": RankTitlePaygrade{officer: true, title: "Rear Admiral (upper half)", paygrade: swagger.O_8},
	"VADM": RankTitlePaygrade{officer: true, title: "Vice Admiral", paygrade: swagger.O_9},
	"ADM":  RankTitlePaygrade{officer: true, title: "Admiral", paygrade: swagger.O_10},
	"WO1":  RankTitlePaygrade{officer: true, title: "Warrant Officer", paygrade: swagger.W_1},
	"CWO2": RankTitlePaygrade{officer: true, title: "Chief Warrant Officer", paygrade: swagger.W_2},
	"CWO3": RankTitlePaygrade{officer: true, title: "Chief Warrant Officer", paygrade: swagger.W_3},
	"CWO4": RankTitlePaygrade{officer: true, title: "Chief Warrant Officer", paygrade: swagger.W_4},
	"CWO5": RankTitlePaygrade{officer: true, title: "Chief Warrant Officer", paygrade: swagger.W_5},
}

var bareEnlistedRanks = map[string]RankTitlePaygrade{
	"AR":    RankTitlePaygrade{officer: false, title: "Airman Recruit", paygrade: swagger.E_1},
	"CR":    RankTitlePaygrade{officer: false, title: "Constructionman Recruit", paygrade: swagger.E_1},
	"FR":    RankTitlePaygrade{officer: false, title: "Fireman Recruit", paygrade: swagger.E_1},
	"HR":    RankTitlePaygrade{officer: false, title: "Hospital Recruit", paygrade: swagger.E_1},
	"SR":    RankTitlePaygrade{officer: false, title: "Seaman Recruit", paygrade: swagger.E_1},
	"AA":    RankTitlePaygrade{officer: false, title: "Airman Apprentice", paygrade: swagger.E_2},
	"CA":    RankTitlePaygrade{officer: false, title: "Constructionman Apprentice", paygrade: swagger.E_2},
	"FA":    RankTitlePaygrade{officer: false, title: "Fireman Apprentice", paygrade: swagger.E_2},
	"HA":    RankTitlePaygrade{officer: false, title: "Hospitalman Apprentice", paygrade: swagger.E_2},
	"SA":    RankTitlePaygrade{officer: false, title: "Seaman Apprentice", paygrade: swagger.E_2},
	"AN":    RankTitlePaygrade{officer: false, title: "Airman", paygrade: swagger.E_3},
	"CN":    RankTitlePaygrade{officer: false, title: "Constructionman", paygrade: swagger.E_3},
	"FN":    RankTitlePaygrade{officer: false, title: "Fireman", paygrade: swagger.E_3},
	"HN":    RankTitlePaygrade{officer: false, title: "Hospitalman", paygrade: swagger.E_3},
	"SN":    RankTitlePaygrade{officer: false, title: "Seaman", paygrade: swagger.E_3},
	"CMDCS": RankTitlePaygrade{officer: false, title: "Command Senior Chief Petty Officer", paygrade: swagger.E_8},
	"CMDCM": RankTitlePaygrade{officer: false, title: "Command Master Chief Petty Officer", paygrade: swagger.E_9},
	"FLTCM": RankTitlePaygrade{officer: false, title: "Fleet Master Chief Petty Officer", paygrade: swagger.E_9},
	"FORCM": RankTitlePaygrade{officer: false, title: "Force Master Chief Petty Officer", paygrade: swagger.E_9},
	"MCPON": RankTitlePaygrade{officer: false, title: "Navy Master Chief Petty Officer", paygrade: swagger.E_9},
}

var enlistedJobs = map[string]string{
	"AB":  "Aviation Boatswain's Mate",
	"ABE": "Aviation Boatswain's Mate",
	"ABF": "Aviation Boatswain's Mate",
	"ABH": "Aviation Boatswain's Mate",
	"AC":  "Air Traffic Controller",
	"AD":  "Aviation Machinist's Mate",
	"AE":  "Aviation Electrician's Mate",
	"AF":  "Aircraft Maintenanceman",
	"AG":  "Aerographer's Mate",
	"AM":  "Aviation Structural Mechanic",
	"AME": "Aviation Structural Mechanic",
	"AO":  "Aviation Ordinanceman",
	"AS":  "Aviation Support Equipment Technician",
	"AT":  "Aviation Electronics Technician",
	"AV":  "Avionics Technician",
	"AWF": "Naval Aircrewman",
	"AWO": "Naval Aircrewman",
	"AWR": "Naval Aircrewman",
	"AWS": "Naval Aircrewman",
	"AWV": "Naval Aircrewman",
	"AZ":  "Aviation Maintenance Administrationman",
	"BM":  "Boatswain's Mate",
	"BU":  "Builder",
	"CE":  "Construction Electrician",
	"CM":  "Construction Mechanic",
	"CS":  "Culinary Specialist",
	"CSS": "Culinary Specialist",
	"CTI": "Cryptologic Technician",
	"CTM": "Cryptologic Technician",
	"CTN": "Cryptologic Technician",
	"CTR": "Cryptologic Technician",
	"CTT": "Cryptologic Technician",
	"CU":  "Constructionman",
	"DC":  "Damage Controlman",
	"EA":  "Engineering Aide",
	"EM":  "Electrician's Mate",
	"EMN": "Electrician's Mate",
	"EN":  "Engineman",
	"EO":  "Equipment Operator",
	"EOD": "Explosive Ordinace Disposal",
	"EQ":  "Equipmentman",
	"ET":  "Electronics Technician",
	"ETN": "Electronics Technician",
	"ETR": "Electronics Technician",
	"ETV": "Electronics Technician",
	"FC":  "Fire Controlman",
	"FCA": "Fire Controlman",
	"FT":  "Fire Control Technician",
	"GM":  "Gunner's Mate",
	"GS":  "Gas Turbine Systems Technician",
	"GSE": "Gas Turbine Systems Technician",
	"GSM": "Gas Turbine Systems Technician",
	"HM":  "Hospital Corpsman",
	"HT":  "Hull Maintenance Technician",
	"IC":  "Interior Communications Electrician",
	"IS":  "Intelligence Specialist",
	"IT":  "Information Systems Technician",
	"ITS": "Information Systems Technician",
	"LN":  "Legalman",
	"LS":  "Logistics Specialist",
	"LSS": "Logistics Specialist",
	"MA":  "Master-At-Arms",
	"MC":  "Mass Communications Specialist",
	"MM":  "Machinist's Mate",
	"MMA": "Machinist's Mate",
	"MMN": "Machinist's Mate",
	"MMW": "Machinist's Mate",
	"MN":  "Mineman",
	"MR":  "Machinery Repairman",
	"MT":  "Missile Technician",
	"MU":  "Musician",
	"NC":  "Navy Counselor",
	"ND":  "Navy Diver",
	"OS":  "Operations Specialist",
	"PR":  "Aircrew Survival Equipmentman",
	"PS":  "Personnel Specialist",
	"QM":  "Quartermaster",
	"RP":  "Religious Program Specialist",
	"RT":  "Repair Technician",
	"SB":  "Special Warfare Boat Operator",
	"SH":  "Ship's Serviceman",
	"SO":  "Special Warfare Operator",
	"STG": "Sonar Technician",
	"STS": "Sonar Technician",
	"SW":  "Steelworker",
	"UC":  "Utilities Constructionman",
	"UT":  "Utilitiesman",
	"YN":  "Yeoman",
	"YNS": "Yeoman",
}

var enlistedRates = map[string]rate{
	"AR": rate{prefix: "", suffix: " Airman Recruit", paygrade: swagger.E_1},
	"CR": rate{prefix: "", suffix: " Constructionman Recruit", paygrade: swagger.E_1},
	"FR": rate{prefix: "", suffix: " Fireman Recruit", paygrade: swagger.E_1},
	"SR": rate{prefix: "", suffix: " Seaman Recruit", paygrade: swagger.E_1},
	"AA": rate{prefix: "", suffix: " Airman Apprentice", paygrade: swagger.E_2},
	"CA": rate{prefix: "", suffix: " Constructionman Apprentice", paygrade: swagger.E_2},
	"FA": rate{prefix: "", suffix: " Fireman Apprentice", paygrade: swagger.E_2},
	"SA": rate{prefix: "", suffix: " Seaman Apprentice", paygrade: swagger.E_2},
	"AN": rate{prefix: "", suffix: " Airman", paygrade: swagger.E_3},
	"CN": rate{prefix: "", suffix: " Constructionman", paygrade: swagger.E_3},
	"FN": rate{prefix: "", suffix: " Fireman", paygrade: swagger.E_3},
	"SN": rate{prefix: "", suffix: " Seaman", paygrade: swagger.E_3},
	"3":  rate{prefix: "", suffix: " Third Class", paygrade: swagger.E_4},
	"2":  rate{prefix: "", suffix: " Second Class", paygrade: swagger.E_5},
	"1":  rate{prefix: "", suffix: " First Class", paygrade: swagger.E_6},
	"C":  rate{prefix: "Chief ", suffix: "", paygrade: swagger.E_7},
	"CS": rate{prefix: "Senior Chief ", suffix: "", paygrade: swagger.E_8},
	"CM": rate{prefix: "Master Chief ", suffix: "", paygrade: swagger.E_9},
}

// RankFromAbbreviation returns the Rank corresponding to the provided abbreviation
func RankFromAbbreviation(abbr string) *RankTitlePaygrade {
	// try the officer and simple enlisted conversions first; they only take one step
	r, ok := officerRanks[abbr]
	if ok {
		return &r
	}

	r, ok = bareEnlistedRanks[abbr]
	if ok {
		return &r
	}

	rank := new(RankTitlePaygrade)

	for key, rate := range enlistedRates {
		if !strings.HasSuffix(abbr, key) {
			continue
		}

		rank.officer = false
		rank.paygrade = rate.paygrade

		job := strings.TrimSuffix(abbr, key)
		title, ok := enlistedJobs[job]
		if ok {
			rank.title = rate.prefix + title + rate.suffix
		} else {
			// Uh-oh, unknown job! Just use the abbreviation as the title
			rank.title = abbr
		}
		return rank
	}

	// Uh-oh, unknown abbreviation!
	return rank
}
