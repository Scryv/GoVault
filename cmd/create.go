
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


func run(){
checkFolder()
var passwd string //just var for passwd
var username string //just var for username
var masterPasswd string

fmt.Println("New users Username: ")
fmt.Scanln(&username)
initUserDB(username)
UserDB.AutoMigrate(&UserData{}) //autocreates tables and updates schema
fmt.Println("new users Password: ") //prompt
fmt.Scanln(&passwd) //scans answer does stop by space tho also & so it can overwrite var
fmt.Println("New users Master Password(for displaying passwds)")
fmt.Scanln(&masterPasswd)
initDB()
VaultDB.AutoMigrate(&Data{}) //autocreates tables and updates schema

salt := genRandoSalt(saltLength) //call and assign genSalt
masterSalt := genRandoSalt(saltLength)
hashedpasswd := hashPasswd(passwd, salt) //call and asign hashPasswd
hashedMasterPasswd := hashPasswd(masterPasswd, masterSalt)
createPost(username, hashedpasswd, hex.EncodeToString(salt), hashedMasterPasswd, hex.EncodeToString(masterSalt))
fmt.Printf("User %s has been created\n", username)
}
