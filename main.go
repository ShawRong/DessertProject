package main

import (
	"dessert/controller"
	"dessert/model"

	"math/rand"
	"strconv"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New() // 实例一个iris对象
	app.Favicon("./favicons/favicon.ico")
	app.RegisterView(iris.HTML("./view", ".html")) //资源库
	app.HandleDir("/static", ("./view/static"))    //静态文件

	controller.HubController(app)

	// 启动服务器
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
	// 监听地址:本服务器上任意id端口8080,设置字符集utf8
}

func RandGenerate(ctx iris.Context) {
	x := rand.Intn(100)
	y := rand.Intn(100)
	sig := rand.Intn(2) //0 for + 1 for -
	res := rand.Intn(2)
	var mark string
	if sig == 0 {
		res = x + y
		mark = "+"
	} else {
		res = x - y
		mark = "-"
	}
	_, _ = ctx.JSON(model.Question{X: strconv.Itoa(x), Y: strconv.Itoa(y), Sig: mark, Res: strconv.Itoa(res)})
}
