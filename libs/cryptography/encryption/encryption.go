package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type encryption struct {
	hashedPass []byte
}

func createEncryption(hashedPass []byte) Encryption {
	out := encryption{
		hashedPass: hashedPass,
	}

	return &out
}

// Encrypt encrypts a message
func (obj *encryption) Encrypt(message []byte) (string, error) {
	block, blockErr := aes.NewCipher(obj.hashedPass)
	if blockErr != nil {
		return "", blockErr
	}

	ciphertext := make([]byte, aes.BlockSize+len(message))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], message)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts a message
func (obj *encryption) Decrypt(encryptedText string) ([]byte, error) {
	cipherText, cipherTextErr := base64.StdEncoding.DecodeString(encryptedText)
	if cipherTextErr != nil {
		return nil, cipherTextErr
	}

	block, blockErr := aes.NewCipher(obj.hashedPass)
	if blockErr != nil {
		return nil, blockErr
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("the encrypted text cannot be decoded using this password: ciphertext block size is too short")
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	// returns the decoded message:
	return cipherText, nil
}
