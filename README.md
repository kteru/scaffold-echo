scaffold-echo
=============

Build & Run
-----------

### Install mattn/gom (if necessary)

```
$ go get github.com/mattn/gom
```

### Install vendor packages

```
$ gom install
```

### Build

```
$ go build
```

### Run

```
$ ./scaffold-echo [-h 127.0.0.1] [-p 8080] [-s secret]
```

Request
-------

```
$ curl -c /tmp/cookie -b /tmp/cookie -i -X POST 'http://127.0.0.1:8080/api/v1/login' -d '{"username":"taro","password":"1234"}'

$ curl -c /tmp/cookie -b /tmp/cookie -i -X GET 'http://127.0.0.1:8080/api/v1/hello'

$ curl -c /tmp/cookie -b /tmp/cookie -i -X POST 'http://127.0.0.1:8080/api/v1/logout'

$ curl -c /tmp/cookie -b /tmp/cookie -i -X GET 'http://127.0.0.1:8080/api/v1/hello'
```
