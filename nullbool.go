package nullable

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"

	"encoding/json"
)

// NullBool is an alias for sql.NullBool data type
type NullBool sql.NullBool

// Scan implements the Scanner interface for NullBool
func (nb *NullBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NullBool{b.Bool, false}
	} else {
		*nb = NullBool{b.Bool, true}
	}

	return nil
}

func (ns NullBool) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Bool)
	}
	return []byte("null"), nil
}

func (ns *NullBool) UnmarshalJSON(b []byte) error {
	var bb bool
	if err := json.Unmarshal(b, &bb); err != nil {
		return err
	}
	ns.Bool = bb
	ns.Valid = true
	return nil
}

func (n *NullBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}

// TODO depricate
type NullBoolString struct {
	String string
	Valid  bool
}

func (ns *NullBoolString) Scan(value interface{}) error {
	fmt.Println("ns", value)
	if value == nil {
		ns.String, ns.Valid = "0", false
		return nil
	}
	if v, ok := value.(bool); ok && v == false {
		ns.String, ns.Valid = "0", false
		return nil
	}
	if v, ok := value.(string); ok && (v == "0" || v == "false") {
		ns.String, ns.Valid = "0", false
		return nil
	}

	ns.Valid = true
	ns.String = "1"
	return nil
}

func (ns NullBoolString) Value() (driver.Value, error) {
	if !ns.Valid || len(ns.String) == 0 {
		return "0", nil
	}
	if ns.Valid && (ns.String == "" || ns.String == "false") {
		return "0", nil
	}
	fmt.Println("ns", ns.String)
	return "1", nil
}

func (ns NullBoolString) MarshalJSON() ([]byte, error) {
	if ns.Valid && ns.String != "" {
		return []byte("1"), nil
	}
	return []byte("0"), nil
}

func (ns *NullBoolString) UnmarshalJSON(b []byte) error {
	var s string
	var bb bool
	var i int64
	if err := json.Unmarshal(b, &s); err != nil {
		if err := json.Unmarshal(b, &bb); err != nil {
			if err := json.Unmarshal(b, &i); err != nil {
				return err
			}
			if i > 0 {
				s = "1"
			} else {
				s = "0"
			}
		}
		if bb {
			s = "1"
		} else {
			s = "0"
		}
	}
	if s == "false" || s == "0" {
		ns.String = "0"
		ns.Valid = true
	}

	ns.String = s
	ns.Valid = true
	return nil
}
