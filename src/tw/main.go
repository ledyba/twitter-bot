package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	mecab "github.com/yukihir0/mecab-go"
	"log"
	"os"
	"strings"
	"regexp"
	"math/rand"
)

const APP_VERSION = "0.1"

func getCred() {
	key, cred,_ := anaconda.AuthorizationURL("")
	fmt.Printf("access: %v\n", key)
	fmt.Printf("Key: ")
	buf := bufio.NewReader(os.Stdin)
	line, _, _ := buf.ReadLine()
	cred, _, _ = anaconda.GetCredentials(cred, string(line))
	fmt.Printf("cred: \n", cred)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

func count(target, s string) int {
	cnt := 0
	for _, a := range []rune(s) {
		if strings.ContainsRune(target, a) {
			cnt++
		}
	}
	return cnt
}

var urlRegexp *regexp.Regexp
func onTweet(api *anaconda.TwitterApi, tw anaconda.Tweet) {
	log.Printf("Tweet: %s", tw.Text)
	var err error
	if urlRegexp == nil{
		urlRegexp, err = regexp.Compile(kUrlRegexp)
		if err != nil{
			log.Fatal(err)
		}
	}
	text := tw.Text
	text = urlRegexp.ReplaceAllString(text, "")
	rs, err := mecab.Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	strs := make([]string, len(rs))
	for i, r := range rs {
		strs[i]=r.Surface
	}
	log.Print(strs)
	alphas := 0
	for _, r := range rs {
		rlen := len([]rune(r.Surface))
		if count(kAlphabet, r.Surface) > (rlen/2) {
			alphas++
		}
	}
	if alphas > (len(rs)/2){
		return
	}
	candidates := make([]string, 0)
	for _, r := range rs {
		if contains(kStopWords, r.Surface){
			return
		}
		rlen := len([]rune(r.Surface))
		syms := count(kSymbols, r.Surface)
		if rlen <= 1 || (syms) > (rlen / 2) {
			continue
		}
		f := r.Feature
		if !(r.Pos == "名詞" && !strings.Contains(f,"接尾") &&
							!strings.Contains(f,"代名詞") &&
							!strings.Contains(f,"接頭詞") &&
							!strings.Contains(f,"形容動詞語幹")){
			continue
		}
		candidates = append(candidates, r.Surface)
	}
	if len(candidates) <= 0{
		continue
	}
	c := candidates[rand.Intn(len(candidates))]
	log.Printf("%s as a Service\n", c)
	t, err := api.PostTweet(c + " as a Service", nil)
	if err != nil{
		log.Println(err)
	}
	log.Println(t)
	os.Exit(0)

}

func main() {
	flag.Parse() // Scan the arguments list

	anaconda.SetConsumerKey(ConsumerKey)
	anaconda.SetConsumerSecret(ConsumerSecret)
	tw := anaconda.NewTwitterApi(OAuthToken, OAuthSecret)
	defer tw.Close()
	stream := tw.UserStream(nil)
	for {
		select {
		case tweet_ := <-stream.C:
			if tweet, ok := tweet_.(anaconda.Tweet); ok {
				onTweet(tw, tweet)
			}
		case <-stream.Quit:
			break
		}
	}
}
