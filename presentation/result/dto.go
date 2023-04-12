package result

// 処理結果を扱うDTO
type ResultDto struct {
	Function string `json:"function"`
	Message  string `json:"message"`
}

// リクエスト実行正常終了結果を扱うDTO
// 2023/04/05
type SuccessDto struct {
	ResultDto
}

// コンストラクタ
func NewSuccessDto(function string, message string) *SuccessDto {
	return &SuccessDto{ResultDto{Function: function, Message: message}}
}

type FailureDto struct {
	ResultDto
}

// コンストラクタ
func NewFailureDto(function string, message string) *FailureDto {
	return &FailureDto{ResultDto{Function: function, Message: message}}
}
