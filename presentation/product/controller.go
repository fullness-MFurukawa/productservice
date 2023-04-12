package product

import (
	"fmt"
	"net/http"
	"sample-service/application/product"
	"sample-service/domain"
	entity "sample-service/domain/product"
	"sample-service/presentation/result"

	"github.com/gin-gonic/gin"
)

// 商品を扱うController
// 2023/03/29
type ProductController struct {
	service   product.ProductService   // 商品サービス
	converter domain.EntitiesConverter // DTO、Entity変換Adapter
}

// コンストラクタ
// 2023/03/29
func NewProductController(service product.ProductService,
	converter domain.EntitiesConverter) *ProductController {
	// ProductServiceを構築する
	return &ProductController{service: service, converter: converter}
}

// 商品一覧を提供するハンドラ
// @Summary 商品一覧を取得する
// @Description 登録されたすべての商品を取得する
// @Accept */*
// @Produce json
// @Success 200 {Object} product.ProductDto
// @Router /list [get]
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

// 商品を指定されたキーワードで検索した結果を提供するハンドラ
// @Summary 指定されたキーワードで検索した結果を取得する
// @Info 2023/03/29
// @Description キーワードで検索した結果を取得する
// @Accept */*
// @Produce json
// @Param keyword query string true "商品キーワード"
// @Success 200 {Object} product.ProductDto
// @Failure 404 {object} result.FailureDto
// @Router /search [get]
func (controller *ProductController) SearchKeyword(context *gin.Context) {
	var keyword string
	// パスにあるパラメータを取得
	keyword = context.Param("keyword")
	if keyword == "" {
		// パスパラメータが無い場合はQueryパラメータを取得
		keyword = context.Query("keyword")
		// Queryパラメータの無い場合はエラーを返す
		if keyword == "" {
			context.JSON(http.StatusBadRequest, gin.H{"parama error": "検索キーワードがありません。"})
		}
	}
	// 指定されたキーワードで商品を検索する
	products, err := controller.service.SearchBykeyword(keyword)
	if err != nil {
		context.JSON(http.StatusNotFound, result.NewFailureDto("キーワード検索", err.Error()))
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

// 新商品の登録を提供するハンドラ
// @Summary 新商品を登録する
// @Info 2023/03/29
// @Description 新商品を登録する
// @Accept */*
// @Produce json
// @Param newproduct body ProductDto true "商品名、単価、カテゴリ番号"
// @Success 200 {Object} result.SuccessDto
// @Router /add [post]
func (controller *ProductController) Add(context *gin.Context) {
	var dto ProductDto
	// ヘッダーのパラメータ(JSON形式)をProductDtoに格納する
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 受信した新商品からProduct Entityを再構築する
	p, err := controller.converter.Restore(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 新しい商品を登録する
	newproduct, _ := p.(*entity.Product)
	if err := controller.service.Add(newproduct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK,
		result.NewSuccessDto("商品の登録", fmt.Sprintf("%sを登録しました。", newproduct.ProductName().Value())))
}

// 商品の更新を提供するハンドラ
// @Summary 商品を更新する
// @Info 2023/03/29
// @Description 商品を更新する
// @Accept */*
// @Produce json
// @Param newproduct body ProductDto true "商品番号、商品名、単価"
// @Success 200 {Object} result.SuccessDto
// @Router /change [put]
func (controller *ProductController) Change(context *gin.Context) {
	var dto ProductDto
	// ヘッダーのパラメータ(JSON形式)をProductDtoに格納する
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 受信した商品からProduct Entityを再構築する
	p, err := controller.converter.Restore(&dto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 商品を更新する
	updateproduct, _ := p.(*entity.Product)
	_, upderr := controller.service.Change(updateproduct)
	if upderr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK,
		result.NewSuccessDto("商品の更新", fmt.Sprintf("%sを更新しました。", updateproduct.ProductName().Value())))
}

// 商品の削除を提供するハンドラ
// @Summary 商品を削除する
// @Info 2023/03/29
// @Description 削除する
// @Accept */*
// @Produce json
// @Param id path string true "商品番号"
// @Success 200 {Object} result.SuccessDto
// @Router /remove/{id} [delete]
func (controller *ProductController) Remove(context *gin.Context) {
	// パスにあるパラメータを取得
	id := context.Param("id")
	productid, err := entity.NewProductId(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 商品を削除する
	_, del_err := controller.service.Remove(productid)
	if del_err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK,
		result.NewSuccessDto("商品の削除", fmt.Sprintf("商品番号:%sの商品を削除しました。", id)))
}
