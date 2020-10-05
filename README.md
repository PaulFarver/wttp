# WTTP

## TL;DR

```shell
$ go get github.com/paulfarver/wttp
$ wttp 502
Bad Gateway
$ wttp -long 502
Bad Gateway

The server was acting as a gateway or proxy and received an invalid response from the upstream server.
$ wttp -long 5xx
Server Error

5xx error codes indicate that an error or unresolvable request occurred on the server side, whether that is a proxy or the origin host.
```