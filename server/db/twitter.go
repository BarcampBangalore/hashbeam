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

func (ctx *DBContext) GetTweetToReview() (*models.Tweet, error) {
	var tweet models.Tweet

	err := ctx.db.Get(&tweet, "SELECT `tweet_id`, `datetime`, `original_tweet_json`, `author_id`, `review_required` FROM `tweets` WHERE `review_required` = true LIMIT 1")
	if err != nil {
		return nil, fmt.Errorf("db GetTweetToReview failed: %v", err)
	}

	return &tweet, nil
}

func (ctx *DBContext) GetTweetsToDisplay() ([]models.Tweet, error) {
	tweets := make([]models.Tweet, 0)

	err := ctx.db.Select(&tweets, "SELECT `tweet_id`, `datetime`, `original_tweet_json`, `author_id`, `review_required` FROM `tweets` WHERE `review_required` = false ORDER BY `datetime` DESC LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("db GetTweetsToDisplay failed: %v", err)
	}

	return tweets, nil
}

func (ctx *DBContext) SaveTweet(t models.Tweet) error {
	_, err := ctx.db.Exec(
		"INSERT INTO `tweets` (`tweet_id`, `datetime`, `original_tweet_json`, `author_id`, `review_required`) VALUES (?, ?, ?, ?, ?)",
		t.TweetID,
		t.Time,
		t.OriginalTweetJSON,
		t.AuthorID,
		t.ReviewRequired,
	)

	if err != nil {
		return fmt.Errorf("db SaveTweet failed: %v", err)
	}
	return nil
}
