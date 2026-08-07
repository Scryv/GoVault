package db

import "gorm.io/gorm"

var VaultDB *gorm.DB
var UserDB *gorm.DB

type Data struct {
	gorm.Model
	Username     string
	Password     string
	Salt         string //gorm:"uniqueIndex:idx_salt"
	MasterPasswd string
	MasterSalt   string
}
type UserData struct {
	gorm.Model
	Service  string
	Username string
	Password string
	Email    string
}
