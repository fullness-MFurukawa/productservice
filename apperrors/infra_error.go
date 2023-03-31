package apperrors

// 内部エラーを表すエラー型
type InternalError struct {
	message string
	err     error // エラーをラップ
}

// エラーメッセージを提供する
func (e *InternalError) Error() string {
	return e.message
}

// エラーをアンラップする
func (e *InternalError) Unwrap() error {
	return e.err
}

// コンストラクタ
func NewInternalError(message string, err error) error {
	return &InternalError{message: message, err: err}
}
