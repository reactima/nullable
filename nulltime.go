package nullable

import (
	"bytes"
	"database/sql/driver"
	"time"

	"encoding/json"
)

// NullTime MarshalJSON ignores ZeroTime
type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	if _, ok := value.(time.Time); ok && !value.(time.Time).IsZero() {
		nt.Time, nt.Valid = value.(time.Time), true
		return nil
	}
	t := new(time.Time) // zero time
	nt.Time, nt.Valid = *t, false
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (ns NullTime) MarshalJSON() ([]byte, error) {
	if ns.Valid && !ns.Time.IsZero() {
		//return json.Marshal(ns.Time.UnixNano() / 1e6)
		return json.Marshal(ns.Time.Format("2006-01-02T15:04:05.000Z"))
		//time.RFC3339
	}
	return []byte("null"), nil
}

func (ns *NullTime) UnmarshalJSON(b []byte) error {

	if bytes.Equal(b, []byte("null")) || string(b) == "" || string(b) == `""` {
		ns.Time = time.Time{}
		ns.Valid = false
		return nil
	}

	var s time.Time
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	ns.Time = s
	ns.Valid = true

	return nil
}
