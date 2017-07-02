# [![gowww](https://avatars.githubusercontent.com/u/18078923?s=20)](https://github.com/gowww) secure [![GoDoc](https://godoc.org/github.com/gowww/secure?status.svg)](https://godoc.org/github.com/gowww/secure) [![Build](https://travis-ci.org/gowww/secure.svg?branch=master)](https://travis-ci.org/gowww/secure) [![Coverage](https://coveralls.io/repos/github/gowww/secure/badge.svg?branch=master)](https://coveralls.io/github/gowww/secure?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/gowww/secure)](https://goreportcard.com/report/github.com/gowww/secure)

Package [secure](https://godoc.org/github.com/gowww/secure) provides utilities to encrypt secures.

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

Use [Encrypt](https://godoc.org/github.com/gowww/secure#Encrypt) to encrypt a secure value:

```Go
ck := &http.Secure{}
ck = secure.Encrypt(ck)
```
