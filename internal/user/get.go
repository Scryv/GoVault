package user

import (
	db "scryv/GoVault/internal/database"
)

func getUser(username string) (string, string, bool) {
	var users []db.Data

	result := db.VaultDB.Find(&users)
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

func getMasterUser(username string) (string, string, bool) {
	var users []db.Data

	result := db.VaultDB.Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}

	for _, user := range users {
		if user.Username == username {
			return user.MasterPasswd, user.MasterSalt, true
		}
	}

	return "", "", false
}
