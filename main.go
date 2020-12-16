package main

import (
	"dessert/controller"
	"dessert/datasource"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Favicon("./favicons/favicon.ico")
	app.RegisterView(iris.HTML("./view", ".html"))
	app.HandleDir("/static", "./view/static")
	controller.HubController(app)

	datasource.DBinit_user()
	datasource.DBinit_quiz()
	datasource.DBinit_topic()

	// 启动服务器
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
	// 监听地址:本服务器上任意id端口8080,设置字符集utf8
	defer datasource.DBclose_user(datasource.Userdb)
	defer datasource.DBclose_quiz(datasource.Quizdb)
	defer datasource.DBclose_topic(datasource.WrongTopicdb)
}
