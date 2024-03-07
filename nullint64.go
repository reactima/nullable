package nullable

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// NullInt64 can accept string or interger from UnmarshalJSON
type NullInt64 struct {
	Int64 int64
	Valid bool
}

func (n *NullInt64) Scan(value interface{}) error {
	if value == nil {
		n.Int64, n.Valid = 0, false
		return nil
	}

	if v, ok := value.(int64); ok && v > 0 {
		n.Int64, n.Valid = v, true
		return nil
	}

	return nil
}

func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

func (n NullInt64) MarshalJSON() ([]byte, error) {
	if n.Valid && n.Int64 > 0 {
		return json.Marshal(n.Int64)
	}
	return []byte("null"), nil
}

func (n *NullInt64) UnmarshalJSON(b []byte) error {
	var s int64
	var ss string
	if err := json.Unmarshal(b, &s); err != nil {
		if err := json.Unmarshal(b, &ss); err != nil {
			n.Int64, n.Valid = 0, true
			return nil
		}
		i, err := strconv.ParseInt(ss, 10, 64)
		if err != nil {
			n.Int64, n.Valid = 0, true
			return nil
		}
		s = i
	}
	n.Int64 = s
	n.Valid = true

	return nil
}
