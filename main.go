package main

import (
	"log"
	"os"
	"project/routes"
	"project/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// SQLite データベースを初期化
	storage.InitDB()

	r := gin.Default()

	// CORS を設定
	r.Use(setupCORS())

	// ルートを設定
	routes.SetupRoutes(r)

	// PORT 環境変数を取得 (Render では PORT が指定される)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ローカル開発時のデフォルト
	}

	// サーバー起動
	log.Println("Starting server on port " + port)
	r.Run(":" + port)
}

func setupCORS() gin.HandlerFunc {
	frontendURL := os.Getenv("FRONTEND_ORIGIN")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	return cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL}, // フロントエンドのURLを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	})
}
