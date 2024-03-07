package nullable

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

func SetEmptyJson() (*NullJSONText, error) {
	bDummy, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return &NullJSONText{JSONText: bDummy, Valid: true}, nil
}

func SetStr(s string) *NullString {
	return &NullString{String: s, Valid: true}
}

func SetStrBool(s string) *NullBoolString {
	return &NullBoolString{String: s, Valid: true}
}

func SetInt64(i int64) *NullInt64 {
	return &NullInt64{Int64: i, Valid: true}
}

func SetDateStrTime(s string) *NullTime {
	//layout := "2018-10-01T15:04:05"
	str := s + "T00:00:00.001Z"
	if len(s) == 4 {
		str = s + "-01-01T00:00:00.001Z"
	}
	t, _ := time.Parse(time.RFC3339, str)

	return &NullTime{Time: t, Valid: true}
}
func SetStrFromInt64(i int64) *NullString {
	s := strconv.FormatInt(i, 10)
	return &NullString{String: s, Valid: true}
}

func SetInt64FromStr(s string) *NullInt64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// TODO
		// fmt.Printf("%d of type %T", n, n)
		return &NullInt64{Int64: 0, Valid: false}
	}
	return &NullInt64{Int64: n, Valid: true}
}

func SetNow() *NullTime {
	return &NullTime{Time: time.Now(), Valid: true}
}

func SetBool(b bool) *NullBool {
	return &NullBool{Bool: b, Valid: true}
}

// TODO add sanitizing function to make it even more safer
// see https://github.com/jackc/pgx
func SqlPrepare(x interface{}) interface{} {
	t := reflect.ValueOf(x).Type().String()

	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var innerValue reflect.Value

	switch t {
	case "*int64":
		return x
	case "*nullable.NullString":
		f := val.Interface().(NullString)
		innerValue = reflect.ValueOf(f.String)
		return innerValue.String()
	case "*nullable.NullBool":
		f := val.Interface().(NullBool)
		innerValue = reflect.ValueOf(f.Bool)
		return innerValue.Bool()
	case "*nullable.NullTime":
		f := val.Interface().(NullTime)
		innerValue = reflect.ValueOf(f.Time)
		return f.Time
	case "*nullable.NullInt64":
		f := val.Interface().(NullInt64)
		innerValue = reflect.ValueOf(f.Int64)
		return innerValue.Int()
	case "*nullable.NullFloat64":
		f := val.Interface().(NullFloat64)
		innerValue = reflect.ValueOf(f.Float64)
		return innerValue.Float()
	case "*nullable.NullJSONText":
		f := val.Interface().(NullJSONText)
		innerValue = reflect.ValueOf(f.JSONText)
		return innerValue.Bytes()
		// TODO depricate
	case "*nullable.NullBoolString":
		f := val.Interface().(NullBoolString)
		innerValue = reflect.ValueOf(f.String)
		return innerValue.String()
	}
	return "" // TODO review return
}
