// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
)

type ExecutionTime struct {
	ID        int32          `json:"id"`
	Parameter sql.NullString `json:"parameter"`
	Test      sql.NullString `json:"test"`
	Value     sql.NullString `json:"value"`
	Deviation sql.NullString `json:"deviation"`
}