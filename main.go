package main

import (
	"sample-service/infrastructure/sqlboiler"
	"sample-service/presentation/gin"
)

func main() {
	// SqlBuilderのコネクションプールを生成
	sqlboiler.NewSqlBiolderInitDB().Init(nil)
	// Ginのセットアップと起動
	gin.SetupGinServer().Run(":8081")
}
