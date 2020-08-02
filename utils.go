package address

import (
	"log"
	"regexp"
	"strings"
)

const (
	// regex
	regexZipcode                 = "[\\d]{5}|[\\d\\-]{10}" //|
	regexAlpha                   = "[^a-zA-Z]+"
	regexAlphaNumeric            = "[^a-zA-Z0-9]+"
	regexStreetNumbers           = "[^(\\da-zA-Z\\-\\/)|(\\d\\-a-zA-Z) ]+" //"[\\d\\-\\/ ]+"
	regexAlphaNumericSpace       = "[^a-zA-Z0-9 ]+"
	regexAlphaNumericSpaceDots   = "[^a-zA-Z0-9\\. ]+"
	regexAlphaNumericSpaceDashes = "[^a-zA-Z0-9\\- ]+"
	regexHyphenatedAddress       = "[^\\da-zA-Z\\-\\da-zA-Z]"
)

func matchStreetNumber(s string) bool {
	// "[\\d\\-\\/a-zA-Z ]+"
	matched, _ := regexp.MatchString(regexStreetNumbers, s)
	return matched
}

// zipCodeCheck cleans variable and returns [^0-9]+
func zipCodeCheck(s string) bool {
	regx, err := regexp.Compile(regexZipcode)
	if err != nil {
		log.Printf("[address package] failed to compile zipCodeCheck regex: %+v", err)
	}
	return regx.Match([]byte(s))
}

// lettersOnly cleans variable and returns [^a-zA-Z]+
func lettersOnly(s string) string {
	regx, err := regexp.Compile(regexAlpha)
	if err != nil {
		log.Printf("[address package] failed to compile lettersOnly regex: %+v", err)
	}
	return regx.ReplaceAllString(s, "")
}

// alphaNumeric cleans variable and returns [^a-zA-Z0-9]+
func alphaNumeric(s string) string {
	regx, err := regexp.Compile(regexAlphaNumeric)
	if err != nil {
		log.Printf("[address package] failed to compile alphanumeric regex: %+v", err)
	}
	return regx.ReplaceAllString(s, "")
}

// alphaNumericSpace cleans variable and returns [^a-zA-Z0-9 ]+
func alphaNumericSpace(s string) string {
	regx, err := regexp.Compile(regexAlphaNumericSpace)
	if err != nil {
		log.Printf("[address package] failed to compile alphaNumericSpace regex: %+v", err)
	}
	return regx.ReplaceAllString(s, "")
}

// alphaNumericSpaceDots cleans variable and returns [^a-zA-Z0-9\. ]+
func alphaNumericSpaceDots(s string) string {
	regx, err := regexp.Compile(regexAlphaNumericSpaceDots)
	if err != nil {
		log.Printf("[address package] failed to compile alphaNumericSpaceDots regex: %+v", err)
	}
	return regx.ReplaceAllString(s, "")
}

// cleanStreetNumber cleans a street number
func cleanStreetNumber(s string) string {
	regx, err := regexp.Compile(regexStreetNumbers)
	if err != nil {
		log.Printf("[address package] failed to compile cleanStreetNumber regex: %+v", err)
	}
	return regx.ReplaceAllString(s, "")
}

// removeSliceByKey removes part of a slice by key
func removeSliceByKey(slice []string, key int) []string {
	return append(slice[:key], slice[key+1:]...)
}

// isHyphenatedAddress checks if the address is hyphenated
// https://pe.usps.com/text/pub28/28ape_003.htm
func isHyphenatedAddress(s string) bool {
	return (strings.Contains(s, "-") || strings.Contains(s, "–")) && (strings.Count(s, "-") == 1 || strings.Count(s, "–") == 1)
}

// isStreetDirection checks if a given string is a cardinal direction
func isStreetDirection(s string) bool {
	switch strings.ToUpper(s) {
	case "N", "NORTH", "NORTE":
		return true
	case "NW", "NORTHWEST", "NOROESTE":
		return true
	case "NE", "NORTHEAST", "NORESTE":
		return true
	case "E", "EAST", "ESTE":
		return true
	case "S", "SOUTH", "SUR":
		return true
	case "SE", "SOUTHEAST", "SURESTE":
		return true
	case "SW", "SOUTHWEST", "SUROESTE":
		return true
	case "W", "WEST", "OESTE":
		return true
	}
	return false
}

// isSuffix checks if a given street segment is a suffix, ie: street, ave, blv, etc...
func isSuffix(segments []string) (string, []string) {
	var streetSuffix string

	// get street suffix
	lastElement := len(segments) - 1
	if ok, suffix := isStreetSuffix(alphaNumeric(segments[lastElement])); ok {
		streetSuffix = suffix
		segments = removeSliceByKey(segments, lastElement)
	}

	return streetSuffix, segments
}

// isStreetSuffix checks if a given string is an English suffix, ie: street, avenue, blv, etc...
func isStreetSuffix(s string) (bool, string) {
	s = strings.ToUpper(s)

	// check for English suffixes
	for suffix, abbreviations := range streetAddressSuffix {
		if s == suffix {
			return true, suffix
		}

		// check if it's an abbreviation
		for _, a := range abbreviations {
			if s == a {
				return true, suffix
			}
		}
	}
	return false, s
}
