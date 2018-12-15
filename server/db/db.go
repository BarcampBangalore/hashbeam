package db

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"server/models"
)

func Init(mongoURL string, user string, password string) (*SessionProvider, error) {
	rootSession, err := mgo.Dial("mongodb://" + user + ":" + password + "@" + mongoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &SessionProvider{rootSession}, nil
}

type SessionProvider struct {
	rootSession *mgo.Session
}

func (s *SessionProvider) GetSession() *Session {
	return &Session{s.rootSession.Copy()}
}

type Session struct {
	session *mgo.Session
}

func (s *Session) GetAnnouncements() ([]models.Announcement, error) {
	var result []models.Announcement

	err := s.session.DB("").C("announcements").Find(bson.M{}).All(&result)
	if err != nil {
		return nil, fmt.Errorf("GetAnnouncements failed: %v", err)
	}

	return result, nil
}

func (s *Session) AddAnnouncement(announcement models.Announcement) error {
	return s.session.DB("").C("announcements").Insert(announcement)
}