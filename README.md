# sqlx-nullable

As we historically stacked with SQLX and PGX wasn't in its full glory, here's the collection of nullable structs for effective communication with database nulls and marshalling/unmarshalling from TS/JS

- https://github.com/jmoiron/sqlx
- https://github.com/jackc/pgx

Each type must have Scan, Value, MarshalJSON and UnmarshalJSON method declared explicitly

Typical model

```go
//go:generate go run ../generator-server/generator.go Cache hh_cache $GOFILE
//go:generate go run ../generator-front/generator.go Cache hh_cache $GOFILE
type Cache struct {
    ID     *int64         `json:"id" db:"cache_id"`

    UserID *nl.NullInt64  `json:"userID,omitempty" db:"cache_user_id"`
    Type   *nl.NullString `json:"type,omitempty" db:"cache_type"`
    TypeID *nl.NullInt64  `json:"typeID,omitempty" db:"cache_type_id"`

    Status   *nl.NullString `json:"status,omitempty" db:"cache_status"`
    Priority *nl.NullInt64  `json:"priority,omitempty" db:"cache_priority"`

    Data *nl.NullJSONText `json:"data,omitempty" db:"cache_data"`
    Meta *nl.NullJSONText `json:"meta,omitempty" db:"cache_meta"`

    EntryDate *nl.NullTime `json:"entryDate,omitempty" db:"cache_entry_date"`
    EditDate  *nl.NullTime `json:"editDate,omitempty" db:"cache_edit_date"`

    // EXTRA FIELDS
    RequestID string `json:"requestID,omitempty" hh:"ignore"`
} // @name Cache
```

