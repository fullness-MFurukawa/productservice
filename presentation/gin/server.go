package gin

import (
	"context"
	"fmt"
	"sample-service/application"
	"sample-service/infrastructure/sqlboiler"
	"sample-service/infrastructure/sqlboiler/category"
	"sample-service/presentation/product"

	"go.uber.org/fx"
)

// 依存関係の定義
// 2023/03/31
var Module = fx.Options(
	application.Module,
	fx.Provide(NewHandler, category.NewCategoryConverter, product.NewProductDtoConverter, product.NewProductController),
	fx.Invoke(RegisterRoutes),
	fx.Invoke(registerHooks),
)

func registerHooks(lifecycle fx.Lifecycle, h *Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting application Post:8080 !!!")
				// SqlBiolderのConnection Poolを生成
				go sqlboiler.NewSqlBiolderInitDB().Init(nil)
				go h.Gin.Run(":8080") //	Ginの起動
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application !!!")
				return nil
			},
		},
	)
}
