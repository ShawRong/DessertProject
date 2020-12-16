package service

import (
	"dessert/datasource"
	"dessert/model"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	aul := new(model.User)
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		_, _ = ctx.JSON(model.Response{Status: false, Data: data})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	userinfo, err := datasource.DBfind_user(datasource.Userdb, aul.Username)
	if err != nil {
		data := "NotFound"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
	} else if userinfo.PassWord == aul.Password {
		data := "Nice"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "Bad"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}

func Registe(ctx iris.Context) {
	aul := new(model.User)
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		_, _ = ctx.JSON(model.Response{Status: false, Data: data})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	var userinfo datasource.UserInfo
	userinfo.Username = aul.Username
	userinfo.PassWord = aul.Password
	datasource.DBcreate_user(datasource.Userdb, userinfo)
	data := "Nice"
	_, _ = ctx.JSON(model.Response{Status: true, Data: data})
}

//about wrongtopic

func Findwrongtopic(ctx iris.Context) {
	aul := new(model.User)
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		_, _ = ctx.JSON(model.Response{Status: false, Data: data})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	username := aul.Username
	wrongtopics, err := datasource.DBfind_topic(datasource.WrongTopicdb, username)
	if err != nil {
		data := "Not Found"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
		return
	} else {
		data := "Found"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
	}
	for _, wrongtopic := range wrongtopics {
		quizinfo, err := datasource.DBfind_quiz(datasource.Quizdb, wrongtopic.QuizNum)
		if err != nil {
			panic(err)
		}
		num := quizinfo.QuizNum
		rank := quizinfo.QuizRank
		x := quizinfo.QuizContent.X
		y := quizinfo.QuizContent.Y
		sig := quizinfo.QuizContent.Sig
		res := quizinfo.QuizContent.Res

		_, _ = ctx.JSON(model.Question{Num: num, Rank: rank, X: x, Y: y, Sig: sig, Res: res})
	}

	return
}

func Buildwrongtopic(ctx iris.Context) {
	type topicer struct {
		model.User
		model.Question
	}
	topic := new(topicer)
	if err := ctx.ReadJSON(&topic); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		_, _ = ctx.JSON(model.Response{Status: false, Data: data})
	}

	ctx.StatusCode(iris.StatusOK)
	var wrongtopic datasource.WrongTopic
	wrongtopic.QuizNum = topic.Num
	wrongtopic.Username = topic.Username
	datasource.DBcreate_topic(datasource.WrongTopicdb, wrongtopic)
	data := "OK"
	_, _ = ctx.JSON(model.Response{Status: true, Data: data})
}
