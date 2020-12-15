package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type quizcontent struct {
	x   string
	y   string
	sig string
	res string
}

type QuizInfo struct {
	gorm.Model
	QuizNum     int
	QuizRank    int
	QuizContent quizcontent
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

func DBfind_quiz(db *gorm.DB, quiznum int) (QuizInfo, error) {
	var quizinfo QuizInfo
	err := db.First(&quizinfo, "quiznum = ?", quiznum).Error
	if err != nil {
		var invalid QuizInfo
		return invalid, err
	}
	return quizinfo, nil
}

func DBfind_quiz_byRank(db *gorm.DB, quizrank int) (QuizInfo, error) {
	var quizinfo QuizInfo
	err := db.First(&quizinfo, "quizrank = ?", quizrank).Error
	if err != nil {
		var invalid QuizInfo
		return invalid, err
	}
	return quizinfo, nil
}

func DBupdate_quiz(db *gorm.DB, quiznum int, quizinfo QuizInfo) error {
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
