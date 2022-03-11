package test

import (
	"diy_ginHello/initRouter"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var eng *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	eng = initRouter.SetupRouter()
}

func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	eng.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
	assert.Contains(t,w.Body.String(),"hello gin get method","返回界面应包含hello gin get method")
}