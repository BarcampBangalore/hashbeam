package db

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"server/models"
)

func (s *Session) GetAnnouncements() ([]models.Announcement, error) {
	var result []models.Announcement

	err := s.session.DB("").C("announcements").Find(bson.M{}).All(&result)
	if err != nil {
		return nil, fmt.Errorf("GetAnnouncements failed: %v", err)
	}

	return result, nil
}

func (s *Session) SaveAnnouncement(announcement models.Announcement) error {
	return s.session.DB("").C("announcements").Insert(announcement)
}
