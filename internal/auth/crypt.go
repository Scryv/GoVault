package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

const SaltLength = 16 //length salt Const cause needs to be a fixed length
func GenRandoSalt(saltLength int) []byte { //func for creating random salt
	var salt = make([]byte, saltLength) // makes a byte slice variable called salt
	rand.Read(salt)                     //reads the slice and fully changes it and ads its own rando value

	return salt //returns salts
}
func HashPasswd(passwd string, salt []byte) string {
	var passwdBytes = []byte(passwd)           //creates byte slice of the passwd str
	passwdBytes = append(passwdBytes, salt...) //appends and the ... is for since salt is a slice
	hash := sha512.Sum512(passwdBytes)         //hashes the slice using sha512
	return hex.EncodeToString(hash[:])         //encodes to readable and [:] to change [64]byte to []byte
}

func Encrypt(plaintext []byte, key []byte) (string, error) { //returns enc string and err
	block, err := aes.NewCipher(key) //telling ais the cipherkey
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block) //Makes ready for getting enc with key
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize()) //uses random number so even same data will be diff
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil) //encrypts
	enc := hex.EncodeToString(ciphertext)                //makes it readable
	return enc, nil
}
func Decrypt(enc string, key []byte) (string, error) {
	decodedCipherText, err := hex.DecodeString(enc) //decodes the encryption
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(decodedCipherText) < nonceSize {
		return "", err
	}

	nonce := decodedCipherText[:nonceSize]
	ciphertext := decodedCipherText[nonceSize:]

	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}

func DoPasswdMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPasswd(currPassword, salt)

	return hashedPassword == currPasswordHash
}
