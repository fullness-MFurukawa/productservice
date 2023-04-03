package main

import (
	"sample-service/presentation"

	"go.uber.org/fx"
)

// @title 商品アクセスAPIサンプル V1
// @version 1.0
// @description 商品および商品カテゴリを管理するAPIサービス
// @termsOfService http://localhost:8081

// @contact.name Fullness,Inc.
// @contact.url https://www.fullness.co.jp
// @contact.email furukawa@fullness.co.jp

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/

// @host localhost:8081
// @BasePath /product
func main() {
	fx.New(
		presentation.Module,
	).Run()
}
