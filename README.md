# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) secure [![GoDoc](https://godoc.org/github.com/gowww/secure?status.svg)](https://godoc.org/github.com/gowww/secure) [![Build](https://travis-ci.org/gowww/secure.svg?branch=master)](https://travis-ci.org/gowww/secure) [![Coverage](https://coveralls.io/repos/github/gowww/secure/badge.svg?branch=master)](https://coveralls.io/github/gowww/secure?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/secure)](https://goreportcard.com/report/github.com/gowww/secure)

Package [secure](https://godoc.org/github.com/gowww/secure) provides security utilities, CSP, HPKP, HSTS and other security wins.

## Installing

1. Get package:

	```Shell
	go get -u github.com/gowww/secure
	```

2. Import it in your code:

	```Go
	import "github.com/gowww/secure"
	```

## Usage

To wrap an [http.Handler](https://golang.org/pkg/net/http/#Handler), use [Handle](https://godoc.org/github.com/gowww/secure#Handle):

```Go
mux := http.NewServeMux()

mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
})

http.ListenAndServe(":8080", secure.Handle(mux, nil))
```

To wrap an [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc), use [HandleFunc](https://godoc.org/github.com/gowww/secure#HandleFunc):

```Go
http.Handle("/", secure.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}, nil))

http.ListenAndServe(":8080", nil)
```

To set custom security options, see [Options](https://godoc.org/github.com/gowww/secure#Options) GoDoc reference.
