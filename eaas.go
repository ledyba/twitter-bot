// Everything as a Service
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/yukihir0/mecab-go"
)

// EaaS is
type EaaS struct {
	nextWord string
}

func newEaaS() *EaaS {
	e := EaaS{}

	return &e
}

// OnTimer ...
func (e *EaaS) OnTimer(api *anaconda.TwitterApi) {
	if e.nextWord == "" {
		return
	}
	tw := fmt.Sprintf("%s as a Service", e.nextWord)
	log.Printf("Everything as a Service / Say: \"%s\"", tw)
	_, err := api.PostTweet(tw, nil)
	if err != nil {
		log.Println(err)
	}
	e.nextWord = ""
}

// OnTweet ...
func (e *EaaS) OnTweet(tw anaconda.Tweet) {
	log.Printf("Tweet: %s", tw.Text)
	var err error
	text := tw.Text
	text = urlRegexp.ReplaceAllString(text, "")
	text = accountRegexp.ReplaceAllString(text, "")
	rs, err := mecab.Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	strs := make([]string, len(rs))
	for i, r := range rs {
		strs[i] = r.Surface
	}
	log.Print(strs)
	alphas := 0
	for _, r := range rs {
		rlen := len([]rune(r.Surface))
		if count(kAlphabet, r.Surface) > (rlen / 2) {
			alphas++
		}
	}
	if alphas > (len(rs) / 2) {
		return
	}
	var candidates []string
	for _, r := range rs {
		if contains(kStopWords, r.Surface) {
			return
		}
		rlen := len([]rune(r.Surface))
		syms := count(kSymbols, r.Surface)
		if rlen <= 1 || (syms) > (rlen/2) {
			continue
		}
		f := r.Feature
		if !(r.Pos == "名詞" && !strings.Contains(f, "接尾") &&
			!strings.Contains(f, "代名詞") &&
			!strings.Contains(f, "接頭詞") &&
			!strings.Contains(f, "形容動詞語幹")) {
			continue
		}
		candidates = append(candidates, r.Surface)
	}
	if len(candidates) <= 0 {
		return
	}
	c := candidates[rand.Intn(len(candidates))]
	e.nextWord = c
	log.Printf("Everything as a Service / Next: \"%s\"", e.nextWord)
}
