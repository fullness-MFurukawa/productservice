package repository

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/infrastructure/sqlboiler/converter"
	"sample-service/infrastructure/sqlboiler/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 全件取得
func TestCategoryFindAll(t *testing.T) {
	ctx, transaction := db.DBInitForTest()

	defer transaction.Rollback()
	repository := NewCategoryRepositiryImpl(converter.NewCategoryConverterImpl())

	categories, err := repository.FindAll(ctx, transaction)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	for _, category := range categories {
		fmt.Println(category.String())
	}
	assert.Equal(t, 3, len(categories))
}

// カテゴリIDで問合せ
func TestCategoryFindById(t *testing.T) {

	ctx, transaction := db.DBInitForTest()

	defer transaction.Rollback()

	id, err := category.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	repository := NewCategoryRepositiryImpl(converter.NewCategoryConverterImpl())
	result, r_err := repository.FindById(ctx, transaction, id)
	if r_err != nil {
		assert.Fail(t, err.Error())
	}
	fmt.Println(result)
	assert.True(t, true)

	// 存在しないカテゴリID
	id, err = category.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c1")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, r_err = repository.FindById(ctx, transaction, id)
	assert.Nil(t, result)
	assert.Nil(t, r_err)
}
