package product

import (
	"sample-service/domain"
	"sample-service/domain/product"
)

// Product Entityと他のモデルを変換するAdapterのコンストラクタ
// 2023/03/27
func NewProductAdapter() domain.EntityConverter {
	return &ProductConverterImpl{}
}

// ProductRepositoryインターフェース実装のコンストラクタ
// 2023/03/27
func NewProductRepository(converter domain.EntityConverter) product.ProductRepositiry {
	return &ProductRepositoryImpl{converter: converter}
}
