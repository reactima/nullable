package nullable

import (
	"database/sql/driver"
	"encoding/json"
)

type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

func (n *NullFloat64) Scan(value interface{}) error {
	if value == nil {
		n.Float64, n.Valid = 0, false
		return nil
	}

	if v, ok := value.(float64); ok {
		n.Float64, n.Valid = v, false
		return nil
	}

	return nil
}

func (n NullFloat64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float64, nil
}

func (ns NullFloat64) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Float64)
	}
	return []byte("null"), nil
}
