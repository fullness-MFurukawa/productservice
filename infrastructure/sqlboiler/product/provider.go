package product

import (
	"sample-service/domain"
	"sample-service/domain/product"
)

// Product Entityと他のモデルを変換するAdapterのコンストラクタ
// 2023/03/27
func NewProductAdapter() domain.EntityAdapter {
	return &ProductAdapterImpl{}
}

// ProductRepositoryインターフェース実装のコンストラクタ
// 2023/03/27
func NewProductRepository(adapter domain.EntityAdapter) product.ProductRepositiry {
	return &ProductRepositoryImpl{adapter: adapter}
}
