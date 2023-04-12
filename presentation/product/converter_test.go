package product

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/domain/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	category, _ := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200), category)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	converter := NewProductDtoConverter()
	dto, _ := converter.Convert(product)
	fmt.Println(dto)
	assert.NotNil(t, dto)
}
