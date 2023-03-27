package application

// アプリケーションサービスエラーを表すエラー型
// 2030/03/27
type ServiceError struct {
	message string
}

// エラーメッセージを提供する
func (e *ServiceError) Error() string {
	return e.message
}

// コンストラクタ
func NewServiceError(message string) error {
	return &ServiceError{message: message}
}
