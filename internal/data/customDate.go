package data

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// CustomDate wraps time.Time
type CustomDate time.Time

// MarshalJSON implements json.Marshaler
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	t := time.Time(cd) // Convert to time.Time
	formatted := fmt.Sprintf(`"%s"`, strings.ToLower(t.Format("02Jan2006")))
	return []byte(formatted), nil
}

// MarshalBSONValue implements bson.ValueMarshaler
func (cd CustomDate) MarshalBSONValue() (bsontype.Type, []byte, error) {
	t := time.Time(cd)
	return bson.MarshalValue(t)
}

// UnmarshalBSONValue extracts the BSON datetime and strips it to Y/M/D
func (cd *CustomDate) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	// Make sure the BSON type is datetime
	if t != bsontype.DateTime {
		return fmt.Errorf("unsupported BSON type for CustomDate: %v", t)
	}

	// Create a BSON value reader for the raw data
	vr := bsonrw.NewBSONValueReader(t, data)

	// Create a decoder from the value reader
	dec, err := bson.NewDecoder(vr)
	if err != nil {
		return err
	}

	// Decode into a time.Time
	var rawTime time.Time
	if err := dec.Decode(&rawTime); err != nil {
		return err
	}

	// Truncate to date only
	*cd = CustomDate(time.Date(rawTime.Year(), rawTime.Month(), rawTime.Day(), 0, 0, 0, 0, rawTime.Location()))
	return nil
}
