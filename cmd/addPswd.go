
package cmd

import (
	"fmt"
	"encoding/hex"
    "crypto/sha256"
	"github.com/spf13/cobra"
)

// addPswdCmd represents the addPswd command
var addPswdCmd = &cobra.Command{
	Use:   "addPswd",
	Short: "Command to add Account/Password",
	Long: `This command makes you able to add accounts usernames and passwords to your database
		encrypted ofcourse :)`,
	Run: func(cmd *cobra.Command, args []string) {
		ruuun()
	},
}

func init() {
	rootCmd.AddCommand(addPswdCmd)
}

func ruuun(){
	var username string
	var passwd string
	initDB()
	VaultDB.AutoMigrate(&Data{})

    fmt.Println("Login: ")
    fmt.Scanln(&username)
    fmt.Println("Password: ")
    fmt.Scanln(&passwd)
    hash := sha256.Sum256([]byte(passwd)) 

	storedHash, storedSalt, found := getUser(username)
	saltBytes, _ := hex.DecodeString(storedSalt)
	if !found {
		fmt.Println("User wasnt found ")
		return
	}

    match := doPasswdMatch(storedHash, passwd, saltBytes)
    if match {
       var choice int
       var AddPasswd string
       var AddUser string
       var AddEmail string
       key := hash[:]
       
       initUserDB(username)
       UserDB.AutoMigrate(&UserData{})
       
       fmt.Println("What do you want to add: [1]Username-password [2]email-password [3]email-passwd-username ")
       fmt.Scanln(&choice)
       switch choice {
       case 1:
           fmt.Println("Username: ")
           fmt.Scanln(&AddUser)
           fmt.Println("Password: ")
           fmt.Scanln(&AddPasswd)
           AddUserByte := []byte(AddUser)
           AddPasswdByte := []byte(AddPasswd)
           AddUser, _ = encrypt(AddUserByte, key)
           AddPasswd, _ = encrypt(AddPasswdByte, key)
           AddData(AddUser, AddPasswd, "")
           
           
           
       case 2:
       	    fmt.Println("Email: ")
       	    fmt.Scanln(&AddEmail)
       	    fmt.Println("Password: ")
       	    fmt.Scanln(&AddPasswd)
       	    AddPasswdByte := []byte(AddPasswd)
       	    AddEmailByte := []byte(AddEmail)
       	    AddPasswd, _ = encrypt(AddPasswdByte, key)
       	    AddEmail, _ = encrypt(AddEmailByte, key)
       	    AddData("", AddPasswd, AddEmail)
       case 3:
           fmt.Println("Email: ")
           fmt.Scanln(&AddEmail)
           fmt.Println("Username: ")
           fmt.Scanln(&AddUser)
           fmt.Println("Password: ")
           fmt.Scanln(&AddPasswd)
           AddUserByte := []byte(AddUser)
           AddEmailByte := []byte(AddEmail)
           AddPasswdByte := []byte(AddPasswd)
           AddUser, _ = encrypt(AddUserByte, key)
           AddPasswd, _ = encrypt(AddPasswdByte, key)
           AddEmail, _ = encrypt(AddEmailByte, key)
           AddData(AddUser, AddPasswd, AddEmail)
       default:
       	 fmt.Println("Please choose an existing option")
       }
    } else {
       fmt.Println("Invalid passwd")
    }
}
