package category

import (
	"sample-service/apperrors"
	"sample-service/domain"
	"sample-service/domain/category"
)

// CategoryDtoとCategory Entityの相互変換Adapter
// 2023/03/28
type CategoryDtoConverter struct{}

// Category EntityからCategoryDtoを生成する
func (adapter *CategoryDtoConverter) Convert(entity any) (any, error) {
	source, ok := entity.(*category.Category)
	if !ok {
		return nil, apperrors.NewDomainError("指定されたEntityはCategoryではありません。")
	}
	dto := CategoryDto{Id: source.CategoryId().Value(), Name: source.CategoryName().Value()}
	return dto, nil
}

// CategoryDtoからCategory Entityを再構築する
func (adapater *CategoryDtoConverter) Restore(model any) (any, error) {
	source, ok := model.(*CategoryDto)
	if !ok {
		return nil, apperrors.NewDomainError("指定されたmodelはCategoryDtoではありません。")
	}
	category, err := category.BuildCategory(source.Id, source.Name)
	if err != nil {
		return nil, err
	}
	return *category, nil
}

// コンストラクタ
func NewCategoryDtoConverter() domain.EntityConverter {
	return &CategoryDtoConverter{}
}
