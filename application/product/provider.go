package product

import "sample-service/domain/product"

// ProductServiceインターフェース実装のインスタンスを生成する
// 2023/03/27
func NewProductService(repository product.ProductRepositiry) ProductService {
	return &ProductServiceIml{repository: repository}
}
