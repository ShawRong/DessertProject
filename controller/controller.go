package controller

import (
	"dessert/service"

	"github.com/kataras/iris/v12"
)

func HubController(app *iris.Application) {

	/*router main*/
	main := app.Party("/")

	/*router home*/
	home := main.Party("/")
	home.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	/*router user*/
	user := main.Party("/user")
	user.Post("/registe", service.Registe)
	user.Post("/login", service.Login)
	user.Get("/findwrongtopic", service.Findwrongtopic)
	user.Get("/buildwrongtopic", service.Buildwrongtopic)
	user.Get("/getquiz", service.Getquiz)

	/*router admin*/
	admin := main.Party("/admin")
	admin.Post("/buildquiz", service.Buildquiz)
	admin.Post("/deletequiz", service.Deletequiz)
}
