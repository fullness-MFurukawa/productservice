package product

import "sample-service/domain/product"

// 商品サービスインターフェース
// 2023/03/27
type ProductService interface {
	// 商品一覧を提供する
	List() ([]product.Product, error)
	// キーワード検索結果を提供する
	SearchBykeyword(keyword string) ([]product.Product, error)
	// 新商品を登録する
	Add(product *product.Product) error
}
