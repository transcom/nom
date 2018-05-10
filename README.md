# Navy Orders Muncher

**nom**, the Navy Orders Muncher, ingests CSV files containing Navy Orders and excretes database updates to the move.mil Orders API.

# Usage

`$ nom <csv input> `

# Input format

**nom** accepts comma-delimited CSV files and expects the following columns:

| Column Name | Description |
| ----------- | ----------- |
| EMPLID | If 9 digits, Social Security Number<br>If 10 digits, EDIPI |
| N_ORDER_CNTL_NBR | Julian day of the year (1-366) followed by last digit of the year (e.g., 8 for 2018). |
| N_ORD_DT | Orders date |
| N_MOD_NBR | Number of modifications made by an Orders Writing System, such as EAIS, OAIS, or NMCMPS. |
| N_MOD_NUM	| Number of modifications made manually via POEMS. |
| N_OBLG_STATUS	| **D**: Cancel Obligation, effectively rescinding these Orders<br>**N**: Initial Mod - amended Orders<br>**P**: Initial Obligation - new Orders|
| N_OBLG_LEG_NBR | Indicates whether any TDY is  included in these Orders.<br>**0** - Perm to Perm<br>**1** - Perm to Temp<br>**5** - Temp to Temp<br>**9** - Temp to Perm |
| N_CIC_PURP | Purpose of the Orders, maps to Orders type |
| N_RATE_RANK | The Navy abbreviation of the rank / title |
| NAME | The sailor's name, in the format LASTNAME,FIRSTNAME (optional MI) (optional suffix) |
| N_UIC_DETACH | Unit Identification Code (UIC) of the detaching activity |
| N_DET_HPORT | Home port of the detaching activity |
| N_PDS_CITY | Detaching activity city |
| N_PDS_STATE | Detaching activity state |
| N_PDS_CNTRY | Detaching activity country |
| N_EST_ARRIVAL_DT | Despite the name, this is the Report No Later Than Date |
| N_UIC_ULT_DTY_STA	| Unit Identification Code (UIC) of the ultimate activity |
| N_ULT_HPORT | Home port of the ultimate activity |
| N_ULT_CITY | Ultimate activity city |
| N_ULT_STATE | Ultimate activity state |
| N_ULT_CNTRY | Ultimate activity country |
| N_NON_ENT_IND | If 'Y', then this is a 'Cost Order' with obligated moving expenses. If 'N', then this is a 'No Cost Order', i.e., a PCA w/o PCS (Permanent Change of Assignment without Permanent Change of Station), and has no moving expenses. |
| N_NUM_DEPN | Despite the name, this column contains either 'Y' or 'N' to indicate whether the sailor has dependents |
| TAC_SDN | Household Goods (HHG) Standard Document Number (SDN), which also incorporates the HHG Transportation Account Code (TAC) as its last four characters |

Columns that do not start with the above headers are ignored.

## Orders number
On printed Navy Orders, the BUPERS Orders number is originally formatted as "`<N_ORDER_CNTL_NBR> <EMPLID>`", for example, "`3108 000-12-3456`". It would be unique (because of the SSN), except that it’s possible for a set of orders to be cut on the same day 10 years later for the same sailor, resulting in a collision.

Because the BUPERS Orders Number contains PII (the SSN) and could potentially not be unique (because it only allows a single digit for the year), the Orders API expects the Orders number to be reformatted as "`{Julian day}{4 digit year} {EDIPI}`", for example, "`3102018 0123456789`". The full 4 digit year can be found in the N_ORD_DT field of the first revision, while the EDIPI can be retrieved using DMDC's Identity Web Services.

## Modification number interpretation
The Orders API has a sequence number to indicate the chronology of amendments to a set of Orders. The input, however, has two modification number fields: `N_MOD_NUM` and `N_MOD_NBR`. Fortunately, these two fields increment atomically, and never decrement.

Therefore, the sequence number is simply the sum of `N_MOD_NUM` and `N_MOD_NBR`.

## Orders type
To determine the effective orders type, lookup the purpose (`N_CIC_PURP`) and community (enlisted or officer) in the following table.

| N_CIC_PURP | Enlisted / Officer | Description | Effective Orders Type |
| ---------- | ------------------ | ----------- | --------------------- |
| 0 | Officer | IPCOT In-place consecutive overseas travel | ipcot |
| 1 | Enlisted | IPCOT In-place consecutive overseas travel | ipcot |
| 2 | Officer | Accession Travel | accession |
| 3 | Officer | Training Travel | training |
| 4 | Officer | Operational Travel | operational |
| 5 | Officer | Separation Travel | separation |
| 6 | Officer | Organized Unit/Homeport Change | unit-move |
| 7 | Officer | Emergency Non-member Evac | emergency-evac |
| 8 | Enlisted | Overseas Tour Extension Incentive Program (OTEIP) | oteip |
| 9 | Enlisted | NAVCAD (Naval Cadet) Training | training |
| A | Enlisted | Accession Travel Recruits | accession |
| B | Enlisted | Non-recruit Accession Travel | accession |
| C | Enlisted | Training Travel | training |
| D | Enlisted | Operational Travel | operational |
| E | Enlisted | Separation Travel | separation |
| F | Enlisted | Organized Unit/Homeport Change | unit-move |
| G | Enlisted | Midshipman Accession Travel | accession |
| H | Both | Special Purpose Reimbursable | special-purpose |
| I | Enlisted | NAVCAD(Naval Cadet) Accession | accession |
| J | Enlisted | Accession Travel Recruits | accession |
| K | Enlisted | Non-recruit Accession Travel | accession |
| L | Enlisted | Training Travel | training |
| M | Enlisted | Rotational Travel | rotational |
| N | Enlisted | Separation Travel | separation |
| O | Enlisted | Organized Unit/Homeport Change | unit-move |
| P | Enlisted | Midshipman Separation Travel | separation |
| Q | Officer | Misc. Rotational Non-member | rotational |
| R | Enlisted | Misc. Operational Non-member | operational |
| S | Officer | Accession Travel | accession |
| T | Officer | Training Travel | training |
| U | Officer | Rotational Travel | rotational |
| V | Officer | Separation Travel | separation |
| W | Officer | Organized Unit/Homeport Change | unit-move |
| X | Enlisted | EMERGENCY NON-MEMBER EVACS | emergency-evac |
| X | Officer | Misc. Rotational Non-member | rotational |
| Y | Enlisted | Misc. Rotational Non-member | rotational |
| Z | Enlisted | NAVCAD(Naval Cadet) Separation | separation |

## Ranks
The `N_RATE_RANK` column contains Navy rank (officer) and rate (enlisted) abbreviations. These abbreviations need to be translated to titles and DoD pay grades.

### Enlisted Rates
To determine the title and pay grade for enlisted rates, if the abbreviation does not have a simple translation, then match the rating suffix to first match the stem of the abbreviation to the job title, and then determine the prefix or suffix of the title using the remainder of the abbreviation in the rates table.

#### Bare enlisted rates and paygrades
The most senior enlisted rates have a simple translation from abbreviation to title and pay grade.

The lowest enlisted paygrades without ratings are also easy to translate.

| Abbreviation | Title | DoD Pay Grade |
| ------------ | ----- | ------------- |
| AR | Airman Recruit | E-1 |
| CR | Constructionman Recruit | E-1 |
| FR | Fireman Recruit | E-1 |
| HR | Hospital Recruit | E-1 |
| SR | Seaman Recruit | E-1 |
| AA | Airman Apprentice | E-2 |
| CA | Constructionman Apprentice | E-2 |
| FA | Fireman Apprentice | E-2 |
| HA | Hospitalman Apprentice | E-2 |
| SA | Seaman Apprentice | E-2 |
| AN | Airman | E-3 |
| CN | Constructionman | E-3 |
| FN | Fireman | E-3 |
| HN | Hospitalman | E-3 |
| SN | Seaman | E-3 |
| CMDCS | Command Senior Chief Petty Officer | E-8 |
| CMDCM | Command Master Chief Petty Officer | E-9 |
| FLTCM | Fleet Master Chief Petty Officer | E-9 |
| FORCM | Force Master Chief Petty Officer | E-9 |
| MCPON | Navy Master Chief Petty Officer | E-9 |

#### Enlisted Ratings
| Abbreviation | Title |
| ----------- | ------ |
| AB<br>ABE<br>ABF<br>ABH | Aviation Boatswain's Mate |
| AC | Air Traffic Controller |
| AD | Aviation Machinist's Mate |
| AE | Aviation Electrician's Mate |
| AF | Aircraft Maintenanceman |
| AG | Aerographer's Mate |
| AM<br>AME | Aviation Structural Mechanic |
| AO | Aviation Ordinanceman |
| AS | Aviation Support Equipment Technician |
| AT | Aviation Electronics Technician |
| AV | Avionics Technician |
| AWO<br>AWF<br>AWV<br>AWS<br>AWR | Naval Aircrewman |
| AZ | Aviation Maintenance Administrationman |
| BM | Boatswain's Mate |
| BU | Builder |
| CE | Construction Electrician |
| CM | Construction Mechanic |
| CS<br>CSS | Culinary Specialist |
| CTI<br>CTM<br>CTN<br>CTR<br>CTT | Cryptologic Technician |
| CU | Constructionman |
| DC | Damage Controlman |
| EA | Engineering Aide |
| EM<br>EMN | Electrician's Mate |
| EN | Engineman |
| EO | Equipment Operator |
| EOD | Explosive Ordinance Disposal |
| EQ | Equipmentman |
| ET<br>ETN<br>ETV<br>ETR | Electronics Technician |
| FC<br>FCA | Fire Controlman |
| FT | Fire Control Technician |
| GM | Gunner's Mate |
| GS<br>GSE<br>GSM | Gas Turbine Systems Technician |
| HM | Hospital Corpsman |
| HT | Hull Maintenance Technician |
| IC | Interior Communications Electrician |
| IS | Intelligence Specialist |
| IT<br>ITS | Information Systems Technician |
| LN | Legalman |
| LS<br>LSS | Logistics Specialist |
| MA | Master-At-Arms |
| MC | Mass Communications Specialist |
| MM<br>MMA<br>MMN<br>MMW | Machinist's Mate |
| MN | Mineman |
| MR | Machinery Repairman |
| MT | Missile Technician |
| MU | Musician |
| NC | Navy Counselor |
| ND | Navy Diver |
| OS | Operations Specialist |
| PR | Aircrew Survival Equipmentman |
| PS | Personnel Specialist |
| QM | Quartermaster |
| RP | Religious Program Specialist |
| RT | Repair Technician |
| SB | Special Warfare Boat Operator |
| SH | Ship's Serviceman |
| SO | Special Warfare Operator |
| STG<br>STS | Sonar Technician |
| SW | Steelworker |
| UC | Utilities Constructionman |
| UT | Utilitiesman |
| YN<br>YNS | Yeoman |

#### Enlisted paygrades
Paygrades E-1 through E-3 can also have a rating abbreviation preceding their paygrade symbol if they are graduates of Class "A" schools; have received the rating designation in a previous enlistment; are assigned to a billet in that specialty as a striker; have passed an advancement examination and not been selected for advancement for reasons of numeric limitations on advancements; or have been reduced in rate because of punishment.

| Abbreviation | Title prefix | Title suffix | DoD Pay Grade |
| ------------ | ------------ | ------------ | -------------- |
| AR | | Airman Recruit | E-1 |
| CR | | Constructionman Recruit | E-1 |
| FR | | Fireman Recruit | E-1 |
| SR | | Seaman Recruit | E-1 |
| AA | | Airman Apprentice | E-2 |
| CA | | Constructionman Apprentice | E-2 |
| FA | | Fireman Apprentice | E-2 |
| SA | | Seaman Apprentice | E-2 |
| AN | | Airman | E-3 |
| CN | | Constructionman | E-3 |
| FN | | Fireman | E-3 |
| SN | | Seaman | E-3 |
| 3 | | Third Class | E-4 |
| 2 | | Second Class | E-5 |
| 1 | | First Class | E-6 |
| C | Chief | | E-7 |
| CS | Senior Chief | | E-8 |
| CM | Master Chief | | E-9 |

### Officer Ranks
| Abbreviation | Title | DoD Pay Grade |
| ------------ | ----- | ------------- |
| OC | Officer Candidate | E-5 |
| ENS | Ensign | O-1 |
| LTJG | Lieutenant Junior Grade | O-2 |
| LT | Lieutenant | O-3 |
| LCDR | Lieutenant Commander | O-4 |
| CDR | Commander | O-5	|
| CAPT | Captain | O-6 |
| RDML | Rear Admiral (lower half) | O-7 |
| RDMU | Rear Admiral (upper half) | O-8 |
| VADM | Vice Admiral | O-9 |
| ADM | Admiral | O-10 |

### Warrant Officer Ranks
| Abbreviation | Title | DoD Pay Grade |
| ----------- | ----- | ------------- |
| WO1 | Warrant Officer | W-1 |
| CWO2 | Chief Warrant Officer | W-2 |
| CWO3 | Chief Warrant Officer | W-3 |
| CWO4 | Chief Warrant Officer | W-4 |
| CWO5 | Chief Warrant Officer | W-5 |

# Output
nom uploads the Orders it reads to the move.mil Orders API.
