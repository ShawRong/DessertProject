package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	Username string `gorm:"primary_key"`
	PassWord string
}

var Userdb *gorm.DB

//dbname = userdb
func DBinit_user() error {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/"+"userdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	db.AutoMigrate(&UserInfo{})
	// ......

	Userdb = db
	return nil
}

func DBcreate_user(db *gorm.DB, userinfo UserInfo) error {
	err := db.Create(&userinfo).Error
	return err
}

func DBfind_user(db *gorm.DB, username string) (UserInfo, error) {
	var userinfo UserInfo
	err := db.Where("username = ?", username).First(&userinfo).Error
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
