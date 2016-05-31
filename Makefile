VERSION := $(shell git rev-parse --short HEAD)-$(shell date '+%Y%m%d')

bootstrap:
	go get -u -v github.com/Masterminds/glide
	glide install

test:
	go test -v

build:
	gox -osarch="linux/amd64"

clean:
	rm -rf dnsforwarder
	rm -rf dnsforwarder_*
	rm -rf *.deb

package-bootstrap:
	@brew install gnu-tar
	sudo gem install fpm
	go get -u -v github.com/mitchellh/gox

package: build
	fpm -s dir -t deb -n dnsforwarder -v 1.0.1 ./dnsforwarder.yml=/etc/dnsforwarder.yml ./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder
