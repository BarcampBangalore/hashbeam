package db

type ContextProvider interface {
	GetContext() Context
}

type Context interface {
	GetTweets() string
}
