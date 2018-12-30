package db

import (
	"fmt"
	"server/models"
)

func (ctx *DBContext) GetWhitelistTwitterIDs() ([]string, error) {
	result := make([]string, 0)

	err := ctx.db.Select(&result, "SELECT `twitter_user_id` FROM `whitelisted_tweeters`")
	if err != nil {
		return nil, fmt.Errorf("db GetWhitelistedTwitterIDs failed: %v", err)
	}

	return result, nil
}

func (ctx *DBContext) SaveTweet(t models.Tweet) error {
	_, err := ctx.db.Exec(
		"INSERT INTO `tweets` (`tweet_id`, `datetime`, `json_body`, `author_id`, `review_required`) VALUES (?, ?, ?, ?, ?)",
		t.TweetID,
		t.Time,
		t.JSONBody,
		t.AuthorID,
		t.ReviewRequired,
	)

	if err != nil {
		return fmt.Errorf("db SaveTweet failed: %v", err)
	}
	return nil
}
