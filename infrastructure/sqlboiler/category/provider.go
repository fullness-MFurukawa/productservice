package category

import (
	"sample-service/domain"
	"sample-service/domain/category"
)

// CategoryAdapterImplとCategoryRepositoryImplのインスタンス生成
// 2023/03/25

// CategoryAdapterImplのインスタンス生成
func NewCategoryAdapater() domain.EntityAdapter {
	return &CategoryAdapterImpl{}
}

// CategoryFRepositoryImplのインスタンス生成
func NewCategoryRepositiry(adapter domain.EntityAdapter) category.CategoryRepository {
	return &CategoryRepositoryImpl{adapter: adapter}
}
