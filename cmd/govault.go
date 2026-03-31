
package cmd

import (
	"fmt"
	_"github.com/glebarez/sqlite"
	_"gorm.io/gorm"
	"github.com/spf13/cobra"
		"encoding/hex"
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


func printUsers() {
	var users []Data  //creates empty list called users []Data cause whole table

	result := VaultDB.Find(&users) //loads all rows from db into userSlice
	if result.Error != nil { //result.Error and not err is gorm doesnt return err
		panic(result.Error) //panic cause if broken state
	}

	csvData := "" //creates variable called csvData

	for _, user := range users { //loops trough every row
		csvData += fmt.Sprintf( //appends all these looped to csvData
			"%s,%s,%s\n",  //Sprintf cause can be saved in var
			user.Username,
			user.Password,
			user.Salt,
		)
	}

	fmt.Println(csvData) //prints csv data
}
func ruun(){
	var username string
    var passwd string
	initDB()
	VaultDB.AutoMigrate(&Data{})

    fmt.Println("Login: ")
    fmt.Scanln(&username)
    fmt.Println("Password: ")
    fmt.Scanln(&passwd)

	storedHash, storedSalt, found := getUser(username)
	saltBytes, _ := hex.DecodeString(storedSalt)
	if !found {
		fmt.Println("User wasnt found ")
		return
	}

    match := doPasswdMatch(storedHash, passwd, saltBytes)
    if match {
       fmt.Println("USERNAME | HASHED_PASSWORD | SALT")
       printUsers()
    } else {
       fmt.Println("Invalid passwd")
    }
}
