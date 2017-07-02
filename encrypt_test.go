package secure

import (
	"testing"
)

var (
	testEncrypterKey       = "secret-key-secret-key-secret-key"
	testEncrypterPlaintext = "foobar"
)

func TestEncrypter(t *testing.T) {
	e, err := NewEncrypter(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	b, err := e.Encrypt([]byte(testEncrypterPlaintext))
	if err != nil {
		panic(err)
	}
	b, err = e.Decrypt(b)
	if err != nil {
		panic(err)
	}
	if testEncrypterPlaintext != string(b) {
		t.Errorf("encode/decode: want %s, got %s", testEncrypterPlaintext, b)
	}
}

func TestEncrypterString(t *testing.T) {
	e, err := NewEncrypter(testEncrypterKey)
	if err != nil {
		panic(err)
	}
	s, err := e.EncryptString(testEncrypterPlaintext)
	if err != nil {
		panic(err)
	}
	s, err = e.DecryptString(s)
	if err != nil {
		panic(err)
	}
	if testEncrypterPlaintext != string(s) {
		t.Errorf("encode/decode string: want %s, got %s", testEncrypterPlaintext, s)
	}
}
