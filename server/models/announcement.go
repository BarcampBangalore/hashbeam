package models

import (
	"time"
)

type Announcement struct {
	ID      int64     `db:"id"`
	Time    time.Time `db:"datetime"`
	Message string    `db:"message"`
}
