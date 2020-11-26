# WTTP

## TL;DR

```shell
$ go get github.com/paulfarver/wttp
$ wttp 502
502 - Bad Gateway
The server was acting as a gateway or proxy and received an invalid response from the upstream server.
$ wttp list
100 - Continue
101 - Switching Protocols
102 - Processing
103 - Early Hints
1xx - Information
200 - OK
201 - Created
202 - Accepted
...
526 - Invalid SSL Certificate
527 - Railgun Listener to Origin Error
530 - Origin DNS Error
598 - Network Read Timeout Error
5xx - Server Error
```
