// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package databases

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int32
	Name       string
	Email      string
	Password   string
	VerifiedAt sql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
}