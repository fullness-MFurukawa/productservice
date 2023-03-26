package domain

// Entityを表すインターフェース
// 2023/02/25
type Entity interface {
	// 等価性検証メソッド
	Equals(obj interface{}) (bool, error)
}
