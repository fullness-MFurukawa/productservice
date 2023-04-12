package product

import (
	"context"
	"fmt"
	"sample-service/apperrors"
	"sample-service/domain/product"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// 商品アプリケーションサービスインターフェースの実装
// 2023/03/27
type ProductServiceIml struct {
	// 商品アクセスRepositoryインターフェース
	repository product.ProductRepositiry
}

// 商品一覧を取得して提供する
func (service *ProductServiceIml) List() ([]product.Product, error) {
	// 空のContextを生成する
	context := context.Background()
	// トランザクションを開始する
	transaction, err := boil.BeginTx(context, nil)
	if err != nil {
		return nil, apperrors.NewServiceError(err.Error())
	}

	// すべての商品を取得する
	products, err := service.repository.FindAll(context, transaction)
	if err != nil {
		return nil, apperrors.NewServiceError(err.Error())
	} else {
		return products, nil
	}
}

// 指定されたキーワードを含む商品検索結果を提供する
func (service *ProductServiceIml) SearchBykeyword(keyword string) ([]product.Product, error) {
	// 空のContextを生成する
	context := context.Background()
	// トランザクションを開始する
	transaction, err := boil.BeginTx(context, nil)
	if err != nil {
		return nil, apperrors.NewServiceError(err.Error())
	}
	// キーワードで商品を検索する
	products, err := service.repository.FindByNameLike(context, transaction, "%"+keyword+"%")
	if err != nil {
		return nil, apperrors.NewServiceError(err.Error())
	}
	if len(products) == 0 { // キーワードを含む商品が存在しない
		return products, apperrors.NewServiceError(fmt.Sprintf("キーワード:%sを含む商品は見つかりませんでした。", keyword))
	} else {
		return products, nil
	}
}

// 新商品を登録する
func (service *ProductServiceIml) Add(product *product.Product) error {
	// 空のContextを生成する
	context := context.Background()
	// トランザクションを開始する
	transaction, err := boil.BeginTx(context, nil)
	if err != nil {
		return apperrors.NewServiceError(err.Error())
	}

	// 同じ商品名が存在するか確認する
	result, ex_err := service.repository.Exist(context, transaction, product.ProductName())
	if ex_err != nil {
		return apperrors.NewServiceError(err.Error())
	}
	if result { // 既に同じ商品が存在する
		return apperrors.NewServiceError(fmt.Sprintf("商品名:%sは既に登録済みです。", product.ProductName().Value()))
	}

	// 新商品を永続化する
	add_err := service.repository.Create(context, transaction, product)
	if add_err != nil { // 永続化でエラー発生
		return apperrors.NewServiceError(add_err.Error())
	} else { // 永続化成功
		transaction.Commit() // トランザクションをコミットする
		return nil
	}
}

// 商品の内容を変更する
func (service *ProductServiceIml) Change(product *product.Product) (bool, error) {
	// 空のContextを生成する
	context := context.Background()
	// トランザクションを開始する
	transaction, err := boil.BeginTx(context, nil)
	if err != nil {
		return false, apperrors.NewServiceError(err.Error())
	}
	// 商品を更新する
	result, upd_err := service.repository.UpdateById(context, transaction, product)
	if upd_err != nil {
		transaction.Rollback() //　トランザクションをロールバックする
		return false, apperrors.NewServiceError(upd_err.Error())
	}
	if !result && upd_err == nil { // 更新対象が見つからない
		return false, apperrors.NewServiceError(
			fmt.Sprintf("商品番号:%sの商品は存在しないため変更できませんでした。", product.ProductId().Value()))
	}
	transaction.Commit() // トランザクションをコミットする
	return result, nil
}

// 商品を削除
func (service *ProductServiceIml) Remove(productId *product.ProductId) (bool, error) {
	// 空のContextを生成する
	context := context.Background()
	// トランザクションを開始する
	transaction, err := boil.BeginTx(context, nil)
	if err != nil {
		return false, apperrors.NewServiceError(err.Error())
	}
	// 商品を削除する
	result, del_err := service.repository.DeleteById(context, transaction, productId)
	if del_err != nil {
		transaction.Rollback() //　トランザクションをロールバックする
		return false, apperrors.NewServiceError(del_err.Error())
	}
	if !result && del_err == nil { // 削除対象が見つからない
		return false, apperrors.NewServiceError(
			fmt.Sprintf("商品番号:%sの商品は存在しないため削除できませんでした。", productId.Value()))
	}
	transaction.Commit() // トランザクションをコミットする
	return result, nil
}

// ProductServiceインターフェース実装のインスタンスを生成する
// 2023/03/27
func NewProductServiceImpl(repository product.ProductRepositiry) ProductService {
	return &ProductServiceIml{repository: repository}
}
