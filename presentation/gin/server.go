package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GinのデフォルトRouter生成とServerの実行
// 2023/03/28
func RunGinServer() {
	router := gin.Default() // デフォルトRouterを利用する
	// CORS設定 (cross-origin sharing standard)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	router.Use(cors.New(config)) // CORS設定をRouterに登録する

	router.Run()
}
