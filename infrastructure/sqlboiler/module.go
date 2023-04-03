package sqlboiler

import (
	"sample-service/infrastructure/sqlboiler/provider"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		provider.NewCategoryRepositiryFx,
		provider.NewProductRepositoryFx,
		fx.Annotated{Name: "categoryconverter", Target: provider.NewCategoryConverterFx},
		fx.Annotated{Name: "productconverter", Target: provider.NewProductConverterFx},
	),
)
