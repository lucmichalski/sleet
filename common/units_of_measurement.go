package common

import "strings"

// ISO Units of Measure codes (http://t3.apptrix.com/syteline/Language/en-US/fields/i/iso_um_ums.htm)
var unitOfMeasurementToCode = map[string]string{
	"115 kilogram drum": "16",
	"100 pound drum":    "17",
	"55 gallon drum":    "18",
	"20 foot container": "20",
	"40 foot container": "21",
	"bin":               "2W",
	"milliroentgen":     "2Y",
	"milliliters per square centimeter second": "35",
	"cubic feet per minute per square foot":    "36",
	"ounces per square foot":                   "37",
	"ounces per square foot per 0.01 inch":     "38",
	"megajoule":                                "3B",
	"milliliter per second":                    "40",
	"milliliter per minute":                    "41",
	"super bulk bag":                           "43",
	"500 kilogram bulk bag":                    "44",
	"300 kilogram bulk bag":                    "45",
	"25 kilogram bulk bag":                     "46",
	"50 pound bag":                             "47",
	"bulk car load":                            "48",
	"20-pack":                                  "4E",
	"100-pack":                                 "4F",
	"microliter":                               "4G",
	"micrometer":                               "4H",
	"milligrams per hour":                      "4M",
	"megabecquerel":                            "4N",
	"pascal":                                   "4S",
	"pounds per hour":                          "4U",
	"cubic meter per hour":                     "4V",
	"ton per hour":                             "4W",
	"kiloliter per hour":                       "4X",
	"actual kilograms":                         "50",
	"net kilograms":                            "58",
	"parts per million":                        "59",
	"batch":                                    "5B",
	"standard cubic foot":                      "5I",
	"percent weight":                           "60",
	"parts per billion":                        "61",
	"bulk pack":                                "AB",
	"centigram":                                "AF",
	"anti-hemophilic factor (ahf) units":       "AQ",
	"suppository":                              "AR",
	"ocular insert system":                     "AU",
	"capsule":                                  "AV",
	"powder-filled vials":                      "AW",
	"british thermal units (btus) per pound":   "AZ",
	"british thermal units (btus) per cubic foot": "B0",
	"cycles":                     "B7",
	"bale":                       "BA",
	"bucket":                     "BC",
	"bundle":                     "BD",
	"bag":                        "BG",
	"band":                       "BJ",
	"book":                       "BK",
	"bulk":                       "BN",
	"bottle":                     "BO",
	"bushel":                     "BU",
	"base weight":                "BW",
	"box":                        "BX",
	"british thermal unit (btu)": "BY",
	"million btu's":              "BZ",
	"composite product pounds (total weight)": "C1",
	"centiliter":                         "C3",
	"carload":                            "C4",
	"cell":                               "C6",
	"cubic decimeter":                    "C8",
	"case":                               "CA",
	"carboy":                             "CB",
	"cubic centimeter":                   "CC",
	"carat":                              "CD",
	"centigrade, celsius":                "CE",
	"cubic feet":                         "CF",
	"card":                               "CG",
	"container":                          "CH",
	"cubic inches":                       "CI",
	"cone":                               "CJ",
	"connector":                          "CK",
	"cylinder":                           "CL",
	"centimeter":                         "CM",
	"can":                                "CN",
	"cubic meters (net)":                 "CO",
	"crate":                              "CP",
	"cartridge":                          "CQ",
	"cubic meter":                        "CR",
	"cassette":                           "CS",
	"carton":                             "CT",
	"cup":                                "CU",
	"cover":                              "CV",
	"hundred pounds (cwt)":               "CW",
	"coil":                               "CX",
	"cubic yard":                         "CY",
	"square decimeter":                   "D3",
	"kilogram per square centimeter":     "D5",
	"dyne per square centimeter":         "D9",
	"days":                               "DA",
	"dry pounds":                         "DB",
	"disk (disc)":                        "DC",
	"degree":                             "DD",
	"deal":                               "DE",
	"dram":                               "DF",
	"decigram":                           "DG",
	"dispenser":                          "DI",
	"decagram":                           "DJ",
	"kilometers":                         "DK",
	"deciliter":                          "DL",
	"decimeter":                          "DM",
	"dollars, u.s.":                      "DO",
	"dozen pair":                         "DP",
	"drum":                               "DR",
	"dry ton":                            "DT",
	"dozen":                              "DZ",
	"inches, fraction--average":          "E3",
	"inches, fraction--minimum":          "E4",
	"inches, fraction--actual":           "E5",
	"inches, decimal--average":           "E7",
	"inches, decimal--actual":            "E8",
	"english, (feet, inches)":            "E9",
	"each":                               "EA",
	"each per month":                     "EC",
	"inches, decimal--nominal":           "ED",
	"inches, fraction-nominal":           "EF",
	"double-time hours":                  "EG",
	"inches, decimal-minimum":            "EM",
	"eleven pack":                        "EP",
	"envelope":                           "EV",
	"feet, inches and fraction":          "EX",
	"feet, inches and decimal":           "EY",
	"feet and decimal":                   "EZ",
	"thousand cubic feet per day":        "F1",
	"international unit":                 "F2",
	"minim":                              "F4",
	"mol":                                "F5",
	"fibers per cubic centimeter of air": "F9",
	"fahrenheit":                         "FA",
	"fields":                             "FB",
	"1000 cubic feet":                    "FC",
	"million particles per cubic foot":   "FD",
	"hundred cubic meters":               "FF",
	"transdermal patch":                  "FG",
	"micromolar":                         "FH",
	"sizing factor":                      "FJ",
	"fibers":                             "FK",
	"flake ton":                          "FL",
	"million cubic feet":                 "FM",
	"fluid ounce":                        "FO",
	"pounds per sq. ft.":                 "FP",
	"foot":                               "FT",
	"fluid ounce (imperial)":             "FZ",
	"u.s. gallons per minute":            "G2",
	"imperial gallons per minute":        "G3",
	"gigabecquerel":                      "G4",
	"gallon":                             "GA",
	"gallons/day":                        "GB",
	"grams per 100 grams":                "GC",
	"gross barrels":                      "GD",
	"pounds per gallon":                  "GE",
	"grams per 100 centimeters":          "GF",
	"great gross (dozen gross)":          "GG",
	"half gallon":                        "GH",
	"imperial gallons":                   "GI",
	"grams per milliliter":               "GJ",
	"grams per kilogram":                 "GK",
	"grams per liter":                    "GL",
	"grams per sq. meter":                "GM",
	"gross gallons":                      "GN",
	"milligrams per square meter":        "GO",
	"milligrams per cubic meter":         "GP",
	"micrograms per cubic meter":         "GQ",
	"gram":                               "GR",
	"gross":                              "GS",
	"gross kilogram":                     "GT",
	"gauss per oersteds":                 "GU",
	"gallons per thousand cubic feet":    "GW",
	"grain":                              "GX",
	"gross yard":                         "GY",
	"gage systems":                       "GZ",
	"half liter":                         "H2",
	"hectoliter":                         "H4",
	"hundred boxes":                      "HB",
	"hundred count":                      "HC",
	"half dozen":                         "HD",
	"hundredth of a carat":               "HE",
	"hundred feet":                       "HF",
	"hectogram":                          "HG",
	"hundred cubic feet":                 "HH",
	"hundred sheets":                     "HI",
	"hundred kilograms":                  "HK",
	"hundred feet - linear":              "HL",
	"millimeters of mercury":             "HN",
	"hundred troy ounces":                "HO",
	"millimeter h20":                     "HP",
	"hours":                              "HR",
	"hundred square feet":                "HS",
	"half hour":                          "HT",
	"hundred":                            "HU",
	"hundred weight (short)":             "HV",
	"hundred weight (long)":              "HW",
	"hundred yards":                      "HY",
	"inch pound":                         "IA",
	"inhaler":                            "IH",
	"impressions":                        "IM",
	"inch":                               "IN",
	"joule per kilogram":                 "J2",
	"joule per kelvin":                   "JE",
	"joule per gram":                     "JG",
	"mega joule per kilogram":            "JK",
	"megajoule/cubic meter":              "JM",
	"joint":                              "JO",
	"jar":                                "JR",
	"kilogram":                           "KG",
	"kiloroentgen":                       "KR",
	"1000 pounds per square inch":        "KS",
	"kilograms per millimeter":           "KW",
	"milliliters per kilogram":           "KX",
	"liters per minute":                  "L2",
	"pounds per cubic inch":              "LA",
	"pound":                              "LB",
	"linear centimeter":                  "LC",
	"linear foot":                        "LF",
	"linear inch":                        "LI",
	"linear meter":                       "LM",
	"liquid pounds":                      "LP",
	"liters per day":                     "LQ",
	"layer(s)":                           "LR",
	"lump sum":                           "LS",
	"liter":                              "LT",
	"linear yards per pound":             "LX",
	"linear yard":                        "LY",
	"milligrams per liter":               "M1",
	"millimeter-actual":                  "M2",
	"microcurie":                         "M5",
	"millibar":                           "M6",
	"micro inch":                         "M7",
	"mega pascals":                       "M8",
	"million british thermal units per one thousand cubic feet": "M9",
	"millimeter-nominal":              "MB",
	"microgram":                       "MC",
	"air dry metric ton":              "MD",
	"milligram":                       "ME",
	"milligram per sq. ft. per side":  "MF",
	"metric gross ton":                "MG",
	"microns (micrometers)":           "MH",
	"metric":                          "MI",
	"minutes":                         "MJ",
	"milligrams per square inch":      "MK",
	"milliliter":                      "ML",
	"millimeter":                      "MM",
	"metric net ton":                  "MN",
	"months":                          "MO",
	"metric ton":                      "MP",
	"1000 meters":                     "MQ",
	"meter":                           "MR",
	"square millimeter":               "MS",
	"metric long ton":                 "MT",
	"millicurie":                      "MU",
	"metric ton kilograms":            "MW",
	"mixed":                           "MX",
	"millimeter-average":              "MY",
	"milligrams per kilogram":         "NA",
	"barge":                           "NB",
	"car":                             "NC",
	"net barrels":                     "ND",
	"net liters":                      "NE",
	"net gallons":                     "NG",
	"net imperial gallons":            "NI",
	"parts per thousand":              "NX",
	"pounds per air-dry metric ton":   "NY",
	"ounces per square yard":          "ON",
	"two pack":                        "OP",
	"ounce - av":                      "OZ",
	"percent":                         "P1",
	"pounds per foot":                 "P2",
	"piece":                           "PC",
	"pad":                             "PD",
	"pounds equivalent":               "PE",
	"pallet (lift)":                   "PF",
	"pounds gross":                    "PG",
	"pounds-percentage":               "PM",
	"pounds net":                      "PN",
	"pounds per inch of length":       "PO",
	"plate":                           "PP",
	"pair":                            "PR",
	"pounds per sq. inch":             "PS",
	"pint":                            "PT",
	"mass pounds":                     "PU",
	"pounds per inch of width":        "PW",
	"pint, imperial":                  "PX",
	"peck, dry u.s.":                  "PY",
	"peck, dry imperial":              "PZ",
	"pint u.s.\r":                     "Q2",
	"quire":                           "QR",
	"quart, dry u.s.":                 "QS",
	"quart":                           "QT",
	"quart, imperial":                 "QU",
	"thousand cubic meters":           "R9",
	"rack":                            "RA",
	"radian":                          "RB",
	"rod (area) - 16.25 square yards": "RC",
	"rod (length) - 5.5 yards":        "RD",
	"reel":                            "RE",
	"roll-metric measure":             "RK",
	"roll":                            "RL",
	"ream":                            "RM",
	"ream-metric measure":             "RN",
	"round":                           "RO",
	"pounds per ream":                 "RP",
	"sixty-fourths of an inch":        "S5",
	"square centimeter":               "SC",
	"solid pounds":                    "SD",
	"square foot":                     "SF",
	"segment":                         "SG",
	"sheet":                           "SH",
	"square inch":                     "SI",
	"sack":                            "SJ",
	"sleeve":                          "SL",
	"square meter":                    "SM",
	"square rod":                      "SN",
	"spool":                           "SO",
	"shelf package":                   "SP",
	"strip":                           "SR",
	"sheet-metric measure":            "SS",
	"set":                             "ST",
	"skid":                            "SV",
	"shipment":                        "SX",
	"square yard":                     "SY",
	"syringe":                         "SZ",
	"thousand pounds gross":           "T1",
	"thousandths of an inch":          "T2",
	"thousand pieces":                 "T3",
	"thousand bags":                   "T4",
	"thousand casings":                "T5",
	"thousand gallons":                "T6",
	"thousand impressions":            "T7",
	"thousand linear inches":          "T8",
	"thousand kilowatt hours":         "T9",
	"tenth cubic foot":                "TA",
	"tube":                            "TB",
	"truckload":                       "TC",
	"therms":                          "TD",
	"tote":                            "TE",
	"ten square yards":                "TF",
	"gross ton":                       "TG",
	"thousand":                        "TH",
	"thousand square inches":          "TI",
	"thousand sq. centimeters":        "TJ",
	"tank":                            "TK",
	"thousand feet (linear)":          "TL",
	"thousand feet (board)":           "TM",
	"net ton (2,000 lb).":             "TN",
	"troy ounce":                      "TO",
	"ten-pack":                        "TP",
	"thousand feet":                   "TQ",
	"ten square feet":                 "TR",
	"thousand square feet":            "TS",
	"thousand linear meters":          "TT",
	"thousand linear yards":           "TU",
	"thousand kilograms":              "TV",
	"thousand sheets":                 "TW",
	"troy pound":                      "TX",
	"tray":                            "TY",
	"thousand cubic feet":             "TZ",
	"treatments":                      "U1",
	"tablet":                          "U2",
	"ten":                             "U3",
	"two hundred fifty":               "U5",
	"ten thousand yards":              "UH",
	"unitless":                        "UL",
	"million units":                   "UM",
	"unit":                            "UN",
	"troche":                          "UP",
	"dosage form":                     "US",
	"inhalation":                      "UT",
	"lozenge":                         "UU",
	"percent topical only":            "UV",
	"milliequivalent":                 "UW",
	"dram (minim)":                    "UX",
	"fifty square feet":               "UY",
	"fifty count":                     "UZ",
	"five hundred":                    "VC",
	"vial":                            "VI",
	"percent volume":                  "VP",
	"visit":                           "VS",
	"work days":                       "WD",
	"pennyweight":                     "WP",
	"wrap":                            "WR",
	"bunch":                           "X2",
	"drop":                            "X4",
	"tablespoon":                      "Y2",
	"teaspoon":                        "Y3",
	"tub":                             "Y4",
	"yard":                            "YD",
	"100 lineal yards":                "YL",
	"ten yards":                       "YT",
	"cask":                            "Z3",
}

// ConvertUnitOfMeasurementToCode returns the codified version of the unit of measurement per
// https://www.namm.org/standards/implementation-guide-/codes-tables/unit-measurement-uom-codes (not yet finalized).
// If no code is found, we return the code for "each" as our best guess.
func ConvertUnitOfMeasurementToCode(unit string) string {
	unitLower := strings.ToLower(unit)
	code, ok := unitOfMeasurementToCode[unitLower]
	if !ok {
		return unitOfMeasurementToCode["each"]
	}
	return code
}