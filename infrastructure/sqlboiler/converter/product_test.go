package converter

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Product EntityからSqlBuilderのモデルに変換
func TestProductConvert(t *testing.T) {
	category, err := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200), category)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	converter := NewProductConverterImpl()
	result, err := converter.Convert(product)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	t.Log(result)
	assert.True(t, true)
}

// SqlBuilderのモデルからProduct Entityを復元
func TestProductRestore(t *testing.T) {

	category := &models.Category{
		ID:    0,
		ObjID: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
		Name:  "文房具"}
	product := &models.Product{
		ID:    0,
		ObjID: "ac413f22-0cf1-490a-9635-7e9ca810e544",
		Name:  "商品-ABC",
		Price: 200}
	// ジョインするカテゴリを設定する
	product.R.NewStruct().Category = category
	fmt.Println(product)
	converter := NewProductConverterImpl()
	result, err := converter.Restore(product)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	fmt.Println(result)
	assert.True(t, true)
}
