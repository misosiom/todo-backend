package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func TestCreateTodoAndGetTodos(t *testing.T) {
	router := setupRouter()

	// 新しいタスクを作成
	newTodo := map[string]interface{}{
		"title":     "Test Task",
		"completed": false,
	}
	jsonValue, _ := json.Marshal(newTodo)

	req, _ := http.NewRequest("POST", "/todos/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// タスク一覧を取得
	req, _ = http.NewRequest("GET", "/todos/", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var todos []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &todos)
	assert.Equal(t, "Test Task", todos[0]["title"])
}
