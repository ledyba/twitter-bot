package main

import (
	"flag"
	"log"
	"regexp"
	"sync"
)

var urlRegexp *regexp.Regexp
var accountRegexp *regexp.Regexp

func main() {
	var err error
	urlRegexp, err = regexp.Compile(kUrlRegexp)
	if err != nil {
		log.Fatal(err)
	}
	accountRegexp, err = regexp.Compile(kAccountRegexp)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse() // Scan the arguments list
	accounts := register()

	var wg sync.WaitGroup
	wg.Add(len(accounts))

	for _, a := range accounts {
		go a.Start(&wg)
	}
	wg.Wait()
}
