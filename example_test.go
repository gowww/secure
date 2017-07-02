package secure_test

import (
	"fmt"

	"github.com/gowww/secure"
)

func ExampleEncrypter() {
	encrypter, _ := secure.NewEncrypter("secret-key-secret-key-secret-key")
	data, _ := encrypter.EncryptString("data to encrypt")
	data, _ = encrypter.DecryptString(data)
	fmt.Println(data)
}
