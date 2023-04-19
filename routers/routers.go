package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//	加载静态页面
	r.LoadHTMLGlob("templates/*")

	r.Use(favicon.New("./favicon.ico"))
	//	加载资源文件
	r.Static("/static", "./static")

	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//	待办事项

		//	添加
		v1Group.POST("/todo", controller.CreateTodo)
		//	查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//	查看某一个待办事项
		v1Group.GET("/todo/:id", controller.GetTodo)
		//	修改

		//	修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//	删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

	}

	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", gin.H{
			"error": "找不到页面",
		})
	})
	return r
}
