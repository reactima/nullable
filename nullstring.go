package nullable

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// NullString MarshalJSON ignores ""
type NullString struct {
	String string
	Valid  bool
}

func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	// TODO review and check perfomance with PGX
	ns.String = fmt.Sprintf("%v", value)
	return nil
}

func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid || len(ns.String) == 0 {
		return "", nil
	}
	return ns.String, nil
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid && ns.String != "" {
		return json.Marshal(ns.String)
	}
	return []byte("null"), nil
}

func (ns *NullString) UnmarshalJSON(b []byte) error {

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	ns.String = s
	ns.Valid = true
	return nil
}
