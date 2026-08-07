package user

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	auth "scryv/GoVault/internal/auth"
	db "scryv/GoVault/internal/database"
	check "scryv/GoVault/internal/utils"

	"golang.org/x/term"
)

func runAdd() {
	var username string
	check.InitDB()
	db.VaultDB.AutoMigrate(&db.Data{})

	fmt.Println("Login: ")
	fmt.Scanln(&username)
	fmt.Println("Password: ")
	passwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	hash := sha256.Sum256(passwd)

	storedHash, storedSalt, found := getUser(username)
	saltBytes, _ := hex.DecodeString(storedSalt)
	if !found {
		fmt.Println("User wasnt found ")
		return
	}

	match := auth.DoPasswdMatch(storedHash, string(passwd), saltBytes)
	if match {
		var choice int
		var AddUser string
		var AddEmail string
		var service string
		key := hash[:]

		check.InitUserDB(username)
		db.UserDB.AutoMigrate(&db.UserData{})

		fmt.Println("What do you want to add: [1]Username-password [2]email-password [3]email-passwd-username ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("What Website/App/Service: ")
			fmt.Scanln(&service)
			fmt.Println("Username: ")
			fmt.Scanln(&AddUser)
			fmt.Println("Password: ")
			AddPasswd, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			AddUserByte := []byte(AddUser)
			AddPasswdByte := AddPasswd
			AddUser, _ = auth.Encrypt(AddUserByte, key)
			AddPswdString, _ := auth.Encrypt(AddPasswdByte, key)
			AddData(service, AddUser, AddPswdString, "")

		case 2:
			fmt.Println("What Website/App/Service: ")
			fmt.Scanln(&service)
			fmt.Println("Email: ")
			fmt.Scanln(&AddEmail)
			fmt.Println("Password: ")
			AddPasswd, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			AddPasswdByte := AddPasswd
			AddEmailByte := []byte(AddEmail)
			AddPswdString, _ := auth.Encrypt(AddPasswdByte, key)
			AddEmail, _ = auth.Encrypt(AddEmailByte, key)
			AddData(service, "", AddPswdString, AddEmail)
		case 3:
			fmt.Println("What Website/App/Service: ")
			fmt.Scanln(&service)
			fmt.Println("Email: ")
			fmt.Scanln(&AddEmail)
			fmt.Println("Username: ")
			fmt.Scanln(&AddUser)
			fmt.Println("Password: ")
			AddPasswd, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			AddUserByte := []byte(AddUser)
			AddEmailByte := []byte(AddEmail)
			AddPasswdByte := AddPasswd
			AddUser, _ = auth.Encrypt(AddUserByte, key)
			AddPswdString, _ := auth.Encrypt(AddPasswdByte, key)
			AddEmail, _ = auth.Encrypt(AddEmailByte, key)
			AddData(service, AddUser, AddPswdString, AddEmail)
		default:
			fmt.Println("Please choose an existing option")
		}
	} else {
		fmt.Println("Invalid passwd")
	}
}
