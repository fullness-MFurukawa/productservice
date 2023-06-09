package gin

import (
	"context"
	"fmt"
	"sample-service/infrastructure/sqlboiler/db"
	"sample-service/presentation/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	_ "sample-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Ginを保持するHandler構造体
// 2023/03/31
type Router struct {
	Gin *gin.Engine
}

// コンストラクタ
// 2023/03/31
func NewRouter() *Router {
	gin := gin.Default() // デフォルトのGinを生成
	// Open API(Swagger)
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// CORS設定 (cross-origin sharing standard)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	gin.Use(cors.New(config)) // CORS設定をRouterに登録する
	router := Router{Gin: gin}
	return &router
}

// ルーティングの設定
// 2023/03/31
func RegisterRouter(router *Router, controller *product.ProductController) {
	// ルーティングの設定とリクエストハンドラのマッピング
	productgrp := router.Gin.Group("/product")
	{
		productgrp.GET("/list", controller.List)
		productgrp.GET("/search", controller.SearchKeyword)
		productgrp.GET("/search/:keyword", controller.SearchKeyword)
		productgrp.POST("/add", controller.Add)
		productgrp.PUT("/change", controller.Change)
		productgrp.DELETE("/remove/:id", controller.Remove)
	}
}

// fxのライフサイクル
// 2023/04/01
func RegisterHooks(lifecycle fx.Lifecycle, router *Router) {
	lifecycle.Append(
		fx.Hook{
			// fx開始時の処理
			OnStart: func(context.Context) error {
				fmt.Println("Starting application Post:8081 !!!")
				// SqlBiolderのConnection Poolを生成
				go db.NewSqlBiolderInitDB().Init(nil)
				go router.Gin.Run(":8081") //	Ginの起動
				return nil
			},
			// fx終了時の処理
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application !!!")
				return nil
			},
		},
	)
}
