package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type Announcement struct {
	ID      string    `bson:"_id"`
	Time    time.Time `bson:"time"`
	Message string    `bson:"message"`
}

func NewAnnouncement(message string) Announcement {
	id := uuid.NewV4().String()
	t := time.Now()
	return Announcement{id, t, message}
}
