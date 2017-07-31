CWD=$(shell pwd)
GOPATH := $(CWD)

build:	fmt bin

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-spr/whosonfirst
	cp whosonfirst/*.go src/github.com/whosonfirst/go-whosonfirst-spr/whosonfirst/
	cp *.go src/github.com/whosonfirst/go-whosonfirst-spr
	cp -r vendor/src/* src/
	cp -r vendor/src/github.com/whosonfirst/go-whosonfirst-geojson-v2/vendor/src/github.com/tidwall src/github.com/
	cp -r vendor/src/github.com/whosonfirst/go-whosonfirst-geojson-v2/vendor/src/github.com/whosonfirst/go-whosonfirst-placetypes src/github.com/whosonfirst/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:   rmdeps
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-geojson-v2"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"

vendor-deps: deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt *.go
	go fmt cmd/*.go
	go fmt whosonfirst/*.go

bin:	self
	@GOPATH=$(GOPATH) go build -o bin/wof-feature-to-spr cmd/wof-feature-to-spr.go
