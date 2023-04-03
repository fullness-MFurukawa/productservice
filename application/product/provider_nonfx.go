package product

import "sample-service/infrastructure/sqlboiler/product"

// fxを使わないProductServiceインスタンス生成
// 2023/03/29
func NewProductServiceNonFx() ProductService {
	// Repositoryで利用するProductAdapterの生成
	converter := product.NewProductConverterImpl()
	// ProductServiceで利用するRepositoryの生成
	repository := product.NewProductRepositoryImpl(converter)
	// ProductRepositoryの生成
	service := NewProductServiceImpl(repository)
	return service
}
