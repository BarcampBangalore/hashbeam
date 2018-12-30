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

func (ctx *DBContext) GetTweetsToReview() ([]models.Tweet, error) {
	tweets := make([]models.Tweet, 0)

	err := ctx.db.Select(&tweets, "SELECT `tweet_id`, `datetime`, `original_tweet_json`, `author_id`, `review_required` FROM `tweets` WHERE `review_required` = true")
	if err != nil {
		return nil, fmt.Errorf("db GetTweetsToReview failed: %v", err)
	}

	return tweets, nil
}

func (ctx *DBContext) GetTweetsToDisplay() ([]models.Tweet, error) {
	tweets := make([]models.Tweet, 0)

	err := ctx.db.Select(&tweets, "SELECT `tweet_id`, `datetime`, `original_tweet_json`, `author_id`, `review_required` FROM `tweets` WHERE `review_required` = false")
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
