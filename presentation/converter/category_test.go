package converter

import (
	"fmt"
	"sample-service/domain/category"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CaregoryEntityからCategoryDtoへの変換
func TestCategoryConvert(t *testing.T) {
	category, err := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	adapter := NewCategoryDtoConverter()
	dto, _ := adapter.Convert(category)
	fmt.Println(dto)
	assert.NotNil(t, dto)
}
