NAME=dnsforwarder
VERSION := $(shell git rev-parse --short HEAD)_$(shell date '+%Y%m%d')

bootstrap:
	go get -u -v github.com/Masterminds/glide
	glide install

test:
	go test -v

build:
	gox -osarch="darwin/amd64 linux/amd64" -ldflags="-X main.version=${VERSION} -w -s"

clean:
	rm -rf dnsforwarder
	rm -rf dnsforwarder_*
	rm -rf *.deb *.rpm

package-bootstrap:
	sudo gem install fpm
	go get -u -v github.com/mitchellh/gox

deb: clean build
	fpm -s dir -t deb -n ${NAME} -v ${VERSION} ./dnsforwarder.yml=/etc/dnsforwarder.yml ./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder

rpm: clean build
	fpm -s dir -t rpm -n ${NAME} -v ${VERSION} ./dnsforwarder.yml=/etc/dnsforwarder.yml ./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder
