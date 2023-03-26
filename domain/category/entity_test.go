package category

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Category生成テスト
func TestNewCategory(t *testing.T) {
	category, err := NewCategory("文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	log.Println(category.String())
	assert.NotNil(t, category)
}

// カテゴリ再構築テスト
func TestBuildCategory(t *testing.T) {
	category, err := BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	log.Println(category.String())
	assert.Equal(t, "b1524011-b6af-417e-8bf2-f449dd58b5c0", category.categoryId.Value())
	assert.Equal(t, "文房具", category.categoryName.Value())
}

// カテゴリの等価性テスト
func TestEquals(t *testing.T) {
	category1, _ := BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	category2, _ := BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	category3, _ := BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c1", "文房具")
	result, _ := category1.Equals(category2)
	assert.True(t, result)
	result, _ = category1.Equals(category3)
	assert.False(t, result)
}
