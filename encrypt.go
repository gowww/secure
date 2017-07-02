// Package secure provides security utilities.
package secure

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// An Encrypter provides encryption and decryption functions.
type Encrypter interface {
	Encrypt([]byte) ([]byte, error)
	EncryptString(string) (string, error)
	Decrypt([]byte) ([]byte, error)
	DecryptString(string) (string, error)
}

// NewEncrypter returns a new AES-128 encrypter for secret key.
// The key must be 32 bytes long.
func NewEncrypter(key string) (Encrypter, error) {
	if len(key) != 32 {
		panic("secure: encrypter key must be 32 bytes long")
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return &encrypter{aead: gcm}, nil
}

type encrypter struct {
	aead cipher.AEAD
}

// Encrypt encrypts a plain text.
func (e *encrypter) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, e.aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return e.aead.Seal(nonce, nonce, plaintext, nil), nil
}

// EncryptString encrypts a plain string.
func (e *encrypter) EncryptString(plaintext string) (string, error) {
	b, err := e.Encrypt([]byte(plaintext))
	return string(b), err
}

// Decrypt decrypts an encrypted text.
func (e *encrypter) Decrypt(ciphertext []byte) ([]byte, error) {
	nonceSize := e.aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return e.aead.Open(nil, nonce, ciphertext, nil)
}

// DecryptString decrypts an encrypted string.
func (e *encrypter) DecryptString(ciphertext string) (string, error) {
	b, err := e.Decrypt([]byte(ciphertext))
	return string(b), err
}
