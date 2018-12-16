package db

import (
	"fmt"
	"github.com/globalsign/mgo"
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
