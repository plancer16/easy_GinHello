package handler

import (
	"diy_ginHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UserRegister(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBind(&user); err != nil {
		log.Panicln("err：", err.Error())
		ctx.String(http.StatusBadRequest, "输入数据不合法")
	}
	passwordAgain := ctx.PostForm("password-again")
	if passwordAgain != user.Password {
		ctx.String(http.StatusBadRequest, "两次密码不一致")
		log.Panicln("两次密码不一致")
	}
	id := user.Save()
	log.Println("id is", id)
	ctx.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(ctx *gin.Context) {
	var user model.User
	if e := ctx.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		ctx.SetCookie("user_cookie", string(u.Id), 1000, "/", "localhost", false, true)
		log.Println("登录成功")
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
			"id":    u.Id,
		})
	}
}

func UserProfile(ctx *gin.Context) {
	id := ctx.Query("id")
	var user model.User
	i, e := strconv.Atoi(id)
	u, err := user.QueryById(i)
	if e != nil || err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
	}
	ctx.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(ctx *gin.Context) {
	var user model.User
	if e := ctx.ShouldBind(&user); e != nil {
		ctx.HTML(http.StatusOK,"error.tmpl",gin.H{
			"error": e.Error(),
		})
		log.Panicln("绑定错误",e.Error())
	}
	err := user.Update(user.Id)
	if err != nil {
		ctx.HTML(http.StatusOK,"error.tmpl",gin.H{
			"error": err.Error(),
		})
		log.Panicln("更新错误",err.Error())
	}
	ctx.Redirect(http.StatusMovedPermanently,"/user/profile?id="+strconv.Itoa(user.Id))
}
