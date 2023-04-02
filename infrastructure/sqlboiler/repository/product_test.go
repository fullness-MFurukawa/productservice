package repository

import (
	"fmt"
	"sample-service/domain/category"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/converter"
	"sample-service/infrastructure/sqlboiler/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Repositoryの生成
func createProductRepository() product.ProductRepositiry {
	converter := &converter.ProductConverterImpl{}
	repository := &ProductRepositoryImpl{converter: converter}
	return repository
}

func TestProductFindAll(t *testing.T) {

	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	results, err := repository.FindAll(ctx, transaction)
	for _, result := range results {
		fmt.Println(result.String())
	}
	assert.NotNil(t, results)
	assert.Nil(t, err)
}

func TestProductFindByNameLike(t *testing.T) {
	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	keyworkd := "%ペン%"
	results, err := repository.FindByNameLike(ctx, transaction, keyworkd)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	for _, result := range results {
		fmt.Println(result.String())
	}
	assert.True(t, true)
	keyworkd = "%ああああ%"
	results, err = repository.FindByNameLike(ctx, transaction, keyworkd)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Nil(t, results)
}

func TestProductExist(t *testing.T) {
	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	name, _ := product.NewProductName("水性ボールペン(黒)")
	result, err := repository.Exist(ctx, transaction, name)
	assert.True(t, result)
	assert.Nil(t, err)

	name, _ = product.NewProductName("水性ボールペン")
	result, err = repository.Exist(ctx, transaction, name)
	assert.False(t, result)
	assert.Nil(t, err)
}

func TestProductCreate(t *testing.T) {

	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	category, err := category.BuildCategory("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	new_product, err := product.NewProduct("商品-ABC", uint32(200), category)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result := repository.Create(ctx, transaction, new_product)
	assert.Nil(t, result)
}

func TestProductUpdateById_OK(t *testing.T) {
	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	update_product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "水性ボールペン(黒)", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, up_err := repository.UpdateById(ctx, transaction, update_product)
	assert.True(t, result)
	assert.Nil(t, up_err)
}

func TestProductUpdateById_NG(t *testing.T) {
	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	update_product, err := product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e544", "水性ボールペン(黒)", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, u_err := repository.UpdateById(ctx, transaction, update_product)
	assert.True(t, result)
	assert.Nil(t, u_err)

	update_product, err = product.BuildProduct("ac413f22-0cf1-490a-9635-7e9ca810e546", "水性ボールペン", uint32(200), nil)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, u_err = repository.UpdateById(ctx, transaction, update_product)
	assert.False(t, result)
	assert.Nil(t, u_err)
}

func TestProductDeleteById(t *testing.T) {
	// Connection PoolとContextの取得
	ctx, transaction := db.DBInitForTest()
	// Repositoryの生成
	repository := createProductRepository()

	defer transaction.Rollback()

	delete_product_id, err := product.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, del_err := repository.DeleteById(ctx, transaction, delete_product_id)
	assert.True(t, result)
	assert.Nil(t, del_err)

	delete_product_id, err = product.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e545")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	result, del_err = repository.DeleteById(ctx, transaction, delete_product_id)
	assert.False(t, result)
	assert.Nil(t, del_err)
}
