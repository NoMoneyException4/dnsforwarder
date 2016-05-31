VERSION := $(shell git rev-parse --short HEAD)-$(shell date '+%Y%m%d')

bootstrap:
	go get -u github.com/Masterminds/glide
	glide install

test:
	go test -v

build:
	go build -v
