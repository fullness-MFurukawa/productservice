package product

import (
	"sample-service/infrastructure/sqlboiler/product"

	"go.uber.org/fx"
)

// ProductServiceの依存定義
var SrvModeul = fx.Provide(
	product.NewProductAdapter,
	product.NewProductRepository,
	NewProductService)
