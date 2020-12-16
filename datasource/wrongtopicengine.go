package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type WrongTopic struct {
	QuizNum  string `gorm:"primary_key"`
	Username string `gorm:"primary_key"`
}

var WrongTopicdb *gorm.DB

//dbname = topicdb
func DBinit_topic() error {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/"+"topicdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	db.AutoMigrate(&WrongTopic{})
	// ......

	WrongTopicdb = db
	return nil
}

func DBcreate_topic(db *gorm.DB, wrongtopic WrongTopic) error {
	err := db.Create(&wrongtopic).Error
	return err
}

func DBfind_topic(db *gorm.DB, username string) ([]WrongTopic, error) {
	var wrongtopics []WrongTopic
	db.Where("username = ?", username).Find(&wrongtopics)
	if len(wrongtopics) == 0 {
		return wrongtopics, errors.New("empty slice")
	} else {
		return wrongtopics, nil
	}
}

func DBdelete_topic(db *gorm.DB, username string, quiznum string) error {
	var wrongtopic WrongTopic
	err := db.Where("username = ? AND quiz_num = ?", username, quiznum).First(&wrongtopic).Error
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
