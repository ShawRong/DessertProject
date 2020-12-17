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
		return
	}
	var quizinfo datasource.QuizInfo
	quizinfo.QuizNum = quiz.Num
	quizinfo.QuizRank = quiz.Rank
	quizinfo.Content = quiz.Content
	quizinfo.Res = quiz.Res

	err := datasource.DBcreate_quiz(datasource.Quizdb, quizinfo)
	if err != nil {
		data := "Exist"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
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
		return
	}
	if info.Num != "" && info.Rank == "" {
		quizinfo, err := datasource.DBfind_quiz(datasource.Quizdb, info.Num)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			type temp struct {
				Status bool              `json:"status"`
				Data   interface{}       `json:"data"`
				Slice  []*model.Question `json:"slice"`
			}
			data := "OK"
			var quiz model.Question
			quiz.Num = quizinfo.QuizNum
			quiz.Rank = quizinfo.QuizRank
			quiz.Content = quizinfo.Content
			quiz.Res = quizinfo.Res
			var slice []*model.Question
			slice = append(slice, &quiz)
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else if info.Num == "" && info.Rank != "" {
		quizinfos, err := datasource.DBfind_quiz_byRank(datasource.Quizdb, info.Rank)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			type temp struct {
				Status bool              `json:"status"`
				Data   interface{}       `json:"data"`
				Slice  []*model.Question `json:"slice"`
			}
			data := "OK"
			var slice []*model.Question
			for _, quizinfo := range quizinfos {
				var quiz model.Question
				quiz.Num = quizinfo.QuizNum
				quiz.Rank = quizinfo.QuizRank
				quiz.Content = quizinfo.Content
				quiz.Res = quizinfo.Res
				slice = append(slice, &quiz)
			}
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else {
		data := "Wrong Format"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}

func Deletequiz(ctx iris.Context) {
	var quiznum string
	if err := ctx.ReadJSON(&quiznum); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	err := datasource.DBdelete_quiz(datasource.Quizdb, quiznum)
	if err != nil {
		data := "delete error"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}
