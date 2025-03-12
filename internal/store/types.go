package store

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func NewNullString(str string) NullString {
	return NullString{sql.NullString{String: str, Valid: true}}
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}
