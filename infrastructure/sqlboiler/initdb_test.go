package sqlboiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// コネクションプール取得テスト
// 2023/03/25
func TestInitDB(t *testing.T) {
	result := NewSqlBiolderInitDB().Init(nil)
	assert.Nil(t, result)
}
