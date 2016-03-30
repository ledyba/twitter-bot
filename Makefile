.PHONY: all get run clean deploy

all:
	gofmt -w .
	go install github.com/ledyba/twitter-bot/...

get:
	go get -u "github.com/yukihir0/mecab-go"
	go get -u "github.com/ChimeraCoder/anaconda"

run:
	$(GOPATH)/bin/twitter-bot

clean:
	go clean github.com/ledyba/twitter-bot/...

deploy: twitter-bot
	ssh ledyba.org mkdir -p /opt/run/twitter-bot
	scp twitter-bot twitter-bot.conf ledyba:/opt/run/twitter-bot

twitter-bot:
	GOOS=linux GOARCH=amd64 go build -o twitter-bot github.com/ledyba/twitter-bot/...
