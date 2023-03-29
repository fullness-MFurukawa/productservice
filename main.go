package main

import (
	"sample-service/infrastructure/sqlboiler"
	"sample-service/presentation/gin"
)

func main() {
	// SqlBuilderのコネクションプールを生成
	sqlboiler.NewSqlBiolderInitDB().Init(nil)
	// Ginの起動
	gin.RunGinServer()
}
