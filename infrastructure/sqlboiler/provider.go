package sqlboiler

import (
	"sample-service/infrastructure/sqlboiler/category"
	"sample-service/infrastructure/sqlboiler/product"

	"go.uber.org/fx"
)

// SqlBoilerを利用するインフラストラクチャ層の依存定義
// 2023/03/31
var Module = fx.Options(
	fx.Provide(category.NewCategoryConverter, category.NewCategoryRepositiry),
	fx.Provide(product.NewProductConverter, product.NewProductRepository),
)
