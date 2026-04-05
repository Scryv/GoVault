package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
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

func run() {
	checkFolder()
	var username string

	fmt.Println("New users Username: ")
	fmt.Scanln(&username)
	initUserDB(username)
	UserDB.AutoMigrate(&UserData{})     //autocreates tables and updates schema
	fmt.Println("new users Password: ") //prompt
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
	initDB()
	VaultDB.AutoMigrate(&Data{}) //autocreates tables and updates schema

	salt := genRandoSalt(saltLength) //call and assign genSalt
	masterSalt := genRandoSalt(saltLength)
	hashedpasswd := hashPasswd(string(passwd), salt) //call and asign hashPasswd
	hashedMasterPasswd := hashPasswd(string(masterPasswd), masterSalt)
	createPost(username, hashedpasswd, hex.EncodeToString(salt), hashedMasterPasswd, hex.EncodeToString(masterSalt))
	fmt.Printf("User %s has been created\n", username)
}
