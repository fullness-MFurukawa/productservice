package category

import (
	"context"
	"database/sql"
	"sample-service/domain"
	"sample-service/domain/category"
	"sample-service/infrastructure"
	"sample-service/infrastructure/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CategoryテーブルアクセスRepositoryインターフェースの実装
// 2023/03/25
type CategoryRepositoryImpl struct {
	converter domain.EntityConverter // Categoryデータ変換Converter
}

// カテゴリを全件取得する
func (rep *CategoryRepositoryImpl) FindAll(ctx context.Context, tran *sql.Tx) ([]category.Category, error) {
	results, err := models.Categories().All(ctx, tran)
	if err != nil {
		return nil, infrastructure.NewInternalError("内部エラー", err)
	}
	var categories []category.Category
	for _, result := range results {
		category_inf, err := rep.converter.Restore(result)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category_inf.(category.Category))
	}
	return categories, nil
}

// 指定された番号のカテゴリを取得する
func (rep *CategoryRepositoryImpl) FindById(ctx context.Context, tran *sql.Tx, id *category.CategoryId) (*category.Category, error) {
	result, err := models.Categories(qm.Where("category.obj_id = ?", id.Value())).One(ctx, tran)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, infrastructure.NewInternalError("内部エラー", err)
		}
	}
	category_inf, err := rep.converter.Restore(result)
	if err != nil {
		return nil, err
	}
	category := category_inf.(category.Category)
	return &category, nil
}
func NewCategoryRepositoryImpl() category.CategoryRepository {
	return &CategoryRepositoryImpl{converter: NewCategoryConverter()}
}
