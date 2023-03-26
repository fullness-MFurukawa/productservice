package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ProductIdのインスタンス生成
func TestNewProductId(t *testing.T) {
	input := "ABC"
	_, err := NewProductId(input)
	assert.Equal(t, "商品IDの長さは36文字でなければなりません。", err.Error())
	input = "b1524011-b6af-417e-8bf2f449dd58b5c0A"
	_, err = NewProductId(input)
	assert.Equal(t, "商品IDはUUIDの形式でなければなりません。", err.Error())
	input = "b1524011-b6af-417e-8bf2-f449dd58b5c0"
	result, _ := NewProductId(input)
	assert.Equal(t, "b1524011-b6af-417e-8bf2-f449dd58b5c0", result.Value())
}

// ProductNameのインスタンス生成
func TestNewProductName(t *testing.T) {
	input := "あいうえおかきくけこさしすせそなにぬねのはひふへほまみむめもや"
	_, err := NewProductName(input)
	assert.Equal(t, "商品名の長さは30文字以内です。", err.Error())
	input = "商品-ABC"
	result, _ := NewProductName(input)
	assert.Equal(t, "商品-ABC", result.Value())
}

// ProductPriceのインスタンス生成
func TestProductPrice(t *testing.T) {
	input := uint32(49)
	_, err := NewProductPrice(input)
	assert.Equal(t, "単価は50以上、10000以下です。", err.Error())
	input = uint32(10001)
	_, err = NewProductPrice(input)
	assert.Equal(t, "単価は50以上、10000以下です。", err.Error())
	input = uint32(50)
	result, _ := NewProductPrice(input)
	assert.Equal(t, uint32(50), result.Value())
	input = uint32(10000)
	result, _ = NewProductPrice(input)
	assert.Equal(t, uint32(10000), result.Value())
}
