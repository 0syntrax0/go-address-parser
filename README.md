# GoAddressParser

## Overview

This is the start of a library for [USPS](https://pe.usps.com/) to validate US addresses.

## License

GoAddressParser is licensed under a MIT License.

## Installation

To install GoAddressParser, simply run `go get github.com/0syntrax0/go-address-parser`

## Segments Of A Street Address

To get the segments of a street address

```go
var street Street
s := street.Parse("201 E Randolph St")
```
