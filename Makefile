.PHONY: all get run clean

all:
	gofmt -w src/
	GOPATH=$(shell pwd):${GOPATH} GOBIN=$(shell pwd)/bin go install -v -gcflags -N ./...
get:
	GOPATH=$(shell pwd):${GOPATH} GOBIN=$(shell pwd)/bin go get github.com/yukihir0/mecab-go
	GOPATH=$(shell pwd):${GOPATH} GOBIN=$(shell pwd)/bin go get github.com/ChimeraCoder/anaconda

run:
	bin/tw

clean:
	GOPATH=$(shell pwd):${GOPATH} GOBIN=$(shell pwd)/bin go clean
