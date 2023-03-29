package category

import (
	"sample-service/domain"
	"sample-service/domain/category"
)

// CategoryConverterImplとCategoryRepositoryImplのインスタンス生成
// 2023/03/25

// CategoryConverterImplのインスタンス生成
func NewCategoryConverter() domain.EntityConverter {
	return &CategoryConverterImpl{}
}

// CategoryFRepositoryImplのインスタンス生成
func NewCategoryRepositiry(converter domain.EntityConverter) category.CategoryRepository {
	return &CategoryRepositoryImpl{converter: converter}
}
