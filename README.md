# DnsForwarder

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/codebear4/dnsforwarder/master/LICENSE)
[![Join the chat at https://gitter.im/codebear4/dnsforwarder](https://badges.gitter.im/codebear4/dnsforwarder.svg)](https://gitter.im/codebear4/dnsforwarder?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://semaphoreci.com/api/v1/codebear4/dnsforwarder/branches/master/shields_badge.svg)](https://semaphoreci.com/codebear4/dnsforwarder)

A dns server, cacher, and forwarder.
* Resolve with multiple upstreams
* Caching the record from upstreams
* Support local hosts files

## Getting started
#### Installation
Make sure you have a correctly configured Go installtion first, then:
```s
$ go get -u -v github.com/codebear4/dnsforwarder
```

#### Configuration
`udpOverTcp`: If set it to `true`, server will use `TCP` connection with upstreams even the request from      client is `UDP` request.

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
* `resolvs`: Hosts files list

`upstreams`: List of upstream dns servers

#### Running
```s
$ sudo dnsforwarder
```

