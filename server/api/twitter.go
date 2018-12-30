package api

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type apiResponseTweet struct {
	TweetID        string        `json:"tweet_id"`
	Time           time.Time     `json:"timestamp"`
	OriginalTweet  twitter.Tweet `json:"original_tweet"`
	AuthorID       string        `json:"author_id"`
	ReviewRequired bool          `json:"review_required"`
}

func (a *API) getTweets() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.URL.Query().Get("review_required") == "1" {
			tweet, err := a.database.GetTweetToReview()
			if err != nil {
				log.Printf("HTTP handler getTweets failed: %v", err)
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			var originalTweet twitter.Tweet
			_ = json.Unmarshal([]byte(tweet.OriginalTweetJSON), &originalTweet)

			apiFormattedTweet := apiResponseTweet{
				TweetID:        tweet.TweetID,
				Time:           tweet.Time,
				OriginalTweet:  originalTweet,
				AuthorID:       tweet.AuthorID,
				ReviewRequired: tweet.ReviewRequired,
			}

			err = json.NewEncoder(w).Encode(apiFormattedTweet)
			if err != nil {
				log.Printf("HTTP handler getTweets encoding results to JSON failed: %v", err)
				http.Error(w, "server error", http.StatusInternalServerError)
			}

			return
		}

		tweets, err := a.database.GetTweetsToDisplay()

		if err != nil {
			log.Printf("HTTP handler getTweets failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		apiFormattedTweets := make([]apiResponseTweet, 0)

		for _, tweet := range tweets {
			var originalTweet twitter.Tweet
			_ = json.Unmarshal([]byte(tweet.OriginalTweetJSON), &originalTweet)

			apiFormattedTweets = append(apiFormattedTweets, apiResponseTweet{
				TweetID:        tweet.TweetID,
				Time:           tweet.Time,
				OriginalTweet:  originalTweet,
				AuthorID:       tweet.AuthorID,
				ReviewRequired: tweet.ReviewRequired,
			})
		}

		err = json.NewEncoder(w).Encode(apiFormattedTweets)
		if err != nil {
			log.Printf("HTTP handler getTweets encoding results to JSON failed: %v", err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	}
}
