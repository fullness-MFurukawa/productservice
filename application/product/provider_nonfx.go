package product

import "sample-service/infrastructure/sqlboiler/product"

// fxを使わないProductServiceインスタンス生成
// 2023/03/29
func NewProductServiceNonFx() ProductService {
	// Repositoryで利用するProductAdapterの生成
	adapter := product.NewProductAdapter()
	// ProductServiceで利用するRepositoryの生成
	repository := product.NewProductRepository(adapter)
	// ProductRepositoryの生成
	service := NewProductService(repository)
	return service
}
