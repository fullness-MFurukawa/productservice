package presentation

import (
	"sample-service/presentation/gin"
	"sample-service/presentation/provider"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		provider.NewProductControllerFx,
		fx.Annotated{Name: "categorydtoconverter", Target: provider.NewCategoryDtoConverterFx},
		fx.Annotated{Name: "productdtoconverter", Target: provider.NewProductDtoConverterFx},
	),
	fx.Invoke(gin.RegisterRoutes),
	fx.Invoke(gin.RegisterHooks),
)
