package db

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"server/models"
)

func Init(mongoURL string, user string, password string) (*ContextProvider, error) {
	rootSession, err := mgo.Dial("mongodb://" + user + ":" + password + "@" + mongoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &ContextProvider{rootSession}, nil
}

type ContextProvider struct {
	rootSession *mgo.Session
}

func (p *ContextProvider) GetContext() *Context {
	return &Context{p.rootSession.Copy()}
}

type Context struct {
	session *mgo.Session
}

func (ctx *Context) GetAnnouncements() ([]models.Announcement, error) {
	var result []models.Announcement

	err := ctx.session.DB("").C("announcements").Find(bson.M{}).All(&result)
	if err != nil {
		return nil, fmt.Errorf("GetAnnouncements failed: %v", err)
	}

	return result, nil
}

func (ctx *Context) AddAnnouncement(announcement models.Announcement) error {
	return ctx.session.DB("").C("announcements").Insert(announcement)
}