package user

import (
	db "scryv/GoVault/internal/database"
)

func createPost(username string, passwd string, salt string, masterPasswd string, masterSalt string) db.Data { //func for creating post and also returns it
	newPost := db.Data{Username: username, Password: passwd, Salt: salt, MasterPasswd: masterPasswd, MasterSalt: masterSalt} //new post with TitleandSlug your input
	if res := db.VaultDB.Create(&newPost); res.Error != nil {                                                                //var of the create func res if res error
		panic(res.Error) //not nil or duplicate it wil give error
	}
	return newPost
}

func AddData(service string, username string, passwd string, email string) db.UserData {
	AddData := db.UserData{Service: service, Username: username, Password: passwd, Email: email}
	if res := db.UserDB.Create(&AddData); res.Error != nil {
		panic(res.Error)
	}
	return AddData
}
