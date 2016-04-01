package main

import (
	"sync"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// Charactor ...
type Charactor interface {
	OnTweet(tw anaconda.Tweet)
	OnTimer(api *anaconda.TwitterApi)
}

// Account ...
type Account struct {
	api       *anaconda.TwitterApi
	character Charactor
	ticker    *time.Ticker
}

// NewAccount ...
func NewAccount(accessToken string, accessTokenSecret string, ch Charactor) *Account {
	account := Account{}
	account.api = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	account.character = ch
	account.ticker = time.NewTicker(30 * time.Minute)
	return &account
}

// Start ...
func (ac *Account) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		ac.loop()
	}
}

func (ac *Account) loop() {
	stream := ac.api.UserStream(nil)
	for {
		select {
		case tweetRaw, received := <-stream.C:
			if !received {
				stream.Stop()
				return
			}
			if tweet, ok := tweetRaw.(anaconda.Tweet); ok {
				ac.character.OnTweet(tweet)
			}
		case <-ac.ticker.C:
			ac.character.OnTimer(ac.api)
		}
	}
}
