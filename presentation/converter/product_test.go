package converter

import (
	"fmt"
	"sample-service/domain/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductConvert(t *testing.T) {
	product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "商品-ABC", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	converter := NewProductDtoConverter()
	dto, _ := converter.Convert(product)
	fmt.Println(dto)
	assert.NotNil(t, dto)
}
