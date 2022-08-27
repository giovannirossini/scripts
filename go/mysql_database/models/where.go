package models

import "database/sql"

// Where stores to locality data from the code
type Where struct {
	Country   string         `json:"country" db:"country"`
	City      sql.NullString `json:"city" db:"city"`
	PhoneCode int            `json:"phonecode" db:"phonecode"`
}
