package db

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"server/models"
	"time"
)

func (s *Session) GetAnnouncements() ([]models.Announcement, error) {
	var result []models.Announcement

	err := s.session.DB("").C("announcements").Find(bson.M{}).All(&result)
	if err != nil {
		return nil, fmt.Errorf("GetAnnouncements failed: %v", err)
	}

	return result, nil
}

func (s *Session) GetAnnouncementsByDate(date time.Time) ([]models.Announcement, error) {
	year, month, day := date.Date()

	st := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	et := st.Add(24 * time.Hour)

	var result []models.Announcement

	err := s.session.DB("").C("announcements").Find(bson.M{
		"time": bson.M{
			"$gt": st,
			"$lt": et,
		},
	}).Sort("-time").All(&result)

	if err != nil {
		return nil, fmt.Errorf("GetAnnouncementsByDate failed: %v", err)
	}

	return result, nil
}

func (s *Session) SaveAnnouncement(announcement models.Announcement) error {
	return s.session.DB("").C("announcements").Insert(announcement)
}
