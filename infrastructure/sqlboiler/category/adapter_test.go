package category

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/domain/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	category, err := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	adapter := NewCategoryAdapater()
	model, _ := adapter.Convert(category)
	fmt.Println(model)
	assert.True(t, true)

	id := "ac413f22-0cf1-490a-9635-7e9ca810e544"
	product, err := product.BuildProduct(id, "商品-ABC", uint32(300), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	_, err = adapter.Convert(product)
	assert.Equal(t, "指定されたEntityはCategoryではありません。", err.Error())
}
