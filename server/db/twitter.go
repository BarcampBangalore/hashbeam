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

func (ctx *DBContext) GetMutedTwitterIDs() ([]string, error) {
	result := make([]string, 0)

	err := ctx.db.Select(&result, "SELECT `twitter_user_id` FROM `muted_tweeters`")
	if err != nil {
		return nil, fmt.Errorf("db GetMutedTwitterIDs failed; %v", err)
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

func (ctx *DBContext) ApproveTweet(t models.Tweet) error {
	tx, err := ctx.db.Begin()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db ApproveTweet failed to start transaction: %v", err)
	}

	_, err = tx.Exec("UPDATE `tweets` SET `review_required` = false WHERE `tweet_id` = ?", t.TweetID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db ApproveTweet failed to set review_required on tweet: %v", err)
	}

	_, err = tx.Exec("INSERT INTO `whitelisted_tweeters` (`twitter_user_id`) VALUES (?)", t.AuthorID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db ApproveTweet failed to add author to whitelisted_tweeters: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db ApproveTweet transaction commit failed: %v", err)
	}

	return nil
}

func (ctx *DBContext) RejectTweet(t models.Tweet) error {
	tx, err := ctx.db.Begin()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db RejectTweet failed to start transaction: %v", err)
	}

	_, err = tx.Exec("DELETE FROM `tweets` where `author_id` = ?", t.AuthorID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db RejectTweet failed to delete tweets belonging to author: %v", err)
	}

	_, err = tx.Exec("INSERT INTO `muted_tweeters` (`twitter_user_id`) VALUES (?)", t.AuthorID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db RejectTweet failed to add author to muted_tweeters: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("db RejectTweet transaction commit failed: %v", err)
	}

	return nil
}
