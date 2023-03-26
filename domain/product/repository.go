package product

import (
	"context"
	"database/sql"
)

// 商品をアクセスするRepositoryインターフェース
// 2023/03/25
type ProductRepositiry interface {
	// すべての商品を取得する
	FindAll(ctx context.Context, tran *sql.Tx) ([]Product, error)
	// 指定された商品名で部分一致検索する
	FindByNameLike(ctx context.Context, tran *sql.Tx, keyword string) ([]Product, error)
	// 商品名で存在確認する
	Exist(ctx context.Context, tran *sql.Tx, name *ProductName) (bool, error)
	// 新しい商品を登録する
	Create(ctx context.Context, tran *sql.Tx, product *Product) error
	// 指定された商品番号で商品内容を変更する
	UpdateById(ctx context.Context, tran *sql.Tx, product *Product) (bool, error)
	// 指定された商品番号で商品を削除する
	DeleteById(ctx context.Context, tran *sql.Tx, id *ProductId) (bool, error)
}
