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

## Handler

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

### Encryption

There are built-in AES-128 encryption helpers.

Use [NewEncrypter](https://godoc.org/github.com/gowww/secure#NewEncrypter) with a 32 bytes long secret key to make a new encrypter:

```Go
encrypter, _ := secure.NewEncrypter("secret-key-secret-key-secret-key")
```

Use [Encrypter.Encrypt](https://godoc.org/github.com/gowww/secure#Encrypter.Encrypt) or [Encrypter.EncryptString](https://godoc.org/github.com/gowww/secure#Encrypter.EncryptString) to encrypt a value:

```Go
encryptedData, _ := encrypter.EncryptString("data to encrypt")
```

Use [Encrypter.Decrypt](https://godoc.org/github.com/gowww/secure#Encrypter.Decrypt) or [Encrypter.DecryptString](https://godoc.org/github.com/gowww/secure#Encrypter.DecryptString) to decrypt a value:

```Go
decryptedData, _ := encrypter.DecryptString("data to encrypt")
```
