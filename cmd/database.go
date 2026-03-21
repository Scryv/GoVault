package cmd

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
		"crypto/sha512"
	"encoding/hex"
	"crypto/rand"
)

var db *gorm.DB
type Data struct {
	gorm.Model
	Username string
	Password string
	Salt     string `gorm:"uniqueIndex:idx_salt"`
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
const saltLength = 14 //length salt Const cause needs to be a fixed length

func genRandoSalt(saltLength int) []byte {  //func for creating random salt
	var salt = make([]byte, saltLength) // makes a byte slice variable called salt
	rand.Read(salt) //reads the slice and fully changes it and ads its own rando value

	return salt //returns salts
}
func hashPasswd(passwd string, salt []byte) string{
	var passwdBytes = []byte(passwd) //creates byte slice of the passwd str
	passwdBytes = append(passwdBytes, salt...) //appends and the ... is for since salt is a slice
	hash := sha512.Sum512(passwdBytes) //hashes the slice using sha512
	return hex.EncodeToString(hash[:]) //encodes to readable and [:] to change [64]byte to []byte
}
