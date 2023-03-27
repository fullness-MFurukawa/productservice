package product

import "go.uber.org/fx"

// ProductAdapterとProductRepositoryの依存定義
var Module = fx.Provide(NewProductAdapter, NewProductRepository)
