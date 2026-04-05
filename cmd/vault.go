package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"os"
	"text/tabwriter"
)

var govaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Command that runs govault",
	Long:  `This Command runs govault it will ask for login and password and when logged in will display the stored passwords with email or username`,
	Run: func(cmd *cobra.Command, args []string) {
		runVault()
	},
}

func init() {
	rootCmd.AddCommand(govaultCmd)
}

func printUsers(password string) {
	var users []UserData //creates empty list called users []Data cause whole table
	hash := sha256.Sum256([]byte(password))
	result := UserDB.Find(&users) //loads all rows from db into userSlice
	if result.Error != nil {      //result.Error and not err is gorm doesnt return err
		panic(result.Error) //panic cause if broken state
	}
	key := hash[:]
	csvData := ""                                        //creates variable called csvData
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0) //values for tab template stylet thingy
	fmt.Fprintln(w, "Service\tUsername\tPassword\tEmail")
	fmt.Fprintln(w, "-------\t--------\t--------\t-----")
	for _, user := range users { //loops trough every row
		UserService := user.Service
		UserUser := user.Username
		UserPasswd := user.Password
		UserEmail := user.Email

		DecUser, _ := decrypt(UserUser, key)
		DecPasswd, _ := decrypt(UserPasswd, key)
		DecEmail, _ := decrypt(UserEmail, key)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", UserService, DecUser, DecPasswd, DecEmail) //\t is there for spacings

	}
	w.Flush()
	fmt.Println(csvData) //prints csv data
}
func runVault() {
	checkFolder()
	initDB()
	VaultDB.AutoMigrate(&Data{})
	var username string

	fmt.Println("Login: ")
	fmt.Scanln(&username)
	fmt.Println("Password: ")
	passwd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Give your master Passwd")
	MasterPasswd, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

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

	masterMatch := doPasswdMatch(storedMasterHash, string(MasterPasswd), masterSaltBytes)
	match := doPasswdMatch(storedHash, string(passwd), saltBytes)
	if match {
		if masterMatch {
			initUserDB(username)
			UserDB.AutoMigrate(&UserData{})
			printUsers(string(passwd))
		} else {
			fmt.Println("Wrong Master Password")
		}
	} else {
		fmt.Println("Invalid passwd")
	}
}
