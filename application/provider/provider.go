package provider

import (
	"sample-service/application/product"
	"sample-service/infrastructure/sqlboiler/provider"

	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	provider.InfraModule,
	fx.Provide(product.NewProductServiceImpl),
)
