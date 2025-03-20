package main

import (
	"project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ルートを設定
	routes.SetupRoutes(r)

	// サーバー起動
	r.Run(":8080")
}
