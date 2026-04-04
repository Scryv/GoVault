
package cmd

import (
	"fmt"
	_"github.com/glebarez/sqlite"
	_"gorm.io/gorm"
	"github.com/spf13/cobra"
		"encoding/hex"
		"crypto/sha256"
)


var govaultCmd = &cobra.Command{
	Use:   "govault",
	Short: "Command that runs govault",
	Long: `This Command runs govault it will ask for login and password and when logged in will display the stored passwords with email or username`,
	Run: func(cmd *cobra.Command, args []string) {
		ruun()
	},
}

func init() {
	rootCmd.AddCommand(govaultCmd)
}


func printUsers(password string) {
	var users []UserData  //creates empty list called users []Data cause whole table
    hash := sha256.Sum256([]byte(password))
	result := UserDB.Find(&users) //loads all rows from db into userSlice
	if result.Error != nil { //result.Error and not err is gorm doesnt return err
		panic(result.Error) //panic cause if broken state
	}
    key := hash[:]
	csvData := "" //creates variable called csvData

	for _, user := range users { //loops trough every row
		UserUser := user.Username
		UserPasswd := user.Password
		UserEmail := user.Email

		DecUser, _ := decrypt(UserUser, key)
		DecPasswd, _ := decrypt(UserPasswd, key)
		DecEmail, _ := decrypt(UserEmail, key)
		csvData += fmt.Sprintf( //appends all these looped to csvData
			"%s,%s,%s\n",  //Sprintf cause can be saved in var
			DecUser,
			DecPasswd,
			DecEmail,
		)
	}

	fmt.Println(csvData) //prints csv data
}
func ruun(){
	var username string
    var passwd string
    var MasterPasswd string
	initDB()
	VaultDB.AutoMigrate(&Data{})

    fmt.Println("Login: ")
    fmt.Scanln(&username)
    fmt.Println("Password: ")
    fmt.Scanln(&passwd)
    fmt.Println("Give your master Passwd")
    fmt.Scanln(&MasterPasswd)

	storedHash, storedSalt, found := getUser(username)
	saltBytes, _ := hex.DecodeString(storedSalt)
	if !found {
		fmt.Println("User wasnt found ")
		return
	}
    storedMasterHash, storedMasterSalt, masterFound := getMasterUser(username)
    masterSaltBytes, _ := hex.DecodeString(storedMasterSalt)
    if !masterFound {
    	fmt.Println("User Not found")
    }

    masterMatch := doPasswdMatch(storedMasterHash, MasterPasswd, masterSaltBytes)
    match := doPasswdMatch(storedHash, passwd, saltBytes)
    if match {
       if masterMatch {
       initUserDB(username)
       UserDB.AutoMigrate(&UserData{})
       fmt.Println("USERNAME | PASSWORD | EMAIL")
       printUsers(passwd)
       } else {
       	fmt.Println("Wrong Master Password")
       }
    } else {
       fmt.Println("Invalid passwd")
    }
}
