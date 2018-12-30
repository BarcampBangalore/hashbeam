package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	twit "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"server/conf"
	"server/db"
	"server/models"
	"sync"
)

type TwitterContext struct {
	database      *db.DBContext
	config        conf.Twitter
	twitterClient *twit.Client
	allowedUserIDsMutex sync.RWMutex
	allowedUserIDs map[string]bool
}

func NewTwitterContext(database *db.DBContext, config conf.Twitter) (*TwitterContext, error) {
	oauthConfig := oauth1.NewConfig(config.ConsumerKey, config.ConsumerSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessTokenSecret)
	twitterClient := twit.NewClient(oauthConfig.Client(oauth1.NoContext, token))

	twitterContext := &TwitterContext{
		database:       database,
		config:         config,
		twitterClient:  twitterClient,
		allowedUserIDs: make(map[string]bool),
	}

	err := twitterContext.updatedAllowedUserIDs()
	if err != nil {
		return nil, fmt.Errorf("twitter NewTwitterContext initial fetch of allowed user IDs failed: %v", err)
	}

	return twitterContext, nil
}

func (ctx *TwitterContext) updatedAllowedUserIDs() error {
	allowedUserIDsSlice, err := ctx.database.GetWhitelistTwitterIDs()
	if err != nil {
		return fmt.Errorf("twitter updateAllowedUserIDs getting allowed user IDs from database failed: %v", err)
	}

	allowedUserIDsMap := make(map[string]bool)
	for _, userID := range allowedUserIDsSlice {
		allowedUserIDsMap[userID] = true
	}

	ctx.allowedUserIDsMutex.Lock()
	ctx.allowedUserIDs = allowedUserIDsMap
	ctx.allowedUserIDsMutex.Unlock()
	return nil
}

func (ctx *TwitterContext) onTweet(t *twit.Tweet) {
	jsonBody := &bytes.Buffer{}
	err := json.NewEncoder(jsonBody).Encode(t)
	if err != nil {
		log.Printf("twitter onTweet encoding tweet body to JSON failed: %v", err)
	}

	reviewRequired := true
	ctx.allowedUserIDsMutex.RLock()
	if _, exists := ctx.allowedUserIDs[t.User.IDStr]; exists {
		reviewRequired = false
	}
	ctx.allowedUserIDsMutex.RUnlock()

	time, _ := t.CreatedAtTime()

	err = ctx.database.SaveTweet(models.Tweet{
		TweetID:        t.IDStr,
		AuthorID:       t.User.IDStr,
		Time:           time,
		JSONBody:       jsonBody.Bytes(),
		ReviewRequired: reviewRequired,
	})

	if err != nil {
		log.Printf("twitter onTweet saving tweet to database failed: %v", err)
	}
}

func (ctx *TwitterContext) StartStream() error {
	params := &twit.StreamFilterParams{
		Track: []string{ctx.config.TextToTrack},
	}

	demux := twit.NewSwitchDemux()
	demux.Tweet = ctx.onTweet

	stream, err := ctx.twitterClient.Streams.Filter(params)
	if err != nil {
		return fmt.Errorf("twitter StartStream failed: %v", err)
	}

	// Receive messages until app quits
	go demux.HandleChan(stream.Messages)
	return nil
}
