package provider

import (
	ap "sample-service/application/product"
	"sample-service/domain"
	"sample-service/presentation/category"
	"sample-service/presentation/product"

	"go.uber.org/fx"
)

type PresenParams struct {
	fx.In
	CategoryDtoConverter domain.EntityConverter   `name:"categorydtoconverter"`
	PeoductDtoConverter  domain.EntitiesConverter `name:"productdtoconverter"`
}

func NewCategoryDtoConverterFx() domain.EntityConverter {
	return category.NewCategoryDtoConverter()
}

func NewProductDtoConverterFx() domain.EntitiesConverter {
	return product.NewProductDtoConverter()
}

func NewProductControllerFx(service ap.ProductService, params PresenParams) *product.ProductController {
	return product.NewProductController(service, params.PeoductDtoConverter)
}
