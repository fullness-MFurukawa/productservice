package category

import (
	"context"
	"database/sql"
)

// カテゴリをアクセスするRepositoryインターフェース
// 2023/03/25
type CategoryRepository interface {
	FindAll(ctx context.Context, tran *sql.Tx) ([]Category, error)
	FindById(ctx context.Context, tran *sql.Tx, id *CategoryId) (*Category, error)
}
