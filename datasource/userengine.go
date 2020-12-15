package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	gorm.Model
	Username string
	PassWord string
}

var Userdb *gorm.DB

//dbname = userdb
func DBinit_user() {
	db, err := gorm.Open("mysql", "dessert:dessert@/"+"userdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&UserInfo{})
	// ......

	Userdb = db
}

func DBcreate_user(db *gorm.DB, userinfo UserInfo) {
	db.Create(&userinfo)
}

func DBfind_user(db *gorm.DB, username string) (UserInfo, error) {
	var userinfo UserInfo
	err := db.First(&userinfo, "username = ?", username).Error
	if err != nil {
		var invalid UserInfo
		return invalid, err
	}
	return userinfo, nil
}

func DBupdate_user(db *gorm.DB, username string, userinfo UserInfo) error {
	finduser, err := DBfind_user(db, username)
	if err != nil {
		return err
	}
	db.Model(&finduser).Updates(userinfo)
	return nil
}

func DBclose_user(db *gorm.DB) {
	db.Close()
}
