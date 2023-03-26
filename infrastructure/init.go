package infrastructure

// データベース初期化処理を表すインターフェース
// 2023/03/25
type InitDB interface {
	Init(interface{}) interface{}
}
