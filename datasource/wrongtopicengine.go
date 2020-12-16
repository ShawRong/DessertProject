package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type WrongTopic struct {
	QuizNum  string `gorm:"primary_key"`
	Username string `gorm:"primary_key"`
}

var WrongTopicdb *gorm.DB

//dbname = topicdb
func DBinit_topic() {
	db, err := gorm.Open("mysql", "dessert:dessert@/"+"topicdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&WrongTopic{})
	// ......

	WrongTopicdb = db
}

func DBcreate_topic(db *gorm.DB, wrongtopic WrongTopic) {
	db.Create(&wrongtopic)
}

func DBfind_topic(db *gorm.DB, username string) ([]WrongTopic, error) {
	var wrongtopics []WrongTopic
	err := db.Where("username = ?", username).Find(&wrongtopics).Error
	if err != nil {
		var invalid []WrongTopic
		return invalid, err
	}
	return wrongtopics, nil
}

func DBdelete_topic(db *gorm.DB, username string, quiznum string) error {
	var wrongtopic WrongTopic
	err := db.Where("username = ? AND quiznum = ?", username, quiznum).First(&wrongtopic).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&wrongtopic).Error
	if err1 != nil {
		return err1
	}
	return nil
}

func DBclose_topic(db *gorm.DB) {
	db.Close()
}
