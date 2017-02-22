# DnsForwarder

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/codebear4/dnsforwarder/master/LICENSE)
[![Join the chat at https://gitter.im/codebear4/dnsforwarder](https://badges.gitter.im/codebear4/dnsforwarder.svg)](https://gitter.im/codebear4/dnsforwarder?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://semaphoreci.com/api/v1/codebear4/dnsforwarder/branches/master/shields_badge.svg)](https://semaphoreci.com/codebear4/dnsforwarder)
[![Go Report Card](https://goreportcard.com/badge/github.com/codebear4/dnsforwarder)](https://goreportcard.com/report/github.com/codebear4/dnsforwarder)
[![Github All Releases](https://img.shields.io/github/downloads/codebear4/dnsforwarder/total.svg?maxAge=2592000)](https://github.com/codebear4/dnsforwarder/releases)

A dns forwarder.
* Resolve with multiple upstreams
* Caching the record from upstreams
* Support local hosts files
* White list support

## Getting started
#### Installation
##### Build from source
Make sure you have a correctly configured Go installtion first, then:
```s
$ make bootstrap
$ go build
```

##### Install with package
Download packages from github [releases](https://github.com/codebear4/dnsforwarder/releases) page.

#### Configuration
`forcetcp`: If set it to `true`, server will use `TCP` connection with upstreams even the request from      client is `UDP` request.

`cache`:
* `enable`: Enable or disable in-memory cache
* `ttl`: Default TTL for cache

`loggers`:
* `console`:
    * `enable`: Enable or disable console logger
    * `level`: Log level for console logger
* `file`:
    * `enable`: Enable or disable file logger
    * `level`: Log level for file logger
    * `path`: Log file path

`timeout`:
* `server`:
    * `read`: Timeout for read from clients
    * `write`: Timeout for write to clients
* `forwarder`:
    * `read`: Timeout for read from upstreams
    * `write`: Timeout for write to upstreams

`host`:
* `enable`: Resolve with hosts files
* `resolves`: Hosts files list

`upstreams`: List of upstream dns servers

#### Running

##### Build from source
```s
$ sudo dnsforwarder
```

##### Install from package
```s
$ service dnsforwarder start # upstart
$ systemctl start dnsforwarder #systemd
```

#### Testing
```s
$ dig domain.tld @your.server.ip
```

## TODO
* Cache persistence
* ~~Packaging~~
* Hot-reload

