package product

import (
	"context"
	"fmt"
	"sample-service/application"
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
		return nil, application.NewServiceError(err.Error())
	}

	// すべての商品を取得する
	products, err := service.repository.FindAll(context, transaction)
	if err != nil {
		return nil, application.NewServiceError(err.Error())
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
		return nil, application.NewServiceError(err.Error())
	}
	// キーワードで商品を検索する
	products, err := service.repository.FindByNameLike(context, transaction, "%"+keyword+"%")
	if err != nil {
		return nil, application.NewServiceError(err.Error())
	}
	if len(products) == 0 { // キーワードを含む商品が存在しない
		return products, application.NewServiceError(fmt.Sprintf("キーワード:%sを含む商品は見つかりませんでした。", keyword))
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
		return application.NewServiceError(err.Error())
	}
	// 新商品を永続化する
	add_err := service.repository.Create(context, transaction, product)
	if add_err != nil { // 永続化でエラー発生
		transaction.Rollback() //　トランザクションをロールバックする
		return application.NewServiceError(add_err.Error())
	} else { // 永続化成功
		transaction.Commit() // トランザクションをコミットする
		return nil
	}
}
