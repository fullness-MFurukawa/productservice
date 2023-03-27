package product

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/domain/product"
	rep "sample-service/infrastructure/sqlboiler/product"
	"sample-service/infrastructure/sqlboiler/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 商品の一覧を取得する
func TestList(t *testing.T) {
	tests.TestDBInit() // コネクションプールを初期化する
	// ProductServiceのインスタンスを取得する
	service := NewProductService(rep.NewProductRepositoryImpl())
	// 商品の一覧を取得する
	products, err := service.List()
	if err != nil {
		assert.Error(t, err)
	}
	// 取得結果をコンソールに出力する
	for _, product := range products {
		fmt.Println(product)
	}
	assert.True(t, len(products) > 0)
}

// 指定されたキーワードを含む商品を取得する
func TestSearchBykeyword(t *testing.T) {
	tests.TestDBInit() // コネクションプールを初期化する
	// ProductServiceのインスタンスを取得する
	service := NewProductService(rep.NewProductRepositoryImpl())

	// キーワードを含む商品が存在しない場合
	products, err := service.SearchBykeyword("あいうえお")
	assert.True(t, len(products) == 0)
	assert.Equal(t, err.Error(), "キーワード:あいうえおを含む商品は見つかりませんでした。")

	// キーワードを含む商品が存在する場合
	products, err = service.SearchBykeyword("ペン")
	if err != nil {
		assert.Error(t, err)
	}
	// 検索結果を出力する
	for _, product := range products {
		fmt.Println(product)
	}
	assert.True(t, len(products) > 0)
}

// 新商品を登録する
func TestAdd(t *testing.T) {
	// テスト用データを作成する
	category, _ := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	product, _ := product.NewProduct("商品-ABC", uint32(200), category)
	tests.TestDBInit() // コネクションプールを初期化する
	// ProductServiceのインスタンスを取得する
	service := NewProductService(rep.NewProductRepositoryImpl())
	// 商品の登録
	err := service.Add(product)
	assert.Nil(t, err)
}

// 商品の内容を変更する
func TestChange(t *testing.T) {
	tests.TestDBInit() // コネクションプールを初期化する
	// ProductServiceのインスタンスを取得する
	service := NewProductService(rep.NewProductRepositoryImpl())

	// 変更対象の商品が見つからない場合
	upd_category, _ := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	upd_product, _ := product.BuildProduct("6a7c0248-45e7-4443-97d8-6454320b6000", "商品-XYZ", uint32(300), upd_category)
	result, err := service.Change(upd_product)
	assert.False(t, result)
	assert.Equal(t, err.Error(), "商品番号:6a7c0248-45e7-4443-97d8-6454320b6000の商品は存在しないため変更できませんでした。")

	// 変更対象の商品が存在する場合
	upd_category, _ = category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	upd_product, _ = product.BuildProduct("6a7c0248-45e7-4443-97d8-6454320b6dbe", "商品-XYZ", uint32(300), upd_category)
	result, err = service.Change(upd_product)
	assert.True(t, result)
	assert.Nil(t, err)
}

// 商品を削除する
func TestRemove(t *testing.T) {
	tests.TestDBInit() // コネクションプールを初期化する
	// ProductServiceのインスタンスを取得する
	service := NewProductService(rep.NewProductRepositoryImpl())
	// 存在しない商品番号を指定した場合
	productId, _ := product.NewProductId("6a7c0248-45e7-4443-97d8-6454320b6000")
	result, err := service.Remove(productId)
	assert.False(t, result)
	assert.Equal(t, err.Error(), "商品番号:6a7c0248-45e7-4443-97d8-6454320b6000の商品は存在しないため削除できませんでした。")
	// 存在する商品番号を指定した場合
	productId, _ = product.NewProductId("6a7c0248-45e7-4443-97d8-6454320b6dbe")
	result, err = service.Remove(productId)
	assert.True(t, result)
	assert.Nil(t, err)
}
