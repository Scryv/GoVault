package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
)

// addPswdCmd represents the addPswd command
var addPswdCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add Account/Password",
	Long: `This command makes you able to add accounts usernames and passwords to your database
		encrypted ofcourse :)`,
	Run: func(cmd *cobra.Command, args []string) {
		runAdd()
	},
}

func init() {
	rootCmd.AddCommand(addPswdCmd)
}

func runAdd() {
	var username string
	initDB()
	VaultDB.AutoMigrate(&Data{})

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

	match := doPasswdMatch(storedHash, string(passwd), saltBytes)
	if match {
		var choice int
		var AddUser string
		var AddEmail string
		var service string
		key := hash[:]

		initUserDB(username)
		UserDB.AutoMigrate(&UserData{})

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
			AddUser, _ = encrypt(AddUserByte, key)
			AddPswdString, _ := encrypt(AddPasswdByte, key)
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
			AddPswdString, _ := encrypt(AddPasswdByte, key)
			AddEmail, _ = encrypt(AddEmailByte, key)
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
			AddUser, _ = encrypt(AddUserByte, key)
			AddPswdString, _ := encrypt(AddPasswdByte, key)
			AddEmail, _ = encrypt(AddEmailByte, key)
			AddData(service, AddUser, AddPswdString, AddEmail)
		default:
			fmt.Println("Please choose an existing option")
		}
	} else {
		fmt.Println("Invalid passwd")
	}
}
