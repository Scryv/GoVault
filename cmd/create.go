/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    _"crypto/rand"
	_"crypto/sha512"
	"encoding/hex"
	"github.com/spf13/cobra"
	_"gorm.io/gorm" //lets me use go structs in plaats van sql
	_"github.com/glebarez/sqlite"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Command to create user for GoVault",
	Long: `This is the command that adds a user to the User.db database so it can access the
		passwd.db database hashed the password with salt and saves it in User.db`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}



func createPost(username string, passwd string, salt string) Data { //func for creating post and also returns it
	newPost := Data{Username: username, Password: passwd, Salt: salt} //new post with TitleandSlug your input
	if res := VaultDB.Create(&newPost); res.Error != nil { //var of the create func res if res error
	panic(res.Error) //not nil or duplicate it wil give error
}
return newPost
}

func run(){
initDB()
var passwd string //just var for passwd
var username string //just var for username
VaultDB.AutoMigrate(&Data{}) //autocreates tables and updates schema

fmt.Println("What username you want to give: ")
fmt.Scanln(&username)
initUserDB(username)
fmt.Println("What password do you want to hash: ") //prompt
fmt.Scanln(&passwd) //scans answer does stop by space tho also & so it can overwrite var

salt := genRandoSalt(saltLength) //call and assign genSalt
hashedpasswd := hashPasswd(passwd, salt) //call and asign hashPasswd

createPost(username, hashedpasswd, hex.EncodeToString(salt))
fmt.Printf("User %s has been created\n", username)
fmt.Println("With this salt: ", hex.EncodeToString(salt)) //Prints salt used
}
