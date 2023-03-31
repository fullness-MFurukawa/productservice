package application

import (
	"sample-service/application/product"
	"sample-service/infrastructure/sqlboiler"

	"go.uber.org/fx"
)

var Module = fx.Options(
	sqlboiler.Module,
	fx.Provide(product.NewProductService),
)
