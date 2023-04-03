package provider

import (
	"sample-service/domain"

	"go.uber.org/fx"
)

type ConverterParams struct {
	fx.In
	CategoryConverter domain.EntityConverter `name:"categoryconverter"`
	ProductConverter  domain.EntityConverter `name:"productconverter"`
}
