package gin

import (
	"net/http"
	"net/http/httptest"
	"sample-service/infrastructure/sqlboiler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// SqlBuilderのコネクションプールを生成
	sqlboiler.NewSqlBiolderInitDB().Init(nil)
	// Serverのセットアップ
	router := SetupGinServer()

	// ResponseRecoderの生成
	recoder := httptest.NewRecorder()
	// テストリクエスト:商品一覧取得の生成
	req, _ := http.NewRequest("GET", "/product/list", nil)
	// テストリクエストの送信
	router.ServeHTTP(recoder, req)

	// レスポンスステータスの評価
	assert.Equal(t, 200, recoder.Code)
}

func TestSearchKeyword(t *testing.T) {
	// SqlBuilderのコネクションプールを生成
	sqlboiler.NewSqlBiolderInitDB().Init(nil)
	// Serverのセットアップ
	router := SetupGinServer()

	// ResponseRecoderの生成
	recoder := httptest.NewRecorder()
	// テストリクエスト:商品一覧取得の生成
	req, _ := http.NewRequest("GET", "/product/search/ペン", nil)
	// テストリクエストの送信
	router.ServeHTTP(recoder, req)

	// レスポンスステータスの評価
	assert.Equal(t, 200, recoder.Code)
}
