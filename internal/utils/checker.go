package utils

import (
	"os"
	"os/user"
	"path/filepath"
	db "scryv/GoVault/internal/database"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	var err error
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	CheckFolder()
	VaultDBFP := filepath.Join(currentUser.HomeDir, "GoVaultDB", "users.db")
	db.VaultDB, err = gorm.Open(sqlite.Open(VaultDBFP), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitUserDB(username string) {
	CheckFolder()
	usersDir := CheckUserFolder()
	UserDBFP := filepath.Join(usersDir, username+".db")

	var err error
	db.UserDB, err = gorm.Open(sqlite.Open(UserDBFP), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func CheckFolder() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	dirPath := filepath.Join(currentUser.HomeDir, "GoVaultDB")

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}
func CheckUserFolder() string {
	currentUser, _ := user.Current()
	dirPath := filepath.Join(currentUser.HomeDir, "GoVaultDB", "Users")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
	return dirPath
}
