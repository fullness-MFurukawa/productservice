package product

import (
	"sample-service/infrastructure/sqlboiler/converter"
	"sample-service/infrastructure/sqlboiler/repository"
)

// fxを使わないProductServiceインスタンス生成
// 2023/03/29
func NewProductServiceNonFx() ProductService {
	// Repositoryで利用するProductAdapterの生成
	converter := converter.NewProductConverterImpl()
	// ProductServiceで利用するRepositoryの生成
	repository := repository.NewProductRepositoryImpl(converter)
	// ProductRepositoryの生成
	service := NewProductServiceImpl(repository)
	return service
}
