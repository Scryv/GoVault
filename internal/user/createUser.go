package user

import (
	"encoding/hex"
	"fmt"
	"os"
	auth "scryv/GoVault/internal/auth"
	db "scryv/GoVault/internal/database"
	check "scryv/GoVault/internal/utils"

	"golang.org/x/term"
)

func CreateUser() {
	check.CheckFolder()
	var username string

	fmt.Println("New users Username: ")
	fmt.Scanln(&username)
	check.InitUserDB(username)
	db.UserDB.AutoMigrate(&db.UserData{}) //autocreates tables and updates schema
	fmt.Println("new users Password: ")   //prompt
	passwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New users Master Password(for displaying passwds)")
	masterPasswd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	check.InitDB()
	db.VaultDB.AutoMigrate(&db.Data{}) //autocreates tables and updates schema

	salt := auth.GenRandoSalt(auth.SaltLength) //call and assign genSalt
	masterSalt := auth.GenRandoSalt(auth.SaltLength)
	hashedpasswd := auth.HashPasswd(string(passwd), salt) //call and asign hashPasswd
	hashedMasterPasswd := auth.HashPasswd(string(masterPasswd), masterSalt)
	createPost(username, hashedpasswd, hex.EncodeToString(salt), hashedMasterPasswd, hex.EncodeToString(masterSalt))
	fmt.Printf("User %s has been created\n", username)
}
