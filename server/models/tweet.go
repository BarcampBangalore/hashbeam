package models

import (
	"time"
)

type Tweet struct {
	TweetID           string    `db:"tweet_id"`
	Time              time.Time `db:"datetime"`
	OriginalTweetJSON string    `db:"original_tweet_json"`
	AuthorID          string    `db:"author_id"`
	ReviewRequired    bool      `db:"review_required"`
}
