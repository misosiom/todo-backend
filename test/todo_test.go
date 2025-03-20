package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"project/models"
	"project/routes"
	"project/storage"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// テスト用のルーターを作成
func setupRouter() *gin.Engine {
	storage.InitDB() // SQLite を初期化
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func TestCreateTodoAndGetTodos(t *testing.T) {
	router := setupRouter()

	// 新しいタスクを作成
	newTodo := models.Todo{
		Title:     "Test Task",
		Completed: false,
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
}
