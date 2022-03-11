package handler

import (
	"diy_ginHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("转换为int类型错误", e.Error())
	}
	article := model.Article{Id: i}
	a := article.FindById(i)
	ctx.JSON(http.StatusOK, gin.H{
		"article": a,
	})
}

func GetAll(ctx *gin.Context) {
	article := model.Article{}
	articles := article.FindAll()
	ctx.JSON(http.StatusOK,gin.H{
		"articles":articles,
	})
}

func Insert(ctx *gin.Context) {
	article := model.Article{}
	var id = -1
	if e := ctx.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	//返回的json的id对应接收者的Body
	ctx.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}

func DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("转换为int类型失败",e.Error())
	}
	article := model.Article{Id: i}
	article.DeleteOne()
}