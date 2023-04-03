package product

import (
	"sample-service/apperrors"
	"sample-service/domain"
	"sample-service/domain/category"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/models"
)

// Entity Product からSqlBoilerのモデルへ変換
// SqlBoilerのModelからEntity Product への変換
// 2023/03/27
type ProductConverterImpl struct{}

// Entity ProductからsqlboilerのModel Productへ変換する
func (converter *ProductConverterImpl) Convert(entity any) (any, error) {
	source, ok := entity.(*product.Product) // Entity Productへ型変換
	if !ok {                                // 型変換できない?
		return nil, apperrors.NewDomainError("指定されたEntityはProductではありません。")
	}
	product := models.Product{
		ID:         0,
		ObjID:      source.ProductId().Value(),
		Name:       source.ProductName().Value(),
		Price:      int(source.ProductPrice().Value()),
		CategoryID: source.Category().CategoryId().Value(),
	}
	return product, nil
}

// sqlboilerのModel Productから任意のEntity Productへ変換する
func (converter *ProductConverterImpl) Restore(model any) (any, error) {
	source, ok := model.(*models.Product) // SqlBoilerのProduct Modelに変換
	if !ok {
		return nil, apperrors.NewDomainError("指定されたmodelはProductではありません。")
	}
	if source.R != nil { // カテゴリあり
		category, err := category.BuildCategory(source.R.Category.ObjID, source.R.Category.Name)
		if err != nil {
			return nil, err
		}
		product, err := product.BuildProduct(source.ObjID, source.Name, uint32(source.Price), category)
		if err != nil {
			return nil, err
		}
		return *product, nil
	} else { // カテゴリ無し
		product, err := product.BuildProduct(source.ObjID, source.Name, uint32(source.Price), nil)
		if err != nil {
			return nil, err
		}
		return *product, nil
	}
}

// コンストラクタ
// 2023/04/01
func NewProductConverterImpl() domain.EntityConverter {
	return &ProductConverterImpl{}
}
