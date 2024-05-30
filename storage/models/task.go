package models

import (
	"time"
)

type Task struct {
	Id          int
	Subject     string
	Description string
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
