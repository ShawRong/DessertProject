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
	err := datasource.DBcreate_user(datasource.Userdb, userinfo)
	if err != nil {
		data := "Exist"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "Nice"
		_, _ = ctx.JSON(model.Response{Status: true, Data: data})
	}
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
		type temp struct {
			Status bool              `json:"status"`
			Data   interface{}       `json:"data"`
			Slice  []*model.Question `json:"slice"`
		}

		data := "Found"
		var slice []*model.Question
		for _, wrongtopic := range wrongtopics {
			quizinfo, err := datasource.DBfind_quiz(datasource.Quizdb, wrongtopic.QuizNum)
			if err != nil {
				panic(err)
			}
			var quiz model.Question
			quiz.Num = quizinfo.QuizNum
			quiz.Rank = quizinfo.QuizRank
			quiz.Content = quizinfo.Content
			quiz.Res = quizinfo.Res
			slice = append(slice, &quiz)
		}
		_, _ = ctx.JSON(temp{Status: true, Data: data, Slice: slice})
	}
	return
}

func Buildwrongtopic(ctx iris.Context) {
	type temp struct {
		QuizNum  string `json:"num"`
		Username string `json:"username"`
	}
	topic := new(temp)
	if err := ctx.ReadJSON(&topic); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		_, _ = ctx.JSON(model.Response{Status: false, Data: data})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	var wrongtopic datasource.WrongTopic
	wrongtopic.QuizNum = topic.QuizNum
	wrongtopic.Username = topic.Username
	datasource.DBcreate_topic(datasource.WrongTopicdb, wrongtopic)
	data := "OK"
	_, _ = ctx.JSON(model.Response{Status: true, Data: data})
}

func Deletetopic(ctx iris.Context) {
	type temp struct {
		QuizNum  string `json:"num"`
		Username string `json:"username"`
	}
	topic := new(temp)
	if err := ctx.ReadJSON(&topic); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	err := datasource.DBdelete_topic(datasource.WrongTopicdb, topic.Username, topic.QuizNum)
	if err != nil {
		data := "delete error"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}
