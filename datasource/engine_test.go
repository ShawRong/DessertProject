package datasource

import (
	"testing"
)

/*
func TestDBinit_user(t *testing.T) {
	if err := DBinit_user(); err != nil {
		t.Errorf("%v\n", err)
	}
}*/

/*func TestDBcreate_user(t *testing.T) {
	test := UserInfo{"name", "name"}
	if err := DBcreate_user(Userdb, test); err != nil {
		t.Errorf("%v\n", err)
	}
}*/

/*
func TestDBfind_user(t *testing.T) {
	var info UserInfo
	var err error
	if info, err = DBfind_user(Userdb, "name1"); err != nil {
		t.Errorf("%v\n", err)
	}
	if info.Username != "name" || info.PassWord != "name" {
		t.Errorf("%s %s", info.Username, info.PassWord)
	}
}

func TestDBupdate_user(t *testing.T) {
	test := UserInfo{"name", "fuck"}
	if err := DBupdate_user(Userdb, "name", test); err != nil {
		t.Errorf("%v\n", err)
	}
}*/
/*
func TestDBinit_quiz(t *testing.T) {
	if err := DBinit_quiz(); err != nil {
		t.Errorf("%v\n", err)
	}
}*/

/*
func TestDBcreate_quiz(t *testing.T) {

	test := QuizInfo{"2", "10", "1", "1", "+", "2"}
	if err := DBcreate_quiz(Quizdb, test); err != nil {
		t.Errorf("%v\n", err)
	}
}*/

/*
func TestDBfind_quiz(t *testing.T) {
	quizinfo, err := DBfind_quiz(Quizdb, "2")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	if quizinfo.QuizNum != "2" {
		t.Errorf("wrong")
	}
}*/

/*
func TestDBfind_quiz_byRank(t *testing.T) {
	quizinfos, err := DBfind_quiz_byRank(Quizdb, "5")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	for _, quizinfo := range quizinfos {
		if quizinfo.QuizRank != "10" {
			t.Errorf("wrong")
		}
	}
}*/

/*
func TestDBupdata_quiz(t *testing.T) {
	err := DBupdate_quiz(Quizdb, "2", QuizInfo{"10", "2", "10", "10", "-", "0"})
	if err != nil {
		t.Errorf("%v\n", err)
	}
}*/

func TestDBinit_topic(t *testing.T) {
	if err := DBinit_topic(); err != nil {
		t.Errorf("%v\n", err)
	}
}

/*

func TestDBcreate_topic(t *testing.T) {
	err := DBcreate_topic(WrongTopicdb, WrongTopic{"2", "name"})
	if err != nil {
		t.Errorf("%v\n", err)
	}
}*/

/*
func TestDBfind_topic(t *testing.T) {
	wrongtopics, err := DBfind_topic(WrongTopicdb, "name")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	for _, wrongtopic := range wrongtopics {
		if wrongtopic.Username != "name" {
			t.Error("wrong")
		}
	}
}*/

/*
func TestDBdelete_topic(t *testing.T) {
	err := DBdelete_topic(WrongTopicdb, "name", "2")
	if err != nil {
		t.Errorf("%v\n", err)
	}
}*/
