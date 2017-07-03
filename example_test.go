package secure_test

import (
	"fmt"
	"net/http"

	"github.com/gowww/secure"
)

func Example() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	http.ListenAndServe(":8080", secure.Handle(mux, &secure.Options{
		SSLForced:      true,
		EnvDevelopment: true,
	}))
}

func ExampleHandle() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})

	http.ListenAndServe(":8080", nil)
}

func ExampleHandleFunc() {
	http.Handle("/", secure.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}, nil))

	http.ListenAndServe(":8080", nil)
}

func ExampleEncrypter() {
	encrypter, _ := secure.NewEncrypter("secret-key-secret-key-secret-key")
	data, _ := encrypter.EncryptString("data to encrypt")
	data, _ = encrypter.DecryptString(data)
	fmt.Println(data)
}
