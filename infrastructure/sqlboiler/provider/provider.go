package provider

import (
	"sample-service/infrastructure/sqlboiler/converter"
	"sample-service/infrastructure/sqlboiler/repository"

	"go.uber.org/fx"
)

var InfraModule = fx.Options(
	fx.Provide(converter.NewCategoryConverterImpl, repository.NewCategoryRepositiryImpl),
	fx.Provide(converter.NewProductConverterImpl, repository.NewProductRepositoryImpl),
)
