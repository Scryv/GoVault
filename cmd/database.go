package cmd

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
		"crypto/sha512"
	"encoding/hex"
	"crypto/rand"
	   "crypto/aes"
   "crypto/cipher"
   "io"
)

var VaultDB *gorm.DB
var UserDB *gorm.DB

type Data struct {
	gorm.Model
	Username string
	Password string
	Salt     string `gorm:"uniqueIndex:idx_salt"`
}
type UserData struct {
	gorm.Model
	Username string
	Password string
	Email    string
}
func initDB() {
	var err error
	VaultDB, err = gorm.Open(sqlite.Open("GoVaultDB/users.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func initUserDB(Username string) {
	var err error
	UserDB, err = gorm.Open(sqlite.Open("GoVaultDB/users/"+Username+".db"), &gorm.Config{})
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

func getUser(username string) (string, string, bool) {
	var users []Data

	result := VaultDB.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}

	for _, user := range users {
		if user.Username == username {
			return user.Password, user.Salt, true
		}
	}

	return "", "", false
}
func doPasswdMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = hashPasswd(currPassword, salt)

	return hashedPassword == currPasswordHash
}

func createPost(username string, passwd string, salt string) Data { //func for creating post and also returns it
	newPost := Data{Username: username, Password: passwd, Salt: salt} //new post with TitleandSlug your input
	if res := VaultDB.Create(&newPost); res.Error != nil { //var of the create func res if res error
	panic(res.Error) //not nil or duplicate it wil give error
}
return newPost
}

func AddData(username string, passwd string, email string,) UserData{
	AddData := UserData{Username: username, Password: passwd, Email: email}
	if res := UserDB.Create(&AddData); res.Error != nil {
		panic(res.Error)
	}
	return AddData
}

func encrypt(plaintext []byte, key []byte) (string, error){ //returns enc string and err
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
     enc := hex.EncodeToString(ciphertext) //makes it readable
     return enc, nil
}
func decrypt(enc string, key []byte) (string, error) {
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
