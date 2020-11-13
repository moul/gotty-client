# gotty-client
:wrench: Terminal client for [GoTTY](https://github.com/yudai/gotty).

![](https://raw.githubusercontent.com/moul/gotty-client/master/resources/gotty-client.png)

[![Build Status](https://travis-ci.org/moul/gotty-client.svg?branch=master)](https://travis-ci.org/moul/gotty-client)
[![GoDoc](https://godoc.org/github.com/moul/gotty-client?status.svg)](https://godoc.org/github.com/moul/gotty-client)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fmoul%2Fgotty-client.svg?type=shield)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fmoul%2Fgotty-client?ref=badge_shield)
[![Sourcegraph](https://sourcegraph.com/github.com/moul/gotty-client/-/badge.svg)](https://sourcegraph.com/github.com/moul/gotty-client?badge)

```ruby
                                                              +----------------+
               +--------------+                +---------+--->|   /bin/bash    |
               |              |                |         |    +----------------+
           +-->|   Browser    |--+             |         |
+-------+  |   |              |  |             |         |
|       |  |   +--------------+  |             |         |    +----------------+
|  Bob  |--+                     +-websocket-->|  Gotty  |--->| emacs /var/www |
|       |  |    XXXXXXXXXXXXXX   |             |         |    +----------------+
+-------+  |   X              X  |             |         |
           +-->X gotty-client X--+             |         |
               X              X                |         |    +----------------+
                XXXXXXXXXXXXXX                 +---------+--->|  tmux attach   |
                                                              +----------------+
```

## Example

Server side ([GoTTY](https://github.com/yudai/gotty))

```console
$ gotty -p 9191 sh -c 'while true; do date; sleep 1; done'
2015/08/24 18:54:31 Server is starting with command: sh -c while true; do date; sleep 1; done
2015/08/24 18:54:31 URL: http://[::1]:9191/
2015/08/24 18:54:34 GET /ws
2015/08/24 18:54:34 New client connected: 127.0.0.1:61811
2015/08/24 18:54:34 Command is running for client 127.0.0.1:61811 with PID 64834
2015/08/24 18:54:39 Command exited for: 127.0.0.1:61811
2015/08/24 18:54:39 Connection closed: 127.0.0.1:61811
...
```

**Client side**

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

## Usage

```console
$ gotty-client -h
NAME:
   gotty-client - GoTTY client for your terminal

USAGE:
   gotty-client [global options] command [command options] GOTTY_URL

AUTHOR:
   Manfred Touron <https://github.com/moul/gotty-client>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -D                  Enable debug mode [$GOTTY_CLIENT_DEBUG]
   --skip-tls-verify            Skip TLS verify [$SKIP_TLS_VERIFY]
   --use-proxy-from-env         Use Proxy from environment [$USE_PROXY_FROM_ENV]
   --detach-keys value          Key sequence for detaching gotty-client (default: "ctrl-p,ctrl-q")
   --v2                         For Gotty 2.0 [$GOTTY_CLIENT_GOTTY2]
   --ws-origin value, -w value  WebSocket Origin URL [$GOTTY_CLIENT_WS_ORIGIN]
   --help, -h                   show help
   --version, -v                print the version
```

## Install

Install latest version using Golang (recommended):

```bash
git clone https://github.com/moul/gotty-client
cd gotty-client
make install # or `go install ./cmd/gotty-client`
```

_PS: I don't know why, but `go get github.com/moul/gotty-client/cmd/gotty-client` is not stable everywhere_

---

Install latest version using Homebrew (Mac OS X)

```console
$ brew install https://raw.githubusercontent.com/moul/gotty-client/master/contrib/homebrew/gotty-client.rb --HEAD
```

or the latest released version

```console
$ brew install https://raw.githubusercontent.com/moul/gotty-client/master/contrib/homebrew/gotty-client.rb
```

## Changelog

See https://github.com/moul/gotty-client/releases

## License

MIT
