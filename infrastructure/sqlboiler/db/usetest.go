package db

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func DBInitForTest() (context.Context, *sql.Tx) {
	NewSqlBiolderInitDB().Init(nil)
	context := context.Background()
	transaction, err := boil.BeginTx(context, nil)
	if err != nil {
		panic(err)
	}
	return context, transaction
}
