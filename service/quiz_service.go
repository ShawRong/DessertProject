package service

import (
	"dessert/datasource"
	"dessert/model"

	"github.com/kataras/iris/v12"
)

func Buildquiz(ctx iris.Context) {
	quiz := new(model.Question)
	if err := ctx.ReadJSON(&quiz); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	var quizinfo datasource.QuizInfo
	quizinfo.QuizNum = quiz.Num
	quizinfo.QuizRank = quiz.Rank
	quizinfo.QuizContent.X = quiz.X
	quizinfo.QuizContent.Y = quiz.Y
	quizinfo.QuizContent.Sig = quiz.Sig
	quizinfo.QuizContent.Res = quiz.Res

	datasource.DBcreate_quiz(datasource.Quizdb, quizinfo)
	data := "OK"
	ctx.JSON(model.Response{Status: true, Data: data})
}

func Getquiz(ctx iris.Context) {
	type Info struct {
		Num  string `json:"num"`
		Rank string `json:"rank"`
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	if info.Num != "" {
		quizinfo, err := datasource.DBfind_quiz(datasource.Quizdb, info.Num)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			data := "OK"
			ctx.JSON(model.Response{Status: false, Data: data})
		}
		var quiz model.Question
		quiz.Num = quizinfo.QuizNum
		quiz.Rank = quizinfo.QuizRank
		quiz.X = quizinfo.QuizContent.X
		quiz.Y = quizinfo.QuizContent.Y
		quiz.Sig = quizinfo.QuizContent.Sig
		quiz.Res = quizinfo.QuizContent.Res
		ctx.JSON(quiz)
	} else {
		quizinfos, err := datasource.DBfind_quiz_byRank(datasource.Quizdb, info.Rank)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			data := "OK"
			ctx.JSON(model.Response{Status: false, Data: data})
		}
		for _, quizinfo := range quizinfos {
			var quiz model.Question
			quiz.Num = quizinfo.QuizNum
			quiz.Rank = quizinfo.QuizRank
			quiz.X = quizinfo.QuizContent.X
			quiz.Y = quizinfo.QuizContent.Y
			quiz.Sig = quizinfo.QuizContent.Sig
			quiz.Res = quizinfo.QuizContent.Res
			ctx.JSON(quiz)
		}
	}
	return
}
