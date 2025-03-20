package main

import (
	"project/routes"
	"project/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// SQLite データベースを初期化
	storage.InitDB()

	r := gin.Default()

	// ルートを設定
	routes.SetupRoutes(r)

	// サーバー起動
	r.Run(":8080")
}
