package test

import (
	"bytes"
	"diy_ginHello/initRouter"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)
var router *gin.Engine
func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}
func TestUserRegister(t *testing.T) {
	value := url.Values{}
	value.Add("email", "1922@qq.com")
	value.Add("password", "123")
	value.Add("password-again","123")
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/user/register", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusMovedPermanently,w.Code)
}

func TestUserLogin(t *testing.T) {
	email:= "1922@qq.com"
	value := url.Values{}
	value.Add("email",email)
	value.Add("password","123")
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/user/login", bytes.NewBufferString(value.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
	assert.Equal(t,strings.Contains(w.Body.String(),email),true)
}