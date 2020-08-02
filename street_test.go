package address

import (
	"testing"
)

// TestStreetStandard the most common type of address in the US
func TestStreetStandard(t *testing.T) {
	var street Street
	s := street.Parse("201 E Randolph St")

	if s.Number != "201" {
		t.Fatal("Failed to get standard street number")
	}
	if s.Direction != "E" {
		t.Fatal("Failed to get standard street direction")
	}
	if s.Name != "RANDOLPH" {
		t.Fatal("Failed to get standard street name")
	}
	if s.Suffix != "STREET" {
		t.Fatal("Failed to get standard street suffix")
	}
}

// TestPOBox PO Boxes
func TestPOBox(t *testing.T) {
	var street Street
	s := street.Parse("12341 (PO Box)")

	if s.Number != "12341" {
		t.Fatal("Failed to get PO Box street number")
	}
	if s.Name != "PO BOX" {
		t.Fatal("Failed to detect PO Box")
	}
	if s.Suffix != "" {
		t.Fatal("Failed to detect no suffix")
	}
}

// TestHyphenatedAddress addresses with dashes (hyphens) in the street number
func TestHyphenatedAddress(t *testing.T) {
	var street Street
	s := street.Parse("111-222 E Randolph St")

	if s.Number != "111-222" {
		t.Fatalf("Failed to get hyphenated street number: %+v", s.Number)
	}
	if s.Direction != "E" {
		t.Fatal("Failed to get hyphenated street direction")
	}
	if s.Name != "RANDOLPH" {
		t.Fatalf("Failed to get hyphenated street name")
	}
	if s.Suffix != "STREET" {
		t.Fatal("Failed to get hyphenated street suffix")
	}
}

// TestGridStyle Usually Found in the Salt Lake City area
func TestGridStyle(t *testing.T) {
	var street Street
	s := street.Parse("11782 Rd 39.4")

	if s.Number != "11782" {
		t.Fatal("Failed to get grid style street number")
	}
	if s.Direction != "" {
		t.Fatal("Failed to get grid style street direction")
	}
	if s.Name != "RD 39.4" {
		t.Fatal("Failed to get grid style street name")
	}
	if s.Suffix != "" {
		t.Fatalf("Failed to get grid style street suffix: %+v", s.Suffix)
	}
}

// TestAlphanumericAddresses This format is usually found in Wisconsin and Northern Illinois
func TestAlphanumericAddresses(t *testing.T) {
	var street Street
	s := street.Parse("N6W23001 BLUEMOUND RD")

	if s.Number != "N6W23001" {
		t.Fatal("Failed to get alphanumeric street number")
	}
	if s.Direction != "" {
		t.Fatal("Failed to get alphanumeric street direction")
	}
	if s.Name != "BLUEMOUND" {
		t.Fatal("Failed to get alphanumeric street name")
	}
	if s.Suffix != "ROAD" {
		t.Fatal("Failed to get alphanumeric street suffix")
	}
}

// TestFractionalAddresses addresses with fraction numbers in them
func TestFractionalAddresses(t *testing.T) {
	var street Street
	s := street.Parse("123 1/2 BLUEMOUND RD")

	if s.Number != "123-1/2" {
		t.Fatal("Failed to get fractional street number")
	}
	if s.Direction != "" {
		t.Fatal("Failed to get fractional street direction")
	}
	if s.Name != "BLUEMOUND" {
		t.Fatal("Failed to get fractional street name")
	}
	if s.Suffix != "ROAD" {
		t.Fatal("Failed to get fractional street suffix")
	}
}

// BenchmarkStreetStandard the most common type of address in the US
func BenchmarkStreetStandard(b *testing.B) {
	for i := 0; i < 10000; i++ {
		var street Street
		street.Parse("201 E Randolph St")
	}
}
