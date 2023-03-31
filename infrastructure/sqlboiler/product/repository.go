package product

import (
	"context"
	"database/sql"
	"sample-service/apperrors"
	"sample-service/domain"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// ProductRepositoryインターフェースの実装
// 2023/03/27
type ProductRepositoryImpl struct {
	converter domain.EntityConverter // Categoryデータ変換Adapter
}

// すべての商品を取得する
func (rep *ProductRepositoryImpl) FindAll(ctx context.Context, tran *sql.Tx) ([]product.Product, error) {
	// 商品を全件取得する
	results, err := models.Products(qm.Load("Category")).All(ctx, tran)
	if err != nil { // エラーなら内部エラーを通知する
		return nil, apperrors.NewInternalError("内部エラー", err)
	}
	var products []product.Product // Entity Productのスライス
	for _, result := range results {
		// modelからEntity Productに変換する
		product_inf, err := rep.converter.Restore(result)
		if err != nil {
			return nil, err
		}
		// 変換したEntity Productをスライスに追加する
		products = append(products, product_inf.(product.Product))
	}
	return products, nil
}

// 指定された商品名で部分一致検索する
func (rep *ProductRepositoryImpl) FindByNameLike(ctx context.Context, tran *sql.Tx, keyword string) ([]product.Product, error) {
	results, err := models.Products(qm.Where("name like ?", keyword), qm.Load("Category")).All(ctx, tran)
	if err != nil { // エラーなら内部エラーを通知する
		return nil, apperrors.NewInternalError("内部エラー", err)
	}
	var products []product.Product // Entity Productのスライス
	for _, result := range results {
		// modelからEntity Productに変換する
		product_inf, err := rep.converter.Restore(result)
		if err != nil {
			return nil, err
		}
		// 変換したEntity Productをスライスに追加する
		products = append(products, product_inf.(product.Product))
	}
	return products, nil
}

// 商品名で存在確認する
func (rep *ProductRepositoryImpl) Exist(ctx context.Context, tran *sql.Tx, name *product.ProductName) (bool, error) {
	_, err := models.Products(qm.Where("name = ?", name.Value())).One(ctx, tran)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, apperrors.NewInternalError("内部エラー", err)
		}
	}
	return true, nil
}

// 新しい商品を登録する
func (rep *ProductRepositoryImpl) Create(ctx context.Context, tran *sql.Tx, product *product.Product) error {
	model, err := rep.converter.Convert(product)
	if err != nil {
		return err
	}
	product_model := model.(models.Product)
	ins_err := product_model.Insert(ctx, tran, boil.Infer())
	if ins_err != nil {
		return apperrors.NewInternalError("内部エラー", err)
	}
	return nil
}

// 指定された商品番号で商品内容を変更する
func (rep *ProductRepositoryImpl) UpdateById(ctx context.Context, tran *sql.Tx, product *product.Product) (bool, error) {
	up_model, err := models.Products(qm.Where("obj_id = ?", product.ProductId().Value())).One(ctx, tran)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, apperrors.NewInternalError("内部エラー", err)
		}
	}
	up_model.Name = product.ProductName().Value()
	up_model.Price = int(product.ProductPrice().Value())
	_, err = up_model.Update(ctx, tran, boil.Infer())
	if err != nil {
		return false, apperrors.NewInternalError("内部エラー", err)
	}
	return true, nil
}

// 指定された商品番号で商品を削除する
func (rep *ProductRepositoryImpl) DeleteById(ctx context.Context, tran *sql.Tx, id *product.ProductId) (bool, error) {
	del_model, err := models.Products(qm.Where("obj_id = ?", id.Value())).One(ctx, tran)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		} else {
			return false, apperrors.NewInternalError("内部エラー", err)
		}
	}
	_, del_err := del_model.Delete(ctx, tran)
	if del_err != nil {
		return false, apperrors.NewInternalError("内部エラー", err)
	}
	return true, nil
}

// コンストラクタ
func NewProductRepositoryImpl() product.ProductRepositiry {
	return &ProductRepositoryImpl{converter: NewProductConverter()}
}
