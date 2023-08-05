NAME=dnsforwarder
TITLE=DnsForwarder
BUILD := $(shell date '+%Y%m%d')
VERSION=1.0.6-$(BUILD)

test:
	go test -v

build:
	go build -ldflags="-X main.version=${VERSION} -w -s"

build-linux:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=${VERSION} -w -s"

clean:
	rm -rf dnsforwarder
	rm -rf dnsforwarder_*
	rm -rf *.deb *.rpm

package-bootstrap:
	sudo gem install fpm --no-ri --no-rdoc

deb: clean build-linux
	fpm -s dir -t deb -n ${NAME} -v ${VERSION} -m "Frank Yang <Yangshifu1024@qq.com>" --url https://github.com/Yangshifu1024/dnsforwarder \
		--license MIT --vendor "Frank Yang" \
		--after-install ./scripts/postinstall \
		--after-remove ./scripts/postuninstall \
		./dnsforwarder.yml=/etc/dnsforwarder.yml \
		./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder \
		./scripts/etc/dnsforwarder=/etc/default/dnsforwarder \
		./scripts/init.d/dnsforwarder=/etc/init.d/dnsforwarder \
		./scripts/systemd/dnsforwarder.service=/lib/systemd/system/dnsforwarder.service

rpm: clean build-linux
	fpm -s dir -t rpm -n ${NAME} -v ${VERSION} -m "Frank Yang <Yangshifu1024@qq.com>" --url https://github.com/Yangshifu1024/dnsforwarder \
		--license MIT --vendor "Frank Yang" \
		--after-install ./scripts/postinstall \
		--after-remove ./scripts/postuninstall \
		./dnsforwarder.yml=/etc/dnsforwarder.yml \
		./dnsforwarder_linux_amd64=/usr/bin/dnsforwarder \
		./scripts/etc/dnsforwarder=/etc/default/dnsforwarder \
		./scripts/init.d/dnsforwarder=/etc/init.d/dnsforwarder \
		./scripts/systemd/dnsforwarder.service=/lib/systemd/system/dnsforwarder.service
