package sqlboiler

import (
	c "sample-service/infrastructure/sqlboiler/category"
	p "sample-service/infrastructure/sqlboiler/product"

	"go.uber.org/fx"
)

// SqlBoilerを利用するインフラストラクチャ層の依存定義
// 2023/03/31
var Module = fx.Options(
	fx.Provide(c.NewCategoryConverter, c.NewCategoryRepositiry),
	fx.Provide(p.NewProductConverter, p.NewProductRepository),
)
