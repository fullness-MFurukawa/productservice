package product

import (
	"fmt"
	"sample-service/domain/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Product EntityからProductDtoへの変換
func TestConvert(t *testing.T) {
	product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	adapter := NewProductDtoConverter()
	dto, _ := adapter.Convert(product)
	fmt.Println(dto)
	assert.NotNil(t, dto)
}

// ProductDtoからProduct Entityを再構築する
func TestRestore(t *testing.T) {
	dto := NewProductDto("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200))
	adapter := NewProductDtoConverter()
	entity, _ := adapter.Restore(dto)
	fmt.Println(entity)
	assert.NotNil(t, entity)
}
