package container

import (
	"context"
	"fmt"
	"log"
	"sample-service/application/product"
	"sample-service/infrastructure/sqlboiler/tests"
	"time"

	"go.uber.org/fx"
)

// 起動時にStart, 終了にStopと出力するライフサイクルフック
func assignLifeCycleLogging(lc fx.Lifecycle) {
	hook := fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Application Start!")
			// Connection Poolの生成
			go tests.TestDBInit()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Application Stop!")
			return nil
		},
	}
	lc.Append(hook)
}

func executeProductServiceList(service product.ProductService) {
	products, err := service.List()
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}

// applicationまでの依存関係を確認する
func Execute() {
	fmt.Println("Connection Poolの生成!")
	app := fx.New(
		fx.Options(
			product.SrvModeul,
		),
		fx.Invoke(assignLifeCycleLogging, executeProductServiceList),
	)

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
