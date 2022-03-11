package test

import (
	"bytes"
	"diy_ginHello/initRouter"
	"diy_ginHello/model"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertArticle(t *testing.T) {
	router := initRouter.SetupRouter()
	article := model.Article{
		Type:    "go",
		Content: "goyyds",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article", bytes.NewBufferString(string(marshal)))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, "{id:-1}", w.Body.String())
}