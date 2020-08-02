package address

import (
	"strings"
)

// Street street information
type Street struct {
	Number    string
	Direction string
	Name      string
	Suffix    string
}

const (
	// street format
	streetFormatErr = "invalid street type"
)

var (
	// streetAddressSuffix list of street suffixes with their abbreviations
	// Provided by USPS: https://pe.usps.com/text/pub28/28apc_002.htm
	// [Primary Suffix Name]:{
	// 		commonly used, <-USPS standard suffix abbreviation
	// 		commonly used,
	// }
	streetAddressSuffix = map[string][]string{
		// English streets suffixes
		"ALLEY": {
			"ALY",
			"ALLY",
		},
		"ANEX": {
			"ANX",
			"ANNEX",
			"ANNX",
		},
		"ARCADE": {
			"ARC",
		},
		"AVENUE": {
			"AVE",
			"AVEN",
			"AVENU",
			"AVN",
			"AVNUE",
		},
		"BAYOU": {
			"BYU",
		},
		"BEACH": {
			"BCH",
		},
		"BEND": {
			"BND",
		},
		"BLUFF": {
			"BLF",
			"BLUF",
		},
		"BLUFFS": {
			"BLFS",
		},
		"BOTTOM": {
			"BTM",
			"BOT",
			"BOTTM",
		},
		"BOULEVARD": {
			"BLVD",
			"BOUL",
			"BOULV",
		},
		"BRANCH": {
			"BR",
			"BRNCH",
		},
		"BRIDGE": {
			"BRG",
			"BRDGE",
		},
		"BROOK": {
			"BRK",
		},
		"BROOKS": {
			"BRKS",
		},
		"BURG": {
			"BG",
		},
		"BURGS": {
			"BGS",
		},
		"BYPASS": {
			"BYP",
			"BYPA",
			"BYPAS",
			"BYPS",
		},
		"CAMP": {
			"CP",
			"CMP",
		},
		"CANYON": {
			"CYN",
			"CNYN",
		},
		"CAPE": {
			"CPE",
		},
		"CAUSEWAY": {
			"CSWY",
			"CAUSWA",
		},
		"CENTER": {
			"CTR",
			"CENT",
			"CENTR",
			"CENTRE",
			"CNTER",
			"CNTR",
		},
		"CENTERS": {
			"CTRS",
		},
		"CIRCLE": {
			"CIR",
			"CIRC",
			"CIRCL",
			"CRCL",
			"CRCLE",
		},
		"CIRCLES": {
			"CIRS",
		},
		"CLIFF": {
			"CLF",
		},
		"CLIFFS": {
			"CLFS",
		},
		"CLUB": {
			"CLB",
		},
		"COMMON": {
			"CMN",
		},
		"COMMONS": {
			"CMNS",
		},
		"CORNER": {
			"COR",
		},
		"CORNERS": {
			"CORS",
		},
		"COURSE": {
			"CRSE",
		},
		"COURT": {
			"CT",
		},
		"COURTS": {
			"CTS",
		},
		"COVE": {
			"CV",
		},
		"COVES": {
			"CVS",
		},
		"CREEK": {
			"CRK",
		},
		"CRESCENT": {
			"CRES",
			"CRSENT",
			"CRSNT",
		},
		"CREST": {
			"CRST",
		},
		"CROSSING": {
			"XING",
			"CRSSNG",
		},
		"CROSSROAD": {
			"XRD",
		},
		"CROSSROADS": {
			"XRDS",
		},
		"CURVE": {
			"CURV",
		},
		"DALE": {
			"DL",
		},
		"DAM": {
			"DM",
		},
		"DIVIDE": {
			"DV",
			"DVD",
			"DIV",
		},
		"DRIVE": {
			"DR",
			"DRIV",
			"DRV",
		},
		"DRIVES": {
			"DRS",
		},
		"ESTATE": {
			"EST",
		},
		"ESTATES": {
			"ESTS",
		},
		"EXPRESSWAY": {
			"EXPY",
			"EXPR",
			"EXPRESS",
			"EXPW",
		},
		"EXTENSION": {
			"EXT",
			"EXTN",
			"EXTNSN",
		},
		"EXTENSIONS": {
			"EXTS",
		},
		"FALL": {
			"FALL",
		},
		"FALLS": {
			"FLS",
		},
		"FERRY": {
			"FRY",
			"FRRY",
		},
		"FIELD": {
			"FLD",
		},
		"FIELDS": {
			"FLDS",
		},
		"FLAT": {
			"FLT",
		},
		"FLATS": {
			"FLTS",
		},
		"FORD": {
			"FRD",
		},
		"FORDS": {
			"FRDS",
		},
		"FOREST": {
			"FRST",
			"FORESTS",
		},
		"FORGE": {
			"FRG",
		},
		"FORGES": {
			"FRGS",
		},
		"FORK": {
			"FRK",
		},
		"FORKS": {
			"FRKS",
		},
		"FORT": {
			"FT",
			"FRT",
		},
		"FREEWAY": {
			"FWY",
			"FREEWY",
			"FRWAY",
			"FRWY",
		},
		"GARDEN": {
			"GDN",
			"GARDN",
			"GRDEN",
			"GRDN",
		},
		"GARDENS": {
			"GDNS",
			"GRDNS",
		},
		"GATEWAY": {
			"GTWY",
			"GATEWY",
			"GATWAY",
			"GTWAY",
		},
		"GLEN": {
			"GLN",
		},
		"GLENS": {
			"GLNS",
		},
		"GREEN": {
			"GRN",
		},
		"GREENS": {
			"GRNS",
		},
		"GROVE": {
			"GRV",
			"GROV",
		},
		"GROVES": {
			"GRVS",
		},
		"HARBOR": {
			"HBR",
			"HARB",
			"HARBR",
			"HRBOR",
		},
		"HARBORS": {
			"HBRS",
		},
		"HAVEN": {
			"HVN",
		},
		"HEIGHTS": {
			"HTS",
		},
		"HIGHWAY": {
			"HWY",
			"HIGHWY",
			"HIWAY",
			"HIWY",
			"HWAY",
		},
		"HILL": {
			"HL",
		},
		"HILLS": {
			"HLS",
		},
		"HOLLOW": {
			"HOLW",
			"HLLW",
			"HOLLOWS",
			"HOLWS",
		},
		"INLET": {
			"INLT",
		},
		"ISLAND": {
			"IS",
			"ISLND",
		},
		"ISLANDS": {
			"ISS",
			"ISLNDS",
		},
		"ISLE": {
			"ISLE",
			"ISLES",
		},
		"JUNCTION": {
			"JCT",
			"JCTION",
			"JCTN",
			"JUNCTN",
			"JUNCTON",
		},
		"JUNCTIONS": {
			"JCTS",
			"JCTNS",
		},
		"KEY": {
			"KY",
		},
		"KEYS": {
			"KYS",
		},
		"KNOLL": {
			"KNL",
			"KNOL",
		},
		"KNOLLS": {
			"KNLS",
		},
		"LAKE": {
			"LK",
		},
		"LAKES": {
			"LKS",
		},
		"LAND": {
			"LAND",
		},
		"LANDING": {
			"LNDG",
			"LNDNG",
		},
		"LANE": {
			"LN",
		},
		"LIGHT": {
			"LGT",
		},
		"LIGHTS": {
			"LGTS",
		},
		"LOAF": {
			"LF",
		},
		"LOCK": {
			"LCK",
		},
		"LOCKS": {
			"LCKS",
		},
		"LODGE": {
			"LDG",
			"LDGE",
			"LODG",
		},
		"LOOP": {
			"LOOP",
			"LOOPS",
		},
		"MALL": {
			"MALL",
		},
		"MANOR": {
			"MNR",
		},
		"MANORS": {
			"MNRS",
		},
		"MEADOW": {
			"MDW",
		},
		"MEADOWS": {
			"MDWS",
			"MEDOWS",
		},
		"MEWS": {
			"MEWS",
		},
		"MILL": {
			"ML",
		},
		"MILLS": {
			"MLS",
		},
		"MISSION": {
			"MSN",
			"MSSN",
			"MISSN",
		},
		"MOTORWAY": {
			"MTWY",
		},
		"MOUNT": {
			"MT",
			"MNT",
		},
		"MOUNTAIN": {
			"MTN",
			"MNTAIN",
			"MNTN",
			"MOUNTIN",
			"MTIN",
		},
		"MOUNTAINS": {
			"MTNS",
			"MNTNS",
		},
		"NECK": {
			"NCK",
		},
		"ORCHARD": {
			"ORCH",
			"ORCHRD",
		},
		"OVAL": {
			"OVAL",
			"OVL",
		},
		"OVERPASS": {
			"OPAS",
		},
		"PARK": {
			"PARK",
			"PRK",
		},
		"PARKS": {
			"PARK",
		},
		"PARKWAY": {
			"PKWY",
			"PARKWY",
			"PKWAY",
			"PKY",
		},
		"PARKWAYS": {
			"PKWY",
			"PKWYS",
		},
		"PASS": {
			"PASS",
		},
		"PASSAGE": {
			"PSGE",
		},
		"PATH": {
			"PATH",
			"PATHS",
		},
		"PIKE": {
			"PIKE",
			"PIKES",
		},
		"PINE": {
			"PNE",
		},
		"PINES": {
			"PNES",
		},
		"PLACE": {
			"PL",
		},
		"PLAIN": {
			"PLN",
		},
		"PLAINS": {
			"PLNS",
		},
		"PLAZA": {
			"PLZ",
			"PLZA",
		},
		"POINT": {
			"PT",
		},
		"POINTS": {
			"PTS",
		},
		"PORT": {
			"PRT",
		},
		"PORTS": {
			"PRTS",
		},
		"PRAIRIE": {
			"PR",
			"PRR",
		},
		"RADIAL": {
			"RADL",
			"RAD",
			"RADIEL",
		},
		"RAMP": {
			"RAMP",
		},
		"RANCH": {
			"RNCH",
			"RANCHES",
			"RNCHS",
		},
		"RAPID": {
			"RPD",
		},
		"RAPIDS": {
			"RPDS",
		},
		"REST": {
			"RST",
		},
		"RIDGE": {
			"RDG",
			"RDGE",
		},
		"RIDGES": {
			"RDGS",
		},
		"RIVER": {
			"RIV",
			"RVR",
			"RIVR",
		},
		"ROAD": {
			"RD",
		},
		"ROADS": {
			"RDS",
		},
		"ROUTE": {
			"RTE",
		},
		"ROW": {
			"ROW",
		},
		"RUE": {
			"RUE",
		},
		"RUN": {
			"RUN",
		},
		"SHOAL": {
			"SHL",
		},
		"SHOALS": {
			"SHLS",
		},
		"SHORE": {
			"SHR",
			"SHOAR",
		},
		"SHORES": {
			"SHRS",
			"SHOARS",
		},
		"SKYWAY": {
			"SKWY",
		},
		"SPRING": {
			"SPG",
			"SPNG",
			"SPRNG",
		},
		"SPRINGS": {
			"SPGS",
			"SPNGS",
			"SPRNGS",
		},
		"SPUR": {
			"SPUR",
		},
		"SPURS": {
			"SPUR",
		},
		"SQUARE": {
			"SQ",
			"SQR",
			"SQRE",
			"SQU",
		},
		"SQUARES": {
			"SQS",
			"SQRS",
		},
		"STATION": {
			"STA",
			"STATN",
			"STN",
		},
		"STRAVENUE": {
			"STRA",
			"STRAV",
			"STRAVEN",
			"STRAVN",
			"STRVN",
			"STRVNUE",
		},
		"STREAM": {
			"STRM",
			"STREME",
		},
		"STREET": {
			"ST",
			"STRT",
			"STR",
		},
		"STREETS": {
			"STS",
		},
		"SUMMIT": {
			"SMT",
			"SUMIT",
			"SUMITT",
		},
		"TERRACE": {
			"TER",
			"TERR",
		},
		"THROUGHWAY": {
			"TRWY",
		},
		"TRACE": {
			"TRCE",
			"TRACES",
		},
		"TRACK": {
			"TRAK",
			"TRACKS",
			"TRK",
			"TRKS",
		},
		"TRAFFICWAY": {
			"TRFY",
		},
		"TRAIL": {
			"TRL",
			"TRAILS",
			"TRLS",
		},
		"TRAILER": {
			"TRLR",
			"TRLRS",
		},
		"TUNNEL": {
			"TUNL",
			"TUNEL",
			"TUNLS",
			"TUNNELS",
			"TUNNL",
		},
		"TURNPIKE": {
			"TPKE",
			"TRNPK",
			"TURNPK",
		},
		"UNDERPASS": {
			"UPAS",
		},
		"UNION": {
			"UN",
		},
		"UNIONS": {
			"UNS",
		},
		"VALLEY": {
			"VLY",
			"VALLY",
			"VLLY",
		},
		"VALLEYS": {
			"VLYS",
		},
		"VIADUCT": {
			"VIA",
			"VDCT",
			"VIADCT",
		},
		"VIEW": {
			"VW",
		},
		"VIEWS": {
			"VWS",
		},
		"VILLAGE": {
			"VLG",
			"VILL",
			"VILLAG",
			"VILLG",
			"VILLIAGE",
		},
		"VILLAGES": {
			"VLGS",
		},
		"VILLE": {
			"VL",
		},
		"VISTA": {
			"VIS",
			"VIST",
			"VST",
			"VSTA",
		},
		"WALK": {
			"WALK",
		},
		"WALKS": {
			"WALK",
		},
		"WALL": {
			"WALL",
		},
		"WAY": {
			"WAY",
		},
		"WAYS": {
			"WAYS",
		},
		"WELL": {
			"WL",
		},
		"WELLS": {
			"WLS",
		},
		// Spanish streets suffixes
		"AVENIDA": {
			"AVE",
		},
		"CALLE": {
			"CLL",
		},
		"CAMINITO": {
			"CMT",
		},
		"CAMINO": {
			"CAM",
		},
		"CERRADA": {
			"CER",
		},
		"CIRCULO": {
			"CIR",
		},
		"ENTRADA": {
			"ENT",
		},
		"PASEO": {
			"PSO",
		},
		"PLACITA": {
			"PLA",
		},
		"RANCHO": {
			"RCH",
		},
		"VEREDA": {
			"VER",
		},
		// Virgen Islands street suffixes
	}

	// spanishPronounds
	spanishPronounds = []string{
		"LA", "LAS", "EL", "LO", "LOS", "DE LA", "DE LAS", "DE LOS",
	}
)

// Parse parses a street
// Follow this pattern: http://help.lucity.com/webhelp/v155/web/35257.htm
func (s *Street) Parse(st string) *Street {
	// get street segments
	segments := strings.Split(st, " ")
	var cleanedSegment string

	// quick check
	if len(segments) < 2 {
		return nil
	}

	// get street number
	s.Number, segments = getStreetNumber(segments)

	// get street direction
	cleanedSegment = lettersOnly(segments[0])
	if isStreetDirection(cleanedSegment) {
		s.Direction = cleanedSegment
		segments = removeSliceByKey(segments, 0)
	}

	// get street suffix
	s.Suffix, segments = isSuffix(segments)

	// what's left is the street name
	s.Name = strings.ToUpper(alphaNumericSpaceDots(strings.Join(segments, " ")))

	// here you go!
	return s
}

// getStreetNumber Parses street numbers: ["123", "123-1/2", "123 1/2"]
func getStreetNumber(segments []string) (string, []string) {
	var streetNumber string

	// check for hyphenated address
	if isHyphenatedAddress(segments[0]) {
		streetNumber = segments[0]
		return streetNumber, removeSliceByKey(segments, 0)
	} else {
		streetNumber = cleanStreetNumber(segments[0])
	}

	// check for fraction number: 636 1/2
	if len(segments[1]) == 3 && strings.Contains(segments[1], "/") {
		fraction := cleanStreetNumber(segments[1])
		segments = removeSliceByKey(segments, 0)

		return streetNumber + "-" + fraction, removeSliceByKey(segments, 0)
	}

	// check for PO Boxes
	if checkPoBox(segments) {
		segments = removeSliceByKey(segments, 0)
		return streetNumber, segments
	}

	// regular address
	return cleanStreetNumber(segments[0]), removeSliceByKey(segments, 0)
}

// checkPoBox
func checkPoBox(segments []string) bool {
	return strings.ToUpper(alphaNumericSpace(segments[1])) == "PO"
}
