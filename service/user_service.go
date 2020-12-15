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
