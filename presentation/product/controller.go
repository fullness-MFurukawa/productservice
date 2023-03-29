package product

import (
	"net/http"
	"sample-service/application/product"
	"sample-service/domain"

	"github.com/gin-gonic/gin"
)

// 商品を扱うController
// 2023/03/29
type ProductController struct {
	service   product.ProductService   // 商品サービス
	converter domain.EntitiesConverter // DTO、Entity変換Adapter
}

// コンストラクタ
func NewProductController() *ProductController {
	// ProductServiceを構築する
	service := product.NewProductServiceNonFx()
	converter := NewProductDtoConverter()
	return &ProductController{service: service, converter: converter}
}

// 商品一覧を提供するリクエストハンドラ
func (controller *ProductController) List(context *gin.Context) {
	// ProductServiceを利用して商品一覧を取得する
	products, err := controller.service.List()
	if err != nil {
		// 何らかのエラーが発生
		// 最終的にエラーハンドラを利用して正確なハンドリングをする
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// ProdustのスライスをProductDtoのスライスに変換する
	dtos, err := controller.converter.MultiConvert(products)
	if err != nil {
		// 何らかのエラーが発生
		// 最終的にエラーハンドラを利用して正確なハンドリングをする
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// ステータス:200でProductDtoのJSONを送信する
	context.JSON(http.StatusOK, dtos)
}
