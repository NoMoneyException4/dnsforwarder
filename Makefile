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
	rm -rf *.deb *.rpm

package-bootstrap:
	sudo gem install fpm
	go get -u -v github.com/mitchellh/gox

deb: clean build
	fpm -s dir -t deb -n dnsforwarder -v 1.0.1 ./dnsforwarder.yml=/etc/dnsforwarder.yml ./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder

rpm: clean build
	fpm -s dir -t rpm -n dnsforwarder -v 1.0.1 ./dnsforwarder.yml=/etc/dnsforwarder.yml ./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder
