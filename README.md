# Usage of "ok"

## Files

- ok.go
  - A simple http server just reply with `{"status":"OK"}` with 200 OK status.

## Description

`ok.go` is a HTTP server which simply replies with `{"status":"OK"}`.
This server is quite simple as it can be used as the target of reverse proxy
for benchmarking or stress tests.

## How to Use

Build the application.

```bash
$ cd <this folder>
$ go build ok.go
```

Run the server.
(Ignore `.exe` extension for linux or mac.)

```bash
$ ./ok.exe
```

Send an HTTP request to the server.
Server's port is `:10001`.

```bash
$ curl -v localhost:10001
*   Trying 127.0.0.1:10001...
* Connected to localhost (127.0.0.1) port 10001 (#0)
> GET / HTTP/1.1
> Host: localhost:10001
> User-Agent: curl/8.0.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Fri, 12 May 2023 10:04:44 GMT
< Content-Length: 15
< Content-Type: text/plain; charset=utf-8
<
{"status":"OK"}
```
