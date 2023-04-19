package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
*
url --> controller --> logic -->model-->dao
请求来了-->控制器-->业务逻辑-->模型层的增删改查
*/
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(context *gin.Context) {
	//	前端页面填写待办事项 点击提交 会发送请求到这里
	//	1.	从请求中把数据拿出来
	var todo models.Todo
	context.BindJSON(&todo)
	//	2.	存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todo)

		//context.JSON(http.StatusOK, gin.H{
		//	"code": 200,
		//	"msg":  "success",
		//	"data": todo,
		//})
	}

	//	3.	返回响应

}

func GetTodoList(context *gin.Context) {

	//	查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		context.JSON(http.StatusOK, todoList)
	}

}

func UpdateATodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}

	todo, err := models.GetATodo(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.BindJSON(todo)
	if err := models.UpdateATodo(todo); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
func GetTodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}

}
