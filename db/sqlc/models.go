// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
)

type Code struct {
	Code           string
	Phone          sql.NullString
	Email          sql.NullString
	Birthday       sql.NullTime
	NameSurname    sql.NullString
	Address        sql.NullString
	CityCode       sql.NullInt32
	RegisteredDate sql.NullTime
	Ip             sql.NullString
	Active         bool
	Banned         bool
	BannedReason   sql.NullString
	TcNo           sql.NullString
	Used           bool
}