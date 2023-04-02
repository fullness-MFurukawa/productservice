package provider

import (
	"sample-service/domain"
	"sample-service/domain/category"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/converter"

	"go.uber.org/fx"
)

type InfraParams struct {
	fx.In
	CategoryConverter domain.EntityConverter `name:"categoryconverter"`
	ProductConverter  domain.EntityConverter `name:"productconverter"`
}

// CategoryConverterImplのインスタンス生成
func NewCategoryConverterImpl() domain.EntityConverter {
	return &converter.CategoryConverterImpl{}
}

// Product Entityと他のモデルを変換するAdapterのコンストラクタ
// 2023/03/27
func NewProductConverterImpl() domain.EntityConverter {
	return &converter.ProductConverterImpl{}
}

// CategoryFRepositoryImplのインスタンス生成
func NewCategoryRepositiryImpl(params InfraParams) category.CategoryRepository {
	//return &repository.CategoryRepositoryImpl{converter: nil}
	return nil
}

// ProductRepositoryインターフェース実装のコンストラクタ
// 2023/03/27
func NewProductRepositoryImpl(params InfraParams) product.ProductRepositiry {
	//return &repository.ProductRepositoryImpl{converter: params.ProductConverter}
	return nil
}
