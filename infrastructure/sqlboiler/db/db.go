package db

import (
	"database/sql"
	"sample-service/infrastructure"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// SqlBoiler データベース接続
// Connection Pool生成
// 2023/03/25
type SqlBiolderInitDB struct{}

// コネクションプールを生成する
func (instance *SqlBiolderInitDB) Init(interface{}) interface{} {
	conn, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/sample_db")
	if err != nil {
		panic(err) // エラーが発生したらパニックにする
	}
	// コネクションプールの設定
	conn.SetMaxIdleConns(10)                   // 初期接続数
	conn.SetMaxOpenConns(100)                  // 最大接続数
	conn.SetConnMaxLifetime(300 * time.Second) // 最大利用生存期間

	boil.SetDB(conn)      // グローバルコネクション設定
	boil.DebugMode = true // デバッグモードに設定 生成されたSQLを出力する
	return nil
}

// コンストラクタ
// 後で全体構成を考慮してリファクタリングする
func NewSqlBiolderInitDB() infrastructure.InitDB {
	return &SqlBiolderInitDB{}
}
