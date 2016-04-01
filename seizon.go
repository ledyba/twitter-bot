// SeizonSenryaku
package main

import (
	"fmt"
	"log"
	"regexp"
	"unicode/utf8"

	"github.com/ChimeraCoder/anaconda"
)

// SeizonSenryaku ...
type SeizonSenryaku struct {
	nextWord   string
	kanjiRegex *regexp.Regexp
}

func newSeizonSenryaku() *SeizonSenryaku {
	var err error
	e := SeizonSenryaku{}
	e.kanjiRegex, err = regexp.Compile(`\p{Han}{4,}`)
	if err != nil {
		log.Fatal(err)
	}
	return &e
}

// OnTimer ...
func (e *SeizonSenryaku) OnTimer(api *anaconda.TwitterApi) {
	if e.nextWord == "" {
		return
	}
	tw := fmt.Sprintf("%s、しましょうか！", e.nextWord)
	log.Printf("SeizonSenryaku / Say: \"%s\"", tw)
	_, err := api.PostTweet(tw, nil)
	if err != nil {
		log.Println(err)
	}
	e.nextWord = ""
}

// OnTweet ...
func (e *SeizonSenryaku) OnTweet(tw anaconda.Tweet) {
	log.Printf("Tweet: %s", tw.Text)
	text := tw.Text
	text = urlRegexp.ReplaceAllString(text, "")
	text = accountRegexp.ReplaceAllString(text, "")
	w := ""
	for _, str := range e.kanjiRegex.FindAllString(text, -1) {
		if utf8.RuneCountInString(str) == 4 {
			w = str
		}
	}
	if w == "" {
		return
	}
	e.nextWord = w
	log.Printf("SeizonSenryaku / Next: \"%s\"", e.nextWord)
}
