/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ruun()
	},
}

func init() {
	rootCmd.AddCommand(govaultCmd)
}


func printUsers() {
	var users []Data  //creates empty list called users []Data cause whole table

	result := db.Find(&users) //loads all rows from db into userSlice
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

func getUser(username string) (string, string, bool) {
	var users []Data

	result := db.Find(&users)
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

func ruun(){
	var username string
    var passwd string
	initDB()
	db.AutoMigrate(&Data{})

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
