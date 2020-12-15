package controller

import (
	"github.com/kataras/iris/v12"
	"dessert/service"
)

func HubController(app *iris.Application){
	
	/*router main*/
	main := app.Party("/")
	
	/*router home*/
	home := main.Party("/")
	home.Get("/", func(ctx iris.Context){
		ctx.View("index.html")
	})

	/*router user*/
	user := main.Party("/user")
	user.Post("/registe",service.Registe)
	user.Post("/login",service.Login)
	
}