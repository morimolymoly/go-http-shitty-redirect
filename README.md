# go-http-shitty-redirect

## Usage
```
make build
make up
```

## `http://localhost:8080/ok`
```
☻  http --follow --all http://localhost:8080/ok
HTTP/1.1 302 Found
Content-Length: 41
Content-Type: text/html; charset=utf-8
Date: Wed, 25 Jul 2018 05:02:54 GMT
Location: http://google.com/

<a href="http://google.com/">Found</a>.

HTTP/1.1 301 Moved Permanently
Cache-Control: public, max-age=2592000
Content-Length: 219
Content-Type: text/html; charset=UTF-8
Date: Wed, 25 Jul 2018 05:02:51 GMT
Expires: Fri, 24 Aug 2018 05:02:51 GMT
Location: http://www.google.com/
Server: gws
X-Frame-Options: SAMEORIGIN
X-XSS-Protection: 1; mode=block

<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="http://www.google.com/">here</A>.
</BODY></HTML>

HTTP/1.1 200 OK
Cache-Control: private, max-age=0
Content-Encoding: gzip
Content-Length: 4987
Google Below...
```

## `http://localhost:8080/shitty/redirect`
```
HTTP/1.1 302 Found
Content-Length: 41
Content-Type: text/html; charset=utf-8
Date: Wed, 25 Jul 2018 05:04:41 GMT
Location: /shitty/google.com

<a href="/shitty/google.com">Found</a>.

HTTP/1.1 404 Not Found
Content-Length: 19
Content-Type: text/plain; charset=utf-8
Date: Wed, 25 Jul 2018 05:04:41 GMT
X-Content-Type-Options: nosniff

404 page not found
```
# LICENSE
MIT
