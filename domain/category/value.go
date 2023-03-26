package category

import (
	"fmt"
	"regexp"
	"sample-service/domain"
	"unicode/utf8"
)

// インスタンスを識別するUUIDを保持する値オブジェクト
// カテゴリ番号
// 2023/02/25
type CategoryId struct {
	value string
}

// コンストラクタ
func NewCategoryId(value string) (*CategoryId, error) {
	const LENGTH int = 36                                                                          // フィールドの長さ
	const REGEXP string = `([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})` // UUIDの正規表現
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, domain.NewDomainError("カテゴリIDの長さは36文字でなければなりません。")
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, domain.NewDomainError("カテゴリIDはUUIDの形式でなければなりません。")
	}
	return &CategoryId{value: value}, nil
}

// ゲッター
func (instance CategoryId) Value() string {
	return instance.value
}

/*
カテゴリ名を保持する値オブジェクト
2023/02/25
*/
type CategoryName struct {
	value string
}

// コンストラクタ
func NewCategoryName(value string) (*CategoryName, error) {
	const LENGTH int = 20 // フィールドの長さ
	if utf8.RuneCountInString(value) > LENGTH {
		return nil, domain.NewDomainError(fmt.Sprintf("カテゴリ名の長さは%d文字以内です。", LENGTH))
	}
	return &CategoryName{value: value}, nil
}

// ゲッター
func (instance CategoryName) Value() string {
	return instance.value
}
