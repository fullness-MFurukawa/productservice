package category

import "go.uber.org/fx"

// CategpryAdapterとCategoryRepositoryの依存関係定義
// 2023/03/25
var Module = fx.Provide(NewCategoryAdapater, NewCategoryRepositiry)
