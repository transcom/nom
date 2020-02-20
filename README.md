# Navy Orders Muncher

**nom**, the Navy Orders Muncher, ingests CSV files containing Navy Orders and excretes database updates to the move.mil Orders API.

## License Information

Works created by U.S. Federal employees as part of their jobs typically are not eligible for copyright in the United
States. In places where the contributions of U.S. Federal employees are not eligible for copyright, this work is in
the public domain. In places where it is eligible for copyright, such as some foreign jurisdictions, the remainder of
this work is licensed under [the MIT License](https://opensource.org/licenses/MIT), the full text of which is included
in the [LICENSE.txt](./LICENSE.txt) file in this repository.

## Usage

`$ nom <csv input file>`

## Building

Building is easy! Once you have the dependencies, run

`$ make`

### Dependencies

**nom** is written in [Go](https://golang.org/). Aside from go, you will need:

- [GNU Make](https://www.gnu.org/software/make/)
- [curl](https://curl.haxx.se/)

Acquiring and installing these is left as an exercise for the reader.

## Input format

**nom** accepts comma-delimited CSV files and expects the following columns:

| Column Name | Description |
| ----------- | ----------- |
| Ssn (obligation) | If 9 digits, Social Security Number<br>If 10 digits, EDIPI |
| TAC | Household Goods (HHG) Transportation Account Code (TAC) |
| Order Create/Modification Date | Orders date, in Excel date format (Day 1 = Dec 31, 1899) |
| Order Modification Number | Number of modifications made by an Orders Writing System, such as EAIS, OAIS, or NMCMPS. |
| Obligation Modification Number | Number of modifications made manually via POEMS. |
| Obligation Status Code | **D**: Cancel Obligation, effectively rescinding these Orders<br>**N**: Initial Mod - amended Orders<br>**P**: Initial Obligation - new Orders|
| Obligation Multi-leg Code | Indicates whether either endpoint is TDY.<br>**0** - Perm to Perm<br>**1** - Perm to Temp<br>**5** - Temp to Temp<br>**9** - Temp to Perm |
| CIC Purpose Information Code (OBLGTN) | Purpose of the Orders, maps to Orders type |
| Paygrade | Three-character DoD Paygrade, e.g., E05, W02, O10 |
| Rank Classification  Description | The Navy rank or rating |
| Service Member Name | The sailor's name, in the format LASTNAME,FIRSTNAME (optional MI) (optional suffix) |
| Detach UIC | Unit Identification Code (UIC) of the detaching activity |
| Detach UIC Home Port | Home port of the detaching activity |
| Detach UIC City Name | Detaching activity city |
| Detach State Code | Detaching activity state |
| Detach Country Code | Detaching activity country |
| Ultimate Estimated Arrival Date | Report No Later Than Date |
| Ultimate UIC | Unit Identification Code (UIC) of the ultimate activity |
| Ultimate UIC Home Port | Home port of the ultimate activity |
| Ultimate UIC City Name | Ultimate activity city |
| Ultimate State Code | Ultimate activity state |
| Ultimate Country Code | Ultimate activity country |
| Entitlement Indicator | If 'Y', then this is a 'Cost Order' with obligated moving expenses. If 'N', then this is a 'No Cost Order'. |
| Count of Dependents Participating in Move (STATIC) | Number of sailor's dependents; needed to determine the correct weight entitlement |
| Count of Intermediate Stops (STATIC) | Number of intermediate activities. If greater than 0, then this move has TDY en route. |
| Primary SDN | The Commercial Travel (CT) Standard Document Number (SDN), which **nom** uses as the unique Orders number |

Columns that do not start with the above headers are ignored.

### Orders number

On printed Navy Orders, the BUPERS Orders number is originally formatted as "`<Order Control Number> <SSN>`", for example, "`3108 000-12-3456`". It would be unique (because of the SSN), except that it’s possible for a set of orders to be cut on the same day 10 years later for the same sailor, resulting in a collision.

Because the BUPERS Orders Number contains PII (the SSN) and could potentially not be unique (because it only allows a single digit for the year), **nom** uses the Primary SDN (aka the Commercial Travel SDN) instead. For what it's worth, Marine Corps orders also use the CT SDN as the unique Orders number.

### Modification number interpretation

The Orders API has a sequence number to indicate the chronology of amendments to a set of Orders. The input, however, has two modification number fields, which track the modification count from different systems. Fortunately, these two fields increment atomically, and never decrement.

Therefore, the sequence number is simply the sum of these two numbers.

### Orders type

To determine the effective orders type, lookup the CIC Purpose Information Code and community (enlisted or officer) in the following table.

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

## Output

nom uploads the Orders it reads to the move.mil Orders API.
