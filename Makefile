NAME=dnsforwarder
TITLE=DnsForwarder
BUILD := $(shell date '+%Y%m%d')
VERSION=1.0.1-$(BUILD)

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
	sudo gem install fpm --no-ri --no-rdoc
	go get -u -v github.com/mitchellh/gox

deb: clean build
	fpm -s dir -t deb -n ${NAME} -v ${VERSION} -m "Frank Yang <codebear4@gmail.com>" --url https://github.com/codebear4/dnsforwarder \
		--license MIT --vendor "Frank Yang" \
		--after-install ./scripts/postinstall \
		--after-remove ./scripts/postuninstall \
		./dnsforwarder.yml=/etc/dnsforwarder.yml \
		./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder \
		./scripts/etc/dnsforwarder=/etc/default/dnsforwarder \
		./scripts/init.d/dnsforwarder=/etc/init.d/dnsforwarder

rpm: clean build
	fpm -s dir -t rpm -n ${NAME} -v ${VERSION} -m "Frank Yang <codebear4@gmail.com>" --url https://github.com/codebear4/dnsforwarder \
		--license MIT --vendor "Frank Yang" \
		--after-install ./scripts/postinstall \
		--after-remove ./scripts/postuninstall \
		./dnsforwarder.yml=/etc/dnsforwarder.yml \
		./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder \
		./scripts/etc/dnsforwarder=/etc/default/dnsforwarder \
		./scripts/init.d/dnsforwarder=/etc/init.d/dnsforwarder
