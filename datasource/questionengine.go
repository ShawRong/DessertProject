package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Quizcontent struct {
	X   string
	Y   string
	Sig string
	Res string
}

type QuizInfo struct {
	QuizNum     string `gorm:"primary_key"`
	QuizRank    string
	QuizContent Quizcontent
}

var Quizdb *gorm.DB

//dbname = quizdb
func DBinit_quiz() {
	db, err := gorm.Open("mysql", "dessert:dessert@/"+"quizdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&QuizInfo{})
	// ......

	Quizdb = db
}

func DBcreate_quiz(db *gorm.DB, quizinfo QuizInfo) {
	db.Create(&quizinfo)
}

func DBfind_quiz(db *gorm.DB, quiznum string) (QuizInfo, error) {
	var quizinfo QuizInfo
	err := db.Where("quiznum = ?", quiznum).First(&quizinfo).Error
	if err != nil {
		var invalid QuizInfo
		return invalid, err
	}
	return quizinfo, nil
}

func DBfind_quiz_byRank(db *gorm.DB, quizrank string) ([]QuizInfo, error) {
	var quizinfos []QuizInfo
	err := db.Where("quizrank = ?", quizrank).Find(&quizinfos).Error
	if err != nil {
		var invalid []QuizInfo
		return invalid, err
	}
	return quizinfos, nil
}

func DBupdate_quiz(db *gorm.DB, quiznum string, quizinfo QuizInfo) error {
	findquiz, err := DBfind_quiz(db, quiznum)
	if err != nil {
		return err
	}
	db.Model(&findquiz).Updates(quizinfo)
	return nil
}

func DBclose_quiz(db *gorm.DB) {
	db.Close()
}
