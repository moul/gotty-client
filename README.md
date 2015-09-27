# gotty-client
:wrench: [GoTTY](https://github.com/yudai/gotty) client for your terminal

[![Build Status](https://travis-ci.org/moul/gotty-client.svg?branch=master)](https://travis-ci.org/moul/gotty-client)
[![GoDoc](https://godoc.org/github.com/moul/gotty-client?status.svg)](https://godoc.org/github.com/moul/gotty-client)

```ruby
                +----------------+       +----------------+      +-------------+
                |                |       |             +--------->  /bin/bash  |
            +--->    Browser    -----+   |     gotty   |  |      +-------------+
+-------+   |   |                |   |   |             |  |
|       |   |   +----------------+   |   |             |  |      +-------------+
|  Bob  +---+                        +--->---websockets+--------->  /bin/bash  |
|       |   |   +================+   |   |             |  |      +-------------+
+-------+   |   |................|   |   |             |  |
            +--->..gotty-client.-----+   |             |  |      +-------------+
                |................|       |             +--------->  /bin/bash  |
                +================+       +----------------+      +-------------+

                  ^  ^  ^  ^  ^
                  |  |  |  |  |
```


## Install

```bash
go get github.com/moul/gotty-client/cmd/gotty-client
```

## Usage

Server side

```console
$ gotty -p 9191 sh -c 'while true; do date; sleep 1; done'
2015/08/24 18:54:31 Server is starting with command: sh -c while true; do date; sleep 1; done
2015/08/24 18:54:31 URL: http://[::1]:9191/
2015/08/24 18:54:31 URL: http://127.0.0.1:9191/
2015/08/24 18:54:31 URL: http://[fe80::1]:9191/
2015/08/24 18:54:31 URL: http://[fe80::ee35:86ff:fe4c:cf5c]:9191/
2015/08/24 18:54:31 URL: http://192.168.86.131:9191/
2015/08/24 18:54:31 URL: http://[2a01:e35:2435:3b00:ee35:86ff:fe4c:cf5c]:9191/
2015/08/24 18:54:31 URL: http://[2a01:e35:2435:3b00:6563:e6a2:dbab:5ade]:9191/
2015/08/24 18:54:31 URL: http://[fe80::a895:4fff:fef2:d4c7]:9191/
2015/08/24 18:54:31 URL: http://[fe80::4882:98b5:853c:15be]:9191/
2015/08/24 18:54:31 URL: http://[fd87:46a0:dbdf:90bb:4882:98b5:853c:15be]:9191/
2015/08/24 18:54:31 URL: http://[fe80::8bf9:b113:3444:98b4]:9191/
2015/08/24 18:54:31 URL: http://192.168.2.1:9191/
2015/08/24 18:54:31 URL: http://[fe80::10dd:b1ff:fe2c:7764]:9191/
2015/08/24 18:54:31 URL: http://172.16.252.11:9191/
2015/08/24 18:54:31 URL: http://192.168.99.1:9191/
2015/08/24 18:54:31 URL: http://[fe80::7bfb:c14b:2534:3bc5]:9191/
2015/08/24 18:54:31 URL: http://[fe80::34a3:95ff:fec7:97d5]:9191/
2015/08/24 18:54:31 URL: http://172.20.10.2:9191/
2015/08/24 18:54:34 GET /ws
2015/08/24 18:54:34 New client connected: 127.0.0.1:61811
2015/08/24 18:54:34 Command is running for client 127.0.0.1:61811 with PID 64834
2015/08/24 18:54:39 Command exited for: 127.0.0.1:61811
2015/08/24 18:54:39 Connection closed: 127.0.0.1:61811
```

Client side
```console
$ gotty-client http://localhost:9191/
INFO[0000] New title: GoTTY - sh -c while true; do date; sleep 1; done (jean-michel-van-damme.local)
WARN[0000] Unhandled protocol message: json pref: 2{}
Mon Aug 24 18:54:34 CEST 2015
Mon Aug 24 18:54:35 CEST 2015
Mon Aug 24 18:54:36 CEST 2015
Mon Aug 24 18:54:37 CEST 2015
Mon Aug 24 18:54:38 CEST 2015
^C
```

## Changelog

### master (unreleased)

* Add 'homebrew support' ([#1](https://github.com/moul/gotty-client/issues/1))
* Add Changelog ([#5](https://github.com/moul/gotty-client/issues/5))
* Add GOXC configuration to build binaries for multiple architectures ([#2](https://github.com/moul/gotty-client/issues/2))

[full commits list](https://github.com/moul/gotty-client/compare/v1.0.1...master)

### [v1.0.1](https://github.com/moul/gotty-client/releases/tag/v1.0.1) (2015-09-27)

* Using party to manage dependencies

[full commits list](https://github.com/moul/gotty-client/compare/v1.0.0...v1.0.1)

### [v1.0.0](https://github.com/moul/gotty-client/releases/tag/v1.0.0) (2015-09-27)

Compatible with [GoTTY](https://github.com/yudai/gotty) version: [v0.0.10](https://github.com/yudai/gotty/releases/tag/v0.0.10)

#### Features

* Support **basic-auth**
* Support **terminal-(re)size**
* Support **write**
* Support **title**
* Support **custom URI**

[full commits list](https://github.com/moul/gotty-client/compare/cf0c1146c7ce20fe0bd65764c13253bc575cd43a...v1.0.0)

## License

MIT
