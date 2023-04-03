package provider

import (
	"sample-service/domain"
	ca "sample-service/domain/category"
	pr "sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/category"
	"sample-service/infrastructure/sqlboiler/product"

	"go.uber.org/fx"
)

type InfraParams struct {
	fx.In
	CategoryConverter domain.EntityConverter `name:"categoryconverter"`
	ProductConverter  domain.EntityConverter `name:"productconverter"`
}

// CategoryConverterImplのインスタンス生成
func NewCategoryConverterFx() domain.EntityConverter {
	return category.NewCategoryConverterImpl()
}

// Product Entityと他のモデルを変換するAdapterのコンストラクタ
// 2023/03/27
func NewProductConverterFx() domain.EntityConverter {
	return product.NewProductConverterImpl()
}

// CategoryFRepositoryImplのインスタンス生成
func NewCategoryRepositiryFx(params InfraParams) ca.CategoryRepository {
	return category.NewCategoryRepositiryImpl(params.CategoryConverter)
}

// ProductRepositoryインターフェース実装のコンストラクタ
// 2023/03/27
func NewProductRepositoryFx(params InfraParams) pr.ProductRepositiry {
	return product.NewProductRepositoryImpl(params.ProductConverter)
}
