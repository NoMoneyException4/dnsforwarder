VERSION := $(shell git rev-parse --short HEAD)-$(shell date '+%Y%m%d')

bootstrap:
	go get -u github.com/Masterminds/glide

build:
	@glide install

