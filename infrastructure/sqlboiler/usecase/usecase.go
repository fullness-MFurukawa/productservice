package usecase

import (
	"context"
	"fmt"
	"log"
	"sample-service/domain/product"
	"sample-service/infrastructure/sqlboiler"
	"sample-service/infrastructure/sqlboiler/db"
	"time"

	"go.uber.org/fx"
)

// 起動時にStart, 終了にStopと出力するライフサイクルフックです。
func lifeCycleLogging(lc fx.Lifecycle) {
	hook := fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("fx Start!")
			go db.NewSqlBiolderInitDB().Init(nil)
			fmt.Println("Connection Pool Create!!!")
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("fx Stop!")
			return nil
		},
	}
	lc.Append(hook)
}

func executeProductRepository(repository product.ProductRepositiry) {
	ctx, tran := db.DBInitForTest()

	products, err := repository.FindAll(ctx, tran)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}

func Execute() {
	app := fx.New(
		sqlboiler.Module,
		fx.Invoke(lifeCycleLogging, executeProductRepository),
	)

	// エラーでないことを確認する
	if err := app.Err(); err != nil {
		log.Fatalf(err.Error())
	}

	startCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := app.Start(startCtx); err != nil { // fx.Appを実行する
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
