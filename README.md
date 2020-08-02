# GoAddressParser

## Overview

This is the start of a library for [USPS](https://pe.usps.com/) to validate US addresses.

## License

GoAddressParser is licensed under a MIT License.

## Installation

To install GoAddressParser, simply run `go get github.com/0syntrax0/go-address-parser`

## Segments Of A Street Address

To get the segments of:

- Standard US style address: `201 E Randolph St`
- PO Boxes: `12341 (PO Box)`
- Addresses with dashes (hyphens): `AB90–AB120 E Randolph St`
- Grid style (Usually Found in the Salt Lake City area): `11782 Rd 39.4`
- Alpha-numeric addresses (usually found in Wisconsin and Northern Illinois): `N6W23001 BLUEMOUND RD`
- Fractional addresses: `123 1/2 BLUEMOUND RD`

Just pass your address through

## Address Segments

```go
var s Street
addressSegments := s.Parse("201 E Randolph St")
```
