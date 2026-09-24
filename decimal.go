package tourmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Decimal carries a precision-sensitive value — currently only monetary
// amounts (see [Price]) — as the exact decimal string the API sends,
// rather than a float64.
//
// openapi.yml documents these fields as `type: string` with a decimal
// pattern (e.g. "1240.00"), specifically so the API can avoid
// floating-point rounding on currency values. [Decimal.UnmarshalJSON] also
// accepts a bare JSON number defensively, preserving its literal token
// instead of round-tripping it through float64, in case any endpoint ever
// sends one unquoted — either way, no precision is lost.
//
// Decimal is a defined string type with no arithmetic of its own: convert
// with a decimal-math library of your choice (this client has no runtime
// dependencies) if you need to compute with it, or strconv.ParseFloat if
// approximate float precision is acceptable for your use case.
type Decimal string

// String returns d's exact decimal string, as sent by the API.
func (d Decimal) String() string { return string(d) }

// UnmarshalJSON accepts either a JSON string (the API's actual encoding)
// or a bare JSON number (in case that ever proves necessary), preserving
// the exact decimal text either way.
func (d *Decimal) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return fmt.Errorf("tourmanager: decimal: %w", err)
		}
		*d = Decimal(s)
		return nil
	}
	// A bare JSON number: keep its literal token verbatim rather than
	// decoding through float64, which would round-trip it through binary
	// floating point and risk losing exact decimal precision.
	*d = Decimal(data)
	return nil
}

// MarshalJSON encodes d as a JSON string, matching the API's actual
// encoding of these fields.
func (d Decimal) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}
