package infrastructure

// 内部エラーを表すエラー型
type InternalError struct {
	message string
}

// エラーメッセージを提供する
func (e *InternalError) Error() string {
	return e.message
}

// コンストラクタ
func NewInternalError(message string) error {
	return &InternalError{message: message}
}
