package sqlboiler

import (
	"sample-service/infrastructure/sqlboiler/provider"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		provider.NewCategoryRepositiryImpl,
		provider.NewProductRepositoryImpl,
		fx.Annotated{Name: "categoryconverter", Target: provider.NewCategoryConverterImpl},
		fx.Annotated{Name: "productconverter", Target: provider.NewProductConverterImpl},
	),
)
