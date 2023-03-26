package tests

import (
	"context"
	"database/sql"
	"sample-service/infrastructure/sqlboiler"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestDBInit(t *testing.T) (context.Context, *sql.Tx) {
	// データベース接続とConenction Poolの生成
	_ = sqlboiler.SqlBiolderInitDB{}.Init(nil)
	ctx := context.Background()
	transaction, tran_err := boil.BeginTx(ctx, nil)
	if tran_err != nil {
		assert.Fail(t, tran_err.Error())
	}
	return ctx, transaction
}
