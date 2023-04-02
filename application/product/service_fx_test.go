package product

import (
	"fmt"
	"sample-service/infrastructure/sqlboiler/db"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func executeLit(service ProductService) {
	products, _ := service.List()
	for _, product := range products {
		fmt.Println(product)
	}
}

func TestProductService(t *testing.T) {
	db.DBInitForTest() // コネクションプールを初期化する
	app := fxtest.New(t,
		//	fx.Options(SrvModeul),
		fx.Invoke(executeLit),
	)
	defer app.RequireStart().RequireStop()
	assert.True(t, true)
}
