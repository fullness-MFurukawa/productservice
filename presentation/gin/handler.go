package gin

/*
// Ginを保持するHandler構造体
// 2023/03/31
type Handler struct {
	Gin *gin.Engine
}

// コンストラクタ
func NewHandler() *Handler {
	gin := gin.Default() // デフォルトのGinを生成
	// CORS設定 (cross-origin sharing standard)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	gin.Use(cors.New(config)) // CORS設定をRouterに登録する
	handler := Handler{Gin: gin}
	return &handler
}

// ルーティングの設定
// 2023/03/31
func RegisterRoutes(handler *Handler, controller *product.ProductController) {
	// ルーティングの設とリクエストハンドラのマッピング
	productgrp := handler.Gin.Group("/product")
	{
		productgrp.GET("/list", controller.List)
		productgrp.GET("/search", controller.SearchKeyword)
		productgrp.GET("/search/:keyword", controller.SearchKeyword)
	}
}
*/
