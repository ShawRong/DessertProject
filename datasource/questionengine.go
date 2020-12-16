package datasource

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type QuizInfo struct {
	QuizNum  string `gorm:"primary_key"`
	QuizRank string
	X        string
	Y        string
	Sig      string
	Res      string
}

var Quizdb *gorm.DB

//dbname = quizdb
func DBinit_quiz() error {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/"+"quizdb"+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	db.AutoMigrate(&QuizInfo{})
	// ......

	Quizdb = db
	return nil
}

func DBcreate_quiz(db *gorm.DB, quizinfo QuizInfo) error {
	err := db.Create(&quizinfo).Error
	return err
}

func DBfind_quiz(db *gorm.DB, quiznum string) (QuizInfo, error) {
	var quizinfo QuizInfo
	err := db.Where("quiz_num = ?", quiznum).First(&quizinfo).Error
	if err != nil {
		var invalid QuizInfo
		return invalid, err
	}
	return quizinfo, nil
}

func DBfind_quiz_byRank(db *gorm.DB, quizrank string) ([]QuizInfo, error) {
	var quizinfos []QuizInfo
	db.Where("quiz_rank = ?", quizrank).Find(&quizinfos)
	if len(quizinfos) == 0 {
		return quizinfos, errors.New("empty slice")
	} else {
		return quizinfos, nil
	}
}

func DBupdate_quiz(db *gorm.DB, quiznum string, quizinfo QuizInfo) error {
	findquiz, err := DBfind_quiz(db, quiznum)
	if err != nil {
		return err
	}
	if quizinfo.QuizNum != findquiz.QuizNum {
		return errors.New("quiznum consistence wrong")
	}
	db.Model(&findquiz).Updates(quizinfo)
	return nil
}

func DBdelete_quiz(db *gorm.DB, quiznum string) error {
	var quiz QuizInfo
	err := db.Where("quiz_num = ?", quiznum).First(&quiz).Error
	if err != nil {
		return err
	}
	err1 := db.Delete(&quiz).Error
	if err1 != nil {
		return err1
	}
	return nil
}

func DBclose_quiz(db *gorm.DB) {
	db.Close()
}
