package product

import (
	"fmt"
	"sample-service/domain/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	converter := NewProductDtoConverter()
	dto, _ := converter.Convert(product)
	fmt.Println(dto)
	assert.NotNil(t, dto)
}

// ProductDtoからProduct Entityを再構築する
func TestRestore(t *testing.T) {
	dto := NewProductDto("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200))
	converter := NewProductDtoConverter()
	entity, _ := converter.Restore(dto)
	fmt.Println(entity)
	assert.NotNil(t, entity)
}

// ProductのスライスからProductDtoのスライスを生成して返す
func TestMultiConvert(t *testing.T) {
	// スライスを生成する無名関数
	f := func(id string, name string, price uint32) product.Product {
		product, err := product.BuildProduct(id, name, price, nil)
		if err != nil {
			assert.Fail(t, err.Error())
		}
		return *product
	}
	// Product Entityのスライスを生成
	var products = make([]product.Product, 0, 3)
	product1 := f("ac413f22-0cf1-490a-9635-7e9ca810e541", "商品-ABC", uint32(200))
	product2 := f("ac413f22-0cf1-490a-9635-7e9ca810e542", "商品-LMN", uint32(300))
	product3 := f("ac413f22-0cf1-490a-9635-7e9ca810e543", "商品-XYZ", uint32(400))
	products = append(products, product1)
	products = append(products, product2)
	products = append(products, product3)

	converter := NewProductDtoConverter()
	dtos, err := converter.MultiConvert(products)
	if err != nil {
		assert.Error(t, err)
	}
	for _, dto := range dtos.([]ProductDto) {
		fmt.Println(dto)
	}
	assert.True(t, len(dtos.([]ProductDto)) > 0)
}
