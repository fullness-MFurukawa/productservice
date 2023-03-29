package tests

import (
	"context"
	"database/sql"
	"sample-service/infrastructure/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestDBInit() (context.Context, *sql.Tx) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.NewSqlBiolderInitDB().Init(nil)
	ctx := context.Background()
	transaction, _ := boil.BeginTx(ctx, nil)
	return ctx, transaction
}
