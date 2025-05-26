package data

import (
	"strings"
	"time"
)

// CustomTime wraps time.Time
type CustomDate struct {
	time.Time
}

// MarshalJSON implements json.Marshaler
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	formatted := cd.Format(`"02Jan2006"`) // Use backticks for raw string literal to avoid escaping quotes
	return []byte(strings.ToLower(formatted)), nil
}
