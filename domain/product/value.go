package product

import (
	"fmt"
	"regexp"
	"sample-service/domain"
	"unicode/utf8"
)

// 商品番号値オブジェクト
// インスタンスを識別するUUIDを保持する値オブジェクト
// 2023/02/25
type ProductId struct {
	value string
}

// コンストラクタ
func NewProductId(value string) (*ProductId, error) {
	const LENGTH int = 36                                                                          // フィールドの長さ
	const REGEXP string = `([0-9a-f]{8})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{4})-([0-9a-f]{12})` // UUIDの正規表現
	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, domain.NewDomainError(fmt.Sprintf("商品IDの長さは%d文字でなければなりません。", LENGTH))
	}
	// 引数の正規表現(UUID)チェック
	if !regexp.MustCompile(REGEXP).Match([]byte(value)) {
		return nil, domain.NewDomainError("商品IDはUUIDの形式でなければなりません。")
	}
	return &ProductId{value: value}, nil
}

// ゲッター
func (val ProductId) Value() string {
	return val.value
}

// 商品名を表す値オブジェクト
// 2023/02/25
type ProductName struct {
	value string
}

// コンストラクタ
func NewProductName(value string) (*ProductName, error) {
	const LENGTH int = 30 // フィールドの長さ
	if utf8.RuneCountInString(value) > LENGTH {
		return nil, domain.NewDomainError(fmt.Sprintf("商品名の長さは%d文字以内です。", LENGTH))
	}
	return &ProductName{value: value}, nil
}

// ゲッター
func (val ProductName) Value() string {
	return val.value
}

// 商品単価を表す値オブジェクト
// 2023/02/25
type ProductPrice struct {
	value uint32
}

// コンストラクタ
func NewProductPrice(value uint32) (*ProductPrice, error) {
	const MIN = 50
	const MAX = 10000
	if value >= MIN && value <= MAX {
		return &ProductPrice{value: uint32(value)}, nil
	} else {
		return nil, domain.NewDomainError(fmt.Sprintf("単価は%d以上、%d以下です。", MIN, MAX))
	}
}

// ゲッター
func (val ProductPrice) Value() uint32 {
	return val.value
}
