package db

import (
	"fmt"
	"github.com/globalsign/mgo"
)


func Init(mongoURL string, user string, password string) (*MongoContextProvider, error) {
	rootSession, err := mgo.Dial("mongodb://" + user + ":" + password + "@" + mongoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &MongoContextProvider{ rootSession }, nil
}

type MongoContextProvider struct {
	rootSession *mgo.Session
}

func (m *MongoContextProvider) GetContext() Context {
	return &MongoContext{ m.rootSession.Copy() }
}

type MongoContext struct {
	session *mgo.Session
}

func (m MongoContext) GetTweets() string {
	return "This string came from MongoContext"
}
