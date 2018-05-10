package main

import "strings"

// Rank correlates a title with its DoD paygrade and indicates whether the rank belongs to a commissioned or warrant officer
type Rank struct {
	officer  bool
	title    string
	paygrade string
}

type rate struct {
	prefix   string
	suffix   string
	paygrade string
}

var officerRanks = map[string]Rank{
	"OC":   Rank{officer: true, title: "Officer Candidate", paygrade: "e-5"},
	"ENS":  Rank{officer: true, title: "Ensign", paygrade: "o-1"},
	"LTJG": Rank{officer: true, title: "Lieutenant Junior Grade", paygrade: "o-2"},
	"LT":   Rank{officer: true, title: "Lieutenant", paygrade: "o-3"},
	"LCDR": Rank{officer: true, title: "Lieutenant Commander", paygrade: "o-4"},
	"CDR":  Rank{officer: true, title: "Commander", paygrade: "o-5"},
	"CAPT": Rank{officer: true, title: "Captain", paygrade: "o-6"},
	"RDML": Rank{officer: true, title: "Rear Admiral (lower half)", paygrade: "o-7"},
	"RDMU": Rank{officer: true, title: "Rear Admiral (upper half)", paygrade: "o-8"},
	"VADM": Rank{officer: true, title: "Vice Admiral", paygrade: "o-9"},
	"ADM":  Rank{officer: true, title: "Admiral", paygrade: "o-10"},
	"WO1":  Rank{officer: true, title: "Warrant Officer", paygrade: "w-1"},
	"CWO2": Rank{officer: true, title: "Chief Warrant Officer", paygrade: "w-2"},
	"CWO3": Rank{officer: true, title: "Chief Warrant Officer", paygrade: "w-3"},
	"CWO4": Rank{officer: true, title: "Chief Warrant Officer", paygrade: "w-4"},
	"CWO5": Rank{officer: true, title: "Chief Warrant Officer", paygrade: "w-5"},
}

var bareEnlistedRanks = map[string]Rank{
	"AR":    Rank{officer: false, title: "Airman Recruit", paygrade: "e-1"},
	"CR":    Rank{officer: false, title: "Constructionman Recruit", paygrade: "e-1"},
	"FR":    Rank{officer: false, title: "Fireman Recruit", paygrade: "e-1"},
	"HR":    Rank{officer: false, title: "Hospital Recruit", paygrade: "e-1"},
	"SR":    Rank{officer: false, title: "Seaman Recruit", paygrade: "e-1"},
	"AA":    Rank{officer: false, title: "Airman Apprentice", paygrade: "e-2"},
	"CA":    Rank{officer: false, title: "Constructionman Apprentice", paygrade: "e-2"},
	"FA":    Rank{officer: false, title: "Fireman Apprentice", paygrade: "e-2"},
	"HA":    Rank{officer: false, title: "Hospitalman Apprentice", paygrade: "e-2"},
	"SA":    Rank{officer: false, title: "Seaman Apprentice", paygrade: "e-2"},
	"AN":    Rank{officer: false, title: "Airman", paygrade: "e-3"},
	"CN":    Rank{officer: false, title: "Constructionman", paygrade: "e-3"},
	"FN":    Rank{officer: false, title: "Fireman", paygrade: "e-3"},
	"HN":    Rank{officer: false, title: "Hospitalman", paygrade: "e-3"},
	"SN":    Rank{officer: false, title: "Seaman", paygrade: "e-3"},
	"CMDCS": Rank{officer: false, title: "Command Senior Chief Petty Officer", paygrade: "e-8"},
	"CMDCM": Rank{officer: false, title: "Command Master Chief Petty Officer", paygrade: "e-9"},
	"FLTCM": Rank{officer: false, title: "Fleet Master Chief Petty Officer", paygrade: "e-9"},
	"FORCM": Rank{officer: false, title: "Force Master Chief Petty Officer", paygrade: "e-9"},
	"MCPON": Rank{officer: false, title: "Navy Master Chief Petty Officer", paygrade: "e-9"},
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
	"AR": rate{prefix: "", suffix: " Airman Recruit", paygrade: "e-1"},
	"CR": rate{prefix: "", suffix: " Constructionman Recruit", paygrade: "e-1"},
	"FR": rate{prefix: "", suffix: " Fireman Recruit", paygrade: "e-1"},
	"SR": rate{prefix: "", suffix: " Seaman Recruit", paygrade: "e-1"},
	"AA": rate{prefix: "", suffix: " Airman Apprentice", paygrade: "e-2"},
	"CA": rate{prefix: "", suffix: " Constructionman Apprentice", paygrade: "e-2"},
	"FA": rate{prefix: "", suffix: " Fireman Apprentice", paygrade: "e-2"},
	"SA": rate{prefix: "", suffix: " Seaman Apprentice", paygrade: "e-2"},
	"AN": rate{prefix: "", suffix: " Airman", paygrade: "e-3"},
	"CN": rate{prefix: "", suffix: " Constructionman", paygrade: "e-3"},
	"FN": rate{prefix: "", suffix: " Fireman", paygrade: "e-3"},
	"SN": rate{prefix: "", suffix: " Seaman", paygrade: "e-3"},
	"3":  rate{prefix: "", suffix: " Third Class", paygrade: "e-4"},
	"2":  rate{prefix: "", suffix: " Second Class", paygrade: "e-5"},
	"1":  rate{prefix: "", suffix: " First Class", paygrade: "e-6"},
	"C":  rate{prefix: "Chief ", suffix: "", paygrade: "e-7"},
	"CS": rate{prefix: "Senior Chief ", suffix: "", paygrade: "e-8"},
	"CM": rate{prefix: "Master Chief ", suffix: "", paygrade: "e-9"},
}

// RankFromAbbreviation returns the Rank corresponding to the provided abbreviation
func RankFromAbbreviation(abbr string) *Rank {
	// try the officer and simple enlisted conversions first; they only take one step
	r, ok := officerRanks[abbr]
	if ok {
		return &r
	}

	r, ok = bareEnlistedRanks[abbr]
	if ok {
		return &r
	}

	rank := new(Rank)

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
