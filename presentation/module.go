package presentation

import (
	"sample-service/application"
	"sample-service/presentation/gin"
	"sample-service/presentation/provider"

	"go.uber.org/fx"
)

var Module = fx.Options(
	application.Module,
	fx.Provide(
		gin.NewRouter,
		provider.NewProductControllerFx,
		fx.Annotated{Name: "categorydtoconverter", Target: provider.NewCategoryDtoConverterFx},
		fx.Annotated{Name: "productdtoconverter", Target: provider.NewProductDtoConverterFx},
	),
	fx.Invoke(gin.RegisterRouter),
	fx.Invoke(gin.RegisterHooks),
)
