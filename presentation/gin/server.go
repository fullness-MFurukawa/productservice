package gin

import (
	"sample-service/presentation/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GinのデフォルトRouter生成とServerの実行
// 2023/03/28
func SetupGinServer() *gin.Engine {
	router := gin.Default() // デフォルトRouterを利用する
	// CORS設定 (cross-origin sharing standard)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	router.Use(cors.New(config)) // CORS設定をRouterに登録する

	// ProductControllerの生成
	procutctrl := product.NewProductController()
	// ルーティングの設とリクエストハンドラのマッピング
	productgrp := router.Group("/product")
	{
		productgrp.GET("/list", procutctrl.List)
		productgrp.GET("/search", procutctrl.SearchKeyword)
		productgrp.GET("/search/:keyword", procutctrl.SearchKeyword)
	}
	return router
}
